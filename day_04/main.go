package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func getRequiredFields() []string {
	return []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" /* "cid" 🙈 */}
}

func findPassports(lines []string) (passports []string) {
	var tmp []string
	for _, line := range lines {
		if len(line) != 0 {
			tmp = append(tmp, line)
		} else {
			passports = append(passports, strings.Join(tmp, " "))
			tmp = tmp[:0]
		}
	}
	return
}

func isValidPassportPartOne(passport string) (valid bool) {
	valid = true
	for _, field := range getRequiredFields() {
		if !strings.Contains(passport, field+":") {
			valid = false
		}
	}
	return valid
}

func partOne(lines []string) (valid int) {
	passports := findPassports(lines)
	for _, passport := range passports {
		if isValidPassportPartOne(passport) {
			valid++
		}
	}

	return
}

// func isValidPassportPartTwo(passport string) (valid bool) {
// 	valid = true
// 	if isValidPassportPartOne(passport) {
// 		fields := strings.Split(passport, " ")
// 		for _, field := range fields {
// 			values := strings.Split(field, ":")
// 			fieldType := values[0]
// 			fieldValue := values[1]
// 			switch fieldType {
// 			case "byr":
// 				break
// 			case "iyr":
// 				break
// 			case "eyr":
// 				break
// 			case "hgt":
// 				break
// 			case "hcl":
// 				break
// 			case "ecl":
// 				break
// 			case "pid":
// 				break
// 			}

// 		}
// 	} else {

// 	}

// 	return valid
// }

// func partTwo(lines []string) (valid int) {
// 	passports := findPassports(lines)
// 	for _, passport := range passports {
// 		if isValidPassportPartTwo(passport) {
// 			valid++
// 		}
// 	}

// 	return
// }

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	start := time.Now()
	fmt.Printf("Part one: %d\n", partOne(lines))
	// fmt.Printf("Part one: %d\n", partTwo(lines))
	end := time.Now()
	fmt.Printf("This solution took: %fms \n", float64(end.Sub(start).Microseconds())/1000)
}
