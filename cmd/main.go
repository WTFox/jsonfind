package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"jsonfind/pkg/scout"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "jf",
		Usage:       "JSONFind",
		UsageText:   "jf <valueToFind> <jsonFile>",
		Version:     "0.0.1",
		Description: "Search a JSON file for a specified value and output full paths of each occurrence found",
		Action:      doSearch,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func doSearch(c *cli.Context) error {
	lookingFor := c.Args().Get(0)
	filename := c.Args().Get(1)

	bytes, err := readFile(filename)
	if err != nil {
		log.Fatalf("Couldn't read file %s\n%v", filename, err)
	}

	var result map[string]interface{}
	if err = json.Unmarshal(bytes, &result); err != nil {
		fmt.Printf("%v\n\nExplanation: JF was unable to parse the JSON file, jf only supports JSON object at the root level for now. (e.g. {})?\n", err)
		os.Exit(1)
	}

	scout := scout.New(lookingFor, result)
	scout.DoSearch()

	for _, occurrence := range scout.Found() {
		fmt.Println(occurrence)
	}

	return nil
}

func readFile(filepath string) ([]byte, error) {
	jsonFile, err := os.Open(filepath)
	defer jsonFile.Close()
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(jsonFile)
}
