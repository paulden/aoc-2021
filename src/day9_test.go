package main

import (
	"testing"
)

func TestDay9ExamplePart1(t *testing.T) {
	// Given
	day8 := readStringsInFile("../data/day9_example.txt")

	// When
	riskLevel := GetSmokeRiskLevel(day8)

	// Then
	if riskLevel != 15 {
		t.Errorf("Day 9 example - Part 1: expected %v, got %v", 15, riskLevel)
	}
}

func TestDay9ExampleFindBasinSize(t *testing.T) {
	// Given
	day8 := readStringsInFile("../data/day9_example.txt")
	heightmap := ParseHeightmap(day8)
	x, y := 2, 2

	// When
	basinSize := GetBasinSize(x, y, heightmap)

	// Then
	if basinSize != 14 {
		t.Errorf("Day 9 basin size: expected %v, got %v", 14, basinSize)
	}
}

func TestDay9ExamplePart2(t *testing.T) {
	// Given
	day8 := readStringsInFile("../data/day9_example.txt")

	// When
	largestBasins := GetBiggestBasins(day8)

	// Then
	if largestBasins != 1134 {
		t.Errorf("Day 9 example - Part 1: expected %v, got %v", 1134, largestBasins)
	}
}
