package main

import (
	"testing"
	"strings"
)

var input string = `939
7,13,x,x,59,x,31,19`
var lines []string = strings.Split(input, "\n")


func TestPartOne(t * testing.T) {
	res := PartOne(lines)
	if res != 295 {
		t.Errorf("PartOne failed, expected: 295, got: %d", res)
	}
}

func TestPartTwo(t * testing.T) {
	res := PartTwo(lines)
	if res != 1068781 {
		t.Errorf("PartOne failed, expected: 1068781, got: %d", res)
	}
}

func BenchmarkPartOne(* testing.B) {
	PartOne(lines);
}

func BenchmarkPartTwo(* testing.B) {
	PartTwo(lines);
}
