package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	var last25 []int

	var file, err = os.Open("day_9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())

		if len(last25) < 25 {
			last25 = append(last25, number)
		} else {
			if !isValid(last25, number) {
				fmt.Printf("%d is not valid.", number)
				break
			}
			last25 = last25[1:]
			last25 = append(last25, number)
		}
	}
}

func isValid(last25 []int, number int) bool {
	last25map := map[int]int{}

	for _, value := range last25 {
		last25map[value]++
	}

	for key := range last25map {
		last25map[key]--
		match := int(math.Abs(float64(number - key)))
		if last25map[match] > 0 {
			return true
		}
		last25map[key]++
	}

	return false
}
