package day09

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay9Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 15

	// When
	riskLevel := GetSmokeRiskLevel(input)

	// Then
	if riskLevel != expected {
		t.Errorf("Day 9 - Expected %v, got %v", expected, riskLevel)
	}
}

func TestDay9ExampleFindBasinSize(t *testing.T) {
	// Given
	day8 := utils.ReadStringsInFile("testdata/example.txt")
	heightmap := parseHeightmap(day8)
	x, y := 2, 2
	expected := 14

	// When
	basinSize := getBasinSize(x, y, heightmap)

	// Then
	if basinSize != expected {
		t.Errorf("Day 9 - Expected %v, got %v", expected, basinSize)
	}
}

func TestDay9Part2Example(t *testing.T) {
	// Given
	day8 := utils.ReadStringsInFile("testdata/example.txt")
	expected := 1134

	// When
	largestBasins := GetBiggestBasins(day8)

	// Then
	if largestBasins != expected {
		t.Errorf("Day 9 - Expected %v, got %v", expected, largestBasins)
	}
}

func TestDay9Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 506

	// When
	riskLevel := GetSmokeRiskLevel(input)

	// Then
	if riskLevel != expected {
		t.Errorf("Day 9 - Expected %v, got %v", expected, riskLevel)
	}
}

func TestDay9Part2(t *testing.T) {
	// Given
	day8 := utils.ReadStringsInFile("testdata/input.txt")
	expected := 931200

	// When
	largestBasins := GetBiggestBasins(day8)

	// Then
	if largestBasins != expected {
		t.Errorf("Day 9 - Expected %v, got %v", expected, largestBasins)
	}
}
