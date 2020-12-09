package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Read input file
	file, err := os.Open("day_2/passwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var validPasswords = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ":")

		policy, password := splitLine[0], splitLine[1]

		splitPolicy := strings.Split(policy, " ")

		policyLetter := splitPolicy[1]

		limits := strings.Split(splitPolicy[0], "-")

		min, err := strconv.Atoi(limits[0])
		if err != nil {
			log.Println(err)
			return
		}

		max, err := strconv.Atoi(limits[1])
		if err != nil {
			log.Println(err)
			return
		}

		password = strings.TrimSpace(password)

		one, two := password[min-1], password[max-1]

		var matches = 0
		if strings.Compare(policyLetter, string(one)) == 0 {
			matches++
		}
		if strings.Compare(policyLetter, string(two)) == 0 {
			matches++
		}
		if matches == 1 {
			validPasswords++
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", validPasswords)
}
