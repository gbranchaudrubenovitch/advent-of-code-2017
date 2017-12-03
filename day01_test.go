package main

import (
	"testing"
)

func Test_Part1_1122_Returns_3(t *testing.T) {
	assertEquals(day1Part1("1122"), 3, t)
}

func Test_Part1_1111_Returns_4(t *testing.T) {
	assertEquals(day1Part1("1111"), 4, t)
}

func Test_Part1_1234_Returns_0(t *testing.T) {
	assertEquals(day1Part1("1234"), 0, t)
}

func Test_Part1_91212129_Returns_9(t *testing.T) {
	assertEquals(day1Part1("91212129"), 9, t)
}

func Test_Part2_1212_Returns_6(t *testing.T) {
	assertEquals(day1Part2("1212"), 6, t)
}

func Test_Part2_1221_Returns_0(t *testing.T) {
	assertEquals(day1Part2("1221"), 0, t)
}

func Test_Part2_123425_Returns_4(t *testing.T) {
	assertEquals(day1Part2("123425"), 4, t)
}

func Test_Part2_123123_Returns_12(t *testing.T) {
	assertEquals(day1Part2("123123"), 12, t)
}

func Test_Part2_12131415_Returns_4(t *testing.T) {
	assertEquals(day1Part2("12131415"), 4, t)
}

func assertEquals(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
