package day3

import (
	"bufio"
	"strings"
	"testing"
)

const challengeData = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestChallenge1(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData))
	fileContent := parseFile(scanner)
	result := challenge1(fileContent)
	if result != 7 {
		t.Errorf("Challenge 1: wrong result %v, wanted %v", result, 7)
	}
}

func TestChallenge2(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData))
	fileContent := parseFile(scanner)
	result := challenge2(fileContent)
	if result != 336 {
		t.Errorf("Challenge 1: wrong result %v, wanted %v", result, 336)
	}
}
