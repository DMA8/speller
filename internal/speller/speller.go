package speller

import (
	"context"
	"log"

	nats "spellCheck/internal/natsClient"
	"spellCheck/internal/storage"

	"github.com/Saimunyz/speller"
)

var s *speller.Speller

func init() {
	s = speller.NewSpeller("config/config.yaml")
	err := s.LoadModel("./models/bin-not-normalized-data.gz")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("speller init done")
}

func AcceptMessage(ctx context.Context, conveyIn <-chan nats.BadMessage, conveyOut chan<- storage.Spelling) {
	var suggest storage.Spelling

	for {
		select {
		case inpQuery := <-conveyIn:
			if inpQuery.Query == "" {
				continue
			}
			//log.Println("Speller got a message from stan client!")
			// suggest.MisSpells = []string{inpQuery.Query}
			suggest.MisSpells = []string{inpQuery.Query}
			suggest.SpellName = s.SpellCorrect(inpQuery.Query)
			if suggest.MisSpells[0] != suggest.SpellName {
				conveyOut <- suggest
			}
		case <-ctx.Done():
			return
		}
	}
}
