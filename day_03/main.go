package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func traverseSlope(lines []string, xStep int, yStep int) (trees int) {
	x, y := 0, 0

	running := true

	for running {
		x += xStep
		y += yStep

		if y < len(lines) && len(lines[y]) > 0 {
			realX := x % len(lines[y])
			if string(lines[y][realX]) == "#" {
				trees++
			}
		} else {
			running = false
		}
	}

	return
}

func partOne(lines []string) (trees int) {
	trees = traverseSlope(lines, 3, 1)
	return
}

func partTwo(lines []string) int {
	return traverseSlope(lines, 1, 1) *
		traverseSlope(lines, 3, 1) *
		traverseSlope(lines, 5, 1) *
		traverseSlope(lines, 7, 1) *
		traverseSlope(lines, 1, 2)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")

	start := time.Now()
	fmt.Printf("Part one: %d\n", partOne(lines))
	fmt.Printf("Part two: %d\n", partTwo(lines))
	end := time.Now()
	fmt.Printf("This solution took: %fms \n", float64(end.Sub(start).Microseconds())/1000)
}
