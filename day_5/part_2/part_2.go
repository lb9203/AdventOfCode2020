package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	var seatIds []int

	//Read input file
	file, err := os.Open("day_5/boarding_passes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rowColString := scanner.Text()
		rowString := rowColString[0:7]
		colString := rowColString[7:10]

		row := findRow(rowString, 0, 127)
		col := findCol(colString, 0, 7)

		seatId := row*8 + col

		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)

	expectedSeatId := seatIds[0]
	for _, seatId := range seatIds {
		if seatId != expectedSeatId {
			fmt.Printf("My seatID is %d\n", seatId-1)
			break
		}
		expectedSeatId++
	}

}

func findRow(rowString string, min int, max int) int {
	if min == max {
		return min
	}

	runes := []rune(rowString)
	curChar := string(runes[0])
	if strings.Compare(curChar, "F") == 0 {
		max -= (max - min + 1) / 2
	} else {
		min += (max - min + 1) / 2
	}

	rowString = string(runes[1:len(rowString)])

	return findRow(rowString, min, max)
}

func findCol(colString string, min int, max int) int {

	if min == max {
		return min
	}

	runes := []rune(colString)
	curChar := string(runes[0])
	if strings.Compare(curChar, "L") == 0 {
		max -= (max - min + 1) / 2
	} else {
		min += (max - min + 1) / 2
	}

	colString = string(runes[1:len(colString)])

	return findCol(colString, min, max)
}
