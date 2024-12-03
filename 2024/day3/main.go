package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

const (
	STATE_NONE int = iota
	STATE_COMMAND_START
	STATE_NUM1
	STATE_SEPERATOR
	STATE_NUM2
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	state := STATE_NONE
	var num1, num2 int
	result := 0
	enabled := true
	result2 := 0
	ScanLoop:
	for scanner.Scan() {
		sym := scanner.Text()
		switch state {
		case STATE_NONE:
			if sym == "m" {
				for _, c := range "ul(" {
					if !scanner.Scan() {
						break ScanLoop
					}
					sym = scanner.Text()
					if sym != string(c) {
						continue ScanLoop
					}
				}
				state = STATE_COMMAND_START
				num1 = 0
				num2 = 0
			} else if sym == "d" {
				if enabled {
					for _, c := range "on't()" {
						if !scanner.Scan() {
							break ScanLoop
						}
						sym = scanner.Text()
						if sym != string(c) {
							continue ScanLoop
						}
					}
					enabled = false
				} else {
					for _, c := range "o()" {
						if !scanner.Scan() {
							break ScanLoop
						}
						sym = scanner.Text()
						if sym != string(c) {
							continue ScanLoop
						}
					}
					enabled = true
				}
			}
		case STATE_COMMAND_START:
			if n, err := strconv.Atoi(sym); err == nil {
				state = STATE_NUM1
				num1 = num1 * 10 + n
			} else {
				state = STATE_NONE
			}
		case STATE_NUM1:
			if n, err := strconv.Atoi(sym); err == nil {
				num1 = num1 * 10 + n
			} else if sym == "," {
				state = STATE_SEPERATOR
			} else {
				state = STATE_NONE
			}
		case STATE_SEPERATOR:
			if n, err := strconv.Atoi(sym); err == nil {
				state = STATE_NUM2
				num2 = num2 * 10 + n
			} else {
				state = STATE_NONE
			}
		case STATE_NUM2:
			if n, err := strconv.Atoi(sym); err == nil {
				num2 = num2 * 10 + n
			} else if sym == ")" {
				fmt.Printf("Got mul(%v,%v)\n", num1, num2)
				result += num1 * num2
				if enabled {
					result2 += num1 * num2
				}
				state = STATE_NONE
			} else {
				state = STATE_NONE
			}
		}
	}
	fmt.Printf("Part 1: %v\nPart 2: %v\n", result, result2)
}
