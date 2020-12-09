package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run(input string) {
	f, _ := os.Open(input)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	numbers := parseData(scanner)

	fmt.Printf("Challenge 1: %v is invalid\n", challenge1(numbers, 25))
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
