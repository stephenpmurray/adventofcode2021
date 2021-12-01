package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(ii []int) (c int) {
	for _, i := range ii {
		c = c + i
	}
	return c
}

func main() {
	b, err := os.ReadFile("/home/smurray/adventofcode2021/day01/input")
	if err != nil {
		panic(err)
	}

	ss := strings.Split(string(b), "\n")

	var vals []int
	for _, s := range ss[:len(ss)-1] {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		vals = append(vals, i)
	}

	// part 1
	c := 0
	for i := 0; i < (len(vals) - 1); i++ {
		if vals[i] < vals[i+1] {
			c++
		}
	}
	fmt.Println(c)

	// part 2
	c = 0
	first := sum(vals[0:3])
	for i := 1; i < (len(vals) - 2); i++ {
		second := sum(vals[i : i+3])
		if first < second {
			c++
		}
		first = second
	}
	fmt.Println(c)
}
