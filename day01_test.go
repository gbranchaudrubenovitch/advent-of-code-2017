package main

import (
	"testing"
)

func Test_d01p1_1122_Returns_3(t *testing.T) {
	assertIntEquals(day1Part1("1122"), 3, t)
}

func Test_d01p1_1111_Returns_4(t *testing.T) {
	assertIntEquals(day1Part1("1111"), 4, t)
}

func Test_d01p1_1234_Returns_0(t *testing.T) {
	assertIntEquals(day1Part1("1234"), 0, t)
}

func Test_d01p1_91212129_Returns_9(t *testing.T) {
	assertIntEquals(day1Part1("91212129"), 9, t)
}

func Test_d01p2_1212_Returns_6(t *testing.T) {
	assertIntEquals(day1Part2("1212"), 6, t)
}

func Test_d01p2_1221_Returns_0(t *testing.T) {
	assertIntEquals(day1Part2("1221"), 0, t)
}

func Test_d01p2_123425_Returns_4(t *testing.T) {
	assertIntEquals(day1Part2("123425"), 4, t)
}

func Test_d01p2_123123_Returns_12(t *testing.T) {
	assertIntEquals(day1Part2("123123"), 12, t)
}

func Test_d01p2_12131415_Returns_4(t *testing.T) {
	assertIntEquals(day1Part2("12131415"), 4, t)
}
