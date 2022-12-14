package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	discoverCave("day14-test.txt")
	discoverCave("day14.txt")
}

func discoverCave(fileName string) {
	file, _ := os.Open("day14/" + fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	coordinates := map[Coordinate]bool{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		addParsedToMap(&coordinates, line)
	}
	fmt.Println("built")
	//printCoordinates(coordinates)
	sand := simulateSandWithGround(coordinates)
	fmt.Println(len(sand))
}

func simulateSand(coordinates map[Coordinate]bool) map[Coordinate]bool {
	sand := map[Coordinate]bool{}
	fallCounter := 0
	newSand := Coordinate{500, 0}
	for fallCounter < 500 {
		_, stoneBelow := coordinates[Coordinate{newSand.x, newSand.y + 1}]
		_, sandBelow := sand[Coordinate{newSand.x, newSand.y + 1}]
		_, stoneLeft := coordinates[Coordinate{newSand.x - 1, newSand.y + 1}]
		_, sandLeft := sand[Coordinate{newSand.x - 1, newSand.y + 1}]
		_, stoneRight := coordinates[Coordinate{newSand.x + 1, newSand.y + 1}]
		_, sandRight := sand[Coordinate{newSand.x + 1, newSand.y + 1}]
		if !stoneBelow && !sandBelow {
			newSand.y++
			fallCounter++
			continue
		}
		if !stoneLeft && !sandLeft {
			newSand.x--
			newSand.y++
			fallCounter++
			continue
		}
		if !stoneRight && !sandRight {
			newSand.x++
			newSand.y++
			fallCounter++
			continue
		}
		sand[newSand] = true
		newSand = Coordinate{500, 0}
		fallCounter = 0
	}
	return sand
}

func simulateSandWithGround(coordinates map[Coordinate]bool) map[Coordinate]bool {
	sand := map[Coordinate]bool{}
	ground := 0
	for stone := range coordinates {
		if stone.y+2 > ground {
			ground = stone.y + 2
		}
	}
	for x := -5; x < 1000; x++ {
		coordinates[Coordinate{x, ground}] = true
	}
	newSand := Coordinate{500, 0}
	for _, sandInMap := sand[Coordinate{500, 0}]; !sandInMap; {
		_, sandInMap := sand[Coordinate{500, 0}]
		if sandInMap {
			break
		}
		_, stoneBelow := coordinates[Coordinate{newSand.x, newSand.y + 1}]
		_, sandBelow := sand[Coordinate{newSand.x, newSand.y + 1}]
		_, stoneLeft := coordinates[Coordinate{newSand.x - 1, newSand.y + 1}]
		_, sandLeft := sand[Coordinate{newSand.x - 1, newSand.y + 1}]
		_, stoneRight := coordinates[Coordinate{newSand.x + 1, newSand.y + 1}]
		_, sandRight := sand[Coordinate{newSand.x + 1, newSand.y + 1}]
		if !stoneBelow && !sandBelow {
			newSand.y++
			continue
		}
		if !stoneLeft && !sandLeft {
			newSand.x--
			newSand.y++
			continue
		}
		if !stoneRight && !sandRight {
			newSand.x++
			newSand.y++
			continue
		}
		sand[newSand] = true
		newSand = Coordinate{500, 0}
	}
	return sand
}

func printCoordinates(coordinates map[Coordinate]bool) {
	for y := 0; y < 50; y++ {
		for x := 400; x < 550; x++ {
			value := "."
			_, inMap := coordinates[Coordinate{x, y}]
			if x == 500 && y == 0 {
				value = "+"
			} else if inMap {
				value = "#"
			}
			fmt.Print(value)
		}
		fmt.Print("\n")
	}
}

type Coordinate struct {
	x, y int
}

func addParsedToMap(coordinateMap *map[Coordinate]bool, line string) {
	lineMap := parseTextToCave(line)
	for key := range lineMap {
		(*coordinateMap)[key] = true
	}
}

func parseTextToCave(s string) map[Coordinate]bool {
	stones := strings.Split(s, " -> ")
	var coordinates []Coordinate
	for _, stone := range stones {
		stoneCoordinates := strings.Split(stone, ",")
		x, _ := strconv.Atoi(stoneCoordinates[0])
		y, _ := strconv.Atoi(stoneCoordinates[1])
		if len(coordinates) > 0 {
			lastCoordinate := coordinates[len(coordinates)-1]
			addXDirection(x, lastCoordinate, &coordinates)
			addYDirection(y, lastCoordinate, &coordinates)
		}
		coordinates = append(coordinates, Coordinate{x, y})
	}
	coordinateMap := map[Coordinate]bool{}
	for _, co := range coordinates {
		coordinateMap[co] = true
	}
	return coordinateMap
}

func addXDirection(x int, lastCoordinate Coordinate, coordinates *[]Coordinate) {
	xDiff, xMultiplier := getDiffAndMultiplier(x, lastCoordinate.x)
	for i := 1; i < xDiff; i++ {
		*coordinates = append(*coordinates, Coordinate{lastCoordinate.x + i*xMultiplier, lastCoordinate.y})
	}
}

func getDiffAndMultiplier(x, lastCoordinateX int) (int, int) {
	diff := x - lastCoordinateX
	multiplier := 1
	if diff < 0 {
		diff *= -1
		multiplier = -1
	}
	return diff, multiplier
}

func addYDirection(y int, lastCoordinate Coordinate, coordinates *[]Coordinate) {
	yDiff, yMultiplier := getDiffAndMultiplier(y, lastCoordinate.y)
	for j := 1; j < yDiff; j++ {
		*coordinates = append(*coordinates, Coordinate{lastCoordinate.x,
			lastCoordinate.y + (j * yMultiplier)})
	}
}
