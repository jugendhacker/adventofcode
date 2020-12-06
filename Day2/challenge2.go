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
		index1, _ := strconv.Atoi(strings.Split(policy, "-")[0])
		index2, _ := strconv.Atoi(strings.Split(strings.Split(policy, "-")[1], " ")[0])
		letter := strings.Split(policy, " ")[1]
		password := strings.TrimSpace(strings.Split(line, ":")[1])
		if (rune(password[index1-1]) == []rune(letter)[0]) != (rune(password[index2-1]) == []rune(letter)[0]) {
			validPasswords++
		}
	}

	fmt.Printf("%v passwords are valid\n", validPasswords)

}
