package main

import (
	"context"
	"log"
	"runtime/pprof"
	"os"
	"os/signal"

	"spellCheck/internal/handlerFastHTTP"
	natsCl "spellCheck/internal/natsClient"
	"spellCheck/internal/speller"
	"spellCheck/internal/storage"

	"github.com/valyala/fasthttp"
)

//2.6 Gb model is up -> 4,5 Gb after stresstest

func main() {
	natsToSpeller := make(chan natsCl.BadMessage)
	spellerToStorage := make(chan storage.Spelling)
	dumpDone := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	
	//Profiling
	p1, _ := os.Create("heap_before.pprof")
	pprof.Lookup("heap").WriteTo(p1, 0)
	p1.Close()
	p2, _ := os.Create("heap_after.pprof")

	//Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		pprof.Lookup("heap").WriteTo(p2, 0)
		cancel()
		p2.Close()
		<-dumpDone
		os.Exit(0)
	}()
	
	//go natsCl.Start(ctx, natsCl.NatsAddress1, natsCl.BadSearchEventSubject, natsToSpeller)
	go natsCl.Start(ctx, "localhost:4222", natsCl.BadSearchEventSubject, natsToSpeller)	
	myStorage := storage.NewStorage("spellcheck.csv")
	r := handlerFastHTTP.ConfiguredRouter(myStorage)
	go myStorage.Dump(ctx, dumpDone, 1) // last arg is a dump cycle
	go speller.AcceptMessage(ctx, natsToSpeller, spellerToStorage)
	go myStorage.AcceptSpellerSuggest(ctx, spellerToStorage)
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
