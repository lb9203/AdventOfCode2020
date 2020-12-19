package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	operation  string
	argument   int
	executions int
}

func newInstruction(operation string, argument int) instruction {
	return instruction{
		operation:  operation,
		argument:   argument,
		executions: 0,
	}
}

func main() {
	code := []instruction{}

	var file, err = os.Open("day_8/code.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		op := s[0]
		arg, _ := strconv.Atoi(s[1])

		code = append(code, newInstruction(op, arg))
	}

	acc := 0
	for i := 0; i < len(code); i++ {
		code[i].executions++
		currentInstruction := code[i]
		if currentInstruction.executions > 1 {
			break
		}

		switch currentInstruction.operation {
		case "acc":
			acc += currentInstruction.argument
			break
		case "jmp":
			i += currentInstruction.argument - 1
			break
		case "nop":
			break
		}
	}

	fmt.Printf("Acc: %d\n", acc)
}
