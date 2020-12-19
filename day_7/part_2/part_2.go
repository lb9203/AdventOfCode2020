package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	numberRegex := regexp.MustCompile("[0-9]+")
	bags := map[string]map[string]int{}
	needleBag := "shiny gold"

	var file, err = os.Open("day_7/rules.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, "contain")
		bagColor, contains := split[0], strings.Split(split[1], ",")

		bagColor = strings.ReplaceAll(bagColor, "bags", "")
		bagColor = strings.TrimSpace(bagColor)

		bags[bagColor] = map[string]int{}
		for _, value := range contains {
			value = strings.ReplaceAll(value, "bags", "")
			value = strings.ReplaceAll(value, "bag", "")
			value = strings.ReplaceAll(value, ".", "")

			number := numberRegex.FindString(value)
			numberAsInt, _ := strconv.Atoi(number)

			value = strings.ReplaceAll(value, number, "")
			value = strings.TrimSpace(value)

			bags[bagColor][value] = numberAsInt
		}
	}

	fmt.Printf("$s requires %d bags.", numberContainedBags(bags, bags[needleBag]))
}

func numberContainedBags(bags map[string]map[string]int, currentBag map[string]int) int {
	if len(currentBag) == 0 {
		return 0
	}
	sum := 0
	for bag, number := range currentBag {
		sum += number + number*numberContainedBags(bags, bags[bag])
	}
	return sum
}
