package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	words, err := ParseWordList("dict.txt")
	if err != nil {
		panic(err)
	}
	length := 5
	excludeRunes := ""
	runesInRightPlace := ""
	runesInWrongPlace := ""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Press enter to start\n")
	for scanner.Scan() {
		excludeRunes = promptString("Enter letters to exclude (e.g. abc), then press Enter.\nExclude letters: ")
		runesInRightPlace = promptString("Enter the letters in their position (e.g. a__c_): ")
		runesInWrongPlace = promptString("Enter the letters that are in the word but not in the right position (e.g. a_b__): ")

		candidates := FilterCandidates(words, length, excludeRunes, runesInRightPlace, runesInWrongPlace)
		words = candidates.Strings()
		candidates.Sort()
		if !candidates.Print(ComputeLetterFrequencies(words), 3) {
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
