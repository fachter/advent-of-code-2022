package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, _ := os.Open("day02/day02.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	totalScore := 0
	for fileScanner.Scan() {
		score := 0
		signs := strings.Fields(fileScanner.Text())
		firstPlayer := int([]byte(signs[0])[0])
		secondPlayer := int([]byte(signs[1])[0]) - 23
		diff := firstPlayer - secondPlayer
		score += secondPlayer - 64
		if diff == 0 {
			score += 3
		} else if diff == 2 || diff == -1 {
			score += 6
		}
		//fmt.Println(signs, firstPlayer, secondPlayer, score)
		totalScore += score
	}
	fmt.Println("TotalScore", totalScore)
	file.Close()
}

func partTwo() {
	file, _ := os.Open("day02/day02.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	totalScore := 0
	for fileScanner.Scan() {
		score := 0
		signs := strings.Fields(fileScanner.Text())
		firstPlayer := int([]byte(signs[0])[0])
		secondPlayer := getSecondPlayer(signs, firstPlayer)
		diff := firstPlayer - secondPlayer
		score += secondPlayer - 64
		if diff == 0 {
			score += 3
		} else if diff == 2 || diff == -1 {
			score += 6
		}
		fmt.Println(signs, firstPlayer, secondPlayer, score)
		totalScore += score
	}
	fmt.Println("TotalScore", totalScore)
	file.Close()
}

func getSecondPlayer(signs []string, firstPlayer int) int {
	var toPlay int
	if signs[1] == "Y" {
		toPlay = firstPlayer
	} else if signs[1] == "X" {
		toPlay = firstPlayer - 1
		if toPlay < 65 {
			toPlay = firstPlayer + 2
		}
	} else {
		toPlay = firstPlayer - 2
		if toPlay < 65 {
			toPlay = firstPlayer + 1
		}
	}
	return toPlay
}
