package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	//	"github.com/protocolbuffers/protobuf-go"
	// "google.golang.org/protobuf/encoding/protojson"
	"github.com/golang/protobuf/proto" 
	"github.com/nats-io/nats.go"
	myproto "spellCheck/natsSendTest/proto"
	//	"github.com/nats-io/nats.go/encoders/protobuf"
)

const (
	Addresses             = "ngx-api-r01-03.dp.wb.ru:4242,ngx-api-r03-03.dl.wb.ru:4242,ngx-api-r04-03.dl.wb.ru:4242,ngx-api-r04-03.dp.wb.ru:4242,ngx-api-r05-03.dp.wb.ru:4242"
	SearchEventSubject    = "wbxsearch.ru.exactmatch.common.searchevent"
	BadSearchEventSubject = "wbxsearch.ru.exactmatch.common.badsearchevent"
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
	fmt.Println("first line")
	conn, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("second line")

	wg := sync.WaitGroup{}

	wg.Add(1)
	_, err = conn.Subscribe(
		SearchEventSubject, func(msg *nats.Msg) {
		//	defer wg.Wait()

			var data myproto.BadSearchEvent
			log.Println("HERE")
			err := proto.Unmarshal(msg.Data, &data)
			fmt.Println(data)
			if err != nil {
				fmt.Println(string(msg.Data))
				log.Fatal("ERROR: " + err.Error())
			}
		})
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	conn.Close()

}
