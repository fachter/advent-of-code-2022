package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	findHumanScream("day21/21-test.txt")
	findHumanScream("day21/21.txt")
}

func findHumanScream(fileName string) {
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
		if monkeyName == "humn" {

		} else if len(fields) == 1 {
			intValue, _ := strconv.Atoi(fields[0])
			values[monkeyName] = intValue
		} else {
			operations[monkeyName] = Operation{firstValue: fields[0], secondValue: fields[2], operation: fields[1]}
		}
	}

	solveForHuman(operations, values)
}
