package natsStreamingClient

import (
	"context"
	"encoding/json"
	"fmt"
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
	Client = "speller"
	Client2 = "client-1"
	TestSubj = "foo"
	NatsAddress1 = "ngx-api-r01-03.dp.wb.ru:4222,ngx-api-r02-03.dl.wb.ru:4222,ngx-api-r03-03.dl.wb.ru:4222,ngx-api-r04-03.dl.wb.ru:4222,ngx-api-r05-03.dp.wb.ru:4222"
	NatsAddressTest = "test-cluster"
	BadSearchEventSubject = "wbxsearch.ru.exactmatch.common.badsearchevent"
	BadSearchEventQueryCapacity = 1024
	SearchEventSubject = "wbxsearch.ru.exactmatch.common.searchevent"
	SearchEventQueryCapacity = 1024
)

func Start(ctx context.Context, address, topic string, channel chan <- BadMessage) {
	conn, err := nats.Connect( address)
	if err != nil {
		log.Fatalln(err)
	}
	sub, err := conn.Subscribe(topic, func(m *nats.Msg){
		var badMessageProto protoType.BadSearchEvent
		var badMessage BadMessage
		err := proto.Unmarshal(m.Data, &badMessageProto)
		if err != nil {
			log.Print(err)
			return
		}
		badMessage.Query= badMessageProto.Query
		channel <- badMessage
	})
	if err != nil {
		log.Println(err)
	}
	log.Println("Nats is established!")
	<-ctx.Done()
	log.Println("Nats is disconnecting!")
	err = sub.Unsubscribe()
	if err != nil {
		log.Println(err)
	}
	conn.Close()
}

func Subscribe(connections nats.Conn, channel chan <- BadMessage) {
	sub, err := connections.Subscribe(BadSearchEventSubject, func(m *nats.Msg){
		var badMessage BadMessage
		err := json.Unmarshal(m.Data, &badMessage)
		if err != nil {
			log.Print(err)
			return
		}
		channel <- badMessage
	})
	if err != nil {
		log.Println(err)
	}
	defer sub.Unsubscribe()
}

func Subscribe2(connections nats.Conn, channel chan <- BadMessage) {
	connections.Subscribe(TestSubj, func(m *nats.Msg){
		log.Println("nats handler caught a message!")
		var badMessage BadMessage
		err := json.Unmarshal(m.Data, &badMessage)
		if err != nil {
			log.Print(err)
			return
		}
		channel <- badMessage
	} )
}

//В каком виде приходит сообщение? какую структуру создавать?
func BadSearchEventHandler(message *nats.Msg) {
}

//Зачем нам этот сабджект?
func SearchEvent(message *nats.Msg) {
}
 
func StanConnect(cluster, client, url string) *nats.Conn {
	sc, err := nats.Connect(
		cluster,
	)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Connected to cluster \"%s\" as client \"%s\"...\n", cluster, client)
	return sc
}


