package main

import (
	"github.com/stephenpmurray/adventofcode2021/day01"
	"github.com/stephenpmurray/adventofcode2021/day02"
)

func main() {

	// day 1
	vals := day01.ReadFile()
	day01.PartOne(vals)
	day01.PartTwo(vals)

	// day 2
	m := day02.ReadFile()
	day02.PartOne(m)
	day02.PartTwo(m)
}
