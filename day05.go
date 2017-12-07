package main

import (
	"strings"
)

type computeOffset func(int) int

func day5Part1(input string) int {
	addOne := func(currentOffset int) int { return currentOffset + 1 }
	return countStepsToExit(input, addOne)
}

func day5Part2(input string) int {
	addOrRemoveOne := func(currentOffset int) int {
		if currentOffset >= 3 {
			return currentOffset - 1
		}
		return currentOffset + 1
	}
	return countStepsToExit(input, addOrRemoveOne)
}

func countStepsToExit(input string, newOffset computeOffset) int {
	offsets := stringsContainingNumberToInts(strings.Split(input, "\n"))
	stepCount := 0

	pc := 0
	for pc < len(offsets) {
		stepCount++

		offset := offsets[pc]
		offsets[pc] = newOffset(offset)

		pc += offset
	}
	return stepCount
}
