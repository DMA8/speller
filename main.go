package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"spellCheck/internal/handlerFastHTTP"
	"spellCheck/internal/natsClient"
	"spellCheck/internal/speller"
	"spellCheck/internal/storage"

	"github.com/valyala/fasthttp"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background()) 
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c      // waiting for Ctrl+C
		cancel() // send signal for every goRoutine with ctx
		time.Sleep(time.Second * 5)
		os.Exit(0)
	}()
	natsToSpeller := make(chan natsClient.BadMessage, 0)
	spellerToStorage := make(chan storage.Spelling, 0)
	natsClient.Start2(natsToSpeller)
	myStorage := storage.NewStorage("spellcheck.csv")
	r := handlerFastHTTP.ConfiguredRouter(myStorage)
	go myStorage.Dump(ctx)
	go speller.AcceptMessage(ctx, natsToSpeller, spellerToStorage)
	go myStorage.AcceptSpellerSuggest(ctx, spellerToStorage)
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
