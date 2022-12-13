package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//sumRightOrderedIndices("day13-test.txt")
	sumRightOrderedIndices("day13.txt")
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
			ordered := packetsAreOrdered(firstItem, secondItem)
			if ordered {
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
	fmt.Println(validPackets)
	fmt.Println(sum)

}

func packetsAreOrdered(firstItem []interface{}, secondItem []interface{}) bool {
	firstLength := len(firstItem)
	secondLength := len(secondItem)
	for i := 0; i < firstLength; i++ {
		if i >= secondLength {
			return false
		}
		firstListInList, firstListOk := firstItem[i].([]interface{})
		secondListInList, secondListOk := secondItem[i].([]interface{})
		if firstListOk == secondListOk && firstListOk {
			if !packetsAreOrdered(firstListInList, secondListInList) {
				return false
			}
		} else if firstListOk == secondListOk {
			if secondItem[i].(float64) != firstItem[i].(float64) {
				return !(secondItem[i].(float64) < firstItem[i].(float64))
			}
		} else {
			if !secondListOk {
				var newSecondListInList []interface{}
				newSecondListInList = append(newSecondListInList, secondItem[i])
				if !packetsAreOrdered(firstListInList, newSecondListInList) {
					return false
				}
			} else if !firstListOk {
				var newFirstListInList []interface{}
				newFirstListInList = append(newFirstListInList, firstItem[i])
				if !packetsAreOrdered(newFirstListInList, secondListInList) {
					return false
				}
			}
		}
	}
	return true
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
