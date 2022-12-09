package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	overlapAtAll()
}

func overlapCompletely() {
	file, _ := os.Open("day04/day04.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	counter := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		tuple := strings.Split(line, ",")
		firstElf := strings.Split(tuple[0], "-")
		secondElf := strings.Split(tuple[1], "-")
		firstElfStart, _ := strconv.Atoi(firstElf[0])
		firstElfEnd, _ := strconv.Atoi(firstElf[1])
		secondElfStart, _ := strconv.Atoi(secondElf[0])
		secondElfEnd, _ := strconv.Atoi(secondElf[1])
		firstElfStartsFirst := firstElfStart <= secondElfStart
		firstElfEndsLast := firstElfEnd >= secondElfEnd
		secondElfStartsFirst := secondElfStart <= firstElfStart
		secondElfEndsLast := secondElfEnd >= firstElfEnd
		if firstElfStartsFirst == firstElfEndsLast || secondElfStartsFirst == secondElfEndsLast {
			counter++
		}
	}
	fmt.Println(counter)
}

func overlapAtAll() {
	file, _ := os.Open("day04/day04.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	counter := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		tuple := strings.Split(line, ",")
		firstElf := strings.Split(tuple[0], "-")
		secondElf := strings.Split(tuple[1], "-")
		firstElfStart, _ := strconv.Atoi(firstElf[0])
		firstElfEnd, _ := strconv.Atoi(firstElf[1])
		secondElfStart, _ := strconv.Atoi(secondElf[0])
		secondElfEnd, _ := strconv.Atoi(secondElf[1])
		secondStartInBetween := secondElfStart >= firstElfStart && secondElfStart <= firstElfEnd
		secondEndInBetween := secondElfEnd >= firstElfStart && secondElfEnd <= firstElfEnd
		secondElfAround := secondElfStart <= firstElfStart && secondElfEnd >= firstElfEnd
		if secondStartInBetween || secondEndInBetween || secondElfAround {
			counter++
		}
	}
	fmt.Println(counter)
}
