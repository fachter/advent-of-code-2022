package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scoreHiddenVisibleTrees("day08/day08-test.txt")
	scoreHiddenVisibleTrees("day08/day08.txt")
}

func scoreHiddenVisibleTrees(fileName string) {
	file, _ := os.Open(fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var treeHeights [][]int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var row []int
		for _, height := range line {
			heightInt, _ := strconv.Atoi(string(height))
			row = append(row, heightInt)
		}
		treeHeights = append(treeHeights, row)
	}
	currentBestScore := 0
	for i := 1; i < len(treeHeights)-1; i++ {
		for j := 1; j < len(treeHeights[i])-1; j++ {
			score := scoreTree(treeHeights, i, j)
			if score > currentBestScore {
				currentBestScore = score
			}
		}
	}
	fmt.Println(currentBestScore)

}

func scoreTree(heights [][]int, i int, j int) int {
	bottomScore := getBottomScore(heights, i, j)
	leftScore := getLeftScore(heights, i, j)
	rightScore := getRightScore(heights, i, j)
	topScore := getTopScore(heights, i, j)
	return leftScore * rightScore * bottomScore * topScore
}

func getBottomScore(heights [][]int, i int, j int) int {
	treeHeight := heights[i][j]
	for bottomI := i + 1; bottomI < len(heights); bottomI++ {
		if heights[bottomI][j] >= treeHeight {
			return bottomI - i
		}
	}
	return len(heights) - 1 - i
}

func getTopScore(heights [][]int, i int, j int) int {
	treeHeight := heights[i][j]
	for topI := i - 1; topI >= 0; topI-- {
		if heights[topI][j] >= treeHeight {
			return i - topI
		}
	}
	return i
}

func getLeftScore(heights [][]int, i int, j int) int {
	treeHeight := heights[i][j]
	for leftJ := j - 1; leftJ >= 0; leftJ-- {
		if heights[i][leftJ] >= treeHeight {
			return j - leftJ
		}
	}
	return j
}

func getRightScore(heights [][]int, i int, j int) int {
	treeHeight := heights[i][j]
	for rightJ := j + 1; rightJ < len(heights[i]); rightJ++ {
		if heights[i][rightJ] >= treeHeight {
			return rightJ - j
		}
	}
	return len(heights[i]) - 1 - j
}
