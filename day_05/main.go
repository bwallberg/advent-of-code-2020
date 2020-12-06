package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"
)

func getRow(ticket string) (row int) {
	min := 0.0
	max := 127.0
	rowID := ticket[0:7]
	for _, letter := range rowID {
		half := math.Ceil((max - min) / 2)
		if string(letter) == "F" {
			max -= half
		} else if string(letter) == "B" {
			min += half
		}
	}
	row = int(min)
	return
}

func getColumn(ticket string) (column int) {
	min := 0.0
	max := 7.0
	colID := ticket[7:]
	for _, letter := range colID {
		half := math.Ceil((max - min) / 2)
		if string(letter) == "L" {
			max -= half
		} else if string(letter) == "R" {
			min += half
		}
	}
	column = int(min)
	return
}

func contains(slice []int, num int) bool {
	for _, value := range slice {
		if value == num {
			return true
		}
	}
	return false
}

func partOne(tickets []string) int {
	highestSeatID := -1
	for _, ticket := range tickets {
		if len(ticket) > 0 {
			seatID := getRow(ticket)*8 + getColumn(ticket)
			if seatID > highestSeatID {
				highestSeatID = seatID
			}
		}
	}
	return highestSeatID
}

func partTwo(tickets []string) (seatID int) {
	var seatIDs []int;
	for _, ticket := range tickets {
		if len(ticket) > 0 {
			seatIDs = append(seatIDs, getRow(ticket)*8 + getColumn(ticket))
		}
	}

	for i := 0; i < 127 * 8 + 7; i++ {
			if (!contains(seatIDs, i) && contains(seatIDs, i+1) && contains(seatIDs, i-1)) {
				seatID = i
				break
			}
	}
	return seatID
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
