package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	var nums []int
	for _, num := range strings.Split(line, ",") {
		numInt, _ := strconv.Atoi(num)
		nums = append(nums, numInt)
	}

	start := time.Now()

	var cycle int
	lastCycle := -1
	for cycle < 80 {
		lowestNum := 9
		for j, num := range nums {
			if num == 0 {
				nums[j] = 6
				nums = append(nums, 8)
			} else {
				nums[j] -= cycle - lastCycle
			}
			if lowestNum > nums[j] {
				lowestNum = nums[j]
			}
		}
		lastCycle = cycle
		if lowestNum == 0 {
			cycle++
		} else {
			cycle += lowestNum
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Nums len: %v, elapsed %s\n", len(nums), elapsed)
}
