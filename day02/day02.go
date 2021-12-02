package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Heading struct {
	Height int
	Length int
	Aim    int
}

type Move struct {
	Dir string
	Val int
}

func ReadFile() []Move {
	b, err := os.ReadFile("/home/smurray/adventofcode2021/day02/input")
	if err != nil {
		panic(err)
	}
	ss := strings.Split(string(b), "\n")

	var mm []Move
	for _, s := range ss[:len(ss)-1] { // ignore empty string at end

		line := strings.Split(s, " ")
		i, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		m := Move{
			Dir: line[0],
			Val: i,
		}
		mm = append(mm, m)
		// fmt.Println(m)
	}
	return mm
}

func PartOne(mm []Move) {
	var h Heading

	for _, m := range mm {
		switch m.Dir {
		case "forward":
			h.Length += m.Val
		case "down":
			h.Height += m.Val
		case "up":
			h.Height -= m.Val
		}
	}
	fmt.Println(h.Height * h.Length)
}

func PartTwo(mm []Move) {
	var h Heading

	for _, m := range mm {
		switch m.Dir {
		case "forward":
			h.Length += m.Val
			h.Height += (m.Val * h.Aim)
		case "down":
			h.Aim += m.Val
		case "up":
			h.Aim -= m.Val
		}
	}

	fmt.Println(h.Height * h.Length)
}
