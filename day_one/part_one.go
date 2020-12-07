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
	var file, err = os.Open("day_one/expense_report.txt")
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

		expenseReport[lineAsInt] = lineAsInt

		var matchingNumber = 2020 - lineAsInt
		var _, exists = expenseReport[matchingNumber]

		if exists {
			fmt.Printf("My matching numbers are %d and %d, their product is %d.", lineAsInt, matchingNumber, lineAsInt*matchingNumber)
			break
		}
	}

}
