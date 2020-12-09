package main

import (
	"testing"
	"strings"
)

var input string = ""
var lines []string = strings.Split(input, "\n")

func BenchmarkPartOne(* testing.B) {
	PartOne(lines);
}

func BenchmarkPartTwo(* testing.B) {
	PartTwo(lines);
}
