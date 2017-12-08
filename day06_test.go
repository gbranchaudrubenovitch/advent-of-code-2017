package main

import (
	"testing"
)

func Test_d06p1_canRedistribute(t *testing.T) {
	assertIntSliceEquals(day6Redistribute([]int{0, 2, 7, 0}), []int{2, 4, 1, 2}, t)
	assertIntSliceEquals(day6Redistribute([]int{2, 4, 1, 2}), []int{3, 1, 2, 3}, t)
	assertIntSliceEquals(day6Redistribute([]int{3, 1, 2, 3}), []int{0, 2, 3, 4}, t)
	assertIntSliceEquals(day6Redistribute([]int{0, 2, 3, 4}), []int{1, 3, 4, 1}, t)
	assertIntSliceEquals(day6Redistribute([]int{1, 3, 4, 1}), []int{2, 4, 1, 2}, t)
}

func Test_d06p1_canCountCycles(t *testing.T) {
	assertIntEquals(day6Part1("0 2 7 0"), 5, t)
	assertIntEquals(day6Part1("0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0"), 16, t)
}
