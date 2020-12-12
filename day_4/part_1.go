package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//Read input file
	file, err := os.Open("day_4/passports.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	validPassports := 0
	curPassport := newPassport()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {

			if validatePassport(curPassport) {
				validPassports++
			}
			curPassport = newPassport()

		} else {

			fields := strings.Split(text, " ")
			for _, field := range fields {
				split := strings.Split(field, ":")
				curPassport.fields[split[0]] = split[1]
			}

		}
	}

	//Validate last passport, because theres no empty line after the last passport (nice input formatting bro)
	if validatePassport(curPassport) {
		validPassports++
	}

	fmt.Printf("Valid passports: %d\n", validPassports)
}

type passport struct {
	fields map[string]string
}

func newPassport() passport {
	return passport{map[string]string{
		"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
		"cid": "",
	}}
}

func validatePassport(p passport) bool {
	for key, value := range p.fields {
		if key != "cid" && len(value) == 0 {
			return false
		}
	}
	return true
}
