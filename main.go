package main

import (
	"context"
	"log"

	"spellCheck/internal/handlerFastHTTP"
	"spellCheck/internal/natsClient"
	"spellCheck/internal/speller"
	"spellCheck/internal/storage"

	"github.com/valyala/fasthttp"
)

func main() {
	natsToSpeller := make(chan natsClient.BadMessage, 0)
	spellerToStorage := make(chan storage.Spelling, 0)
	myStorage := storage.NewStorage("spellcheck.csv")
	r := handlerFastHTTP.ConfiguredRouter(myStorage)
	natsClient.Start(natsToSpeller)
	go speller.AcceptMessage(context.Background(), natsToSpeller, spellerToStorage)
	go myStorage.AcceptSpellerSuggest(context.Background(), spellerToStorage)
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
