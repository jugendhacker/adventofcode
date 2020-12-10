package day10

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Run(input string) {
	f, _ := os.Open(input)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	adapters := parseData(scanner)

	fmt.Printf("Challenge 1: result %v\n", challenge1(adapters))
}

func parseData(input *bufio.Scanner) (adapters []int) {
	for input.Scan() {
		line := input.Text()
		adapter, _ := strconv.Atoi(line)
		adapters = append(adapters, adapter)
	}
	adaptersSorted := (sort.IntSlice)(adapters)
	adaptersSorted.Sort()
	adapters = adaptersSorted
	builtInAdapter := adapters[len(adapters)-1] + 3
	adapters = append(adapters, builtInAdapter)
	return adapters
}

func challenge1(adapters []int) (result int) {
	var oneStep, threeStep int
	switch adapters[0] {
	case 1:
		oneStep++
	case 3:
		threeStep++
	}
	for i, adapter := range adapters {
		if i < len(adapters)-1 {
			switch adapters[i+1] - adapter {
			case 1:
				oneStep++
			case 3:
				threeStep++
			}
		}
	}
	return oneStep * threeStep
}
