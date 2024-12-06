package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var board [][]int
	var x, y int
	direction := 0

	for scanner.Scan() {
		scan_line := scanner.Text()
		line := make([]int, len(scan_line))
		for i, c := range scan_line {
			if c == '#' {
				line[i] = 1
			} else if c == '^' {
				y = i
				x = len(board) + 1
				line[i] = 2
			}
		}
		board = append(board, line)
	}

	for {
		new_x, new_y := x, y
		// Determin new position
		switch direction {
		case 0:
			new_x -= 1
		case 1:
			new_y += 1
		case 2:
			new_x += 1
		case 3:
			new_y -= 1
		}

		// Out of bounds check
		if new_x < 0 || new_x >= len(board) {
			break
		}
		if new_y < 0 || new_y >= len(board[0]) {
			break
		}

		//Check for blockage
		if board[new_x][new_y] == 1 {
			direction = (direction + 1) % 4
		} else { // Do movement
			x = new_x
			y = new_y
			board[x][y] = 2
		}
	}

	count := 0

	for _, line := range board {
		for _, field := range line {
			if field == 2 {
				count++
			}
		}
	}

	fmt.Printf("Part 1: %v\n", count)
}
