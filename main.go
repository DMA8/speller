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

// TODO:
// 2. Переписать дамп. Протестировать дамп
// 3. Добавить комментарии. Попробовать swagger


func main() {
	natsToSpeller := make(chan natsClient.BadMessage)
	spellerToStorage := make(chan storage.Spelling)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c      // waiting for Ctrl+C
		cancel() // send signal for every goRoutine with ctx
		time.Sleep(time.Second * 5)
		os.Exit(0)
	}()

	myStorage := storage.NewStorage("spellcheck.csv")
	r := handlerFastHTTP.ConfiguredRouter(myStorage)
	go natsClient.Start(ctx, "test-cluster", "client1", "foo", natsToSpeller)
	go myStorage.Dump(ctx, 1)
	go speller.AcceptMessage(ctx, natsToSpeller, spellerToStorage)
	go myStorage.AcceptSpellerSuggest(ctx, spellerToStorage)
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
