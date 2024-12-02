package day6

import (
	"strings"
	"testing"
)

const challenge1Data = `abc

a
b
c

ab
ac

a
a
a
a

b`

const challenge2Data = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestChallenge1(t *testing.T) {
	groups := strings.Split(challenge1Data, "\n\n")
	result := challenge1(groups)
	if result != 11 {
		t.Errorf("Challenge 1: wrong result %v, want %v", result, 11)
	}
}

func TestChallenge2(t *testing.T) {
	groups := strings.Split(challenge2Data, "\n\n")
	result := challenge2(groups)
	if result != 6 {
		t.Errorf("Challenge 2: wrong result %v, want %v", result, 6)
	}
}
