package main

import (
	"testing"
	"strings"
)

var input string = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

var lines []string = strings.Split(input, "\n")

func TestPartOne(t * testing.T) {
	count := PartOne(lines)
	if count != 37 {
		t.Errorf("PartOne expected: 37, got %d", count)
	}
}

func TestPartTwo(t * testing.T) {
	count := PartTwo(lines)
	if count != 26 {
		t.Errorf("PartTwo expected: 26, got %d", count)
	}
}

func BenchmarkPartOne(* testing.B) {
	PartOne(lines);
}

func BenchmarkPartTwo(* testing.B) {
	PartTwo(lines);
}
