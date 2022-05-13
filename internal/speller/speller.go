package speller

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	nats "spellCheck/internal/natsClient"
	"spellCheck/internal/storage"
	"time"

	"github.com/Saimunyz/speller"
)

var s *speller.Speller

var AllTests, SuccessSpellerAct int

func init() {
	s = speller.NewSpeller("config/config.yaml")
	err := s.LoadModel("./models/Ildar_AllRu-model.gz")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("speller init done")
}

func AcceptMessage(ctx context.Context, conveyIn <-chan nats.BadMessage, conveyOut chan<- storage.Spelling) {
	var suggest storage.Spelling
	fileSuccess, err := os.Create("successAlter.txt")
	if err != nil {
		panic (err)
	}
	fileFail, err := os.Create("failAlter.txt")
	if err != nil {
		panic(err)
	}
	for {
		select {
		case inpQuery := <-conveyIn:
			if inpQuery.Query == "" {
				continue
			}
			suggest.MisSpells = []string{inpQuery.Query}
			if inpQuery.Query == "" {
				continue
			}
			t := time.Now()
			suggest.SpellName = s.SpellCorrect2(inpQuery.Query)
			t2 := time.Since(t)
			if suggest.MisSpells[0] != suggest.SpellName {
				AllTests++
				conveyOut <- suggest
				if SuccessAlter(suggest.SpellName) {
					SuccessSpellerAct++
					fmt.Fprintf(fileSuccess, "(%v) %s -> %s\n", t2, inpQuery, suggest.SpellName)
				} else {
					fmt.Fprintf(fileFail, "(%v) %s -> %s\n", t2, inpQuery, suggest.SpellName)
				}
				fmt.Printf("%v Q: %s -> S: %s\n",t2, inpQuery.Query, suggest.SpellName)
			}
		case <-ctx.Done():
			return
		}
	}
}

func SuccessAlter(inp string) bool {
	url := fmt.Sprintf("http://exactmatch-common.wbx-ru.svc.k8s.3dcat/v2/search?query=%s", inp)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	_, err = io.ReadAll(resp.Body)
	return err == nil
}


// func filterQuery(inp string) string {
// 	splt := strings.Fields(inp)
// 	blckList := []
// 	for i := range splt {
// 		for _, v := range splt[i] {
// 			if !unicode.IsLetter(v) {

// 			}
// 		}
// 	}
// }
