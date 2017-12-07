package main

import (
	"strings"
)

func day5Part1(input string) int {
	offsets := stringsContainingNumberToInts(strings.Split(input, "\n"))
	stepCount := 0

	pc := 0
	for pc < len(offsets) {
		stepCount++

		offset := offsets[pc]
		offsets[pc] = offset + 1

		pc += offset
	}
	return stepCount
}
