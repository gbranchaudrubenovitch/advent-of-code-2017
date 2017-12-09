package main

import (
	"testing"
)

func Test_d07p1_canParseTowers(t *testing.T) {
	towers := day7Parse("siyms (98)\njgtvhkv (1885) -> ykqcpiv, gvupyd, vuyxvq")

	assertIntEquals(len(towers), 2, t)

	assertStringEquals(towers[0].name, "siyms", t)
	assertStringEquals(towers[1].name, "jgtvhkv", t)

	assertIntEquals(towers[0].weight, 98, t)
	assertIntEquals(towers[1].weight, 1885, t)

	assertIntEquals(len(towers[0].children), 0, t)
	assertIntEquals(len(towers[1].children), 3, t)

	assertStringEquals(towers[1].children[0], "ykqcpiv", t)
	assertStringEquals(towers[1].children[1], "gvupyd", t)
	assertStringEquals(towers[1].children[2], "vuyxvq", t)
}

func Test_d07p1_canFindBottomProgram(t *testing.T) {
	bottomProgram := day7Part1(`
		pbga (66)
		xhth (57)
		ebii (61)
		havc (66)
		ktlj (57)
		fwft (72) -> ktlj, cntj, xhth
		qoyq (66)
		padx (45) -> pbga, havc, qoyq
		tknk (41) -> ugml, padx, fwft
		jptl (61)
		ugml (68) -> gyxo, ebii, jptl
		gyxo (61)
		cntj (57)`)
	assertStringEquals(bottomProgram, "tknk", t)
}
