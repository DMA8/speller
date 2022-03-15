package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)

//SpellStorage
type SpellStorage struct {
	Storage map[string][]string
	mu      sync.Mutex
}

//Spelling: SpellName - correct word; MisSpells - incorrect variants of SpellName
type Spelling struct {
	SpellName string   `json:"spellName"`
	MisSpells []string `json:"misSpells"`
}

//NewStorage creates new storage
func NewStorage(fileName string) *SpellStorage {
	var storage SpellStorage
	var err error
	storage.Storage, err = CSVReader(fileName)
	storage.mu = sync.Mutex{}
	if err != nil {
		log.Fatal(err)
	}
	return &storage
}

//AcceptSpellerSuggest gets message from the speller and creates or adds new spelling in storage
func (s *SpellStorage) AcceptSpellerSuggest(ctx context.Context, convey <-chan Spelling) {
	for {
		select {
		case msg := <-convey:
			log.Println("Storage got a message from speller!")
			s.createOrAdd(msg)
		case <-ctx.Done():
			return
		}
	}
}

func (s *SpellStorage) createOrAdd(spelling Spelling) {
	log.Println(spelling)
	if err := s.CreateSpell(&spelling); err != nil {
		log.Println("storage:", err)
		err = s.AddSpell(&spelling)
		if err != nil {
			log.Print(err)
		}
	}
}

//CreateSpell creates pair spellWord - misSpells in storage's map
func (s *SpellStorage) CreateSpell(spelling *Spelling) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Storage[spelling.SpellName]; ok {
		return errors.New(spelling.SpellName + " is already created")
	}
	s.Storage[spelling.SpellName] = append(s.Storage[spelling.SpellName], spelling.MisSpells...)

	return nil
}

//ReadSpell return misSpells for given key
func (s *SpellStorage) ReadSpell(spelling string) (*Spelling, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Storage[spelling]; !ok {
		return nil, fmt.Errorf("`%s` %s", spelling, " is not created")
	}
	return &Spelling{spelling, s.Storage[spelling]}, nil
}

//AddSpell adds given pair spellName - misSpells
func (s *SpellStorage) AddSpell(spelling *Spelling) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Storage[spelling.SpellName]; !ok {
		return errors.New(spelling.SpellName + " is not added")
	}
	if len(spelling.MisSpells) < 1 {
		return errors.New(spelling.SpellName + " provide incorrect words")
	}
	for _, v := range spelling.MisSpells {
		if !in(v, s.Storage[spelling.SpellName]) {
			s.Storage[spelling.SpellName] = append(s.Storage[spelling.SpellName], v)
		}
	}

	return nil
}

//DeleteSpell deletes given key from storage's map
func (s *SpellStorage) DeleteSpell(spelling string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Storage[spelling]; !ok {
		return errors.New(spelling + " is not added")
	}
	delete(s.Storage, spelling)
	return nil
}

//DeleteParticularSpellings deletes
func (s *SpellStorage) DeleteParticularSpellings(spelling *Spelling) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Storage[spelling.SpellName]; !ok {
		return errors.New(spelling.SpellName + " is not added")
	}
	if len(spelling.MisSpells) < 1 {
		return errors.New(spelling.SpellName + " provide spellings to delete")
	}
	for _, v := range spelling.MisSpells {
		for i := 0; i < len(s.Storage[spelling.SpellName]); i++ {
			if v == s.Storage[spelling.SpellName][i] {
				s.Storage[spelling.SpellName][i] = s.Storage[spelling.SpellName][len(s.Storage[spelling.SpellName])-1]
				s.Storage[spelling.SpellName] = s.Storage[spelling.SpellName][:len(s.Storage[spelling.SpellName])-1]
			}
		}
	}
	return nil
}
