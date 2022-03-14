package storage

import (
	"bytes"
	"context"
	"log"
	"os"
	"strings"
	"time"
)

func (s *SpellStorage)Dump(ctx context.Context) {
	ticker := time.NewTicker(time.Minute * 1)
	for {
	select{
	case <-ctx.Done():
		s.SaveFile()
		return
	case <-ticker.C :
		s.SaveFile()
	}
	}
}

func (s *SpellStorage)SaveFile() {
	var line bytes.Buffer
	err := os.Mkdir("Dump", 0777)
	if err != nil {
		log.Println(err)
	}
	file, err := os.Create("Dump/spellcheck.csv")
	if err != nil {
		log.Println(err)
		return
	}
	for key, value := range s.Storage{
		line.Write([]byte(key + ";" + strings.Join(value, "|") + ";\n"))
	}
	file.Write(line.Bytes())
}