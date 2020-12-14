package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func mapToInt(array []string) (intArray []int) {
	for _, value := range array {
		num, err := strconv.Atoi(value)
		if err == nil {
			intArray = append(intArray, num)
		} else {
			intArray = append(intArray, -1)
		}
	}
	return intArray
}

func PartOne(lines []string) int {
	time, _ := strconv.Atoi(lines[0])
	buses := mapToInt(strings.Split(lines[1], ","))
	nextBus := -1
	last := -1.
	for _, bus := range buses {
		if bus != -1 {
			val := float64(time)/float64(bus) - float64(int(time/bus))
			if val > last {
				last = val
				nextBus = bus
			}
		}
	}
	waitTime := int(math.Ceil(float64(time)/float64(nextBus)))*nextBus - time
	return nextBus * waitTime
}

func PartTwo(lines []string) int {
	buses := mapToInt(strings.Split(lines[1], ","))
	step := 1
	t := 0
	for index, bus := range buses {
		if bus == -1 {
			continue
		}
		for (t+index)%bus != 0 {
			t += step
		}
		step *= bus
	}
	return t
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
