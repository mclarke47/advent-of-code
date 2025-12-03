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

const depthLimit = 12

func main() {
	file, err := os.ReadFile("/Users/mclarke/repos/adventofcode2025/three/part2/input.txt")
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

	sum := 0
	for _, bank := range array {
		r := findMax(bank)
		sum += r.Int
	}
	fmt.Println(sum)

}

func findNextMax(bank []Val, start, end int) int {
	maxIndex := start
	for i := start + 1; i <= end; i++ {
		if bank[i].Int > bank[maxIndex].Int {
			maxIndex = i
		}
	}
	return maxIndex
}
func findMax(bank []Val) Val {

	digits := []Val{}
	currentMaxIndex := 0
	for i := range 12 {
		remaining := 12 - i
		lastAllowed := len(bank) - remaining
		nextMaxIndex := findNextMax(bank, currentMaxIndex, lastAllowed)
		digits = append(digits, bank[nextMaxIndex])
		currentMaxIndex = nextMaxIndex + 1
	}
	var result Val
	for _, digit := range digits {
		result = result.ConcatWith(digit)
	}
	return result

}
