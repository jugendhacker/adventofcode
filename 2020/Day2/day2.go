package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pwPolicy struct {
	policy   string
	value1   int
	value2   int
	letter   rune
	password string
}

func Run(input string) {
	f, _ := os.Open(input)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	policies := parseFile(scanner)

	fmt.Printf("Challenge 1: %v passwords are valid\n", challenge1(policies))
	fmt.Printf("Challenge 2: %v passwords are valid\n", challenge2(policies))

}

func parseFile(input *bufio.Scanner) (policies []pwPolicy) {
	for input.Scan() {
		line := input.Text()
		policy := strings.Split(line, ":")[0]
		min, _ := strconv.Atoi(strings.Split(policy, "-")[0])
		max, _ := strconv.Atoi(strings.Split(strings.Split(policy, "-")[1], " ")[0])
		letter := []rune(strings.Split(policy, " ")[1])[0]
		password := strings.TrimSpace(strings.Split(line, ":")[1])
		policies = append(policies, pwPolicy{
			policy:   policy,
			value1:   min,
			value2:   max,
			letter:   letter,
			password: password,
		})
	}
	return policies
}

func challenge1(policies []pwPolicy) (validPasswords int) {
	for _, policy := range policies {
		var letterCounter int
		for _, char := range policy.password {
			if char == policy.letter {
				letterCounter++
			}
		}
		if letterCounter >= policy.value1 && letterCounter <= policy.value2 {
			validPasswords++
		}
	}
	return validPasswords
}

func challenge2(policies []pwPolicy) (validPasswords int) {
	for _, policy := range policies {
		if (rune(policy.password[policy.value1-1]) == policy.letter) != (rune(policy.password[policy.value2-1]) == policy.letter) {
			validPasswords++
		}
	}
	return validPasswords
}
