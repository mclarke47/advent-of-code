package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Val struct {
	String string
	Int    int
}

func (v *Val) ConcatWith(other Val) Val {

	t := v.String + other.String
	num, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}

	return Val{
		String: t,
		Int:    num,
	}
}

func main() {
	file, err := os.ReadFile("/Users/mclarke/repos/adventofcode2025/three/part1/input.txt")
	if err != nil {
		panic(err)
	}

	array := [][]Val{}
	for _, line := range strings.Split(string(file), "\n") {
		bank := []Val{}
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			bank = append(bank, Val{
				String: string(char),
				Int:    num,
			})
		}
		array = append(array, bank)
	}

	maxes := []Val{}
	for _, bank := range array {
		currMax := Val{}
		for i, num := range bank {
			for j := i + 1; j < len(bank); j++ {
				t := num.ConcatWith(bank[j])
				if t.Int > currMax.Int {
					currMax = t
				}
			}
		}
		maxes = append(maxes, currMax)
	}

	sum := 0
	for _, max := range maxes {
		sum += max.Int
	}
	fmt.Println(sum)
}
