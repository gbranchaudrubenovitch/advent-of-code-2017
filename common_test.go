package main

import (
	"testing"
)

func assertIntEquals(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
