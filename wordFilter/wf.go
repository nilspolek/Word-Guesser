package wordfilter

import (
	"strings"
)

type Words []string

func New(words []string) *Words {
	var w Words
	for _, word := range words {
		w = append(w, word)
	}
	return &w
}

func (w *Words) FilterLength(length int) {
	var result Words
	for _, candidate := range *w {
		if len(candidate) == length {
			result = append(result, candidate)
		}
	}
	*w = result
}

func (w *Words) ToLower() {
	var result Words
	for _, candidate := range *w {
		result = append(result, string([]rune(candidate)))
	}
	*w = result
}

func (w *Words) FilterWrongPlace(runesWrongPlace string) {
	var result Words
	for _, candidate := range *w {
		valid := true	
		for i := range min(len(runesWrongPlace)) {
			if runesWrongPlace[i] == '_' || runesWrongPlace[i] == ' ' {
				continue
			}
			if runesWrongPlace[i] == byte(candidate[i]) {
				valid = false // Buchstabe an dieser Stelle ist falsch
				
			}
			if !strings.ContainsRune(candidate, rune(runesWrongPlace[i])) {
				valid = false // Buchstabe muss im Wort enthalten sein
			}
		}
		if valid {
		result = append(result, candidate)
		}
	}
	*w = result	
	
}

func (w *Words) ExcludeLetters(exclude string) {
	excluded := make(map[rune]bool)
	for _, r := range exclude {
		excluded[r] = true
	}

	var result Words
	for _, candidate := range *w {
		valid := true
		for _, ch := range candidate {
			if excluded[ch] {
				valid = false
				break
			}
		}
		if valid {
			result = append(result, candidate)
		}
	}
	*w = result
}


func (w *Words) FilterRightPlace(runesRightInPlace string) {
	var result Words
	for _, candidate := range *w {
		valid := true
		for i, ch := range candidate {
			if runesRightInPlace[i] == '_' || runesRightInPlace[i] == ' ' || runesRightInPlace[i] == byte(ch) {
				continue // Buchstabe passt oder ist nicht festgelegt
			}
			valid = false
			break
		}
		if valid {
			result = append(result, candidate)
		} 
	}
	*w = result
}
