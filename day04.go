package main

import (
	"sort"
	"strings"
)

type phraseValidator func(wordsInPhrase []string) bool

func day4Part1(input string) int {
	return day4SecurityAnalyzer(input, noDuplicateWords)
}

func day4Part2(input string) int {
	return day4SecurityAnalyzer(input, noAnagrams)
}

func day4SecurityAnalyzer(input string, phraseIsValid phraseValidator) int {
	passphrases := strings.Split(input, "\n")

	validCount := 0
	for _, phrase := range passphrases {
		words := strings.Fields(phrase)

		if phraseIsValid(words) {
			validCount++
		}
	}
	return validCount
}

func noDuplicateWords(wordsInPhrase []string) bool {
	uniqueWords := make(map[string]bool)

	for _, word := range wordsInPhrase {
		uniqueWords[word] = true
	}

	return len(uniqueWords) == len(wordsInPhrase)
}

func noAnagrams(wordsInPhrase []string) bool {
	for i, word := range wordsInPhrase {
		for _, candidateWord := range wordsInPhrase[i+1:] {
			if len(word) != len(candidateWord) {
				continue
			}

			sortedWord := []rune(word)
			sort.Slice(sortedWord, func(i int, j int) bool { return sortedWord[i] < sortedWord[j] })

			sortedCandidate := []rune(candidateWord)
			sort.Slice(sortedCandidate, func(i int, j int) bool { return sortedCandidate[i] < sortedCandidate[j] })

			if string(sortedWord) == string(sortedCandidate) {
				return false
			}
		}
	}
	return true
}
