package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	command string
	value   int
	ran     bool
}

func Run(input string) {
	f, _ := os.Open(input)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	instructions := parseData(scanner)
	instrctionsCopy := make(map[int]instruction)
	for i, inst := range instructions {
		instrctionsCopy[i] = inst
	}
	fmt.Printf("Challenge 1: accumulator value is %v\n", challenge1(instructions, 0, 0))
	_, accumulator := challenge2(instrctionsCopy, 0, false)
	fmt.Printf("Challenge 2: accumulator value is %v\n", accumulator)
}

func parseData(input *bufio.Scanner) (instructions map[int]instruction) {
	instructions = make(map[int]instruction)
	var i int
	for input.Scan() {
		line := strings.Split(input.Text(), " ")
		cmd := line[0]
		val, _ := strconv.Atoi(line[1])
		instructions[i] = instruction{
			command: cmd,
			value:   val,
			ran:     false,
		}
		i++
	}
	return instructions
}

func challenge1(instructions map[int]instruction, index int, accumulator int) (accumulatorResult int) {
	if instructions[index].ran {
		return accumulator
	}
	inst := instructions[index]
	inst.ran = true
	instructions[index] = inst
	switch instructions[index].command {
	case "acc":
		accumulator += instructions[index].value
		return challenge1(instructions, index+1, accumulator)
	case "nop":
		return challenge1(instructions, index+1, accumulator)
	case "jmp":
		return challenge1(instructions, index+instructions[index].value, accumulator)
	}
	return accumulator
}

func challenge2(instructions map[int]instruction, index int, changed bool) (noLoop bool, accumulator int) {
	if instructions[index].ran {
		return false, 0
	} else if instructions[index].command == "" {
		return true, 0
	}
	inst := instructions[index]
	inst.ran = true
	instructions[index] = inst
	switch instructions[index].command {
	case "acc":
		noLoop, accumulator = challenge2(instructions, index+1, changed)
		accumulator += instructions[index].value
		return noLoop, accumulator
	case "nop":
		noLoop, accumulator = challenge2(instructions, index+1, changed)
		if !noLoop && !changed {
			changed = true
			noLoop, accumulator = challenge2(instructions, index+instructions[index].value, changed)
		}
		return noLoop, accumulator
	case "jmp":
		noLoop, accumulator = challenge2(instructions, index+instructions[index].value, changed)
		if !noLoop && !changed {
			changed = true
			noLoop, accumulator = challenge2(instructions, index+1, changed)
		}
		return noLoop, accumulator
	}
	return false, 0
}
