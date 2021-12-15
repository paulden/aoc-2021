package main

import (
	"testing"
)

func TestDay15Part1Example(t *testing.T) {
	// Given
	day15 := readStringsInFile("../data/day15_example.txt")
	expected := 40

	// When
	lowestTotalRisk := GetLowestTotalRiskPath(day15)

	// Then
	if lowestTotalRisk != expected {
		t.Errorf("Day 15 example - Part 1 example: expected %v, got %v", expected, lowestTotalRisk)
	}
}

func TestDay15Part1Real(t *testing.T) {
	// Given
	day15 := readStringsInFile("../data/day15.txt")
	expected := 581

	// When
	lowestTotalRisk := GetLowestTotalRiskPath(day15)

	// Then
	if lowestTotalRisk != expected {
		t.Errorf("Day 15 example - Part 1 real: expected %v, got %v", expected, lowestTotalRisk)
	}
}

func TestDay15Part2Example(t *testing.T) {
	// Given
	day15 := readStringsInFile("../data/day15_example.txt")
	expected := 315

	// When
	lowestTotalRisk := GetLowestTotalRiskPathInRealMap(day15)

	// Then
	if lowestTotalRisk != expected {
		t.Errorf("Day 15 example - Part 2 example: expected %v, got %v", expected, lowestTotalRisk)
	}
}

func TestDay15Part2Real(t *testing.T) {
	// Given
	day15 := readStringsInFile("../data/day15.txt")
	expected := 2916

	// When
	lowestTotalRisk := GetLowestTotalRiskPathInRealMap(day15)

	// Then
	if lowestTotalRisk != expected {
		t.Errorf("Day 15 example - Part 1 example: expected %v, got %v", expected, lowestTotalRisk)
	}
}


