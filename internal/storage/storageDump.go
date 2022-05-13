package storage

import (
	"bytes"
	"context"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// Dump saves storage in spellcheck.csv file each N minutes or when ctrl+c pressed
func (s *SpellStorage) Dump(ctx context.Context, done chan<- struct{}, everyMinutes int) {
	ticker := time.NewTicker(time.Minute * time.Duration(everyMinutes))
	for {
		select {
		case <-ctx.Done():
			s.saveFile()
			s.saveFileBlackList()
			log.Println("Dump at exit is done")
			done <- struct{}{}
			return
		case <-ticker.C:
			s.saveFile()
			s.saveFileBlackList()
			log.Println("Dump is done")
			//runtime.GC()
		}
	}
}

//Should be ordered!!!!
func (s *SpellStorage) saveFile() {
	s.mu.Lock()
	defer s.mu.Unlock()
	//Should be ordered!!!!
	sortedKeys := make([]string, 0, len(s.Storage))
	for key := range s.Storage {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)
	os.Mkdir("Dump", 0777)
	file, err := os.Create("Dump/spellcheck.csv")
	if err != nil {
		log.Println(err)
		return
	}

	// newAdded := make([]string, 0)
	// oldAdded := make([]string, 0)
	// str := strings.Builder{}
	for _, key := range sortedKeys {
		file.WriteString(strings.Join([]string{key, ";", strings.Join(s.Storage[key], "|"), ";\n"}, ""))
	}
	file.Close()
}

func (s *SpellStorage) saveFileBlackList() {
	s.mu.Lock()
	defer s.mu.Unlock()
	//Should be ordered!!!!
	sortedKeys := make([]string, 0, len(s.Storage))
	for key := range s.StorageOnlyNewKeysAdded {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)
	os.Mkdir("Dump", 0777)
	file, err := os.Create("Dump/spellcheck_DontTouchOld.csv")
	if err != nil {
		log.Println(err)
		return
	}
	file2, err := os.Create("Dump/spellcheck_OnlyNew.csv")
	if err != nil {
		log.Println(err)
		return
	}
	for _, key := range sortedKeys {
		if _, ok := s.BlackList[key]; !ok {
			file.WriteString(strings.Join([]string{key, ";", strings.Join(s.StorageOnlyNewKeysAdded[key], "|"), ";(!!!)\n"}, ""))
			file2.WriteString(strings.Join([]string{key, ";", strings.Join(s.StorageOnlyNewKeysAdded[key], "|"), ";\n"}, ""))

		} else {
			file.WriteString(strings.Join([]string{key, ";", strings.Join(s.StorageOnlyNewKeysAdded[key], "|"), ";\n"}, ""))
		}
	}
	file.Close()
}

func (s *SpellStorage) saveFileOld() {
	var line bytes.Buffer
	os.Mkdir("Dump", 0777)
	file, err := os.Create("Dump/spellcheck.csv")
	if err != nil {
		log.Println(err)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for key, value := range s.Storage {
		line.Write([]byte(strings.Join([]string{key, ";", strings.Join(value, "|"), ";\n"}, "")))
	}
	file.Write(line.Bytes())
	file.Close()
}
