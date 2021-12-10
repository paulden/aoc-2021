package main

import (
	"testing"
)

func TestDay9ExamplePart1(t *testing.T) {
	// Given
	day8 := readStringsInFile("../data/day9_example.txt")
	expected := 15

	// When
	riskLevel := GetSmokeRiskLevel(day8)

	// Then
	if riskLevel != expected {
		t.Errorf("Day 9 example - Part 1: expected %v, got %v", expected, riskLevel)
	}
}

func TestDay9ExampleFindBasinSize(t *testing.T) {
	// Given
	day8 := readStringsInFile("../data/day9_example.txt")
	heightmap := ParseHeightmap(day8)
	x, y := 2, 2
	expected := 14

	// When
	basinSize := GetBasinSize(x, y, heightmap)

	// Then
	if basinSize != expected {
		t.Errorf("Day 9 basin size: expected %v, got %v", expected, basinSize)
	}
}

func TestDay9ExamplePart2(t *testing.T) {
	// Given
	day8 := readStringsInFile("../data/day9_example.txt")
	expected := 1134

	// When
	largestBasins := GetBiggestBasins(day8)

	// Then
	if largestBasins != expected {
		t.Errorf("Day 9 example - Part 1: expected %v, got %v", expected, largestBasins)
	}
}
