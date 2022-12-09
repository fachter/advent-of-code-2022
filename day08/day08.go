package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	countHiddenVisibleTrees("day08/day08-test.txt")
	countHiddenVisibleTrees("day08/day08.txt")
}

func countHiddenVisibleTrees(fileName string) {
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
	notVisible := 0
	for i := 1; i < len(treeHeights)-1; i++ {
		for j := 1; j < len(treeHeights[i])-1; j++ {
			if treeIsNotVisible(treeHeights, i, j) {
				notVisible++
			}
		}
	}
	totalTrees := len(treeHeights) * len(treeHeights[0])
	fmt.Println(totalTrees - notVisible)

}

func treeIsNotVisible(heights [][]int, i int, j int) bool {
	treeHeight := heights[i][j]
	visibleTop := isVisibleTop(heights, i, j, treeHeight)
	visibleBottom := isVisibleBottom(heights, i, j, treeHeight)
	visibleLeft := isVisibleLeft(heights, i, j, treeHeight)
	visibleRight := getVisibleRight(heights, i, j, treeHeight)
	return !visibleTop && !visibleBottom && !visibleLeft && !visibleRight
}

func isVisibleBottom(heights [][]int, i int, j int, treeHeight int) bool {
	for bottomI := i + 1; bottomI < len(heights); bottomI++ {
		if heights[bottomI][j] >= treeHeight {
			return false
		}
	}
	return true
}

func isVisibleTop(heights [][]int, i int, j int, treeHeight int) bool {
	for topI := i - 1; topI >= 0; topI-- {
		if heights[topI][j] >= treeHeight {
			return false
		}
	}
	return true
}

func isVisibleLeft(heights [][]int, i int, j int, treeHeight int) bool {
	for leftJ := j - 1; leftJ >= 0; leftJ-- {
		if heights[i][leftJ] >= treeHeight {
			return false
		}
	}
	return true
}

func getVisibleRight(heights [][]int, i int, j int, treeHeight int) bool {
	for rightJ := j + 1; rightJ < len(heights[i]); rightJ++ {
		if heights[i][rightJ] >= treeHeight {
			return false
		}
	}
	return true
}
