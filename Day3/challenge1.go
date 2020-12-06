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

	var treeCounter int
	y, x := 0, 0
	for {
		y++
		if y > len(fileContent) - 1{
			break
		}
		x += 3
		if x > len(fileContent[y]) - 1{
			x = x - len(fileContent[y])
		}
		if fileContent[y][x] == rune('#') {
			treeCounter++
		}
	}

	fmt.Printf("Encountered %v trees\n", treeCounter)
}