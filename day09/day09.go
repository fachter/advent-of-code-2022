package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func main() {
	//countVisitedPlaces("day09-test.txt")
	countVisitedPlaces("day09.txt")
}

func countVisitedPlaces(fileName string) {
	file, _ := os.Open("day09/" + fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	headPosition := Position{0, 0}
	tailPosition := Position{0, 0}
	visitedPlaces := map[Position]bool{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fields := strings.Fields(line)
		direction := fields[0]
		steps, _ := strconv.Atoi(fields[1])
		for i := 0; i < steps; i++ {
			headStep(direction, &headPosition)
			tailStep(headPosition, &tailPosition)
			visitedPlaces[tailPosition] = true
			//fmt.Println(headPosition)
			//fmt.Println(tailPosition)
			//fmt.Println()
		}

	}
	fmt.Println(len(visitedPlaces))
}

func tailStep(headPosition Position, tailPosition *Position) {
	xDiff := headPosition.x - tailPosition.x
	yDiff := headPosition.y - tailPosition.y
	if absValue(xDiff) > 1 || absValue(yDiff) > 1 {
		xStep := 0
		if xDiff > 0 {
			xStep = 1
		} else if xDiff < 0 {
			xStep = -1
		}
		yStep := 0
		if yDiff > 0 {
			yStep = 1
		} else if yDiff < 0 {
			yStep = -1
		}
		tailPosition.x += xStep
		tailPosition.y += yStep

	}
}

func absValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func headStep(direction string, headPosition *Position) {
	switch direction {
	case "R":
		headPosition.x++
		break
	case "L":
		headPosition.x--
		break
	case "U":
		headPosition.y++
		break
	case "D":
		headPosition.y--
	default:
		fmt.Println("Should not happen!")
		break
	}
}

// 5127 to low
