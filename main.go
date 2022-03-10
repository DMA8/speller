package main

import (
	"log"
	"speller/internal/handlerFastHTTP"

	//	"./internal/handler"
	"speller/internal/storage"

	"github.com/valyala/fasthttp"
)

func main() {
	a := storage.NewStorage("spellcheck.csv")
	r := handlerFastHTTP.ConfiguredRouter(a)
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
