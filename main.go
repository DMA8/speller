package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"spellCheck/internal/handlerFastHTTP"
	natsCl "spellCheck/internal/natsClient"
	"spellCheck/internal/speller"
	"spellCheck/internal/dumpS3"

	"spellCheck/internal/storage"

	"github.com/valyala/fasthttp"
)

//2.6 Gb model is up -> 4,5 Gb after stresstest

func main() {
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
		os.Exit(0)
	}()
	dumpS3.Dump()
	go natsCl.Start(ctx, natsToSpeller, natsDisconnected)
	// go natsCl.Start(ctx, "localhost:4222", natsCl.BadSearchEventSubject, natsToSpeller)
	myStorage := storage.NewStorage("spellcheck.csv")
	r := handlerFastHTTP.ConfiguredRouter(myStorage)
	go myStorage.Dump(ctx, dumpDone, 1) // last arg is a dump cycle
	go speller.AcceptMessage(ctx, natsToSpeller, spellerToStorage)
	go myStorage.AcceptSpellerSuggest(ctx, spellerToStorage)
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
