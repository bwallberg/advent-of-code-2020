package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func partOne(expenses []int) int {
	for _, expense := range expenses {
		for _, expense2 := range expenses {
			if expense+expense2 == 2020 {
				return expense * expense2
			}
		}
	}
	return -1
}

func partTwo(expenses []int) int {
	for _, expense := range expenses {
		for _, expense2 := range expenses {
			for _, expense3 := range expenses {
				if expense+expense2+expense3 == 2020 {
					return expense * expense2 * expense3
				}
			}
		}
	}
	return -1
}

func mapToInt(array []string) []int {
	intArray := make([]int, 0)
	for _, value := range array {
		num, err := strconv.Atoi(value)
		if err == nil && num != 0 {
			intArray = append(intArray, num)
		}
	}
	return intArray
}

func main() {
	expenseReport, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	expenses := mapToInt(strings.Split(string(expenseReport), "\n"))
	start := time.Now();
	fmt.Printf("Part one: %d\n", partOne(expenses))
	fmt.Printf("Part two: %d\n", partTwo(expenses))
	end := time.Now();
	fmt.Printf("This solution took: %fms \n", float64(end.Sub(start).Microseconds()) / 1000)
}
