package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	wordfilter "wordle/wordFilter"
)

func main() {
	words, err := parse("dict.txt")

	if err != nil {
		panic(err)
	}
	length := 5
	excludeRundes := ""
	runesInRightPlace := ""
	runesInWrongPlace := ""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Press enter to start\n")
	for scanner.Scan() {
		excludeRundes = promptString("Enter letters to exclude (e.g. abc), then press Enter.\nExclude letters: ")

		runesInRightPlace = promptString("Enter the letters in thir position (e.g. a__c_): ")
		runesInWrongPlace = promptString("Enter the letters that are in the word but not in the right position (e.g. a_b__): ")
		data := filterWords(words, length, excludeRundes, runesInRightPlace, runesInWrongPlace)
		words = data.Strings()
		fmt.Println(words)
		sortCandidates(data)
		if !printCandidates(data, computeLetterFrequencies(words), 3) {
			return
		}
	}
}

// Liest Wörter aus Datei
func parse(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			words = append(words, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

func promptString(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToLower(scanner.Text())
	return input
}

// Zählt Buchstabenhäufigkeiten
func computeLetterFrequencies(words []string) map[rune]int {
	letterFreq := make(map[rune]int)
	for _, word := range words {
		for _, r := range strings.ToLower(word) {
			if r >= 'a' && r <= 'z' {
				letterFreq[r]++
			}
		}
	}
	return letterFreq
}

// Filtert Wörter nach Kriterien
type wordData struct {
	word        string
	uniqueRunes map[rune]bool
}

type wordDataSlice []wordData

func (w wordDataSlice) Strings() []string {
	var result []string
	for _, wd := range w {
		result = append(result, wd.word)
	}
	return result
}

func filterWords(words []string, length int, excluded, runesRightInPlace, runesWrongPlace string) wordDataSlice {
	var candidates []wordData

	var wds = wordfilter.New(words)
	wds.FilterLength(length)
	wds.ToLower()

	if len(runesWrongPlace) == length {
		wds.FilterWrongPlace(runesWrongPlace)
	}

	if len(runesRightInPlace) == length {
		wds.FilterRightPlace(runesRightInPlace)
	}
	wds.ExcludeLetters(excluded)

	words = *wds
	for _, word := range words {

		// Wort akzeptieren
		seen := make(map[rune]bool)
		for _, r := range word {
			if r >= 'a' && r <= 'z' {
				seen[r] = true
			}
		}
		candidates = append(candidates, wordData{word: word, uniqueRunes: seen})
	}
	return candidates
}

// Sortiert Kandidaten nach Anzahl einzigartiger Buchstaben, dann alphabetisch
func sortCandidates(candidates []wordData) {
	sort.Slice(candidates, func(i, j int) bool {
		a, b := len(candidates[i].uniqueRunes), len(candidates[j].uniqueRunes)
		if a == b {
			return candidates[i].word < candidates[j].word
		}
		return a > b
	})
}

// Gibt die Top-N Kandidaten aus
func printCandidates(candidates []wordData, letterFreq map[rune]int, topN int) bool {
	mn := min(topN, len(candidates))
	if mn == 0 {
		fmt.Println("No candidates found.")
		return false
	}

	for i := 0; i < mn; i++ {
		w := candidates[i]

		type letterWithFreq struct {
			letter rune
			freq   int
		}
		letters := make([]letterWithFreq, 0, len(w.uniqueRunes))
		for letter := range w.uniqueRunes {
			letters = append(letters, letterWithFreq{letter, letterFreq[letter]})
		}
		sort.Slice(letters, func(i, j int) bool {
			return letters[i].freq > letters[j].freq
		})

		fmt.Printf("%d. %s (%d unique letters)\n   Letters sorted by global frequency: ", i+1, w.word, len(w.uniqueRunes))
		for _, lf := range letters {
			fmt.Printf("%c ", lf.letter)
		}
		fmt.Println()
	}
	return len(candidates) > topN
}
