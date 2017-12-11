package main

import (
	"testing"
)

func Test_d07p1_canParseTowers(t *testing.T) {
	towers := day7Parse("siyms (98)\njgtvhkv (1885) -> ykqcpiv, gvupyd, vuyxvq")

	assertIntEquals(len(towers), 2, t)

	assertIntEquals(towers["siyms"].weight, 98, t)
	assertIntEquals(towers["jgtvhkv"].weight, 1885, t)

	assertIntEquals(len(towers["siyms"].children), 0, t)
	assertIntEquals(len(towers["jgtvhkv"].children), 3, t)

	assertStringEquals(towers["jgtvhkv"].children[0], "ykqcpiv", t)
	assertStringEquals(towers["jgtvhkv"].children[1], "gvupyd", t)
	assertStringEquals(towers["jgtvhkv"].children[2], "vuyxvq", t)
}

func Test_d07p1_canFindBottomProgram(t *testing.T) {
	bottomProgram := day7Part1(`pbga (66)
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

func Test_d07p2_canComputeWeightOfDisc(t *testing.T) {
	towers := day7Parse("pbga (66)\nhavc (66)\nqoyq (66)\npadx (45) -> pbga, havc, qoyq")

	assertIntEquals(day7ComputeWeight("padx", towers), 243, t)
}
