package main

import (
	"testing"
)

func Test_d02p1_spreadsheetChecksumIs18(t *testing.T) {
	assertIntEquals(
		day2Part1(`	5 1 9 5
								7 5 3
								2 4 6 8`), 18, t)
}

func Test_d02p2_spreadsheetSumIs9(t *testing.T) {
	assertIntEquals(
		day2Part2(`	5 9 2 8
								9 4 7 3
								3 8 6 5`), 9, t)
}
