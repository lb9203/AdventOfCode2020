package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)


func main() {
	//Read input file
	var file, err = os.Open("day_one/part_two/expense_report.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Parse input into int array
	var expenseReport = map[int]int{}
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()

		var lineAsInt, err = strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		//Fill map with input values
		expenseReport[lineAsInt] = lineAsInt
	}

	var answers, exists = numbersWithSum(expenseReport,2020,3)
	if exists {
		var product int = 1
		for i, value := range answers {
			fmt.Printf("Answer %d: %d\n",i,value)
			product *= value
		}
		fmt.Printf("Product of answers is %d", product)
	} else {
		fmt.Println("Answer does not exist.")
	}
}


func numbersWithSum(numbers map[int]int, sum int, numberOfTerms int) ([]int,bool) {
	if numberOfTerms == 1 {
		var answer, exists = numbers[sum]
		return []int{answer}, exists
	}

	for key, _ := range numbers {
		if key <= sum {
			delete(numbers, key)

			var newNumbers = copyOfMap(numbers)
			var newSum = sum-key

			var answers, exists = numbersWithSum(newNumbers, newSum, numberOfTerms-1)
			if exists {
				return append(answers, key), exists
			}
		}
	}

	return nil, false
}

func copyOfMap(src map[int]int) map[int]int {
	var newMap = map[int]int{}
	for key, value := range src {
		newMap[key] = value
	}
	return newMap
}