package storage

import (
	"testing"
)

func in(word string, collection []string) bool {
	for _, v := range collection {
		if word == v {
			return true
		}
	}
	return false
}

func TestNewStorage(t *testing.T) {
	testMap := NewStorage("spells.csv")

	if _, ok := testMap.Storage["товар"]; !ok {
		t.Error("fail in creating map!")
	}
	if !in("тавар", testMap.Storage["товар"]) {
		t.Error("fail in adding wrong spellings")
	}
	if _, ok := testMap.Storage["спички"]; !ok {
		t.Error("fail in creating map!")
	}
	if !in("спичьки", testMap.Storage["спички"]) {
		t.Error("fail in adding wrong spellings")
	}
	if !in("спитчки", testMap.Storage["спички"]) {
		t.Error("fail in adding wrong spellings")
	}
	if !in("спишки", testMap.Storage["спички"]) {
		t.Error("fail in adding wrong spellings")
	}
	if in("фывцвфыв", testMap.Storage["спички"]) {
		t.Error("fail in adding wrong spellings")
	}
	if _, ok := testMap.Storage[";dsfnoefow"]; ok {
		t.Error("fail in creating map!")
	}
}

func TestCreateSpell(t *testing.T) {
	var testWord string
	testMap := NewStorage("spells.csv")

	testWord = "арбуз;арбус|орбуз;"
	testKey := "арбуз"
	if _, ok := testMap.Storage[testKey]; !ok {
		err := testMap.CreateSpell(testWord)
		if _, ok := testMap.Storage[testKey]; !ok && err != nil{
			t.Error("new spell wasn't created")
		}
	}
	if !in("орбуз", testMap.Storage["арбуз"]) && !in("арбус", testMap.Storage["арбуз"]) {
		t.Error("failed to add spelling")
	}
	err := testMap.CreateSpell(testWord)
	if err == nil {
		t.Errorf("%s added second time", testKey)
	}
}

func TestReadSpell(t *testing.T) {
	testMap := NewStorage("spells.csv")
	readedWords, err := testMap.ReadSPell("спички")
	if err != nil {
		t.Error("спичка not found")
	}
	if !in("спичьки", readedWords) || !in("спишки", readedWords) || !in("спитчки", readedWords) {
		t.Error("reading error!")
	}
}

func TestAddSpell(t *testing.T) {
	testMap := NewStorage("spells.csv")
	err := testMap.AddSpell("спички;спеттчки")
	if err != nil {
		t.Error("error while adding to the storage")
	}
	if !in("спеттчки", testMap.Storage["спички"]) {
		t.Error("added word didn't found")
	}
}

func TestDeleteSpell(t *testing.T) {
	testMap := NewStorage("spells.csv")
	err := testMap.DeleteSpell("спички")
	if err != nil {
		t.Error("Error while deleting!")
	}
	if _, ok := testMap.Storage["спички"]; ok {
		t.Error("delete didn't succeed")
	}
}

func TestDeleteParticularSpelling(t *testing.T) {
	testMap := NewStorage("spells.csv")
	err := testMap.DeleteParticularSpellings("спички;спитчки")
	if err != nil {
		t.Error("Error while deleting particular spelling")
	}
	if in("спитчки", testMap.Storage["спички"]) {
		t.Error("Particular spelling wasn't deleted")
	}
}