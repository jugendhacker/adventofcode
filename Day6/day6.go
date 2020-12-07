package day6

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Run(input string) {
	fileContent, _ := ioutil.ReadFile(input)
	fileText := string(fileContent)

	groups := strings.Split(fileText, "\n\n")

	fmt.Printf("Challenge 1 sum: %v\n", challenge1(groups))
	fmt.Printf("Challenge 2 sum: %v\n", challenge2(groups))
}

func challenge1(groups []string) int {
	var answerSum int
	for _, group := range groups {
		replacer := strings.NewReplacer("\n", "")
		group = replacer.Replace(group)

		chars := make(map[rune]bool)

		for _, char := range group {
			if _, ok := chars[char]; !ok {
				chars[char] = true
			}
		}

		answerSum += len(chars)
	}

	return answerSum
}

func challenge2(groups []string) int {
	var answerSum int
	for _, group := range groups {
		persons := strings.Split(group, "\n")

		answers := make(map[rune]bool)
		for _, char := range persons[0] {
			answers[char] = true
		}
		for _, person := range persons[1:] {
			newAnswers := make(map[rune]bool)
			for _, char := range person {
				if _, ok := answers[char]; ok {
					newAnswers[char] = true
				}
			}
			answers = newAnswers
		}

		answerSum += len(answers)
	}

	return answerSum
}
