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
	fmt.Printf("Challenge 1: accumulator value is %v\n", challenge1(instructions, 0, 0))
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
	return 0
}
