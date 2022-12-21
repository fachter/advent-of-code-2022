package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	findRootMonkeyScream("day21/21-test.txt")
	findRootMonkeyScream("day21/21.txt")
}

type Operation struct {
	firstValue, secondValue string
	operation               string
}

func findRootMonkeyScream(fileName string) {
	file, _ := os.Open(fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	values := map[string]int{}
	operations := map[string]Operation{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		monkey := strings.Split(line, ": ")
		monkeyName := monkey[0]
		fields := strings.Fields(monkey[1])
		if len(fields) == 1 {
			intValue, _ := strconv.Atoi(fields[0])
			values[monkeyName] = intValue
		} else {
			operations[monkeyName] = Operation{firstValue: fields[0], secondValue: fields[2], operation: fields[1]}
		}
	}

	rootResult := solveByName(operations, values, "root")
	fmt.Println(rootResult)
}

func performOperation(val1, val2 int, op string) int {
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	return ops[op](val1, val2)
}

func solveByName(operations map[string]Operation, values map[string]int, name string) int {
	value, exists := values[name]
	if exists {
		return value
	}
	operation, _ := operations[name]
	return performOperation(
		solveByName(operations, values, operation.firstValue),
		solveByName(operations, values, operation.secondValue),
		operation.operation)
}
