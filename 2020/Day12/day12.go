package day12

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type moveCommand struct {
	command rune
	value   int
}

type wayPoint struct {
	relNorth int
	relEast  int
}

func Run(input string) {
	f, _ := os.Open(input)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	moveCommands := parseData(scanner)

	fmt.Printf("Challenge 1: manhattan distance is %v\n", challenge1(moveCommands))
	fmt.Printf("Challenge 2: manhattan distance is %v\n", challenge2(moveCommands))
}

func parseData(input *bufio.Scanner) (moveCommands []moveCommand) {
	for input.Scan() {
		line := input.Text()
		var command moveCommand
		command.command = []rune(line)[0]
		command.value, _ = strconv.Atoi(line[1:])
		moveCommands = append(moveCommands, command)
	}
	return
}

func challenge1(moveCommands []moveCommand) (manhattanDistance int) {
	var north, east, direction int
	for _, command := range moveCommands {
		switch command.command {
		case 'N':
			north += command.value
		case 'E':
			east += command.value
		case 'S':
			north -= command.value
		case 'W':
			east -= command.value
		case 'L':
			direction -= command.value
			for direction < 0 {
				direction = direction + 360
			}
		case 'R':
			direction += command.value
			for direction >= 360 {
				direction = direction - 360
			}
		case 'F':
			switch direction {
			case 0:
				east += command.value
			case 90:
				north -= command.value
			case 180:
				east -= command.value
			case 270:
				north += command.value
			}
		}
	}
	return int(math.Abs(float64(north)) + math.Abs(float64(east)))
}

func challenge2(moveCommands []moveCommand) (manhattanDistance int) {
	waypoint := wayPoint{
		relNorth: 1,
		relEast:  10,
	}
	var north, east int
	for _, command := range moveCommands {
		switch command.command {
		case 'N':
			waypoint.relNorth += command.value
		case 'S':
			waypoint.relNorth -= command.value
		case 'E':
			waypoint.relEast += command.value
		case 'W':
			waypoint.relEast -= command.value
		case 'L':
			n := int(math.Abs(float64(command.value)) / 90)
			waypoint = rotateWaypoint(false, waypoint, n)
		case 'R':
			n := int(math.Abs(float64(command.value)) / 90)
			waypoint = rotateWaypoint(true, waypoint, n)
		case 'F':
			for i := 0; i < command.value; i++ {
				north += waypoint.relNorth
				east += waypoint.relEast
			}
		}
	}
	return int(math.Abs(float64(north)) + math.Abs(float64(east)))
}

func rotateWaypoint(right bool, waypoint wayPoint, n int) (newWaypoint wayPoint) {
	if right {
		newWaypoint.relEast = waypoint.relNorth
		newWaypoint.relNorth = -waypoint.relEast
	} else {
		newWaypoint.relNorth = waypoint.relEast
		newWaypoint.relEast = -waypoint.relNorth
	}
	if n > 1 {
		newWaypoint = rotateWaypoint(right, newWaypoint, n-1)
	}
	return
}
