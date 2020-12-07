package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bagWithCount struct {
	bag   string
	count int
}

func Run(input string) {
	f, _ := os.Open(input)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rules := make(map[string][]bagWithCount)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " bags contain ")

		var contents []bagWithCount
		if !strings.HasSuffix(line, "contain no other bags.") {
			for _, bag := range strings.Split(lineSplit[1], ", ") {
				bagName := strings.Join(strings.Split(bag, " ")[1:], " ")
				bagName = strings.TrimSuffix(bagName, ".")
				bagName = strings.TrimSuffix(bagName, " bag")
				bagName = strings.TrimSuffix(bagName, " bags")
				bagCount, _ := strconv.Atoi(strings.Split(bag, " ")[0])
				bagStruct := bagWithCount{
					bag:   bagName,
					count: bagCount,
				}
				contents = append(contents, bagStruct)
			}
		}

		rules[lineSplit[0]] = contents
	}

	fmt.Printf("Challenge 1: counted %v gold\n", challenge1(rules))
	fmt.Printf("Challenge 2: need %v bags\n", challenge2(rules, "shiny gold")-1)
}

func challenge1(rules map[string][]bagWithCount) int {
	var goldCounter int
	for name, _ := range rules {
		if findGold(rules, name) {
			goldCounter++
		}
	}

	return goldCounter
}

func findGold(rules map[string][]bagWithCount, color string) bool {
	for _, bag := range rules[color] {
		if bag.bag == "shiny gold" {
			return true
		} else {
			if findGold(rules, bag.bag) {
				return true
			}
		}
	}
	return false
}

func challenge2(rules map[string][]bagWithCount, color string) int {
	var count int
	for _, bag := range rules[color] {
		count += bag.count * challenge2(rules, bag.bag)
	}
	return count + 1
}
