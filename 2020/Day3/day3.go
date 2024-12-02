package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Run(input string) {
	f, _ := os.Open(input)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	fileContent := parseFile(scanner)

	fmt.Printf("Challenge 1: encountered %v trees\n", challenge1(fileContent))
	fmt.Printf("Challenge 2: the solution is %v\n", challenge2(fileContent))
}

func parseFile(input *bufio.Scanner) (fileContent [][]rune) {
	for input.Scan() {
		line := input.Text()
		var lineArray []rune
		for _, char := range line {
			lineArray = append(lineArray, char)
		}
		fileContent = append(fileContent, lineArray)
	}
	return fileContent
}

func challenge1(fileContent [][]rune) (treeCounter int) {
	y, x := 0, 0
	for {
		y++
		if y > len(fileContent)-1 {
			break
		}
		x += 3
		if x > len(fileContent[y])-1 {
			x = x - len(fileContent[y])
		}
		if fileContent[y][x] == rune('#') {
			treeCounter++
		}
	}
	return treeCounter
}

func challenge2(fileContent [][]rune) (solution int) {
	solution = countTrees(fileContent, 1, 1) * countTrees(fileContent, 1, 3) * countTrees(fileContent, 1, 5) * countTrees(fileContent, 1, 7) * countTrees(fileContent, 2, 1)
	return solution
}

func countTrees(fileContent [][]rune, yAdd int, xAdd int) int {
	var treeCounter int
	y, x := 0, 0
	for {
		y += yAdd
		if y > len(fileContent)-1 {
			break
		}
		x += xAdd
		if x > len(fileContent[y])-1 {
			x = x - len(fileContent[y])
		}
		if fileContent[y][x] == rune('#') {
			treeCounter++
		}
	}
	return treeCounter
}
