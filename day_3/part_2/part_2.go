package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var step = [][]int{
	{1, 1},
	{3, 1},
	{5, 1},
	{7, 1},
	{1, 2},
}

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

	var treeProduct = 1
	for _, steps := range step {
		var totalTrees = 0
		var curX = 0
		for index, slopeLine := range slope {
			if index%steps[1] == 0 {
				if slopeLine[curX] {
					totalTrees++
				}
				curX = (curX + steps[0]) % len(slopeLine)
			}
		}
		treeProduct *= totalTrees
	}

	fmt.Printf("Product of trees: %d\n", treeProduct)
}
