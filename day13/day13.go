package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//sumRightOrderedIndices("day13-test.txt")
	//sumRightOrderedIndices("day13.txt")

	sortPackages("day13-test.txt")
	sortPackages("day13.txt")
}

func sortPackages(fileName string) {
	file, _ := os.Open("day13/" + fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var packages [][]interface{}
	divider2 := parseLineToStruct("[[2]]")
	divider6 := parseLineToStruct("[[6]]")
	packages = append(packages, divider2, divider6)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			packages = append(packages, parseLineToStruct(line))
		}
	}
	for i := len(packages); i > 0; i-- {
		for j := 1; j < i; j++ {
			ordered, same := packetsAreOrdered(packages[j-1], packages[j])
			if !ordered && !same {
				temp := packages[j]
				packages[j] = packages[j-1]
				packages[j-1] = temp
			}
		}
	}
	s2Idx := 0
	s6Idx := 0
	for i, pack := range packages {
		_, same2 := packetsAreOrdered(pack, divider2)
		if same2 {
			s2Idx = i + 1
			fmt.Println("Same2 Index", i)
		}
		_, same6 := packetsAreOrdered(pack, divider6)
		if same6 {
			s6Idx = i + 1
			fmt.Println("Same6 Index", i)
		}
	}
	fmt.Println(s2Idx * s6Idx)
}

func sumRightOrderedIndices(fileName string) {
	file, _ := os.Open("day13/" + fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var firstItem []interface{}
	var secondItem []interface{}
	writeToFirst := true
	var validPackets []int
	packetIdx := 1
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			continue
		}
		if writeToFirst {
			firstItem = parseLineToStruct(line)
			writeToFirst = false
		} else {
			secondItem = parseLineToStruct(line)
			ordered, same := packetsAreOrdered(firstItem, secondItem)
			if ordered || same {
				validPackets = append(validPackets, packetIdx)
			}
			writeToFirst = true
			packetIdx++
		}
	}
	sum := 0
	for _, packet := range validPackets {
		sum += packet
	}
	fmt.Println(sum)

}

func packetsAreOrdered(firstItem []interface{}, secondItem []interface{}) (bool, bool) {
	firstLength := len(firstItem)
	secondLength := len(secondItem)
	for i := 0; i < firstLength; i++ {
		if i >= secondLength {
			return false, false
		}
		firstListInList, firstListOk := firstItem[i].([]interface{})
		secondListInList, secondListOk := secondItem[i].([]interface{})
		if firstListOk == secondListOk && firstListOk {
			ordered, same := packetsAreOrdered(firstListInList, secondListInList)
			if !same {
				return ordered, false
			}
		} else if firstListOk == secondListOk {
			same := secondItem[i].(float64) == firstItem[i].(float64)
			if !same {

				return !(secondItem[i].(float64) < firstItem[i].(float64)), same
			}
		} else {
			if !secondListOk {
				var newSecondListInList []interface{}
				newSecondListInList = append(newSecondListInList, secondItem[i])
				ordered, same := packetsAreOrdered(firstListInList, newSecondListInList)
				if !same {
					return ordered, false
				}
			} else if !firstListOk {
				var newFirstListInList []interface{}
				newFirstListInList = append(newFirstListInList, firstItem[i])
				ordered, same := packetsAreOrdered(newFirstListInList, secondListInList)
				if !same {
					return ordered, false
				}
			}
		}
	}
	if firstLength < secondLength {
		return true, false
	}
	return true, true
}

func parseLineToStruct(line string) []interface{} {
	var parts []interface{}
	err := json.Unmarshal([]byte(line), &parts)
	if err != nil {
		return nil
	}
	return parts
}

// 751 to low
// 3798 to low
