package main

import (
	"fmt"
	"math"
)

type point struct{ X, Y int }

var invalidPoint = point{math.MinInt32, math.MinInt32}

func day3Part1(input int) int {
	return 0
}

func day3SpiralSquaresUpTo(finalSquare int) []point {
	squares := make([]point, 0, finalSquare)
	squares = append(squares, invalidPoint)
	squares = append(squares, point{0, 0})

	// spiral algo is based on https://stackoverflow.com/a/6824451/26605
	var a = 1.0
	var b = 1.0
	for i := 0; i < finalSquare; i++ {
		angle := float64(i) * 2 * math.Pi / 16.0
		x := 0 - (a+b*angle)*math.Cos(angle)
		y := 0 + (a+b*angle)*math.Sin(angle)

		nextPoint := point{int(-math.Floor(x + 0.5)), int(math.Floor(y + 0.5))}
		squares = append(squares, nextPoint)

		fmt.Println(i+1, ":", -x, y, " angle: ", angle)
	}

	return squares
}
