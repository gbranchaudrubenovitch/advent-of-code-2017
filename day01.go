package main

func day1Part1(inputString string) int {
	aToI := func (b byte) int { return int(b - '0') }

	sum := 0
	for i := 0; i < len(inputString); i ++ {
		currentDigit := aToI(inputString[i])

		var nextDigit int
		if i < len(inputString) - 1 {
			nextDigit = aToI(inputString[i + 1])
		} else {
			nextDigit = aToI(inputString[0])
		}

		if currentDigit == nextDigit {
			sum += currentDigit
		}
	}

	return sum
}
