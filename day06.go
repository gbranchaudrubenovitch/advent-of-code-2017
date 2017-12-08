package main

import (
	"strings"
)

func day6Part1(input string) int {
	banks := stringsContainingNumberToInts(strings.Fields(input))

	count, _ := day6CountCycles(banks)
	return count
}

func day6Part2(input string) int {
	banks := stringsContainingNumberToInts(strings.Fields(input))

	_, banksAtStartOfLoop := day6CountCycles(banks)
	count, _ := day6CountCycles(banksAtStartOfLoop)
	return count
}

func day6CountCycles(banks []int) (int, []int) {
	previouslySeenBanks := make(map[string]bool)
	previouslySeenBanks[intSliceToString(banks)] = true

	cycleCount := 0
	for true {
		banks = day6Redistribute(banks)
		cycleCount++

		banksID := intSliceToString(banks)
		if previouslySeenBanks[banksID] {
			return cycleCount, banks
		}
		previouslySeenBanks[banksID] = true
	}

	return -1, banks
}

func day6Redistribute(banks []int) []int {
	biggestBank := 0
	for i, blocks := range banks {
		if blocks > banks[biggestBank] {
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
