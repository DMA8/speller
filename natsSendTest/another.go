package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	Addresses          = "ngx-api-r01-03.dp.wb.ru:4242,ngx-api-r03-03.dl.wb.ru:4242,ngx-api-r04-03.dl.wb.ru:4242,ngx-api-r04-03.dp.wb.ru:4242,ngx-api-r05-03.dp.wb.ru:4242"
	SearchEventSubject = "wbxsearch.ru.exactmatch.common.searchevent"
)

type ResponseType int

const (
	ResponseTypeCatalog ResponseType = iota
	ResponseTypePreset
	ResponseTypeExtendSearch
	ResponseTypeOnlineSearch
)

type SearchResult struct {
	URL          string
	Resource     string
	ShardKey     string
	Query        string
	ResponseType ResponseType
}

type SearchEvent struct {
	Timestamp time.Time
	Query     string
	Category  string
	Result    SearchResult
}

type BadSearchEvent struct {
	Timestamp time.Time
	Query     string
	Error     string
}

func main() {
	conn, err := nats.Connect(Addresses)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	_, err = conn.Subscribe(
		SearchEventSubject, func(msg *nats.Msg) {
			defer wg.Wait()

			var data SearchEvent

			err := json.Unmarshal(msg.Data, &data)
			if err != nil {
				fmt.Println(string(msg.Data))
				log.Fatal("ERROR: " + err.Error())
			}

			fmt.Println(data)
		})
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	conn.Close()

}
