package storage

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strings"
)

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

func in(word string, collection []string) bool {
	for _, v := range collection {
		if word == v {
			return true
		}
	}
	return false
}
