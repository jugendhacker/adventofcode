package day9

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
	numbers := parseData(scanner)

	challenge1Result := challenge1(numbers, 25)
	fmt.Printf("Challenge 1: %v is invalid\n", challenge1Result)
	fmt.Printf("Challenge 2: result: %v\n", challenge2(numbers, challenge1Result))
}

func parseData(input *bufio.Scanner) (numbers []int) {
	for input.Scan() {
		line := input.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	return numbers
}

func challenge1(numbers []int, preambleLength int) (wrongNumber int) {
	for index := preambleLength; index < len(numbers); index++ {
		valid := false
		for i := 0; i <= preambleLength; i++ {
			for j := i + 1; j <= preambleLength; j++ {
				if sum := numbers[index-i] + numbers[index-j]; sum == numbers[index] {
					valid = true
				}
			}
		}
		if !valid {
			return numbers[index]
		}
	}
	return -1
}

func challenge2(numbers []int, invalidNumber int) (result int) {
	for i := 0; i < len(numbers)-1; i++ {
		sum := numbers[i] + numbers[i+1]
		if sum == invalidNumber {
			return numbers[i] + numbers[i+1]
		}
		var j int
		for j = 2; sum < invalidNumber; j++ {
			sum += numbers[i+j]
		}
		if sum == invalidNumber {
			row := (sort.IntSlice)(numbers[i : i+j])
			row.Sort()
			return row[0] + row[len(row)-1]
		}
	}
	return 0
}
