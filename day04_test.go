package main

import (
	"testing"
)

func Test_d04p1_validPassphrases(t *testing.T) {
	assertIntEquals(day4Part1("aa bb cc dd ee\naa bb cc dd aaa"), 2, t)
}

func Test_d04p1_invalidPassphrase(t *testing.T) {
	assertIntEquals(day4Part1("aa bb cc dd aa"), 0, t)
}
