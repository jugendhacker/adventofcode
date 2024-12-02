package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	var counter int
	var numbers []int
	scanner.Scan()
	last, _ := strconv.Atoi(scanner.Text())
	numbers = append(numbers, last)
	for scanner.Scan() {
		new, _ := strconv.Atoi(scanner.Text())
		if new > last {
			counter++
		}
		numbers = append(numbers, new)
		last = new
	}

	fmt.Printf("Counter for part 1 is: %v\n", counter)

	counter = 0
	last = numbers[0] + numbers[1] + numbers[2]
	fmt.Printf("First: %v\n", last)
	for i := 3; i < len(numbers); i++ {
		new := numbers[i-2] + numbers[i-1] + numbers[i]
		fmt.Printf("New: %v, Old: %v\n", new, last)
		if new > last {
			counter++
		}
		last = new
	}

	fmt.Printf("Counter for part 2 is: %v\n", counter)
}
