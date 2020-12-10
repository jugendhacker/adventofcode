package day10

import (
	"bufio"
	"strings"
	"testing"
)

const challengeData = `16
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

const challengeData2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestChallenge1(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData))
	adapters := parseData(scanner)
	result := challenge1(adapters)
	if result != 35 {
		t.Errorf("Challenge 1: wrong result %v, wanted %v", result, 35)
	}
}

func TestChallenge2(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(challengeData2))
	adapters := parseData(scanner)
	result := challenge2(adapters, 0, make(map[int]int))
	if result != 19208 {
		t.Errorf("Challenge 2: wrong result %v, wanted %v", result, 19208)
	}
}
