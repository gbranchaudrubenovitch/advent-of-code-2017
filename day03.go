package main

import (
	"math"
)

type point struct{ X, Y int }
type direction string

var invalidPoint = point{math.MinInt32, math.MinInt32}
var up = direction("up")
var down = direction("down")
var left = direction("left")
var right = direction("right")

func day3Part1(lastSquare int) int {
	squares := day3SpiralSquaresUpTo(lastSquare)
	theSquare := squares[lastSquare]

	return intAbs(theSquare.X) + intAbs(theSquare.Y)
}

func day3SpiralSquaresUpTo(finalSquare int) []point {
	origin := point{0, 0}
	squares := make([]point, 0, finalSquare)
	squares = append(squares, invalidPoint)
	squares = append(squares, origin)

	gridIsOccupiedAt := make(map[point]bool)
	currentPoint := origin
	currentDirection := right
	gridIsOccupiedAt[currentPoint] = true

	for i := 0; i < finalSquare; i++ {
		thereWasCollision := false
		originalDirection := currentDirection
		candidatePoint := computeNextPoint(currentPoint, currentDirection)

		for gridIsOccupiedAt[candidatePoint] {
			thereWasCollision = true
			currentDirection = computeNextDirectionAfterCollision(currentDirection)
			candidatePoint = computeNextPoint(currentPoint, currentDirection)
		}

		if thereWasCollision {
			currentDirection = originalDirection
		} else {
			currentDirection = computeNextDirection(originalDirection)
		}

		currentPoint = candidatePoint
		gridIsOccupiedAt[currentPoint] = true
		squares = append(squares, currentPoint)
	}

	return squares
}

func computeNextPoint(currentPoint point, nextDirection direction) point {
	if nextDirection == up {
		return point{currentPoint.X, currentPoint.Y + 1}
	} else if nextDirection == left {
		return point{currentPoint.X - 1, currentPoint.Y}
	} else if nextDirection == down {
		return point{currentPoint.X, currentPoint.Y - 1}
	}
	return point{currentPoint.X + 1, currentPoint.Y}
}

func computeNextDirection(currentDirection direction) direction {
	if currentDirection == up {
		return left
	} else if currentDirection == left {
		return down
	} else if currentDirection == down {
		return right
	}
	return up
}

func computeNextDirectionAfterCollision(currentDirection direction) direction {
	if currentDirection == up {
		return right
	} else if currentDirection == left {
		return up
	} else if currentDirection == down {
		return left
	}
	return down
}
