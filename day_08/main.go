package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	Name       string
	Argument   int
	Executions int
}

func parseLines(lines []string) (instructions []instruction) {
	for _, line := range lines {
		parts := strings.Split(line, " ")
		name := parts[0]
		argument, _ := strconv.Atoi(strings.ReplaceAll(parts[1], "+", ""))
		instructions = append(instructions, instruction{
			Name:       name,
			Argument:   argument,
			Executions: 0,
		})
	}
	return
}

func PartOne(lines []string) (accumulator int) {
	instructions := parseLines(lines)
	i := 0
	for true {
		instruction := &instructions[i]
		if instruction.Executions > 0 {
			break
		}
		instruction.Executions = instruction.Executions + 1
		switch instruction.Name {
		case "acc":
			accumulator += instruction.Argument
			break
		case "jmp":
			i += instruction.Argument
			break
		}
		if i < 0 {
			i = len(instructions) - i
		} else if i > len(instructions)-1 {
			i = i % len(instructions)
		} else if instruction.Name != "jmp" {
			i++
		}
	}
	return accumulator
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
