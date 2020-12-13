package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type earliestDeparture struct {
	BusID     int
	Departure int
}

func Run(input string) {
	f, _ := os.Open(input)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	currentTime, busIDs := parseData(scanner)

	fmt.Printf("Challenge 1: result %v\n", challenge1(currentTime, busIDs))
}

func parseData(input *bufio.Scanner) (currentTime int, busIDs []int) {
	input.Scan()
	currentTime, _ = strconv.Atoi(input.Text())

	input.Scan()
	line := input.Text()
	for _, busID := range strings.Split(line, ",") {
		if busID != "x" {
			busIDInt, _ := strconv.Atoi(busID)
			busIDs = append(busIDs, busIDInt)
		}
	}
	return
}

func challenge1(currentTime int, busIDs []int) int {
	results := make(chan earliestDeparture)
	for _, busID := range busIDs {
		go calculateEarliestDeparture(busID, currentTime, results)
	}
	var resultCounter int
	var earliestBus earliestDeparture
	for result := range results {
		if earliestBus.BusID != 0 {
			if earliestBus.Departure > result.Departure {
				earliestBus = result
			}
		} else {
			earliestBus = result
		}
		resultCounter++
		if resultCounter >= len(busIDs) {
			close(results)
		}
	}
	return earliestBus.BusID * (earliestBus.Departure - currentTime)
}

func calculateEarliestDeparture(busID int, time int, results chan earliestDeparture) {
	var departure int
	for departure < time {
		departure += busID
	}
	result := earliestDeparture{
		BusID:     busID,
		Departure: departure,
	}
	results <- result
}
