package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
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

func isValidRange(input, start, end int) bool {
	return input >= start && input <= end
}

func isValidHeight(height string) bool {
	if strings.HasSuffix(height, "cm") {
		number, err := strconv.Atoi(strings.TrimSuffix(height, "cm"))
		if err != nil {
			return false
		}
		return isValidRange(number, 150, 193)
	} else if strings.HasSuffix(height, "in") {
		number, err := strconv.Atoi(strings.TrimSuffix(height, "in"))
		if err != nil {
			return false
		}
		return isValidRange(number, 59, 76)
	} else {
		return false
	}
}

func isValidColor(color string) bool {
	regex := regexp.MustCompile(`#[0-9A-Fa-f]{6}`)
	return regex.MatchString(color)
}

func isValidEyeColor(color string) bool {
	regex := regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
	return regex.MatchString(color)
}

func isValidPid(pid string) bool {
	regexp := regexp.MustCompile(`^([0-9]{9})$`)
	return regexp.MatchString(pid)
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

func isValidPassportPartTwo(passport string) (valid bool) {
	valid = isValidPassportPartOne(passport)
	if valid {
		fields := strings.Split(passport, " ")
		for _, field := range fields {
			values := strings.Split(field, ":")
			fieldType := values[0]
			fieldValue := values[1]
			switch fieldType {
			case "byr":
				year, _ := strconv.Atoi(fieldValue)
				valid = isValidRange(year, 1920, 2002)
				break
			case "iyr":
				year, _ := strconv.Atoi(fieldValue)
				valid = isValidRange(year, 2010, 2020)
				break
			case "eyr":
				year, _ := strconv.Atoi(fieldValue)
				valid = isValidRange(year, 2020, 2030)
				break
			case "hgt":
				valid = isValidHeight(fieldValue)
				break
			case "hcl":
				valid = isValidColor(fieldValue)
				break
			case "ecl":
				valid = isValidEyeColor(fieldValue)
				break
			case "pid":
				valid = isValidPid(fieldValue)
				break
			}
			if !valid {
				return valid
			}
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

func partTwo(lines []string) (valid int) {
	passports := findPassports(lines)
	for _, passport := range passports {
		if isValidPassportPartTwo(passport) {
			valid++
		}
	}

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
	fmt.Printf("Part two: %d\n", partTwo(lines))
	end := time.Now()
	fmt.Printf("This solution took: %fms \n", float64(end.Sub(start).Microseconds())/1000)
}
