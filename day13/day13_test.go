package main

import (
	"adventOfCode"
	"testing"
)

func TestParseLineToStruct(t *testing.T) {
	t.Run("givenEmptyList", func(t *testing.T) {
		result := parseLineToStruct("[]")
		adventOfCode.AssertExpectedOrFail(t, len(result), 0)
	})

	t.Run("givenOneItemInList", func(t *testing.T) {
		result := parseLineToStruct("[1]")
		adventOfCode.AssertExpectedOrFail(t, len(result), 1)
		adventOfCode.AssertExpectedOrFail(t, result[0], 1.)
	})

	t.Run("givenMultipleItemsInList", func(t *testing.T) {
		result := parseLineToStruct("[1, 2, 3]")
		adventOfCode.AssertExpectedOrFail(t, len(result), 3)
		adventOfCode.AssertExpectedOrFail(t, result[0], 1.)
		adventOfCode.AssertExpectedOrFail(t, result[1], 2.)
		adventOfCode.AssertExpectedOrFail(t, result[2], 3.)
	})

	t.Run("givenListInList", func(t *testing.T) {
		result := parseLineToStruct("[[10]]")
		adventOfCode.AssertExpectedOrFail(t, len(result), 1)
		adventOfCode.AssertExpectedOrFail(t, len(result[0].([]interface{})), 1)
		adventOfCode.AssertExpectedOrFail(t, result[0].([]interface{})[0], 10.)
		//as
	})

	t.Run("givenMoreComplexStructure", func(t *testing.T) {
		result := parseLineToStruct("[1,[2,[3,[4,[5,6,7]]]],8,9]")
		adventOfCode.AssertExpectedOrFail(t, len(result), 4)
		adventOfCode.AssertExpectedOrFail(t, result[0], 1.)
		adventOfCode.AssertExpectedOrFail(t, result[2], 8.)
		adventOfCode.AssertExpectedOrFail(t, result[3], 9.)
		secondEntry := result[1].([]interface{})
		secondSecondEntry := secondEntry[1].([]interface{})
		secondSecondSecondEntry := secondSecondEntry[1].([]interface{})
		adventOfCode.AssertExpectedOrFail(t, len(secondEntry), 2)
		adventOfCode.AssertExpectedOrFail(t, secondEntry[0], 2.)
		adventOfCode.AssertExpectedOrFail(t, len(secondSecondEntry), 2)
		adventOfCode.AssertExpectedOrFail(t, secondSecondEntry[0], 3.)
		adventOfCode.AssertExpectedOrFail(t, len(secondSecondSecondEntry), 2)
		adventOfCode.AssertExpectedOrFail(t, secondSecondSecondEntry[0], 4.)
	})
}

func TestPacketsAreOrdered(t *testing.T) {
	t.Run("givenInvalidIntList", func(t *testing.T) {
		result, _ := packetsAreOrdered(
			parseLineToStruct("[3]"),
			parseLineToStruct("[1]"))
		adventOfCode.AssertExpectedOrFail(t, result, false)
	})
	t.Run("givenIntLists", func(t *testing.T) {
		result, _ := packetsAreOrdered(
			parseLineToStruct("[1,1,3,1,1]"),
			parseLineToStruct("[1,1,5,1,1]"))
		adventOfCode.AssertExpectedOrFail(t, result, true)
	})

	t.Run("givenListInList", func(t *testing.T) {
		result, _ := packetsAreOrdered(
			parseLineToStruct("[[1],[2,3,4]]"),
			parseLineToStruct("[[1],4]"))
		adventOfCode.AssertExpectedOrFail(t, result, true)
	})

	t.Run("givenInvalidOrder", func(t *testing.T) {
		result, _ := packetsAreOrdered(
			parseLineToStruct("[9]"),
			parseLineToStruct("[[8,7,6]]"))
		adventOfCode.AssertExpectedOrFail(t, result, false)
	})
	t.Run("givenRightSideRanOutOfItems", func(t *testing.T) {
		result, _ := packetsAreOrdered(
			parseLineToStruct("[7, 7, 7, 7]"),
			parseLineToStruct("[7, 7, 7]"))
		adventOfCode.AssertExpectedOrFail(t, result, false)
	})
	t.Run("givenRightSideRanOutOfItemsWithEmptyLists", func(t *testing.T) {
		result, _ := packetsAreOrdered(
			parseLineToStruct("[[[]]]"),
			parseLineToStruct("[[]]"))
		adventOfCode.AssertExpectedOrFail(t, result, false)
	})
	t.Run("givenLeftSideRanOutOfItems", func(t *testing.T) {
		result, _ := packetsAreOrdered(
			parseLineToStruct("[]"),
			parseLineToStruct("[3]"))
		adventOfCode.AssertExpectedOrFail(t, result, true)
	})
	t.Run("givenLeftSideRanOutOfItemsWithItemsAfterwards", func(t *testing.T) {
		result, _ := packetsAreOrdered(
			parseLineToStruct("[[], 2]"),
			parseLineToStruct("[[3], 1]"))
		adventOfCode.AssertExpectedOrFail(t, result, true)
	})
	t.Run("complexExample", func(t *testing.T) {
		result, _ := packetsAreOrdered(
			parseLineToStruct("[1,[2,[3,[4,[5,6,7]]]],8,9]"),
			parseLineToStruct("[1,[2,[3,[4,[5,6,0]]]],8,9]"))
		adventOfCode.AssertExpectedOrFail(t, result, false)

	})
	t.Run("givenSameListsWithIntAfterwards", func(t *testing.T) {
		result, _ := packetsAreOrdered(
			parseLineToStruct("[[1, 2, 3, 4], 2]"),
			parseLineToStruct("[[1, 2, 3, 4], 1]"))
		adventOfCode.AssertExpectedOrFail(t, result, false)
	})
	t.Run("givenSameLists", func(t *testing.T) {
		_, same := packetsAreOrdered(
			parseLineToStruct("[[1, 2, 3, 4], 2]"),
			parseLineToStruct("[[1, 2, 3, 4], 2]"))
		adventOfCode.AssertExpectedOrFail(t, same, true)
	})
}
