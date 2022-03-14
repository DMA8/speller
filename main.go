package main

import (
	"log"

	"spellCheck/internal/handlerFastHTTP"
	"spellCheck/internal/natsClient"
	"spellCheck/internal/storage"

	"github.com/valyala/fasthttp"
)

func main() {
	a := storage.NewStorage("spellcheck.csv")
	r := handlerFastHTTP.ConfiguredRouter(a)
	natsClient.Start()
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
