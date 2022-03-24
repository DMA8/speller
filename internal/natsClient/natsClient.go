package natsStreamingClient

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	protoType "spellCheck/internal/proto"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"gopkg.in/yaml.v2"
)

type BadMessage struct {
	Query string
	Error string
}

type natsConfig struct {
	NatsAddress                 string `yaml:"natsAddress"`
	BadSearchEventSubject       string `yaml:"badSearchEventSubjectCommon"`
	SearchEventSubject          string `yaml:"searchEventSubject"`
	BadSearchEventQueryCapacity int    `yaml:"badSearchEventQueryCapacity"`
	SearchEventQueryCapacity    int    `yaml:"searchEventQueryCapacity"`
}

// Config - contains all configuration parameters in config package
type config struct {
	NatsConfig natsConfig `yaml:"nats_config"`
}

func ReadConfigYML(filePath string) (cfg *config, err error) {
	file, err := os.Open(filepath.Clean(filePath))
	defer file.Close()
	if err != nil {
		return cfg, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func Start(ctx context.Context, channel chan<- BadMessage, done chan struct{}) {
	cfg, err := ReadConfigYML("config/config.yaml")
	if err != nil {
		panic(err)
	}
	conn, err := nats.Connect(cfg.NatsConfig.NatsAddress)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("nats CONNECTED: %s", cfg.NatsConfig.NatsAddress)
	sub, err := conn.Subscribe(cfg.NatsConfig.BadSearchEventSubject, func(m *nats.Msg) {
		var badMessageProto protoType.BadSearchEvent
		var badMessage BadMessage
		err := proto.Unmarshal(m.Data, &badMessageProto)
		if err != nil {
			log.Print(err)
			return
		}
		
		badMessage.Query = badMessageProto.Query
		badMessage.Error = badMessageProto.Error
		if !strings.HasPrefix(badMessageProto.Error, "only") {
			fmt.Println("query: ", badMessageProto.Query)
			fmt.Println("error: ", badMessageProto.Error)
			fmt.Println("--------------------------------------------")
		} 


		channel <- badMessage
	})
	if err != nil {
		log.Println(err)
	}
	log.Printf("nats SUB: %s", cfg.NatsConfig.BadSearchEventSubject)
	<-ctx.Done()
	err = sub.Unsubscribe()
	log.Printf("nats UNSUB: %s", cfg.NatsConfig.BadSearchEventSubject)
	if err != nil {
		log.Println(err)
	}
	conn.Close()
	log.Printf("nats DISCONNECTED: %s ", cfg.NatsConfig.NatsAddress)
	done <- struct{}{}
}
