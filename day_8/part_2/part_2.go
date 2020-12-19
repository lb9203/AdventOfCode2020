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
	opChangeMap := map[string]string{
		"nop": "jmp",
		"jmp": "nop",
	}

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

	code = append(code, newInstruction("end", 0))

	for index := range code {
		value, didChangeOp := opChangeMap[code[index].operation]

		if didChangeOp {
			oldValue := code[index].operation
			code[index].operation = value
			if acc, didExecute := canExecute(code); didExecute {
				fmt.Printf("Changed instruction %d, executed with acc=%d\n", index, acc)
			}
			code[index].operation = oldValue
		}

	}
}

func canExecute(code []instruction) (acc int, finishedExecution bool) {
	for i := 0; ; i++ {
		code[i].executions++
		currentInstruction := code[i]
		if currentInstruction.executions > 1 {
			for index := range code {
				code[index].executions = 0
			}
			return acc, false
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
		case "end":
			return acc, true
		}
	}
}
