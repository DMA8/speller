package speller

import (
	"log"
	"testing"

	"github.com/Saimunyz/speller"
)

var myspeller *speller.Speller

func init() {
	myspeller = speller.NewSpeller()
	// err := s.LoadModel("../../models/bin-not-normalized-data.gz")
	err := myspeller.LoadModel("models/bin-not-normalized-data.gz")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("speller init done")
}

func BenchmarkSpellCorrect(b *testing.B) {
	
	for i := 0; i < b.N; i++ {
		myspeller.SpellCorrect("очинь длинное саабщение просто штобы пасматреть как профилирование работает в данном случае как насчет езе парачки мслов может быть ты уже упадешь а удачи ")
	}
}
