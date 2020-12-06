package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	/*scanner.Scan()
	line := scanner.Text()
	lenLine := len(line)*/

	var fileContent [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []rune
		for _, char := range line{
			lineArray = append(lineArray, char)
		}
		fileContent = append(fileContent, lineArray)
	}

	trees := countTrees(fileContent, 1, 1) * countTrees(fileContent, 1, 3) * countTrees(fileContent, 1, 5) * countTrees(fileContent, 1, 7) * countTrees(fileContent, 2, 1)

	fmt.Printf("Encountered %v trees\n", trees)
}

func countTrees(fileContent [][]rune, yAdd int, xAdd int) int{
	var treeCounter int
	y, x := 0, 0
	for {
		y += yAdd
		if y > len(fileContent) - 1{
			break
		}
		x += xAdd
		if x > len(fileContent[y]) - 1{
			x = x - len(fileContent[y])
		}
		if fileContent[y][x] == rune('#') {
			treeCounter++
		}
	}
	return treeCounter
}
