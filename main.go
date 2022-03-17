package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"spellCheck/internal/handlerFastHTTP"
	//natsStream "spellCheck/internal/natsStreamingClient"
	nats "spellCheck/internal/natsClient"
	"spellCheck/internal/speller"
	"spellCheck/internal/storage"

	"github.com/valyala/fasthttp"
)

// TODO:
// 2. Переписать дамп. Протестировать дамп
// 3. Добавить комментарии. Попробовать swagger

func main() {
	// natsToSpeller := make(chan natsStream.BadMessage)
	natsToSpeller := make(chan nats.BadMessage)
	
	spellerToStorage := make(chan storage.Spelling)
	dumpDone := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c      // waiting for Ctrl+C
		cancel() // send signal for every goRoutine with ctx
		<-dumpDone
		os.Exit(0)
	}()

	//go natsStream.Start(ctx, "ngx-api-r01-03.dp.wb.ru:4242,ngx-api-r03-03.dl.wb.ru:4242,ngx-api-r04-03.dl.wb.ru:4242,ngx-api-r04-03.dp.wb.ru:4242,ngx-api-r05-03.dp.wb.ru:4242", "client-123", "wbxsearch.ru.exactmatch.common.searchevent", natsToSpeller)
	go nats.Start(ctx, "localhost:4222", nats.BadSearchEventSubject, natsToSpeller)
	go nats.Start(ctx, "localhost:4222", nats.BadSearchEventSubject, natsToSpeller)
	go nats.Start(ctx, "localhost:4222", nats.BadSearchEventSubject, natsToSpeller)
	// go nats.Start(ctx, "localhost:4222", nats.BadSearchEventSubject, natsToSpeller)
	// go nats.Start(ctx, "localhost:4222", nats.BadSearchEventSubject, natsToSpeller)
	// go nats.Start(ctx, "localhost:4222", nats.BadSearchEventSubject, natsToSpeller)
	// go nats.Start(ctx, "localhost:4222", nats.BadSearchEventSubject, natsToSpeller)
	// go nats.Start(ctx, "localhost:4222", nats.BadSearchEventSubject, natsToSpeller)
	// go nats.Start(ctx, "localhost:4222", nats.BadSearchEventSubject, natsToSpeller)
	
	myStorage := storage.NewStorage("spellcheck.csv")
	r := handlerFastHTTP.ConfiguredRouter(myStorage)
	// go natsClient.Start(ctx, "test-cluster", "client1", "foo", natsToSpeller)
	go myStorage.Dump(ctx, dumpDone, 1)
	go speller.AcceptMessage(ctx, natsToSpeller, spellerToStorage)
	go myStorage.AcceptSpellerSuggest(ctx, spellerToStorage)
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
