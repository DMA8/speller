package main

import (
	"context"
	"log"

	//	"log"
	"os"
	"os/signal"

	//	"spellCheck/internal/handlerFastHTTP"
	natsCl "spellCheck/internal/natsClient"
	"spellCheck/internal/speller"

	//"spellCheck/internal/dumpS3"

	"spellCheck/internal/storage"
	//	"github.com/valyala/fasthttp"
)

//2.6 Gb model is up -> 4,5 Gb after stresstest

func main() {
	wait := make(chan struct{})
	natsToSpeller := make(chan natsCl.BadMessage)
	spellerToStorage := make(chan storage.Spelling)
	dumpDone := make(chan struct{})
	natsDisconnected := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	//Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		cancel()
		<-dumpDone
		<-natsDisconnected
		log.Println(speller.AllTests, speller.SuccessSpellerAct, float64(speller.SuccessSpellerAct)/float64(speller.AllTests))
		os.Exit(0)
	}()
	//dumpS3.Dump()
	go natsCl.Start(ctx, natsToSpeller, natsDisconnected)

	myStorage := storage.NewStorage("spellcheck1.csv")
	//r := handlerFastHTTP.ConfiguredRouter(myStorage)
	go myStorage.Dump(ctx, dumpDone, 30)
	go speller.AcceptMessage(ctx, natsToSpeller, spellerToStorage)
	go myStorage.AcceptSpellerSuggest(ctx, spellerToStorage)
	<-wait
	//log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
