package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
	"sort"
)

func main() {
	var left []int
	var right []int
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		num, _ := strconv.Atoi(nums[0])
		left = append(left, num)
		num, _ = strconv.Atoi(nums[1])
		right = append(right, num)
	}
	sort.Ints(left)
	sort.Ints(right)

	diff_sum := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff *= -1
		}
		diff_sum += diff
	}
	fmt.Printf("Part 1: %v\n", diff_sum)

	m := make(map[int]int)
	for _, n := range left {
		m[n] = 0
	}
	for _, n := range right {
		m[n] += 1
	}

	mul_sum := 0
	for _, n := range left {
		mul_sum += n * m[n]
	}
	fmt.Printf("Part 2: %v\n", mul_sum)
}
