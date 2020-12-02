package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func parseEntry(entry string) (policy, letter, password string) {
	parts := strings.Split(entry, " ")
	policy, password = parts[0], parts[2]
	letter = strings.Replace(parts[1], ":", "", 1)

	return
}

func parsePolicy(policy string) (min, max int) {
	rules := strings.Split(policy, "-")
	min, minErr := strconv.Atoi(rules[0])
	max, maxErr := strconv.Atoi(rules[1])

	if minErr != nil || maxErr != nil {
		log.Fatal(minErr, maxErr)
	}

	return
}

func partOne(passwords []string) int {
	validPasswords := 0
	for _, entry := range passwords {
		if isPasswordEntryValidPartOne(entry) {
			validPasswords++
		}
	}
	return validPasswords
}

func isPasswordEntryValidPartOne(entry string) bool {
	if len(entry) != 0 {
		policy, letter, password := parseEntry(entry)
		min, max := parsePolicy(policy)
		letterCount := strings.Count(password, letter)
		if letterCount >= min && letterCount <= max {
			return true
		}
	}
	return false
}

func partTwo(passwords []string) int {
	validPasswords := 0
	for _, entry := range passwords {
		if isPasswordEntryValidPartTwo(entry) {
			validPasswords++
		}
	}
	return validPasswords
}

func isPasswordEntryValidPartTwo(entry string) bool {
	if len(entry) != 0 {
		policy, letter, password := parseEntry(entry)
		i1, i2 := parsePolicy(policy)

		letters := string(password[i1-1]) + string(password[i2-1])
		letterCount := strings.Count(letters, letter)
		if letterCount == 1 {
			return true
		}
	}
	return false
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	passwords := strings.Split(string(input), "\n")

	start := time.Now()
	fmt.Printf("Part one: %d\n", partOne(passwords))
	fmt.Printf("Part one: %d\n", partTwo(passwords))
	end := time.Now()
	fmt.Printf("This solution took: %fms \n", float64(end.Sub(start).Microseconds())/1000)
}
