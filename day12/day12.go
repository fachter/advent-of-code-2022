package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	showShortestPath("day12/day12-test.txt")
	showShortestPath("day12/day12.txt")
}

type Node struct {
	row, col int
}

type NodeQueueItem struct {
	node     Node
	distance int
}

func showShortestPath(fileName string) {
	grid := getGrid(fileName)
	startNode := getIndexOf("E", &grid, "z")
	var queue []NodeQueueItem
	queue = append(queue, NodeQueueItem{startNode, 0})
	var visitedNodes []NodeQueueItem
	for len(queue) > 0 {
		nItem := queue[0]
		queue = queue[1:]
		neighbors := getNeighbors(grid, nItem, visitedNodes)
		for _, nb := range neighbors {
			if grid[nb.row][nb.col] == "a" {
				fmt.Println("DONE with distance", nItem.distance+1)
				return
			} else {
				if !inQueue(queue, nb) {
					queue = append(queue, NodeQueueItem{nb, nItem.distance + 1})
				}
			}
		}
		if !inQueue(visitedNodes, nItem.node) {
			visitedNodes = append(visitedNodes, nItem)
		}
	}
}

func getNeighbors(grid [][]string, nItem NodeQueueItem, visited []NodeQueueItem) []Node {
	n := nItem.node
	var neighbors []Node
	possibleNode := []Node{{n.row - 1, n.col}, {n.row + 1, n.col}, {n.row, n.col - 1}, {n.row, n.col + 1}}
	for _, pN := range possibleNode {
		if pN.row < 0 || pN.row >= len(grid) || pN.col < 0 || pN.col >= len(grid[0]) {
			continue
		}
		if inQueue(visited, pN) {
			continue
		}
		if isNotWalkable(grid, n, pN) {
			continue
		}
		neighbors = append(neighbors, pN)

	}
	return neighbors
}

func inQueue(walkedNodes []NodeQueueItem, node Node) bool {
	for _, n := range walkedNodes {
		if node == n.node {
			return true
		}
	}
	return false
}

func isNotWalkable(grid [][]string, n Node, neighbor Node) bool {
	neighborString := grid[neighbor.row][neighbor.col]
	neighborValue := []byte(neighborString)[0]
	nodeString := grid[n.row][n.col]
	nodeValue := []byte(nodeString)[0]
	diff := int(nodeValue) - int(neighborValue)
	return diff > 1
}

func getGrid(fileName string) [][]string {
	file, _ := os.Open(fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var grid [][]string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}
	//fmt.Println(grid)
	return grid
}

func getIndexOf(value string, grid *[][]string, replace string) Node {
	for x := 0; x < len(*grid); x++ {
		for y := 0; y < len((*grid)[x]); y++ {
			if (*grid)[x][y] == value {
				(*grid)[x][y] = replace
				return Node{x, y}
			}
		}
	}
	return Node{-1, -1}
}
