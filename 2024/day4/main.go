package main

import (
	"bufio"
	"fmt"
	"os"
)

var searchStringPart1 = []rune("XMAS")
var searchStringPart2 = []rune("MAS")

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var board [][]rune

	for scanner.Scan() {
		scan_line := scanner.Text()
		line := make([]rune, len(scan_line))
		for i, c := range scan_line {
			line[i] = c
		}
		board = append(board, line)
	}

	count := 0
	count_part2 := 0
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == 'X' {
				if checkRight(board, x, y, searchStringPart1) {
					count++
				}
				if checkLeft(board, x, y, searchStringPart1) {
					count++
				}
				if checkDown(board, x, y, searchStringPart1) {
					count++
				}
				if checkUp(board, x, y, searchStringPart1) {
					count++
				}
				if checkRightDown(board, x, y, searchStringPart1) {
					count++
				}
				if checkLeftUp(board, x, y, searchStringPart1) {
					count++
				}
				if checkRightUp(board, x, y, searchStringPart1) {
					count++
				}
				if checkLeftDown(board, x, y, searchStringPart1) {
					count++
				}
			}
			if y < len(board)-1 && y > 0 && x < len(board[y])-1 && x > 0 {
				if checkRightDown(board, x-1, y-1, searchStringPart2) || checkLeftUp(board, x+1, y+1, searchStringPart2) {
					if checkRightUp(board, x-1, y+1, searchStringPart2) || checkLeftDown(board, x+1, y-1, searchStringPart2) {
						count_part2++
					}
				}
			}
		}
	}

	fmt.Printf("Part 1: %v\n", count)
	fmt.Printf("Part 2: %v\n", count_part2)
}

func checkRight(board [][]rune, x, y int, searchString []rune) bool {
	for i, c := range searchString {
		if x+i < len(board[y]) {
			if board[y][x+i] != c {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func checkLeft(board [][]rune, x, y int, searchString []rune) bool {
	for i, c := range searchString {
		if x-i >= 0 {
			if board[y][x-i] != c {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func checkDown(board [][]rune, x, y int, searchString []rune) bool {
	for i, c := range searchString {
		if y+i < len(board) {
			if board[y+i][x] != c {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func checkUp(board [][]rune, x, y int, searchString []rune) bool {
	for i, c := range searchString {
		if y-i >= 0 {
			if board[y-i][x] != c {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func checkRightDown(board [][]rune, x, y int, searchString []rune) bool {
	for i, c := range searchString {
		if y+i < len(board) && x+i < len(board[y]) {
			if board[y+i][x+i] != c {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func checkLeftUp(board [][]rune, x, y int, searchString []rune) bool {
	for i, c := range searchString {
		if y-i >= 0 && x-i >= 0 {
			if board[y-i][x-i] != c {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func checkRightUp(board [][]rune, x, y int, searchString []rune) bool {
	for i, c := range searchString {
		if y-i >= 0 && x+i < len(board[y]) {
			if board[y-i][x+i] != c {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func checkLeftDown(board [][]rune, x, y int, searchString []rune) bool {
	for i, c := range searchString {
		if y+i < len(board) && x-i >= 0 {
			if board[y+i][x-i] != c {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
