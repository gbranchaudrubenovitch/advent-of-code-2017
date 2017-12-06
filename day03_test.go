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

	grid := day3SpiralSquaresUpTo(25, false)
	assertSquareCoordEquals(grid[0], invalidPoint, t)
	assertSquareCoordEquals(grid[1], point{0, 0}, t)
	assertSquareCoordEquals(grid[2], point{1, 0}, t)
	assertSquareCoordEquals(grid[3], point{1, 1}, t)
	assertSquareCoordEquals(grid[4], point{0, 1}, t)
	assertSquareCoordEquals(grid[5], point{-1, 1}, t)
	assertSquareCoordEquals(grid[6], point{-1, 0}, t)
	assertSquareCoordEquals(grid[7], point{-1, -1}, t)
	assertSquareCoordEquals(grid[8], point{0, -1}, t)
	assertSquareCoordEquals(grid[9], point{1, -1}, t)
	assertSquareCoordEquals(grid[10], point{2, -1}, t)
	assertSquareCoordEquals(grid[25], point{2, -2}, t)
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

func Test_d03p2_canBuildValuedGridOf25Points(t *testing.T) {
	// Note: origin is at square 1 (middile)
	// 147  142  133  122   59
	// 304    5    4    2   57
	// 330   10    1    1   54
	// 351   11   23   25   26
	// 362  747  806  854  905

	grid := day3SpiralSquaresUpTo(25, true)
	assertSquareValueEquals(grid[0], 0, t)
	assertSquareValueEquals(grid[1], 1, t)
	assertSquareValueEquals(grid[2], 1, t)
	assertSquareValueEquals(grid[3], 2, t)
	assertSquareValueEquals(grid[4], 4, t)
	assertSquareValueEquals(grid[5], 5, t)
}

func Test_d03p2_firstValueLargerThan25Is26(t *testing.T) {
	assertIntEquals(day3Part2(25), 26, t)
}

func assertSquareCoordEquals(actual square, expectedCoord point, t *testing.T) {
	if actual.Coord != expectedCoord {
		t.Error("Expected", expectedCoord, "but got", actual.Coord)
	}
}

func assertSquareValueEquals(actual square, expectedValue int, t *testing.T) {
	if actual.Value != expectedValue {
		t.Error("Expected", expectedValue, "but got", actual.Value)
	}
}
