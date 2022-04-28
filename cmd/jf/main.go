package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"jsonfind/pkg/scout"

	"github.com/urfave/cli/v2"
)

const flagParentPaths = "parent-paths"
const flagFirstOnly = "first-only"
const flagUseRegex = "use-regex"

func main() {
	app := &cli.App{
		Name:        "jf",
		Usage:       "JSONFind",
		UsageText:   "jf <valueToFind> <jsonFile>",
		Version:     "1.1.0",
		Description: "Search a JSON file for a specified value and output full paths of each occurrence found",
		Action:      entrypoint,
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
			&cli.BoolFlag{
				Name:    flagUseRegex,
				Usage:   "Use pattern matching via regex expression",
				Aliases: []string{"r"},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func entrypoint(c *cli.Context) error {
	lookingFor := c.Args().Get(0)
	filename := c.Args().Get(1)

	bytes, err := readInput(filename)
	if err != nil {
		cli.ShowAppHelpAndExit(c, 1)
	}

	var result any
	if err = json.Unmarshal(bytes, &result); err != nil {
		return fmt.Errorf("%v\njf was unable to parse the JSON file", err)
	}

	scout := scout.New(lookingFor, result, c.Bool(flagUseRegex))
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
			// exit early if the user only wants the first found path
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

func inputIsFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func readInput(filepath string) ([]byte, error) {
	if inputIsFromPipe() {
		return ioutil.ReadAll(os.Stdin)
	}
	return readFromFile(filepath)
}

func readFromFile(filepath string) ([]byte, error) {
	if filepath == "" {
		return []byte{}, errors.New("no file provided")
	}

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
