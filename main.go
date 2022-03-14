package main

import (
	"fmt"
	"log"

	"spellCheck/internal/handlerFastHTTP"
	"spellCheck/internal/natsClient"
	"spellCheck/internal/storage"

	"github.com/Saimunyz/speller"
	"github.com/valyala/fasthttp"
)

func main() {
	a := storage.NewStorage("spellcheck.csv")
	r := handlerFastHTTP.ConfiguredRouter(a)
	s := speller.NewSpeller()
	err := s.LoadModel("./models/small-data.gz")
	if err != nil {
		log.Fatal(err)
	}
	ans := s.SpellCorrect("акно")
	fmt.Println(ans)
	natsClient.Start()
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
