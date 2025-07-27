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

	solveForHuman(operations, values)
}

func solveForHuman(operations map[string]Operation, values map[string]int) {
	firstValue, humanFirst := solveByName(operations, values, operations["root"].firstValue)
	secondValue, humanSecond := solveByName(operations, values, operations["root"].secondValue)
	fmt.Println(firstValue, humanFirst)
	fmt.Println(secondValue, humanSecond)
	if humanFirst {

	}

}

func findHumanValue(operations map[string]Operation, values map[string]int) int {

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

func oppositeOperation(op string) string {
	if op == "+" {
		return "-"
	}
	if op == "-" {
		return "+"
	}
	if op == "/" {
		return "*"
	}
	return "/"
}

func solveByName(operations map[string]Operation, values map[string]int, name string) (int, bool) {
	isHuman := name == "humn"
	value, exists := values[name]
	if exists {
		return value, isHuman
	}
	operation, _ := operations[name]
	first, humanFirst := solveByName(operations, values, operation.firstValue)
	second, humanSecond := solveByName(operations, values, operation.secondValue)
	return performOperation(
		first,
		second,
		operation.operation), isHuman || humanFirst || humanSecond
}
