package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	bitCount := make([]int64, 12)
	var lines []uint16
	var lineCount int64
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()

		var lineBits uint16
		for i, char := range line {
			if char == '1' {
				bitCount[i]++
				lineBits |= uint16(1 << (11 - i))
			}
		}
		lines = append(lines, lineBits)
	}

	var gamma, epsilon uint16
	for i, count := range bitCount {
		if count >= (lineCount / 2) {
			gamma |= uint16(1 << (11 - i))
		} else {
			epsilon |= uint16(1 << (11 - i))
		}
	}

	result := uint64(gamma) * uint64(epsilon)

	fmt.Printf("PART 1: Gamma: %016b, Epsilon: %016b, result: %v\n", gamma, epsilon, result)

	var oxygenLines, output []uint16
	oxygenLines = lines
	bitPosition := 5
	for len(oxygenLines) > 1 {
		var oneCount int
		for _, line := range oxygenLines {
			if (line & (1 << (16 - bitPosition))) == (1 << (16 - bitPosition)) {
				oneCount++
			}
		}
		var compareMask uint16
		if oneCount >= (len(oxygenLines) / 2) {
			compareMask = (1 << (16 - bitPosition))
		}
		for _, line := range oxygenLines {
			if (line & (1 << (16 - bitPosition))) == compareMask {
				output = append(output, line)
			}
		}
		bitPosition++
		oxygenLines = output
		output = []uint16{}
	}

	var co2Lines []uint16
	co2Lines = lines
	bitPosition = 5
	for len(co2Lines) > 1 {
		var oneCount int
		for _, line := range co2Lines {
			if (line & (1 << (16 - bitPosition))) == (1 << (16 - bitPosition)) {
				oneCount++
			}
		}
		var compareMask uint16
		if oneCount < (len(co2Lines) / 2) {
			compareMask = (1 << (16 - bitPosition))
		}
		for _, line := range co2Lines {
			if (line & (1 << (16 - bitPosition))) == compareMask {
				output = append(output, line)
			}
		}
		bitPosition++
		co2Lines = output
		output = []uint16{}
	}

	result = uint64(oxygenLines[0]) * uint64(co2Lines[0])
	fmt.Printf("Ogygen: %012b, CO2 Rating: %012b, result: %v\n", oxygenLines[0], co2Lines[0], result)
}
