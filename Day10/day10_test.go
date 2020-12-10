package day10

import (
	"bufio"
	"strings"
	"testing"
)

const challenge1Data = `16
10
15
5
1
11
7
19
6
12
4`

func TestChallenge1(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challenge1Data))
	adapters := parseData(scanner)
	result := challenge1(adapters)
	if result != 35 {
		t.Errorf("Challenge 1: wrong result %v, wanted %v", result, 35)
	}
}
