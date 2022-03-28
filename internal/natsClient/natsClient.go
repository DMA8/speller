package natsStreamingClient

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	protoType "spellCheck/internal/proto"
	"strings"
	"unicode"

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
	SearchEventSubjectCommon    string `yaml:"searchEventSubjectCommon"`
	SearchEventSubjectMale      string `yaml:"searchEventSubjectMale"`
	SearchEventSubjectFemale    string `yaml:"searchEventSubjectFemale"`
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
	regEx := regexp.MustCompile(`^[а-яА-Яa-zA-Z]*[\.\,-]?$`)
	cfg, err := ReadConfigYML("config/config.yaml")
	if err != nil {
		panic(err)
	}
	conn, err := nats.Connect(cfg.NatsConfig.NatsAddress)
	if err != nil {
		panic(err)
	}
	natsSubjects := []string{cfg.NatsConfig.SearchEventSubjectCommon,
		cfg.NatsConfig.SearchEventSubjectMale,
		cfg.NatsConfig.SearchEventSubjectFemale,
	}
	natsSubsObjects := make([]*nats.Subscription, 0, len(natsSubjects))
	log.Printf("nats CONNECTED: %s", cfg.NatsConfig.NatsAddress)
	for _, subject := range natsSubjects {
		sub, err := conn.Subscribe(subject, func(m *nats.Msg) {
			var badMessageProto protoType.SearchEvent
			var badMessage BadMessage
			err := proto.Unmarshal(m.Data, &badMessageProto)
			if err != nil {
				log.Print(err)
				return
			}
			
			if badMessageProto.ShardKey == ""{ //|| badMessageProto.ShardKey == "merger"
			filterdMsg, ok := filterMsg(badMessageProto.Query, regEx)
				if ok {
					fmt.Println(filterdMsg)
					badMessage.Query = filterdMsg
					channel <- badMessage
				}
			}
		})
		if err != nil {
			log.Fatal(err)
			return
		}
		natsSubsObjects = append(natsSubsObjects, sub)
		log.Printf("nats SUB: %s", subject)
	}
	<-ctx.Done()
	for _, s := range natsSubsObjects {
		err = s.Unsubscribe()
		log.Printf("nats UNSUB: %s", s.Subject)
		if err != nil {
			log.Println(err)
		}
	}
	conn.Close()
	log.Printf("nats DISCONNECTED: %s ", cfg.NatsConfig.NatsAddress)
	done <- struct{}{}
}


//если в запросе много грязи, то мы не пропускаем его. Его в запросе несколько плохих слов, мы вернем запрос без плохих слов
func filterMsg(msg string, regEng *regexp.Regexp) (string, bool){
	allWords := strings.Split(msg, " ")
	outWords := make([]string, 0, len(allWords))
	for _, word := range allWords {
		if regEng.Match([]byte(word)){
			outWords = append(outWords, word)
		}
	}
	if len(outWords) * 2 >= len(allWords) {
		fmt.Println(msg, "->", strings.Join(outWords, " "))
		return strings.Join(outWords, " "), true
	}
	return "", false
}