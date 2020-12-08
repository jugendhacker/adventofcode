package day8

import (
	"bufio"
	"strings"
	"testing"
)

const challenge1Data = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestChallenge1(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challenge1Data))
	instructions := parseData(scanner)
	result := challenge1(instructions, 0, 0)
	if result != 5 {
		t.Errorf("Challenge 1: wrong result %v, wanted %v", result, 5)
	}
}
