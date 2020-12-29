package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	const invalid = 1398413738

	runningSum := 0
	var answerSet []int
	min, max := int(^uint(0)>>1), 0

	var file, err = os.Open("day_9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		runningSum += number
		answerSet = append(answerSet, number)

		for runningSum > invalid {
			takeOut := answerSet[0]
			runningSum -= takeOut
			answerSet = answerSet[1:]
		}

		if runningSum == invalid {

			for _, value := range answerSet {
				if value >= max {
					max = value
				}
				if value <= min {
					min = value
				}
			}
			fmt.Printf("Answer reached! min = %d, max = %d, min+max = %d\n", min, max, min+max)
			break
		}
	}
}
