package natsClient

import (
	"fmt"
	"log"

	stan "github.com/nats-io/stan.go"
	
)


const (
	Client = "speller"
	NatsAddress1 = "ngx-api-r01-03.dp.wb.ru:4222"
	NatsAddress2 = "ngx-api-r02-03.dl.wb.ru:4222"
	NatsAddress3 = "ngx-api-r03-03.dl.wb.ru:4222"
	NatsAddress4 = "ngx-api-r04-03.dl.wb.ru:4222"
	NatsAddress5 = "ngx-api-r05-03.dp.wb.ru:4222"

	BadSearchEventSubject = "wbxsearch.ru.exactmatch.common.badsearchevent"
	BadSearchEventQueryCapacity = 1024
	SearchEventSubject = "wbxsearch.ru.exactmatch.common.searchevent"
	SearchEventQueryCapacity = 1024
)

func Start() {
	SubscribeAllConnections(ConnectAllClusters())
}

func SubscribeAllConnections(connections []stan.Conn) {
	for _, v := range connections {
		v.Subscribe(BadSearchEventSubject, BadSearchEventHandler)
		v.Subscribe(SearchEventSubject, BadSearchEventHandler)
	}
}

//В каком виде приходит сообщение? какую структуру создавать?
func BadSearchEventHandler(message *stan.Msg) {
}

//Зачем нам этот сабджект?
func SearchEvent(message *stan.Msg) {
}

func ConnectAllClusters()[]stan.Conn {
	out := make([]stan.Conn,5)
	out[0] = StanConnect(NatsAddress1, Client, "")
	out[1] = StanConnect(NatsAddress2, Client, "")
	out[2] = StanConnect(NatsAddress3, Client, "")
	out[3] = StanConnect(NatsAddress4, Client, "")
	out[4] = StanConnect(NatsAddress5, Client, "")
	return out
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
