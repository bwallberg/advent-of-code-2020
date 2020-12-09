package main

import (
	"testing"
	"strings"
)

var input string = `35
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

var lines []int = MapToInt(strings.Split(input, "\n"))

func TestPartOne(t * testing.T) {
	num := PartOne(5, lines);
	if num != 127 {
		t.Errorf("PartOne, expected: 127, got: %d", num)
	}
}

func TestPartTwo(t * testing.T) {
	num := PartTwo(5, lines);
	if num != 62 {
		t.Errorf("PartTwo, expected: 62, got: %d", num)
	}
}

func BenchmarkPartOne(* testing.B) {
	PartOne(25, lines);
}

func BenchmarkPartTwo(* testing.B) {
	PartTwo(25, lines);
}
