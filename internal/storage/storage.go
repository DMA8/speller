package storage

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type SpellStorage struct {
	Storage map[string][]string
}



func NewStorage(fileName string) *SpellStorage {
	var storage SpellStorage
	var err error
	storage.Storage, err = CsvReader(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return &storage
}

func (s *SpellStorage) CreateSpell(csvLine string) error {
	words := strings.Split(csvLine, ";")
	key := words[0]
	if _, ok := s.Storage[key]; ok {
		return errors.New(csvLine + "is already created")
	}
	if len(words) < 2 {
		return errors.New(csvLine + "provide incorrect words")
	}
	errorWords := strings.Split(words[1], "|")
	s.Storage[key] = append(s.Storage[key], errorWords...)
	return nil
}

func (s *SpellStorage) ReadSPell(csvLine string) ([]string, error) {
	words := strings.Split(csvLine, ";")
	key := words[0]
	if _, ok := s.Storage[key]; !ok {
		return nil, fmt.Errorf("\"%s\" %s", csvLine," is not created")
	}
	return s.Storage[key], nil
}

func (s *SpellStorage) AddSpell(csvLine string) error {
	words := strings.Split(csvLine, ";")
	key := words[0]
	if _, ok := s.Storage[key]; !ok {
		return errors.New(csvLine + "is not added")
	}
	if len(words) < 2 {
		return errors.New(csvLine + "provide incorrect words")
	}
	errorWords := strings.Split(words[1], "|")
	s.Storage[key] = append(s.Storage[key], errorWords...)
	return nil
}

func (s *SpellStorage) DeleteSpell(csvLine string) error {
	words := strings.Split(csvLine, ";")
	key := words[0]
	if _, ok := s.Storage[key]; !ok {
		return errors.New(csvLine + "is not added")
	}
	delete(s.Storage, key)
	return nil
}

func (s *SpellStorage) DeleteParticularSpellings(csvLine string) error {
	words := strings.Split(csvLine, ";")
	key := words[0]
	if _, ok := s.Storage[key]; !ok {
		return errors.New(key + "is not added")
	}
	if len(words) < 2 || words[1] == "" {
		return errors.New(csvLine + "provide spellings to delete")
	}
	fmt.Println(words)
	wordsToDelete := strings.Split(words[1], "|")
	for _, v := range wordsToDelete {
		for i := 0; i < len(s.Storage[key]); i++ {
			if v == s.Storage[key][i] {
				s.Storage[key][i] = s.Storage[key][len(s.Storage[key])-1]
				s.Storage[key] = s.Storage[key][:len(s.Storage[key])-1]
			}
		}
	}
	fmt.Println(s.Storage[key])
	return nil
}

func CsvReader(fileName string) (map[string][]string, error) {
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
