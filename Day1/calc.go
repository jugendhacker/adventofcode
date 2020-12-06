package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()

	var numbers []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if (numbers[i] + numbers[j] + numbers[k]) == 2020 {
					fmt.Printf("The Answer is: %v\n", (numbers[i] * numbers[j] * numbers[k]))
				}
			}
		}
	}
}
