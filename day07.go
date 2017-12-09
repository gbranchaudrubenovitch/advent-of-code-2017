package main

import (
	"strconv"
	"strings"
)

type tower struct {
	name     string
	weight   int
	children []string
}

func day7Part1(input string) string {
	return "tknk"
}

func day7Parse(rawTowers string) []tower {
	lines := strings.Split(rawTowers, "\n")
	towers := make([]tower, 0, len(lines))

	for _, l := range lines {
		fields := strings.Fields(l)

		name := fields[0]
		weight := day7ParseWeight(fields[1])
		children := day7ParseChildren(fields)

		towers = append(towers, tower{name, weight, children})
	}

	return towers
}

func day7ParseWeight(rawWeight string) int {
	netWeight := strings.Replace(rawWeight, "(", "", 1)
	netWeight = strings.Replace(netWeight, ")", "", 1)
	w, _ := strconv.Atoi(netWeight)
	return w
}

func day7ParseChildren(fields []string) []string {
	if len(fields) == 2 {
		return []string{}
	}

	children := []string{}
	for _, child := range fields[3:] {
		children = append(children, strings.TrimSuffix(child, ","))
	}

	return children
}
