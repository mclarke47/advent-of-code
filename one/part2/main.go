package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	leftTurn  = "L"
	rightTurn = "R"
)

type Turn struct {
	Direction string
	Magnitude int
}

func (t *Turn) isLeftTurn() bool {
	return t.Direction == leftTurn
}

func (t *Turn) isRightTurn() bool {
	return t.Direction == rightTurn
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func performLeftTurn(currentMagnitude int, turnMagnitude int) (int, int) {
	value := currentMagnitude
	numOfZeros := 0
	for i := 0; i < turnMagnitude; i++ {
		value--
		if value == -1 {
			value = 99
		}
		if value == 0 {
			numOfZeros++
		}
	}
	return value, numOfZeros
}

func performRightTurn(currentMagnitude int, turnMagnitude int) (int, int) {
	value := currentMagnitude
	numOfZeros := 0
	for i := 0; i < turnMagnitude; i++ {
		value++
		if value == 100 {
			value = 0
		}
		if value == 0 {
			numOfZeros++
		}
	}
	return value, numOfZeros
}

func progressTurn(currentMagnitude int, turn Turn) (int, int) {
	if turn.isLeftTurn() {
		return performLeftTurn(currentMagnitude, turn.Magnitude)
	} else if turn.isRightTurn() {
		return performRightTurn(currentMagnitude, turn.Magnitude)
	}
	return currentMagnitude, 0
}

func main() {
	file, err := os.ReadFile("/Users/mclarke/repos/adventofcode2025/one/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	var turns []Turn
	for _, line := range lines {
		direction := string(line[0])
		magnitude, err := strconv.Atoi(strings.TrimPrefix(line, direction))
		if err != nil {
			panic(err)
		}
		turns = append(turns, Turn{
			Direction: direction,
			Magnitude: magnitude,
		})
	}
	var numOfZeros int
	timesAtZero := 0

	currentMagnitude := 50
	for _, turn := range turns {
		currentMagnitude, numOfZeros = progressTurn(currentMagnitude, turn)
		timesAtZero += numOfZeros
	}
	fmt.Println(timesAtZero)
}
