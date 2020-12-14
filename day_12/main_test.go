package main

import (
	"strings"
	"testing"
)

var input string = `F10
N3
F7
R90
F11`

var lines []string = strings.Split(input, "\n")

func TestPartOne(t * testing.T) {
	res := PartOne(lines)
	if res != 25 {
		t.Errorf("PartOne failed, expected: 25, got: %d", res)
	}
}

func TestPartTwo(t * testing.T) {
	res := PartTwo(lines)
	if res != 286 {
		t.Errorf("PartOne failed, expected: 286, got: %d", res)
	}
}

func BenchmarkPartOne(* testing.B) {
	PartOne(lines);
}

func BenchmarkPartTwo(* testing.B) {
	PartTwo(lines);
}
