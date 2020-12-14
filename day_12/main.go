package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	north   = "N"
	east    = "E"
	south   = "S"
	west    = "W"
	forward = "F"
	left    = "L"
	right   = "R"
)

var direction [4]string = [4]string{north, east, south, west}

type instruction struct {
	action string
	value  int
}

func getInstructions(lines []string) (instructions []instruction) {
	for _, line := range lines {
		val, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, instruction{
			action: string(line[:1]),
			value:  val,
		})
	}
	return
}

func navigate(steps int, dir string) (x, y int) {
	switch dir {
	case north:
		y += steps
		break
	case east:
		x += steps
		break
	case south:
		y -= steps
		break
	case west:
		x -= steps
		break
	}
	return
}

func PartOne(lines []string) int {
	x, y := 0, 0
	dir := 1
	instructions := getInstructions(lines)
	for _, instruction := range instructions {
		switch instruction.action {
		case north, east, south, west:
			xx, yy := navigate(instruction.value, instruction.action)
			x += xx
			y += yy
			break
		case forward:
			xx, yy := navigate(instruction.value, direction[dir])
			x += xx
			y += yy
			break
		case left:
			dir = (dir - instruction.value/90)
			break
		case right:
			dir = (dir + instruction.value/90)
			break
		}

		if dir > len(direction)-1 {
			dir -= 4
		} else if dir < 0 {
			dir += 4
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func PartTwo(lines []string) int {
	x, y := 0, 0
	wx, wy := 10, 1
	instructions := getInstructions(lines)
	for _, instruction := range instructions {
		switch instruction.action {
		case north, east, south, west:
			xx, yy := navigate(instruction.value, instruction.action)
			wx += xx
			wy += yy
			break
		case left:
			for i := 0; i < instruction.value/90; i++ {
				wx, wy = -wy, wx
			}
			break
		case right:
			for i := 0; i < instruction.value/90; i++ {
				wx, wy = wy, -wx
			}
			break
		case forward:
			for i := 0; i < instruction.value; i++ {
				x += wx
				y += wy				
			}
			break
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
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
