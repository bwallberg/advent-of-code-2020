package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func countUniqueYes(group string) (count int) {
	questions := make(map[string]bool)
	for _, letter := range group {
		if (string(letter) != "\n" && !questions[string(letter)]) {
			count++
			questions[string(letter)] = true
		}
	}
	return
}

func countSharedYes(group string) (count int) {
	questions := make(map[string]int)
	people := len(strings.Split(group, "\n"))
	for _, letter := range group {
		if (string(letter) != "\n") {
			questions[string(letter)]++
		}
	}
	for _, letter := range questions {
		if (letter == people) {
			count++
		}
	}
	return
}

func partOne(lines []string) (sum int) {
	for _, line := range lines {
		sum += countUniqueYes(line);
	}
	return
}

func partTwo(lines []string) (sum int) {
	for _, line := range lines {
		sum += countSharedYes(line);
	}
	return
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n\n")

	start := time.Now()
	fmt.Printf("Part one: %d\n", partOne(lines))
	fmt.Printf("Part two: %d\n", partTwo(lines))
	end := time.Now()
	fmt.Printf("This solution took: %fms \n", float64(end.Sub(start).Microseconds())/1000)
}
