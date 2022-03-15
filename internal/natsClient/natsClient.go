package natsClient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	stan "github.com/nats-io/stan.go"
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

func Start(ctx context.Context, address, client, cluster string, channel chan <- BadMessage) {
	conn, err := stan.Connect(
		address,
		client,
		stan.Pings(1, 3),
		stan.NatsURL(""),
	)
	if err != nil {
		log.Fatalln(err)
	}
	sub, err := conn.Subscribe(cluster, func(m *stan.Msg){
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
	log.Println("Nats is established!")
	<-ctx.Done()
	log.Println("Nats is disconnecting!")
	err = sub.Unsubscribe()
	if err != nil {
		log.Println(err)
	}
	err = conn.Close()
	if err != nil {
		log.Println(err)
	}
}

func Subscribe(connections stan.Conn, channel chan <- BadMessage) {
	sub, err := connections.Subscribe(BadSearchEventSubject, func(m *stan.Msg){
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

func Subscribe2(connections stan.Conn, channel chan <- BadMessage) {
	connections.Subscribe(TestSubj, func(m *stan.Msg){
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
func BadSearchEventHandler(message *stan.Msg) {
}

//Зачем нам этот сабджект?
func SearchEvent(message *stan.Msg) {
}
 
func StanConnect(cluster, client, url string) stan.Conn {
	sc, err := stan.Connect(
		cluster,
		client,
		stan.Pings(1, 3),
		stan.NatsURL(""),
	)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Connected to cluster \"%s\" as client \"%s\"...\n", cluster, client)
	return sc
}


