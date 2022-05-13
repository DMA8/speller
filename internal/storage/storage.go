package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
)

//SpellStorage
type SpellStorage struct {
	Storage					map[string][]string
	StorageOnlyNewKeysAdded	map[string][]string
	BlackList				map[string]struct{}
	NewAddedErrors			map[string]struct{}
	mu						sync.Mutex
}
var a int
//Spelling: SpellName - correct word; MisSpells - incorrect variants of SpellName
type Spelling struct {
	SpellName string   `json:"spellName"` //то, как исправил спеллер
	MisSpells []string `json:"misSpells"` //запрос, который надо было исправить
}

//NewStorage creates new storage
func NewStorage(fileName string) *SpellStorage {
	var storage SpellStorage
	var err, err2 error
	storage.Storage, err = CSVReader(fileName)
	storage.StorageOnlyNewKeysAdded, err2 = CSVReader(fileName)
	storage.BlackList = make(map[string]struct{})
	storage.NewAddedErrors = make(map[string]struct{})
	for key := range storage.Storage {
		storage.BlackList[key] = struct{}{}
	}
	storage.mu = sync.Mutex{}
	if err != nil{
		log.Fatal(err)
	} else if err2 != nil {
		log.Fatal(err2)
	}
	return &storage
}

//AcceptSpellerSuggest gets message from the speller and creates or adds new spelling in storage
func (s *SpellStorage) AcceptSpellerSuggest(ctx context.Context, convey <-chan Spelling) {
	for {
		select {
		case msg := <-convey:
			//fmt.Printf("spellerSuggest: %s\nError: %s\n\n", msg.SpellName, msg.MisSpells[0])
			a++
			s.createOrAdd(msg)
		case <-ctx.Done():
			return
		}
	}
}

func (s *SpellStorage) createOrAdd(spelling Spelling) {
	spellerSuggestSplitted := strings.Fields(spelling.SpellName)
	rawCustomerQuerySplitted := strings.Fields(spelling.MisSpells[0])
	for i := range rawCustomerQuerySplitted {
		rawCustomerQuerySplitted[i] = strings.Trim(rawCustomerQuerySplitted[i], ".,-/!%#$^:&?*()")
	}
	if len(spellerSuggestSplitted) != len(rawCustomerQuerySplitted) {
		return
	}
	for i := range spellerSuggestSplitted {
		if spellerSuggestSplitted[i] == rawCustomerQuerySplitted[i] {
			continue
		}
		addObject := Spelling{
			SpellName: spellerSuggestSplitted[i],
			MisSpells: []string{rawCustomerQuerySplitted[i]},
		}
		if err := s.CreateSpell(&addObject); err != nil {
			err = s.AddSpell(&addObject)
			if err != nil {
				log.Print(err)
			}
		}
		if err := s.CreateSpellBlackList(&addObject); err != nil {
			err = s.AddSpellBlackList(&addObject)
			if err == nil {
				s.NewAddedErrors[addObject.MisSpells[0]] = struct{}{}
			}
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

//CreateSpell creates pair spellWord - misSpells in storage's map
func (s *SpellStorage) CreateSpellBlackList(spelling *Spelling) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Storage[spelling.SpellName]; ok {
		return errors.New(spelling.SpellName + " is already created")
	}
	s.StorageOnlyNewKeysAdded[spelling.SpellName] = append(s.StorageOnlyNewKeysAdded[spelling.SpellName], spelling.MisSpells...)

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

//AddSpell adds given pair spellName - misSpells
func (s *SpellStorage) AddSpellBlackList(spelling *Spelling) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Storage[spelling.SpellName]; !ok {
		return errors.New(spelling.SpellName + " is not added")
	}
	if _, ok := s.BlackList[spelling.SpellName]; ok {
		return errors.New(spelling.SpellName + " is in the black list")
	}
	if len(spelling.MisSpells) < 1 {
		return errors.New(spelling.SpellName + " provide incorrect words")
	}
	for _, v := range spelling.MisSpells {
		if !in(v, s.Storage[spelling.SpellName]) {
			s.StorageOnlyNewKeysAdded[spelling.SpellName] = append(s.StorageOnlyNewKeysAdded[spelling.SpellName],  v)
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
