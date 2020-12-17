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

		count := strings.Count(password, policyLetter)

		password = strings.TrimSpace(password)

		if count <= max && count >= min {
			validPasswords++
		}

	}

	fmt.Printf("Number of valid passwords: %d\n", validPasswords)
}
