package main

import (
	"strconv"
	"strings"
)

func aToI(b byte) int {
	return int(b - '0')
}

func stringsContainingNumberToInts(numberStrings []string) []int {
	numbers := make([]int, 0)
	for _, currentNumberAsString := range numberStrings {
		number, _ := strconv.Atoi(currentNumberAsString)
		numbers = append(numbers, number)
	}
	return numbers
}

func intSliceToString(values []int) string {
	valuesText := []string{}
	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}
	return strings.Join(valuesText, " ")
}

func intAbs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
