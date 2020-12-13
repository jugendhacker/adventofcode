package day13

import (
	"bufio"
	"strings"
	"testing"
)

const challenge1Data = `939
7,13,x,x,59,x,31,19`

func TestChallenge1(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challenge1Data))
	currentTime, busIDs := parseData(scanner)
	result := challenge1(currentTime, busIDs)
	if result != 295 {
		t.Errorf("Challenge 1: wrong result %v, want %v", result, 295)
	}
}
