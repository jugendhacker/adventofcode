package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	UNCLEAR Direction = iota
	UP
	DOWN
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	count := 0
	part2_count := 0
	scanner := bufio.NewScanner(file)
	LineLoop:
	for scanner.Scan() {
		line := scanner.Text()
		nums_str := strings.Fields(line)
		nums := make([]int, len(nums_str))
		for i, n := range nums_str {
			nums[i], _ = strconv.Atoi(n)
		}
		if checkLine(nums) {
			count++
			part2_count++
		} else {
			for i := 0; i < len(nums); i++ {
				cutArray := make([]int, len(nums)-1)
				for j := 0; j < i; j++ {
					cutArray[j] = nums[j]
				}
				for j := i + 1; j < len(nums); j++ {
					cutArray[j-1] = nums[j]
				}
				if checkLine(cutArray) {
					part2_count++
					continue LineLoop
				}
			}
		}
	}
	fmt.Printf("Part 1: %v\n", count)
	fmt.Printf("Part 2: %v\n", part2_count)
}

func checkLine(nums []int) bool {
	state := UNCLEAR
	for i := 1; i < len(nums); i++ {
		switch state {
		case UNCLEAR:
			diff := nums[i] - nums[i-1]
			if diff < 0 {
				state = DOWN
				diff *= -1
			} else {
				state = UP
			}
			if diff < 1 || diff > 3 {
				return false
			}
		case UP:
			diff := nums[i] - nums[i-1]
			if diff < 1 || diff > 3 {
				return false
			}
		case DOWN:
			diff := nums[i-1] - nums[i]
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
}
