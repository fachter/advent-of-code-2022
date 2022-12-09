package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	checkForFirstStartMarker("day06-test1.txt")
	checkForFirstStartMarker("day06-test2.txt")
	checkForFirstStartMarker("day06-test3.txt")
	checkForFirstStartMarker("day06-test4.txt")
	checkForFirstStartMarker("day06.txt")
}

func checkForFirstStartMarker(fileName string) {
	file, _ := os.Open(fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanBytes)
	fmt.Println(fileName)
	fmt.Println(getFirstMarker(fileScanner))
	fmt.Println()
}

func getFirstMarker(fileScanner *bufio.Scanner) int {
	var lastWords []string
	counter := 0
	for fileScanner.Scan() {
		counter++
		token := fileScanner.Text()
		lastWords = append(lastWords, token)
		if len(lastWords) > 13 {
			if allItemsAreDifferent(lastWords) {
				return counter
			}
			lastWords = lastWords[1:]
		}
	}
	return -1
}

func allItemsAreDifferent(words []string) bool {
	existingWords := map[string]bool{}
	for _, word := range words {
		_, exists := existingWords[word]
		if exists {
			return false
		}
		existingWords[word] = true
	}
	return true
}
