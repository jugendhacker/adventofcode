package day11

import (
	"bufio"
	"strings"
	"testing"
)

const countOccupiedData = `#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`

const challengeData = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

func TestCountOccupied(t *testing.T) {
	room := parseData(bufio.NewScanner(strings.NewReader(countOccupiedData)))
	result := countOccupied(room)
	if result != 37 {
		t.Errorf("CountOccupied: wrong result %v, wanted %v", result, 37)
	}
}

func TestChallenge1(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData))
	room := parseData(scanner)
	result := challenge1(room)
	if result != 37 {
		t.Errorf("Challenge 1: wrong result %v, wanted %v", result, 37)
	}
}

func TestChallenge2(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData))
	room := parseData(scanner)
	result := challenge2(room)
	if result != 26 {
		t.Errorf("Challenge 2: wrong result %v, wanted %v", result, 26)
	}
}
