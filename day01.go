package main

type indexProducer func(int, int) int

func day1Part1(input string) int {
	takeTheNextOne := func(currentIndex int, length int) int {
		return (currentIndex + 1) % length
	}

	return loopAndSum(input, takeTheNextOne)
}

func day1Part2(input string) int {
	takeTheOneHalfwayAround := func(currentIndex int, length int) int {
		halfOfLength := length / 2
		return (currentIndex + halfOfLength) % length
	}

	return loopAndSum(input, takeTheOneHalfwayAround)
}

func loopAndSum(input string, nextIndex indexProducer) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		currentDigit := aToI(input[i])
		nextDigit := aToI(input[nextIndex(i, len(input))])

		if currentDigit == nextDigit {
			sum += currentDigit
		}
	}
	return sum
}

func aToI(b byte) int { return int(b - '0') }
