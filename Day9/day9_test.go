package day9

import (
	"bufio"
	"strings"
	"testing"
)

const challengeData = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestChallenge1(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData))
	numbers := parseData(scanner)
	result := challenge1(numbers, 5)
	if result != 127 {
		t.Errorf("Challenge 1: wrong result %v, wanted %v", result, 127)
	}
}
