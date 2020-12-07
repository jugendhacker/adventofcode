package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Run(input string) {
	f, _ := os.Open(input)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	seatIDs := parseData(scanner)

	fmt.Printf("Challenge 1: Highest SeatID: %v\n", challenge1(seatIDs))
	fmt.Printf("Challenge 2: Own SeatID: %v\n", challenge2(seatIDs))
}

func parseData(input *bufio.Scanner) (seatIDs map[int]bool) {
	seatIDs = make(map[int]bool)
	for input.Scan() {
		line := input.Text()
		seatID := findBinary(0, 127, line[:7], 'F', 'B')*8 + findBinary(0, 7, line[7:], 'L', 'R')
		seatIDs[seatID] = true
	}
	return seatIDs
}

func challenge1(seatIDs map[int]bool) int {
	var highestSeatID int
	for seatID, _ := range seatIDs {
		if highestSeatID < seatID {
			highestSeatID = seatID
		}
	}

	return highestSeatID
}

func challenge2(seatIDs map[int]bool) int {
	for i := 8; i < 1016; i++ {
		if _, ok := seatIDs[i]; !ok {
			if _, before := seatIDs[i-1]; before {
				if _, after := seatIDs[i+1]; after {
					return i
				}
			}
		}
	}
	return -1
}

func findBinary(min int, max int, actions string, lower rune, upper rune) int {
	if len(actions) == 1 {
		if rune(actions[0]) == lower {
			return min
		} else if rune(actions[0]) == upper {
			return max
		}
	} else {
		middle := (float64(max-min) / 2.0) + float64(min)
		if rune(actions[0]) == lower {
			return findBinary(min, int(math.Floor(middle)), actions[1:], lower, upper)
		} else if rune(actions[0]) == upper {
			return findBinary(int(math.Ceil(middle)), max, actions[1:], lower, upper)
		}
	}
	return -1
}
