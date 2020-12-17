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
	var file, err = os.Open("day_3/slope.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var slope [][]bool

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var slopeLine []bool

		var line = scanner.Text()
		var lineLength = len(line)
		for i := 0; i < lineLength; i++ {
			var charAtIndex = string(line[i])
			if strings.Compare(charAtIndex, ".") == 0 {
				slopeLine = append(slopeLine, false)
			} else {
				slopeLine = append(slopeLine, true)
			}
		}
		slope = append(slope, slopeLine)
	}

	var curX = 0
	var xStep = 3
	var totalTrees = 0

	for _, array := range slope {
		if array[curX%len(array)] {
			totalTrees++
		}
		curX += xStep
	}

	fmt.Printf("Total trees: %d\n", totalTrees)
}
