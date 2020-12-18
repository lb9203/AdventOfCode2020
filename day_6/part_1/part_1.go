package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var file, err = os.Open("day_6/customs_forms.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)

	group := map[string]bool{}
	sum := 0

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			sum += len(group)
			group = map[string]bool{}
		} else {
			for _, value := range scanner.Text() {
				group[string(value)] = true
			}
		}
	}

	sum += len(group)

	fmt.Printf("sumOfCounts: %d\n", sum)
}
