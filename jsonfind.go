package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readFile(filepath string) ([]byte, error) {
	jsonFile, err := os.Open(filepath)
	defer jsonFile.Close()
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(jsonFile)
}

func main() {
	lookingFor := os.Args[1]
	filename := os.Args[2]

	bytes, err := readFile(filename)
	if err != nil {
		log.Fatalf("Couldn't read file %s\n%v", filename, err)
	}

	var parsedBytes jsonData
	json.Unmarshal(bytes, &parsedBytes)

	scout := NewScout(lookingFor, parsedBytes)
	scout.DoSearch()

	for _, occurrence := range scout.Found() {
		fmt.Println(occurrence)
	}
}
