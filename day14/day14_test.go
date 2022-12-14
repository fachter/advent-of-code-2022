package main

import (
	"adventOfCode"
	"testing"
)

func TestParseTextToCave(t *testing.T) {
	t.Run("givenTwoEntries", func(t *testing.T) {
		result := parseTextToCave("498,4 -> 498,5")
		adventOfCode.AssertExpectedOrFail(t, len(result), 2)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{498, 4}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{498, 5}], true)
	})

	t.Run("givenTwoEntriesWithOneInBetween", func(t *testing.T) {
		result := parseTextToCave("498,4 -> 498,6")
		adventOfCode.AssertExpectedOrFail(t, len(result), 3)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{498, 4}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{498, 5}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{498, 6}], true)
	})
	t.Run("givenTwoEntriesWithOneInBetweenOppositeDirection", func(t *testing.T) {
		result := parseTextToCave("498,6 -> 498,4")
		adventOfCode.AssertExpectedOrFail(t, len(result), 3)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{498, 6}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{498, 5}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{498, 4}], true)
	})
	t.Run("givenXEntriesBetween", func(t *testing.T) {
		result := parseTextToCave("496,4 -> 498,4")
		adventOfCode.AssertExpectedOrFail(t, len(result), 3)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{496, 4}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{497, 4}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{498, 4}], true)
	})
	t.Run("givenXEntriesBetween", func(t *testing.T) {
		result := parseTextToCave("498,4 -> 496,4")
		adventOfCode.AssertExpectedOrFail(t, len(result), 3)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{498, 4}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{497, 4}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{496, 4}], true)
	})
	t.Run("givenLargerExample", func(t *testing.T) {
		result := parseTextToCave("503,4 -> 502,4 -> 502,9 -> 494,9")
		adventOfCode.AssertExpectedOrFail(t, len(result), 15)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{503, 4}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{502, 8}], true)
		adventOfCode.AssertExpectedOrFail(t, result[Coordinate{494, 9}], true)
	})
}

func TestAddParsedToMap(t *testing.T) {
	t.Run("givenEmptyMap", func(t *testing.T) {
		emptyMap := map[Coordinate]bool{}

		addParsedToMap(&emptyMap, "498,4 -> 498,6 -> 496,6")

		adventOfCode.AssertExpectedOrFail(t, len(emptyMap), 5)
	})
	t.Run("givenExistingMap", func(t *testing.T) {
		emptyMap := map[Coordinate]bool{}

		addParsedToMap(&emptyMap, "498,4 -> 498,6 -> 496,6")
		addParsedToMap(&emptyMap, "503,4 -> 502,4 -> 502,9 -> 494,9")

		adventOfCode.AssertExpectedOrFail(t, len(emptyMap), 20)
	})
}

func TestSimulateSand(t *testing.T) {
	t.Run("givenNoStones", func(t *testing.T) {
		result := simulateSand(map[Coordinate]bool{})
		adventOfCode.AssertExpectedOrFail(t, len(result), 0)
	})

	t.Run("givenOneStoneBelow", func(t *testing.T) {
		stones := map[Coordinate]bool{
			Coordinate{499, 10}: true,
			Coordinate{500, 10}: true,
			Coordinate{501, 10}: true,
		}

		result := simulateSand(stones)

		adventOfCode.AssertExpectedOrFail(t, len(result), 1)
	})

	t.Run("givenOneStoneBelow", func(t *testing.T) {
		stones := map[Coordinate]bool{
			Coordinate{498, 10}: true,
			Coordinate{499, 10}: true,
			Coordinate{500, 10}: true,
			Coordinate{501, 10}: true,
			Coordinate{502, 10}: true,
		}

		result := simulateSand(stones)

		adventOfCode.AssertExpectedOrFail(t, len(result), 4)
	})
}
