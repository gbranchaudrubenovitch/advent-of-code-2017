package main

import (
	"testing"
)

func Test_d05p1_exitsIn1Step(t *testing.T) {
	assertIntEquals(day5Part1("1"), 1, t)
}

func Test_d05p1_exitsIn2Steps(t *testing.T) {
	assertIntEquals(day5Part1("0"), 2, t)
}

func Test_d05p1_exitsIn3Steps(t *testing.T) {
	assertIntEquals(day5Part1("1\n-1"), 3, t)
}

func Test_d05p1_exitsIn5Steps(t *testing.T) {
	assertIntEquals(day5Part1("0\n3\n0\n1\n-3"), 5, t)
}

func Test_d05p2_exitsIn10Steps(t *testing.T) {
	assertIntEquals(day5Part2("0\n3\n0\n1\n-3"), 10, t)
}
