package natsStreamingClient

import (
	"context"
	"log"
	protoType "spellCheck/natsSendTest/proto"

	"github.com/golang/protobuf/proto"

	//"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/nats-io/nats.go"
)

type BadMessage struct {
	Query string `json:"query"`
}

const (
	Client                      = "speller"
	Client2                     = "client-1"
	TestSubj                    = "foo"
	NatsAddress1                = "ngx-api-r01-03.dp.wb.ru:4242,ngx-api-r02-03.dl.wb.ru:4242,ngx-api-r03-03.dl.wb.ru:4242,ngx-api-r04-03.dl.wb.ru:4242,ngx-api-r05-03.dp.wb.ru:4242"
	NatsAddressTest             = "test-cluster"
	BadSearchEventSubject       = "wbxsearch.ru.exactmatch.common.badsearchevent"
	BadSearchEventQueryCapacity = 1024
	SearchEventSubject          = "wbxsearch.ru.exactmatch.common.searchevent"
	SearchEventQueryCapacity    = 1024
)

func Start(ctx context.Context, address, topic string, channel chan<- BadMessage) {
	conn, err := nats.Connect(address)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("nats CONNECTED: %s", address)
	sub, err := conn.Subscribe(topic, func(m *nats.Msg) {
		var badMessageProto protoType.BadSearchEvent
		var badMessage BadMessage
		err := proto.Unmarshal(m.Data, &badMessageProto)
		if err != nil {
			log.Print(err)
			return
		}
		badMessage.Query = badMessageProto.Query
		channel <- badMessage
	})
	if err != nil {
		log.Println(err)
	}
	log.Printf("nats SUB: %s", topic)
	<-ctx.Done()
	err = sub.Unsubscribe()
	log.Printf("nats UNSUB: %s", topic)
	if err != nil {
		log.Println(err)
	}
	conn.Close()
	log.Printf("nats DISCONNECTED: %s ", address)
}
