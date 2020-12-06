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
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var validPasswords int
	for scanner.Scan() {
		line := scanner.Text()
		policy := strings.Split(line, ":")[0]
		min, _ := strconv.Atoi(strings.Split(policy, "-")[0])
		max, _ := strconv.Atoi(strings.Split(strings.Split(policy, "-")[1], " ")[0])
		letter := strings.Split(policy, " ")[1]
		password := strings.TrimSpace(strings.Split(line, ":")[1])
		fmt.Printf("min: %v, max: %v, letter: %v, password: %v\n", min, max, letter, password)
		var letterCounter int
		for _, char := range password {
			if char == []rune(letter)[0] {
				letterCounter++
			}
		}
		if letterCounter >= min && letterCounter <= max {
			validPasswords++
		}
	}

	fmt.Printf("%v passwords are valid\n", validPasswords)

}
