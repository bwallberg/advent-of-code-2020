package main

import(
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func getExpenseSum(expenses []int) int {
	for _, expense := range expenses {
		for _, expense2 := range expenses {
			if expense + expense2 == 2020  {
				return expense * expense2
			}
		}
	} 
	return -1;
}

func mapToInt(array []string) []int  {
	intArray := make([]int, len(array))	
	for _, value := range array {
		num, err := strconv.Atoi(value)
		if err == nil {
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
	fmt.Printf("%d\n", getExpenseSum(expenses))
}