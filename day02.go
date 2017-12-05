package main

import (
	"math"
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

func day2Part2(input string) int {
	evenDivisionResult := 0

	rows := strings.Split(input, "\n")
	for _, currentRowString := range rows {
		currentRow := numberStringsToInts(strings.Fields(currentRowString))
		sort.Sort(sort.Reverse(sort.IntSlice(currentRow)))

		foundEvenDivisionForRow := false
		for i, big := range currentRow {
			for _, small := range currentRow[i+1:] {
				if math.Remainder(float64(big), float64(small)) == 0 {
					foundEvenDivisionForRow = true
					evenDivisionResult += big / small
					break
				}
			}

			if foundEvenDivisionForRow {
				break
			}
		}
	}
	return evenDivisionResult
}

func numberStringsToInts(numberStrings []string) []int {
	numbers := make([]int, 0)
	for _, currentNumberAsString := range numberStrings {
		number, _ := strconv.Atoi(currentNumberAsString)
		numbers = append(numbers, number)
	}
	return numbers
}
