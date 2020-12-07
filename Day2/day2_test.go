package day2

import (
	"bufio"
	"strings"
	"testing"
)

const challengeData = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`

func TestChallenge1(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData))
	policies := parseFile(scanner)
	result := challenge1(policies)
	if result != 2 {
		t.Errorf("Challenge 1: wrong result %v, wanted %v", result, 2)
	}
}

func TestChallenge2(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData))
	policies := parseFile(scanner)
	result := challenge2(policies)
	if result != 1 {
		t.Errorf("Challenge 2: wrong result %v, wanted %v", result, 1)
	}
}
