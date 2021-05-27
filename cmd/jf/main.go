package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"jsonfind/pkg/scout"

	"github.com/urfave/cli/v2"
)

const flagParentPaths = "parent-paths"
const flagFirstOnly = "first-only"

func main() {
	app := &cli.App{
		Name:        "jf",
		Usage:       "JSONFind",
		UsageText:   "jf <valueToFind> <jsonFile>",
		Version:     "1.0.3",
		Description: "Search a JSON file for a specified value and output full paths of each occurrence found",
		Action:      doSearch,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagParentPaths,
				Usage:   "Renders the parent paths only",
				Aliases: []string{"p"},
			},
			&cli.BoolFlag{
				Name:    flagFirstOnly,
				Usage:   "Returns the first occurrence only",
				Aliases: []string{"f"},
			},
		},
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
		if c.Bool(flagParentPaths) {
			fmt.Println(splitToParentPath(occurrence))
		} else {
			fmt.Println(occurrence)
		}
		if c.Bool(flagFirstOnly) {
			return nil
		}
	}

	return nil
}

func splitToParentPath(path string) string {
	pathElements := strings.Split(path, ".")
	if len(pathElements) > 0 && len(pathElements) < 3 {
		return "."
	}

	lastElement := pathElements[len(pathElements)-1]
	if len(lastElement) > 0 && lastElement[len(lastElement)-1] == ']' {
		lastElement = strings.Split(lastElement, "[")[0]
		pathElements = append(pathElements[:len(pathElements)-1], lastElement)
		return strings.Join(pathElements, ".")
	}

	return strings.Join(pathElements[:len(pathElements)-1], ".")
}

func readFile(filepath string) ([]byte, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(jsonFile)
}
