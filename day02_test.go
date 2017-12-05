package main

import (
	"testing"
)

func Test_d02p1_spreadsheetChecksumIs18(t *testing.T) {
	assertEquals(
		day2Part1(`	5 1 9 5
								7 5 3
								2 4 6 8`), 18, t)
}
