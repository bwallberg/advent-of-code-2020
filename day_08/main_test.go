package main

import (
	"testing"
	"strings"
)

var input string = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

var lines []string = strings.Split(input, "\n")

func TestPartOne(t * testing.T) {
	count := PartOne(lines);
	if count != 5 {
		t.Errorf("PartOne, expected: 5, got: %d", count)
	}
}

// func TestPartTwo(* testing.T) {
// 	PartTwo(lines);
// }

func BenchmarkPartOne(* testing.B) {
	PartOne(lines);
}

func BenchmarkPartTwo(* testing.B) {
	PartTwo(lines);
}
