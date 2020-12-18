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

	group := [](map[string]bool){}
	sum := 0

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			for len(group) > 1 {
				group[1] = intersect(group[0], group[1])
				group = group[1:]
			}
			sum += len(group[0])
			group = [](map[string]bool){}
		} else {
			person := map[string]bool{}
			for _, val := range scanner.Text() {
				person[string(val)] = true
			}
			group = append(group, person)
		}
	}

	for len(group) > 1 {
		group[1] = intersect(group[0], group[1])
		group = group[1:]
	}
	sum += len(group[0])

	fmt.Printf("Sum: %d\n", sum)
}

func intersect(map1 map[string]bool, map2 map[string]bool) map[string]bool {
	ret := map[string]bool{}
	for key := range map1 {
		if _, exists := map2[key]; exists {
			ret[key] = true
		}
	}
	return ret
}
