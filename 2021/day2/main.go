package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	var depth, horizontal, depth2, horizontal2, aim int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(line[1])

		switch line[0] {
		case "forward":
			horizontal += value
			horizontal2 += value
			depth2 += value * aim
		case "down":
			depth += value
			aim += value
		case "up":
			depth -= value
			aim -= value
		}
	}

	fmt.Printf("Result part 1: %v\n", depth*horizontal)
	fmt.Printf("Result part 2: %v\n", depth2*horizontal2)
}
