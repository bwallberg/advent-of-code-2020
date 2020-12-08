package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type bagsMap map[string]int
type rulesMap map[string]bagsMap

func contains(ruleMap map[string]int, input string) bool {
	_, ok := ruleMap[input]
	return ok
}

func parseRule(rule string) (bags bagsMap) {
	bags = make(bagsMap)
	parts := strings.Split(rule, ", ")
	for _, part := range parts {
		if strings.Contains(rule, "no other bags") {
			break
		}
		bag := strings.Split(strings.Trim(part, " "), " ")
		color := bag[1]+" "+bag[len(bag)-2]
		number, _ := strconv.Atoi(bag[0]) 
		bags[color] = number
	}
	return
}

func parseLine(line string) (color string, bags bagsMap) {
	parts := strings.Split(line, " bags contain ")
	color, bags = parts[0], parseRule(parts[1])
	return
}

func canContainShinyGold(color string, rulesMap rulesMap) bool {
	if len(rulesMap[color]) == 0 {
		return false
	} else if contains(rulesMap[color], "shiny gold") {
		return true
	} else {
		for subColor := range rulesMap[color] {
			if canContainShinyGold(subColor, rulesMap) {
				return true;
			}
		}
	}
	return false
}

func requiredNumberOfBags(color string, rulesMap rulesMap) (count int) {
	for color, number := range rulesMap[color] {
		count += number
		for i := 0; i < number; i++ {
			count += requiredNumberOfBags(color, rulesMap)
		}
	}
	return
}

func partOne(lines []string) (count int) {
	rulesMap := make(rulesMap)
	for _, line := range lines {
		color, bags := parseLine(line)
		rulesMap[color] = bags
	}
	for color := range rulesMap {
		if canContainShinyGold(color, rulesMap) {
			count++
		}
	}
	return
}

func partTwo(lines []string) (count int) {
	rulesMap := make(rulesMap)
	for _, line := range lines {
		color, bags := parseLine(line)
		rulesMap[color] = bags
	}

	count += requiredNumberOfBags("shiny gold", rulesMap)
	return
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")

	start := time.Now()
	fmt.Printf("Part one: %d\n", partOne(lines))
	fmt.Printf("Part one: %d\n", partTwo(lines))
	end := time.Now()
	fmt.Printf("This solution took: %fms \n", float64(end.Sub(start).Microseconds())/1000)
}
