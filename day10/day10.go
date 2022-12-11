package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	detectSignalStrengths("day10-test.txt")
	detectSignalStrengths("day10.txt")
}

func detectSignalStrengths(fileName string) {
	file, _ := os.Open("day10/" + fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	cycle := 1
	x := 1
	strength := map[int]int{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		commands := strings.Fields(line)
		if commands[0] == "addx" {
			cycle++
			checkCycleStrength(&strength, cycle, x)
			printCrt(cycle, x)
			cycle++
			printCrt(cycle, x)
			value, _ := strconv.Atoi(commands[1])
			x += value
			checkCycleStrength(&strength, cycle, x)
		} else {
			cycle++
			printCrt(cycle, x)
			checkCycleStrength(&strength, cycle, x)
		}
	}
	totalStrength := 0
	for _, v := range strength {
		totalStrength += v
	}
	fmt.Println()
	fmt.Println(totalStrength)
	fmt.Println()
	fmt.Println()
}

func printCrt(cycle int, x int) {
	crtPosition := cycle - 2
	if crtPosition != 0 && crtPosition % 40 == 0 {
		fmt.Print("\n")
	}
	for crtPosition > 39 {
		crtPosition -= 40
	}
	diff := crtPosition - x
	if diff >= -1 && diff <= 1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}

func checkCycleStrength(strength *map[int]int, cycle int, x int) {
	if (cycle - 20) % 40 == 0 {
		(*strength)[cycle] = cycle * x
	}
}
