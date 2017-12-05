package main

import (
	"sort"
	"strconv"
	"strings"
)

func day2Part1(input string) int {
	checkSum := 0

	rows := strings.Split(input, "\n")
	for _, currentRowString := range rows {
		currentRow := numberStringsToInts(strings.Fields(currentRowString))
		sort.Ints(currentRow)

		checkSum += currentRow[len(currentRow)-1] - currentRow[0]
	}

	return checkSum
}

func numberStringsToInts(numberStrings []string) []int {
	numbers := make([]int, 0)
	for _, currentNumberAsString := range numberStrings {
		number, _ := strconv.Atoi(currentNumberAsString)
		numbers = append(numbers, number)
	}
	return numbers
}
