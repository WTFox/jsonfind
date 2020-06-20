package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"jsonfind/pkg/scout"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:        "jf",
		Usage:       "JSONFind",
		UsageText:   "jf <valueToFind> <jsonFile>",
		Version:     "1.0.1",
		Description: "Search a JSON file for a specified value and output full paths of each occurrence found",
		Action:      doSearch,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func doSearch(c *cli.Context) error {
	if c.NArg() != 2 {
		cli.ShowAppHelpAndExit(c, 1)
	}

	lookingFor := c.Args().Get(0)
	filename := c.Args().Get(1)

	bytes, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("couldn't read file %s\n%v", filename, err)
	}

	var result interface{}
	if err = json.Unmarshal(bytes, &result); err != nil {
		return fmt.Errorf("%v\njf was unable to parse the JSON file", err)
	}

	scout := scout.New(lookingFor, result)
	found, err := scout.DoSearch()
	if err != nil {
		return err
	}

	for _, occurrence := range found {
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
