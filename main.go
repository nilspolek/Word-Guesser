package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/nilspolek/Word-Guesser/parser"
	wordfilter "github.com/nilspolek/Word-Guesser/wordFilter"
)

//go:embed dict.txt
var dict []byte

var (
	length   int
	filePath string
)

func main() {
	flag.StringVar(&filePath, "f", "", "Path to the dictionary file")
	flag.IntVar(&length, "l", 5, "Length of the words to guess")
	flag.Parse()

	var (
		words []string
		err   error
	)

	if filePath != "" {
		words, err = parser.ParseWordList(filePath)
	} else {
		words, err = parser.ParseByteSlice(dict)
	}
	if err != nil {
		fmt.Printf("Error reading word list: %v\n", err)
		return
	}

	excludeRunes := ""
	runesInRightPlace := ""
	runesInWrongPlace := ""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Press enter to start\n")
	for scanner.Scan() {
		excludeRunes = promptString("Enter letters to exclude (e.g. abc), then press Enter.\nExclude letters: ")
		runesInRightPlace = promptString("Enter the letters in their position (e.g. a__c_): ")
		runesInWrongPlace = promptString("Enter the letters that are in the word but not in the right position (e.g. a_b__): ")

		candidates := wordfilter.FilterCandidates(words, length, excludeRunes, runesInRightPlace, runesInWrongPlace)
		words = candidates.Strings()
		candidates.Sort()
		if !candidates.Print(parser.ComputeLetterFrequencies(words), 3) {
			return
		}
	}
}

func promptString(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.ToLower(scanner.Text())
}
