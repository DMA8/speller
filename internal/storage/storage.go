package storage

import (
	"bufio"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//SpellStorage
type SpellStorage struct {
	Storage map[string][]string
}

//Spelling: SpellName - correct word; MisSpells - incorrect variants of SpellName
type Spelling struct {
	SpellName	string `json:"spellName"`
	MisSpells	[]string `json:"misSpells"`
}

//NewStorage creates new storage
func NewStorage(fileName string) *SpellStorage {
	var storage SpellStorage
	var err error
	storage.Storage, err = CSVReader(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return &storage
}

func (s *SpellStorage)AcceptSpellerSuggest(ctx context.Context, convey <- chan Spelling) {
	for {
		select {
		case msg := <- convey:
			s.CreateOrAdd(msg)
		case <- ctx.Done():
			return
		}
	}
}

func (s *SpellStorage)CreateOrAdd(spelling Spelling) {
	if err := s.CreateSpell(&spelling); err != nil {
		err = s.AddSpell(&spelling)
		if err != nil {
			log.Print(err)
		}
	}
}

//CreateSpell creates pair spellWord - misSpells in storage's map
func (s *SpellStorage) CreateSpell(spelling *Spelling) error {
	if _, ok := s.Storage[spelling.SpellName]; ok {
		return errors.New(spelling.SpellName + "is already created")
	}
	s.Storage[spelling.SpellName] = append(s.Storage[spelling.SpellName], spelling.MisSpells...)
	return nil
}

//ReadSpell return misSpells for given key
func (s *SpellStorage) ReadSpell(spelling string) (*Spelling, error) {
	if _, ok := s.Storage[spelling]; !ok {
		return nil, fmt.Errorf("`%s` %s", spelling," is not created")
	}
	return &Spelling{spelling, s.Storage[spelling]}, nil
}

//AddSpell adds given pair spellName - misSpells
func (s *SpellStorage) AddSpell(spelling *Spelling) error {
	if _, ok := s.Storage[spelling.SpellName]; !ok {
		return errors.New(spelling.SpellName + " is not added")
	}
	if len(spelling.MisSpells) < 1 {
		return errors.New(spelling.SpellName + " provide incorrect words")
	}
	s.Storage[spelling.SpellName] = append(s.Storage[spelling.SpellName], spelling.MisSpells...)
	return nil
}

//DeleteSpell deletes given key from storage's map
func (s *SpellStorage) DeleteSpell(spelling string) error {
	if _, ok := s.Storage[spelling]; !ok {
		return errors.New(spelling + " is not added")
	}
	delete(s.Storage, spelling)
	return nil
}

//DeleteParticularSpellings deletes 
func (s *SpellStorage) DeleteParticularSpellings(spelling *Spelling) error {
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

//CSVReader creates map[firstColumnCSV] and splitted by "|" second column
func CSVReader(fileName string) (map[string][]string, error) {
	outputMap := make(map[string][]string)
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if line == nil {
			continue
		}
		splittedLine := strings.Split(line[0], ";")
		if len(splittedLine) < 2 {
			continue
		}
		key := strings.TrimSuffix(strings.Split(line[0], ";")[0], ";")
		errWords := strings.Split(strings.Split(line[0], ";")[1], "|")
		errWords[len(errWords)-1] = strings.TrimSuffix(errWords[len(errWords)-1], ";")
		outputMap[key] = append(outputMap[key], errWords...)
	}
	return outputMap, nil
}

// func csvReaderRegex(fileName string) map[string][]string { //пытаемся прочитать csv и сделать привычную для нас мапу "правильноеСлово":["неправильные слова"...]
// 	outputMap := make(map[string][]string)
// 	csvFile, err := os.Open(fileName)
// 	patt := `^[а-яА-Я]`
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	reader := csv.NewReader(bufio.NewReader(csvFile))
// 	for {
// 		line, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if len(line) < 2 {
// 			continue
// 		}
// 		key := strings.TrimSuffix(strings.Split(line[0], ";")[0], ";")
// 		if ok, _ := regexp.Match(patt, []byte(key)); !ok {
// 			continue
// 		}
// 		errWords := strings.Split(line[1], "|")
// 		for _, v := range errWords {
// 			v := strings.TrimSuffix(v, ";")
// 			if ok, _ := regexp.Match(patt, []byte(v)); ok {
// 				if len(strings.Fields(key)) == len(strings.Fields(v)) {
// 					outputMap[key] = append(outputMap[key], v)
// 				}
// 			}
// 		}
// 	}
// 	return outputMap
// }
