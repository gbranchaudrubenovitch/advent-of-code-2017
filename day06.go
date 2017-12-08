package main

import (
	"strings"
)

func day6Part1(input string) int {
	banks := stringsContainingNumberToInts(strings.Fields(input))

	previouslySeenBanks := make(map[string]bool)
	previouslySeenBanks[intSliceToString(banks)] = true

	cycleCount := 0
	for true {
		banks = day6Redistribute(banks)
		cycleCount++

		banksID := intSliceToString(banks)
		if previouslySeenBanks[banksID] {
			return cycleCount
		}
		previouslySeenBanks[banksID] = true
	}

	return -1
}

func day6Redistribute(banks []int) []int {
	biggestBank := 0
	for i, blocks := range banks {
		if blocks > banks[0] {
			biggestBank = i
		}
	}

	blocksToRedistribute := banks[biggestBank]
	banks[biggestBank] = 0

	theOneReceivingTheNextBlock := biggestBank + 1
	for blocksToRedistribute > 0 {
		theOneReceivingTheNextBlock = theOneReceivingTheNextBlock % len(banks)
		banks[theOneReceivingTheNextBlock]++
		blocksToRedistribute--
		theOneReceivingTheNextBlock++
	}

	return banks
}
