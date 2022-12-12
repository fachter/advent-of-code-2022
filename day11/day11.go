package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items []int
	operation string
	secondValue int
	testDivider int
	trueNext int
	falseNext int
	totalInspections int
}


func main() {

	partOne("day11-test.txt")
	partOne("day11.txt")
}

func partOne(fileName string) {
	monkeys := getMonkeys(fileName)
	SimulateRounds(monkeys)
	printScore(monkeys)
}

func printScore(monkeys map[int]*Monkey) {
	for _, m := range monkeys {
		fmt.Println(m.totalInspections)
	}
	fmt.Println()
}

func getMonkeys(fileName string) map[int]*Monkey {
	file, _ := os.Open("day11/" + fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	monkeys := map[int]*Monkey{}
	currentIndex := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, "Monkey") {
			indexWithColon := strings.Fields(line)[1]
			index, _ := strconv.Atoi(indexWithColon[:len(indexWithColon)-1])
			currentIndex = index
			monkeys[currentIndex] = &Monkey{}
		} else if strings.Contains(line, "Starting items") {
			startingItems := strings.Split(strings.Split(line, ": ")[1], ", ")
			for _, item := range startingItems {
				score, _ := strconv.Atoi(item)
				monkeys[currentIndex].items = append(monkeys[currentIndex].items, score)
			}
		} else if strings.Contains(line, "Operation") {
			monkeys[currentIndex].operation = strings.Split(line, "= ")[1]
		} else if strings.Contains(line, "Test") {
			monkeys[currentIndex].testDivider = getLastFieldAsInt(line)
		} else if strings.Contains(line, "If true") {
			monkeys[currentIndex].trueNext = getLastFieldAsInt(line)
		} else if strings.Contains(line, "If false") {
			monkeys[currentIndex].falseNext = getLastFieldAsInt(line)
		}
	}
	return monkeys
}

func SimulateRounds(monkeys map[int]*Monkey) {
	mod := 1
	for _, m := range monkeys {
		mod *= m.testDivider
	}
	for i := 0; i < 10_000; i++ {
		for index := 0; index < len(monkeys); index++ {
			monkeyItems := monkeys[index].items
			for _, item := range monkeyItems {
				monkeys[index].totalInspections++
				newValue := calc(monkeys[index].operation, item)
				newValue %= mod
				//newValue /= 3
				if newValue%monkeys[index].testDivider == 0 {
					monkeys[monkeys[index].trueNext].items = append(monkeys[monkeys[index].trueNext].items, newValue)
				} else {
					monkeys[monkeys[index].falseNext].items = append(monkeys[monkeys[index].falseNext].items, newValue)
				}
				monkeys[index].items = monkeys[index].items[1:]
			}
		}
		if i == 0 || i == 999 || i == 19 || i == 4999 {
			fmt.Println(i + 1)
			for k, m := range monkeys {
				fmt.Println(k, m.totalInspections)
			}
			fmt.Println()
		}
	}
}

func getLastFieldAsInt(line string) int {
	fields := strings.Fields(line)
	lastFieldInt, _ := strconv.Atoi(fields[len(fields)-1])
	return lastFieldInt
}

func calc(val string, x int) int {
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	fields := strings.Fields(val)
	val1 := x
	val2 := x
	if fields[0] != "old" {
		val1, _ = strconv.Atoi(fields[0])
	}
	if fields[2] != "old" {
		val2, _ = strconv.Atoi(fields[2])
	}
	return ops[fields[1]](val1, val2)
}



