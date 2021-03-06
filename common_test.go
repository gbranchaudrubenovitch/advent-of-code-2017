package main

import (
	"testing"
)

func assertIntEquals(actual int, expected int, t *testing.T) {
	t.Helper()
	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func assertStringEquals(actual string, expected string, t *testing.T) {
	t.Helper()
	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func assertIntSliceEquals(actual []int, expected []int, t *testing.T) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Error("Expected 2 arrays of the same length, but got", expected, "and", actual)
		return
	}
	for i, v := range actual {
		if v != expected[i] {
			t.Error("Expected", expected, "but got", actual, "Slices differ at index", i)
			return
		}
	}
}
