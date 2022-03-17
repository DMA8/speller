package storage

import (
	"bytes"
	"context"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// Dump saves storage in spellcheck.csv file each N minutes or when ctrl+c pressed
func (s *SpellStorage)Dump(ctx context.Context, done chan <- struct{}, everyMinutes int) {
	ticker := time.NewTicker(time.Minute * time.Duration(everyMinutes))
	for {
	select{
	case <-ctx.Done():
		s.saveFile()
		log.Println("Dump at exit is done")
		done <- struct{}{}
		return
	case <-ticker.C :
		s.saveFile()
		log.Println("Dump is done")
		runtime.GC()
	}
	}
}

func (s *SpellStorage)saveFile() {
	os.Mkdir("Dump", 0777)
	file, err := os.Create("Dump/spellcheck.csv")
	if err != nil {
		log.Println(err)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for key, value := range s.Storage{
		file.WriteString(strings.Join([]string{key, ";", strings.Join(value, "|"), ";\n"}, ""))
		//line.Write([]byte(strings.Join([]string{key, ";", strings.Join(value, "|"), ";\n"}, "")))
	}
	file.Close()
}

func (s *SpellStorage)saveFileOld() {
	var line bytes.Buffer
	os.Mkdir("Dump", 0777)
	file, err := os.Create("Dump/spellcheck.csv")
	if err != nil {
		log.Println(err)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for key, value := range s.Storage{
		line.Write([]byte(strings.Join([]string{key, ";", strings.Join(value, "|"), ";\n"}, "")))
	}
	file.Write(line.Bytes())
	file.Close()
}