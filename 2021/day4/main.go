package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type boardField struct {
	Number int
	Marked bool
}

type boardWin struct {
	Board      [5][5]boardField
	Turns      int
	LastNumber int
}

func main() {
	start := time.Now()
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	var numbers []int
	for _, num := range strings.Split(line, ",") {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, numInt)
	}

	var boards [][5][5]boardField
	var board [5][5]boardField
	var i int
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards = append(boards, board)
			board = [5][5]boardField{}
			i = 0
		} else {
			for j, num := range strings.Fields(line) {
				numInt, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				board[i][j] = boardField{
					Number: numInt,
					Marked: false,
				}
			}
			i++
		}
	}
	boards = append(boards, board)

	var wg sync.WaitGroup
	output := make(chan boardWin)
	for _, board := range boards {
		wg.Add(1)
		go MarkAndCheckBoard(board, numbers, &wg, output)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	currentBest := boardWin{
		Turns: len(numbers) + 1,
	}
	currentWorst := boardWin{
		Turns: 0,
	}
	for winnerBoard := range output {
		if winnerBoard.Turns <= currentBest.Turns {
			currentBest = winnerBoard
		}
		if winnerBoard.Turns >= currentWorst.Turns {
			currentWorst = winnerBoard
		}
	}
	var unmarkedSum int
	for _, row := range currentBest.Board {
		for _, cell := range row {
			if !cell.Marked {
				unmarkedSum += cell.Number
			}
		}
	}
	result := unmarkedSum * currentBest.LastNumber
	fmt.Printf("Fastest win with: %v in %v turns, result is: %v\n", currentBest.Board, currentBest.Turns, result)

	unmarkedSum = 0
	for _, row := range currentWorst.Board {
		for _, cell := range row {
			if !cell.Marked {
				unmarkedSum += cell.Number
			}
		}
	}
	result = unmarkedSum * currentWorst.LastNumber
	fmt.Printf("Slowest win with: %v in %v turns, result is: %v\n", currentWorst.Board, currentWorst.Turns, result)

	elapsed := time.Since(start)
	fmt.Printf("This took %s\n", elapsed)
}

func MarkAndCheckBoard(board [5][5]boardField, nums []int, wg *sync.WaitGroup, output chan boardWin) {
	for numIndex, num := range nums {
	OuterLoop:
		for i, row := range board {
			for j, _ := range row {
				if board[i][j].Number == num {
					board[i][j].Marked = true
					break OuterLoop
				}
			}
		}

		for i, row := range board {
			var rowCount int
			for j := range row {
				if board[i][j].Marked {
					rowCount++
				}
			}
			if rowCount == 5 {
				output <- boardWin{
					Board:      board,
					Turns:      numIndex + 1,
					LastNumber: num,
				}
				wg.Done()
				return
			}
		}

		for j := 0; j < len(board[0]); j++ {
			var colCount int
			for i := 0; i < len(board); i++ {
				if board[i][j].Marked {
					colCount++
				}
			}
			if colCount == 5 {
				output <- boardWin{
					Board:      board,
					Turns:      numIndex + 1,
					LastNumber: num,
				}
				wg.Done()
				return
			}
		}
	}
	wg.Done()
}
