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

func contains(slice []int, num int) bool {
	for _, value := range slice {
		if value == num {
			return true
		}
	}
	return false
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

func runPartTwo(from int, history map[int]int, lines []int) int {
	if _, ok := history[from]; ok {
		return history[from]
	}
	if from == len(lines) - 1 {
		return 1
	} 
	count := 0;
	for i := from + 1; i < len(lines); i++ {
		diff := lines[i] - lines[from] 
		if diff <= 3 {
			count += runPartTwo(i, history, lines);
		}
	}
	history[from] = count
	return count 
}

func PartTwo(lines []int) (variations int) {
	sort.Sort(sort.IntSlice(lines))
	lines = append([]int{0}, lines...)
	lines = append(lines, lines[len(lines)-1]+3)
	return runPartTwo(0, make(map[int]int), lines);
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
