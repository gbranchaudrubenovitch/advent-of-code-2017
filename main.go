package main

import (
	"fmt"

	"github.com/gbranchaudrubenovitch/advent-of-code-2017/inputs"
)

func main() {
	fmt.Println("--- Day 1: Inverse Captcha ---")
	fmt.Println(fmt.Sprintf("  1) solution to captcha is %d", day1Part1(inputs.Day01)))
	fmt.Println(fmt.Sprintf("  2) solution to new captcha is %d", day1Part2(inputs.Day01)))
	fmt.Println()

	fmt.Println("--- Day 2: Corruption Checksum ---")
	fmt.Println(fmt.Sprintf("  1) checksum for the spreadsheet is %d", day2Part1(inputs.Day02)))
	fmt.Println(fmt.Sprintf("  2) sum of each row's factors is %d", day2Part2(inputs.Day02)))
	fmt.Println()

	fmt.Println("--- Day 3: Spiral Memory ---")
	fmt.Println(fmt.Sprintf("  1) # of steps required to carry data is %d", day3Part1(inputs.Day03)))
	fmt.Println(fmt.Sprintf("  2) first bigger value written is %d", day3Part2(inputs.Day03)))
	fmt.Println()

	fmt.Println("--- Day 4: High-Entropy Passphrases ---")
	fmt.Println(fmt.Sprintf("  1) # of valid passphrases is %d", day4Part1(inputs.Day04)))
	fmt.Println(fmt.Sprintf("  2) # of valid passphrases under new policy is %d", day4Part2(inputs.Day04)))
	fmt.Println()

	fmt.Println("--- Day 5: A Maze of Twisty Trampolines, All Alike ---")
	fmt.Println(fmt.Sprintf("  1) # of steps to reach exit is %d", day5Part1(inputs.Day05)))
	fmt.Println(fmt.Sprintf("  2) new # of steps to reach exit is %d", day5Part2(inputs.Day05)))
	fmt.Println()

	fmt.Println("--- Day 6: Memory Reallocation ---")
	fmt.Println(fmt.Sprintf("  1) # of redistribution cycles is %d", day6Part1(inputs.Day06)))
	fmt.Println(fmt.Sprintf("  2) # of cycles in loop is %d", day6Part2(inputs.Day06)))
	fmt.Println()
}
