package main

import (
	"testing"
)

func Test_1122_Returns_3(t *testing.T) {
	assertEquals(day1Part1("1122"), 3, t)
}

func Test_1111_Returns_4(t *testing.T) {
	assertEquals(day1Part1("1111"), 4, t)
}

func Test_1234_Returns_0(t *testing.T) {
	assertEquals(day1Part1("1234"), 0, t)
}

func Test_91212129_Returns_9(t *testing.T) {
	assertEquals(day1Part1("91212129"), 9, t)
}

func assertEquals(actual int, expected int, t *testing.T) {
	if (actual != expected) {
		t.Error("Expected", actual, "but got", expected)
	}
}
