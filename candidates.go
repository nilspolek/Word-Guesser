package main

import (
	"fmt"
	"sort"

	wordfilter "github.com/nilspolek/Word-Guesser/wordFilter"
)

type Candidate struct {
	Word        string
	UniqueRunes map[rune]bool
}

type CandidateList []Candidate

func (cl CandidateList) Strings() []string {
	var result []string
	for _, c := range cl {
		result = append(result, c.Word)
	}
	return result
}

func (cl CandidateList) Sort() {
	sort.Slice(cl, func(i, j int) bool {
		a, b := len(cl[i].UniqueRunes), len(cl[j].UniqueRunes)
		if a == b {
			return cl[i].Word < cl[j].Word
		}
		return a > b
	})
}

func (cl CandidateList) Print(letterFreq map[rune]int, topN int) bool {
	mn := min(topN, len(cl))
	if mn == 0 {
		fmt.Println("No candidates found.")
		return false
	}
	for i := 0; i < mn; i++ {
		w := cl[i]
		type letterWithFreq struct {
			letter rune
			freq   int
		}
		letters := make([]letterWithFreq, 0, len(w.UniqueRunes))
		for letter := range w.UniqueRunes {
			letters = append(letters, letterWithFreq{letter, letterFreq[letter]})
		}
		sort.Slice(letters, func(i, j int) bool {
			return letters[i].freq > letters[j].freq
		})
		fmt.Println(w.Word)
	}
	return len(cl) > topN
}

func FilterCandidates(words []string, length int, excluded, runesRightInPlace, runesWrongPlace string) CandidateList {
	var candidates CandidateList
	wds := wordfilter.New(words)
	wds.FilterLength(length)
	wds.ToLower()
	if len(runesWrongPlace) == length {
		wds.FilterWrongPlace(runesWrongPlace)
	}
	if len(runesRightInPlace) == length {
		wds.FilterRightPlace(runesRightInPlace)
	}
	wds.ExcludeLetters(excluded)
	for _, word := range *wds {
		seen := make(map[rune]bool)
		for _, r := range word {
			if r >= 'a' && r <= 'z' {
				seen[r] = true
			}
		}
		candidates = append(candidates, Candidate{Word: word, UniqueRunes: seen})
	}
	return candidates
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
