package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var highestSeatID int
	for scanner.Scan(){
		line := scanner.Text()
		seatID := findBinary(0, 127, line[:7], 'F', 'B') * 8 + findBinary(0, 7, line[7:], 'L', 'R')
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Printf("Highest SeatID: %v\n", highestSeatID)
}

func findBinary(min int, max int, actions string, lower rune, upper rune) int{
	if len(actions) == 1 {
		if rune(actions[0]) == lower {
			return min
		} else if rune(actions[0]) == upper {
			return max
		}
	} else {
		middle := (float64(max - min) / 2.0) + float64(min)
		if rune(actions[0]) == lower {
			return findBinary(min, int(math.Floor(middle)), actions[1:], lower, upper)
		} else if rune(actions[0]) == upper {
			return findBinary(int(math.Ceil(middle)), max, actions[1:], lower, upper)
		}
	}
	return -1
}