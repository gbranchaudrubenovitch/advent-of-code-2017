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

func Test_d04p2_validPassphrases(t *testing.T) {
	assertIntEquals(day4Part2("abcde fghij\na ab abc abd abf abj\niiii oiii ooii oooi oooo"), 3, t)
}

func Test_d04p2_invalidPassphrases(t *testing.T) {
	assertIntEquals(day4Part2("abcde xyz ecdab\noiii ioii iioi iiio"), 0, t)
}
