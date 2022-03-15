package speller

import (
	"context"
	"log"
	"spellCheck/internal/natsClient"
	"spellCheck/internal/storage"

	"github.com/Saimunyz/speller"
)

var s *speller.Speller

func init() {
	s = speller.NewSpeller()
	err := s.LoadModel("./models/small-data.gz")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("speller init done")
}

func AcceptMessage(ctx context.Context, conveyIn <-chan natsClient.BadMessage, conveyOut chan<- storage.Spelling) {
	var suggest storage.Spelling
	for {
		select {
		case inpQuery := <-conveyIn:
			log.Println("Speller got a message from stan client!")
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
