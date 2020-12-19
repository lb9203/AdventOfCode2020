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

	numberCanContain := 0
	for _, contains := range bags {
		if containsBag(bags, contains, needleBag) {
			numberCanContain++
		}
	}

	fmt.Printf("%d bags can contain \"%s\".\n", numberCanContain, needleBag)
}

func containsBag(bags map[string]map[string]int, currentBag map[string]int, needleBag string) bool {
	if _, exists := currentBag[needleBag]; exists {
		return true
	} else {
		for containedBag := range currentBag {
			if containsBag(bags, bags[containedBag], needleBag) {
				return true
			}
		}
	}
	return false
}
