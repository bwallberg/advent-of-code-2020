package main

import (
	"strings"
	"testing"
)

var input string = `28
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

var lines []int = MapToInt(strings.Split(input, "\n"))

func TestPartOne(t *testing.T) {
	result := PartOne(lines)
	if result != 220 {
		t.Errorf("PartOne expected 220, got: %d", result)
	}
}

func TestPartTwo(t *testing.T) {
}

func BenchmarkPartOne(*testing.B) {
	PartOne(lines)
}

func BenchmarkPartTwo(*testing.B) {
	PartTwo(lines)
}
