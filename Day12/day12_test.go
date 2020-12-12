package day12

import (
	"bufio"
	"strings"
	"testing"
)

const challengeData = `F10
N3
F7
R90
F11`

func TestChallenge1(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData))
	moveCommands := parseData(scanner)
	result := challenge1(moveCommands)
	if result != 25 {
		t.Errorf("Challenge 1: wrong result %v, wanted %v", result, 25)
	}
}

func TestChallenge2(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData))
	moveCommands := parseData(scanner)
	result := challenge2(moveCommands)
	if result != 286 {
		t.Errorf("Challenge 1: wrong result %v, wanted %v", result, 286)
	}
}
