package main

import (
	"bufio"
	"os"
	"strings"
)

// ParseWordList reads words from a file, one per line.
func ParseWordList(filename string) ([]string, error) {
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

// ComputeLetterFrequencies counts the frequency of each letter in the given words.
func ComputeLetterFrequencies(words []string) map[rune]int {
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