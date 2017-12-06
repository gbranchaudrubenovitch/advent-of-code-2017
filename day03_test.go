package main

import (
	"testing"
)

func Test_d03p1_canBuildGridOf25Points(t *testing.T) {
	// Note: origin is at square 1
	//	17  16  15  14  13
	//	18   5   4   3  12
	//	19   6   1   2  11
	//	20   7   8   9  10
	//	21  22  23  24  25

	grid := day3SpiralSquaresUpTo(25)
	assertPointEquals(grid[0], invalidPoint, t)
	assertPointEquals(grid[1], point{0, 0}, t)
	assertPointEquals(grid[2], point{1, 0}, t)
	assertPointEquals(grid[3], point{1, 1}, t)
	assertPointEquals(grid[4], point{0, 1}, t)
	assertPointEquals(grid[5], point{-1, 1}, t)
	assertPointEquals(grid[6], point{-1, 0}, t)
	assertPointEquals(grid[7], point{-1, -1}, t)
	assertPointEquals(grid[8], point{0, -1}, t)
	assertPointEquals(grid[9], point{1, -1}, t)
	assertPointEquals(grid[10], point{2, -1}, t)
	assertPointEquals(grid[25], point{2, -2}, t)
}

func Test_d03p1_square1Needs0Step(t *testing.T) {
	assertIntEquals(day3Part1(1), 0, t)
}

func Test_d03p1_square12Needs3Steps(t *testing.T) {
	assertIntEquals(day3Part1(12), 3, t)
}

func Test_d03p1_square23Needs2Steps(t *testing.T) {
	assertIntEquals(day3Part1(23), 2, t)
}

func Test_d03p1_square1024Needs31Steps(t *testing.T) {
	assertIntEquals(day3Part1(1024), 31, t)
}

func assertPointEquals(actual point, expected point, t *testing.T) {
	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
