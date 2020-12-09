package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func MapToInt(array []string) (intArray []int) {
	for _, value := range array {
		num, err := strconv.Atoi(value)
		if err == nil {
			intArray = append(intArray, num)
		}
	}
	return intArray
}

func sum(slice []int) int {
	result := 0
	for _, v := range slice {
		result += v
	}
	return result
}

func PartOne(preamble int, lines []int) (result int) {
	for i := preamble; i < len(lines); i++ {
		prev := lines[i-preamble : i]
		sum := lines[i]
		found := false
		for _, num := range prev {
			for _, num2 := range prev {
				if num+num2 == sum {
					found = true
				}
			}
		}
		if !found {
			result = sum
			break
		}
	}
	return
}

func PartTwo(preamble int, lines []int) (result int) {
	target := PartOne(preamble, lines)
	found := false
	for start := range lines {
		for end := start; end < len(lines); end++ {
			value := sum(lines[start:end])
			if value == target {
				sequence := lines[start:end]
				sort.Ints(sequence)
				result = sequence[0] + sequence[len(sequence)-1]
				found = true
			}
		}
		if found {
			break
		}
	}
	return
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := MapToInt(strings.Split(string(input), "\n"))

	fmt.Printf("Part one: %d\n", PartOne(25, lines))
	fmt.Printf("Part two: %d\n", PartTwo(25, lines))
}
