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

		switch key {
		case "byr":
			match, _ := regexp.MatchString("^[0-9]{4}$", value)
			if !match {
				return false
			}
			asInt, _ := strconv.Atoi(value)
			if asInt < 1920 || asInt > 2002 {
				return false
			}
			break

		case "iyr":
			match, _ := regexp.MatchString("^[0-9]{4}$", value)
			if !match {
				return false
			}
			asInt, _ := strconv.Atoi(value)
			if asInt < 2010 || asInt > 2020 {
				return false
			}
			break

		case "eyr":
			match, _ := regexp.MatchString("^[0-9]{4}$", value)
			if !match {
				return false
			}
			asInt, _ := strconv.Atoi(value)
			if asInt < 2020 || asInt > 2030 {
				return false
			}
			break

		case "hgt":
			match, _ := regexp.MatchString("^([0-9]+)((in)|(cm))$", value)
			if !match {
				return false
			}
			numberRegEx := regexp.MustCompile("^[0-9]+")
			unitRegEx := regexp.MustCompile("(cm)|(in)$")
			number := numberRegEx.FindString(value)
			numberAsInt, _ := strconv.Atoi(number)
			unit := unitRegEx.FindString(value)
			switch unit {
			case "cm":
				if numberAsInt < 150 || numberAsInt > 193 {
					return false
				}
				break
			case "in":
				if numberAsInt < 59 || numberAsInt > 76 {
					return false
				}
				break
			}
			break

		case "hcl":
			match, _ := regexp.MatchString("^#([0-9]|[a-f]){6}$", value)
			if !match {
				return false
			}
			break

		case "ecl":
			match, _ := regexp.MatchString("^((amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)|){1}$", value)
			if !match {
				return false
			}
			break

		case "pid":
			match, _ := regexp.MatchString("^[0-9]{9}$", value)
			if !match {
				return false
			}
			break

		case "cid":
			break

		default:
			return false
		}
	}
	return true
}
