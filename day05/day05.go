package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	moveAllItemsAtOnce()
}

func moveItemsOneByOne() {
	file, _ := os.Open("day05.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	items := map[int][]string{
		0: {},
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == " 1   2   3   4   5   6   7   8   9" {
			fileScanner.Scan()
			break
		}
		for i := 0; i < 9; i++ {
			valueAtI := getValueAtI(line, i)
			if valueAtI != " " {
				items[i] = append(items[i], valueAtI)
			}
		}
	}
	for fileScanner.Scan() {
		commands := strings.Fields(fileScanner.Text())
		n, _ := strconv.Atoi(commands[1])
		from, _ := strconv.Atoi(commands[3])
		from--
		to, _ := strconv.Atoi(commands[5])
		to--
		for i := 0; i < n; i++ {
			last := items[from][0]
			items[to] = append([]string{last}, items[to]...)
			items[from] = items[from][1:]
		}
	}
	codeWord := ""
	for i := 0; i < 9; i++ {
		codeWord += items[i][0]
	}
	print(codeWord)
}

func moveAllItemsAtOnce() {
	file, _ := os.Open("day05.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	items := map[int][]string{
		0: {},
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == " 1   2   3   4   5   6   7   8   9" {
			fileScanner.Scan()
			break
		}
		for i := 0; i < 9; i++ {
			valueAtI := getValueAtI(line, i)
			if valueAtI != " " {
				items[i] = append(items[i], valueAtI)
			}
		}
	}
	for fileScanner.Scan() {
		commands := strings.Fields(fileScanner.Text())
		n, _ := strconv.Atoi(commands[1])
		from, _ := strconv.Atoi(commands[3])
		from--
		to, _ := strconv.Atoi(commands[5])
		to--
		fmt.Println(items[to])
		fmt.Println(n)
		itemsToMove := make([]string, len(items[from][:n]))
		copy(itemsToMove, items[from][:n])
		fmt.Println(itemsToMove)
		newTo := append(itemsToMove, items[to]...)
		items[to] = newTo
		fmt.Println(items[to])
		items[from] = items[from][n:]
		fmt.Println()
	}
	codeWord := ""
	for i := 0; i < 9; i++ {
		codeWord += items[i][0]
	}
	print(codeWord)
}

func getValueAtI(line string, i int) string {
	return string(line[i*4+1])
}
