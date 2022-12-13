package main

import "testing"

func TestParseLineToStruct(t *testing.T) {
	t.Run("givenEmptyList", func(t *testing.T) {
		result := parseLineToStruct("[]")
		assertExpectedOrFail(t, len(result), 0)
	})

	t.Run("givenOneItemInList", func(t *testing.T) {
		result := parseLineToStruct("[1]")
		assertExpectedOrFail(t, len(result), 1)
		assertExpectedOrFail(t, result[0], 1.)
	})

	t.Run("givenMultipleItemsInList", func(t *testing.T) {
		result := parseLineToStruct("[1, 2, 3]")
		assertExpectedOrFail(t, len(result), 3)
		assertExpectedOrFail(t, result[0], 1.)
		assertExpectedOrFail(t, result[1], 2.)
		assertExpectedOrFail(t, result[2], 3.)
	})

	t.Run("givenListInList", func(t *testing.T) {
		result := parseLineToStruct("[[10]]")
		assertExpectedOrFail(t, len(result), 1)
		assertExpectedOrFail(t, len(result[0].([]interface{})), 1)
		assertExpectedOrFail(t, result[0].([]interface{})[0], 10.)
		//as
	})

	t.Run("givenMoreComplexStructure", func(t *testing.T) {
		result := parseLineToStruct("[1,[2,[3,[4,[5,6,7]]]],8,9]")
		assertExpectedOrFail(t, len(result), 4)
		assertExpectedOrFail(t, result[0], 1.)
		assertExpectedOrFail(t, result[2], 8.)
		assertExpectedOrFail(t, result[3], 9.)
		secondEntry := result[1].([]interface{})
		secondSecondEntry := secondEntry[1].([]interface{})
		secondSecondSecondEntry := secondSecondEntry[1].([]interface{})
		assertExpectedOrFail(t, len(secondEntry), 2)
		assertExpectedOrFail(t, secondEntry[0], 2.)
		assertExpectedOrFail(t, len(secondSecondEntry), 2)
		assertExpectedOrFail(t, secondSecondEntry[0], 3.)
		assertExpectedOrFail(t, len(secondSecondSecondEntry), 2)
		assertExpectedOrFail(t, secondSecondSecondEntry[0], 4.)
	})
}

func TestPacketsAreOrdered(t *testing.T) {
	t.Run("givenInvalidIntList", func(t *testing.T) {
		result := packetsAreOrdered(
			parseLineToStruct("[3]"),
			parseLineToStruct("[1]"))
		assertExpectedOrFail(t, result, false)
	})
	t.Run("givenIntLists", func(t *testing.T) {
		result := packetsAreOrdered(
			parseLineToStruct("[1,1,3,1,1]"),
			parseLineToStruct("[1,1,5,1,1]"))
		assertExpectedOrFail(t, result, true)
	})

	t.Run("givenListInList", func(t *testing.T) {
		result := packetsAreOrdered(
			parseLineToStruct("[[1],[2,3,4]]"),
			parseLineToStruct("[[1],4]"))
		assertExpectedOrFail(t, result, true)
	})

	t.Run("givenInvalidOrder", func(t *testing.T) {
		result := packetsAreOrdered(
			parseLineToStruct("[9]"),
			parseLineToStruct("[[8,7,6]]"))
		assertExpectedOrFail(t, result, false)
	})
	t.Run("givenRightSideRanOutOfItems", func(t *testing.T) {
		result := packetsAreOrdered(
			parseLineToStruct("[7, 7, 7, 7]"),
			parseLineToStruct("[7, 7, 7]"))
		assertExpectedOrFail(t, result, false)
	})
	t.Run("givenRightSideRanOutOfItemsWithEmptyLists", func(t *testing.T) {
		result := packetsAreOrdered(
			parseLineToStruct("[[[]]]"),
			parseLineToStruct("[[]]"))
		assertExpectedOrFail(t, result, false)
	})
	t.Run("givenLeftSideRanOutOfItems", func(t *testing.T) {
		result := packetsAreOrdered(
			parseLineToStruct("[]"),
			parseLineToStruct("[3]"))
		assertExpectedOrFail(t, result, true)
	})
	t.Run("complexExample", func(t *testing.T) {
		result := packetsAreOrdered(
			parseLineToStruct("[1,[2,[3,[4,[5,6,7]]]],8,9]"),
			parseLineToStruct("[1,[2,[3,[4,[5,6,0]]]],8,9]"))
		assertExpectedOrFail(t, result, false)

	})
	t.Run("givenSameListsWithIntAfterwards", func(t *testing.T) {
		result := packetsAreOrdered(
			parseLineToStruct("[[1, 2, 3, 4], 2]"),
			parseLineToStruct("[[1, 2, 3, 4], 1]"))
		assertExpectedOrFail(t, result, false)
	})
}

func assertExpectedOrFail(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("FAILED. Expected %s, got %s", expected, actual)
	}
}
