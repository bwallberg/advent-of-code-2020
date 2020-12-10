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

func PartOne(lines []int) int {
	diffMap := make(map[int]int)
	sort.Sort(sort.IntSlice(lines))
	lines = append(lines, lines[len(lines)-1]+3)
	rating := 0
	for _, i := range lines {
		diff := i - rating
		if diff <= 3 {
			diffMap[diff]++
			rating = i
		}
	}
	return diffMap[1] * diffMap[3]
}

func PartTwo(lines []int) int {
	return 0
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := MapToInt(strings.Split(string(input), "\n"))

	fmt.Printf("Part one: %d\n", PartOne(lines))
	fmt.Printf("Part two: %d\n", PartTwo(lines))
}
