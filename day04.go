package main

import (
	"strings"
)

func day4Part1(input string) int {
	passphrases := strings.Split(input, "\n")

	validCount := 0
	for _, phrase := range passphrases {
		words := strings.Fields(phrase)

		uniqueWords := make(map[string]bool)
		for _, word := range words {
			uniqueWords[word] = true
		}

		if len(uniqueWords) == len(words) {
			validCount++
		}
	}
	return validCount
}
