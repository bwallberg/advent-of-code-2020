package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func PartOne(lines []string) int {
	return 0
}

func PartTwo(lines []string) int {
	return 0
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")

	fmt.Printf("Part one: %d\n", PartOne(lines))
	fmt.Printf("Part two: %d\n", PartTwo(lines))
}
