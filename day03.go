package main

import (
	"math"
)

type direction int
type point struct{ X, Y int }
type square struct {
	Coord point
	Value int
}

var invalidPoint = point{math.MinInt32, math.MinInt32}
var up = direction(0)
var left = direction(1)
var down = direction(2)
var right = direction(3)

func day3Part1(lastSquare int) int {
	squares := day3SpiralSquaresUpTo(lastSquare, false)
	theSquare := squares[lastSquare]

	return intAbs(theSquare.Coord.X) + intAbs(theSquare.Coord.Y)
}

func day3Part2(lastSquare int) int {
	squares := day3SpiralSquaresUpTo(lastSquare, true)
	return squares[len(squares)-1].Value
}

func day3SpiralSquaresUpTo(finalSquare int, computeValueAndStopWhenSquareReached bool) []square {
	startingSquare := square{point{0, 0}, 1}
	squares := make([]square, 0, finalSquare)
	squares = append(squares, square{invalidPoint, 0})
	squares = append(squares, startingSquare)

	currentSquare := startingSquare
	currentDirection := right

	grid := make(map[point]int)
	gridIsOccupied := func(s square) bool { return grid[s.Coord] > 0 }
	grid[currentSquare.Coord] = currentSquare.Value

	for i := 0; i < finalSquare; i++ {
		thereWasCollision := false
		originalDirection := currentDirection
		candidateSquare := computeNextSquare(currentSquare, currentDirection)

		for gridIsOccupied(candidateSquare) {
			thereWasCollision = true
			currentDirection = computeNextDirectionAfterCollision(currentDirection)
			candidateSquare = computeNextSquare(currentSquare, currentDirection)
		}

		if thereWasCollision {
			currentDirection = originalDirection
		} else {
			currentDirection = computeNextDirection(originalDirection)
		}

		currentSquare = candidateSquare
		if computeValueAndStopWhenSquareReached {
			grid[currentSquare.Coord] = computeValue(currentSquare.Coord, grid)
		} else {
			grid[currentSquare.Coord] = 1
		}
		currentSquare.Value = grid[currentSquare.Coord]

		squares = append(squares, currentSquare)

		if computeValueAndStopWhenSquareReached && currentSquare.Value > finalSquare {
			return squares
		}
	}

	return squares
}

func computeNextSquare(currentSquare square, nextDirection direction) square {
	currentPoint := currentSquare.Coord
	if nextDirection == up {
		return square{point{currentPoint.X, currentPoint.Y + 1}, 0}
	} else if nextDirection == left {
		return square{point{currentPoint.X - 1, currentPoint.Y}, 0}
	} else if nextDirection == down {
		return square{point{currentPoint.X, currentPoint.Y - 1}, 0}
	}
	return square{point{currentPoint.X + 1, currentPoint.Y}, 0}
}

func computeNextDirection(currentDirection direction) direction {
	return (currentDirection + 1) % 4
}

func computeNextDirectionAfterCollision(currentDirection direction) direction {
	return (currentDirection - 1) % 4
}

func computeValue(location point, grid map[point]int) int {
	return grid[point{location.X, location.Y + 1}] +
		grid[point{location.X + 1, location.Y + 1}] +
		grid[point{location.X + 1, location.Y}] +
		grid[point{location.X + 1, location.Y - 1}] +
		grid[point{location.X, location.Y - 1}] +
		grid[point{location.X - 1, location.Y - 1}] +
		grid[point{location.X - 1, location.Y}] +
		grid[point{location.X - 1, location.Y + 1}]
}
