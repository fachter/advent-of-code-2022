package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//readDuplicatesForOneElf()
	readDuplicatesForThreeElves()
}

func readDuplicatesForThreeElves() {
	file, _ := os.Open("day03.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	prioritySum := 0
	count := 0
	rucksacks := [3][]byte{}
	for fileScanner.Scan() {
		rucksack := fileScanner.Text()
		rucksacks[count] = []byte(rucksack)
		count ++
		if count == 3 {

			duplicates := map[int]bool{}
			for _, f := range rucksacks[0] {
				for _, s := range rucksacks[1] {
					for _, t := range rucksacks[2] {
						if f == s && s == t {
							duplicates[int(f)] = true
						}
					}
				}
			}
			priority := 0
			for k := range duplicates {
				if k >= 97 {
					priority += k - 96
				} else {
					priority += k - 38
				}
			}
			prioritySum += priority
			rucksacks[0] = nil
			rucksacks[1] = nil
			rucksacks[2] = nil
			count = 0
		}
	}
	fmt.Println(prioritySum)
}

func readDuplicatesForOneElf() {
	file, _ := os.Open("day03.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	prioritySum := 0
	for fileScanner.Scan() {
		rucksack := fileScanner.Text()
		firstHalf := []byte(rucksack[:(len(rucksack) / 2)])
		secondHalf := []byte(rucksack[(len(rucksack) / 2):])
		duplicates := map[int]bool{}
		for _, f := range firstHalf {
			for _, s := range secondHalf {
				if f == s {
					duplicates[int(f)] = true
				}
			}
		}
		priority := 0
		for k, _ := range duplicates {
			if k >= 97 {
				priority += k - 96
			} else {
				priority += k - 38
			}
		}
		prioritySum += priority
	}
	fmt.Println(prioritySum)
}
