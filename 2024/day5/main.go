package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rules_finished := false
	rules := make(map[int][]int)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if !rules_finished {
			if line == "" {
				rules_finished = true
				continue
			}
			line_split := strings.Split(line, "|")
			rule_l, _ := strconv.Atoi(line_split[0])
			rule_r, _ := strconv.Atoi(line_split[1])
			if _, ok := rules[rule_l]; ok {
				rules[rule_l] = append(rules[rule_l], rule_r)
			} else {
				rules[rule_l] = []int{rule_r}
			}
			continue
		} else {
			line_split := strings.Split(line, ",")
			already := make([]int, len(line_split))
			correct := true
		LineSplit:
			for i, n := range line_split {
				n_num, _ := strconv.Atoi(n)
				rules_r := rules[n_num]
				for _, r := range rules_r {
					if slices.Contains(already, r) {
						correct = false
						break LineSplit
					}
				}
				already[i] = n_num
			}
			if correct {
				mid := int(len(line_split) / 2)
				mid_e, _ := strconv.Atoi(line_split[mid])
				sum += mid_e
			}
		}

	}

	fmt.Println(sum)
}
