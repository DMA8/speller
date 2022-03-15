package storage

import (
	"testing"
)



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
	var testWord Spelling
	testMap := NewStorage("spells.csv")

	testWord.MisSpells = []string{"арбус", "орбуз"}
	testWord.SpellName = "арбуз"
	testKey := "арбуз"
	if _, ok := testMap.Storage[testKey]; !ok {
		err := testMap.CreateSpell(&testWord)
		if _, ok := testMap.Storage[testKey]; !ok && err != nil{
			t.Error("new spell wasn't created")
		}
	}
	if !in("орбуз", testMap.Storage["арбуз"]) && !in("арбус", testMap.Storage["арбуз"]) {
		t.Error("failed to add spelling")
	}
	err := testMap.CreateSpell(&testWord)
	if err == nil {
		t.Errorf("%s added second time", testKey)
	}
}

func TestReadSpell(t *testing.T) {
	testMap := NewStorage("spells.csv")
	readedWords, err := testMap.ReadSpell("спички")
	if err != nil {
		t.Error("спичка not found")
	}
	content := readedWords.MisSpells
	if !in("спичьки", content) || !in("спишки", content) || !in("спитчки", content) {
		t.Error("reading error!")
	}
}

func TestAddSpell(t *testing.T) {
	var testWord Spelling
	testWord.SpellName = "спички"
	testWord.MisSpells = []string{"спеттчки"}
	testMap := NewStorage("spells.csv")
	err := testMap.AddSpell(&testWord)
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
	var testWord Spelling
	testMap := NewStorage("spells.csv")
	testWord.SpellName = "спички"
	testWord.MisSpells = []string{"спитчки"}
	err := testMap.DeleteParticularSpellings(&testWord)
	if err != nil {
		t.Error("Error while deleting particular spelling")
	}
	if in("спитчки", testMap.Storage["спички"]) {
		t.Error("Particular spelling wasn't deleted")
	}
}