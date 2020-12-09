package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	Name     string
	Argument int
}

func contains(slice []int, num int) bool {
	for _, value := range slice {
		if value == num {
			return true
		}
	}
	return false
}

func parseLines(lines []string) (instructions []instruction) {
	for _, line := range lines {
		parts := strings.Split(line, " ")
		name := parts[0]
		argument, _ := strconv.Atoi(strings.ReplaceAll(parts[1], "+", ""))
		instructions = append(instructions, instruction{
			Name:     name,
			Argument: argument,
		})
	}
	return
}

func runInstructions(instructions []instruction) (accumulator, i int) {
	var log []int
	for i < len(instructions) {
		instruction := &instructions[i]
		if contains(log, i) {
			break
		}
		log = append(log, i)
		switch instruction.Name {
		case "acc":
			accumulator += instruction.Argument
			break
		case "jmp":
			i += instruction.Argument
			break
		}
		if instruction.Name != "jmp" {
			i++
		}
	}
	return accumulator, i
}

func PartOne(lines []string) (accumulator int) {
	instructions := parseLines(lines)
	accumulator, _ = runInstructions(instructions)
	return
}

func PartTwo(lines []string) (accumulator int) {
	instructions := parseLines(lines)
	var potential []int
	for i, instruction := range instructions {
		if instruction.Name == "jmp" || instruction.Name == "nop" {
			potential = append(potential, i)
		}
	}

	// 🤮 not happy with this
	for _, i := range potential {
		instruction := &instructions[i]
		if instruction.Name == "jmp" {
			instruction.Name = "nop"
		} else if instruction.Name == "nop" {
			instruction.Name = "jmp"
		} else {
			continue
		}
		tmpAccumulator, index := runInstructions(instructions)
		if index == len(instructions) {
			accumulator = tmpAccumulator
			break
		}
		if instruction.Name == "nop" {
			instruction.Name = "jmp"
		} else if instruction.Name == "jmp" {
			instruction.Name = "nop"
		}
	}
	return
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
