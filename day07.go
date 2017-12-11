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
	towers := day7Parse(input)
	return day7FindRootProgramName(towers)
}

func day7Part2(input string) int {
	return -1
}

func day7ComputeWeight(programName string, towers map[string]tower) int {
	currentTower := towers[programName]
	if len(currentTower.children) == 0 {
		return currentTower.weight
	}

	weight := currentTower.weight
	for _, child := range currentTower.children {
		weight += day7ComputeWeight(child, towers)
	}
	return weight
}

func day7FindRootProgramName(towers map[string]tower) string {
	allChildren := make(map[string]bool, 0)
	for _, tower := range towers {
		for _, child := range tower.children {
			allChildren[child] = true
		}
	}

	for name := range towers {
		if !allChildren[name] {
			return name
		}
	}
	return ""
}

func day7Parse(rawTowers string) map[string]tower {
	lines := strings.Split(rawTowers, "\n")
	towers := make(map[string]tower)

	for _, l := range lines {
		fields := strings.Fields(l)

		name := fields[0]
		weight := day7ParseWeight(fields[1])
		children := day7ParseChildren(fields)

		towers[name] = tower{name, weight, children}
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
