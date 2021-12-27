package day15

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay15Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 40

	// When
	lowestTotalRisk := GetLowestTotalRiskPath(input)

	// Then
	if lowestTotalRisk != expected {
		t.Errorf("Day 15 - Expected %v, got %v", expected, lowestTotalRisk)
	}
}

func TestDay15Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 315

	// When
	lowestTotalRisk := GetLowestTotalRiskPathInRealMap(input)

	// Then
	if lowestTotalRisk != expected {
		t.Errorf("Day 15 - Expected %v, got %v", expected, lowestTotalRisk)
	}
}

func TestDay15Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 581

	// When
	lowestTotalRisk := GetLowestTotalRiskPath(input)

	// Then
	if lowestTotalRisk != expected {
		t.Errorf("Day 15 - Expected %v, got %v", expected, lowestTotalRisk)
	}
}

func TestDay15Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 2916

	// When
	lowestTotalRisk := GetLowestTotalRiskPathInRealMap(input)

	// Then
	if lowestTotalRisk != expected {
		t.Errorf("Day 15 - Expected %v, got %v", expected, lowestTotalRisk)
	}
}
