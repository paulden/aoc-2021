package day14

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay14Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 1588

	// When
	polymerCountsDifference := CountPolymerCountsDifferenceNaive(input, 10)

	// Then
	if polymerCountsDifference != expected {
		t.Errorf("Day 14 - Expected %v, got %v", expected, polymerCountsDifference)
	}
}

func TestDay14Part1ExampleOptimized(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 1588

	// When
	polymerCountsDifference := CountPolymerCountsDifferenceOptimized(input, 10)

	// Then
	if polymerCountsDifference != expected {
		t.Errorf("Day 14 - Expected %v, got %v", expected, polymerCountsDifference)
	}
}

func TestDay14Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 2188189693529

	// When
	polymerCountsDifference := CountPolymerCountsDifferenceOptimized(input, 40)

	// Then
	if polymerCountsDifference != expected {
		t.Errorf("Day 14 - Expected %v, got %v", expected, polymerCountsDifference)
	}
}

func TestDay14Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 2590

	// When
	polymerCountsDifference := CountPolymerCountsDifferenceNaive(input, 10)

	// Then
	if polymerCountsDifference != expected {
		t.Errorf("Day 14 - Expected %v, got %v", expected, polymerCountsDifference)
	}
}

func TestDay14Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 2875665202438

	// When
	polymerCountsDifference := CountPolymerCountsDifferenceOptimized(input, 40)

	// Then
	if polymerCountsDifference != expected {
		t.Errorf("Day 14 - Expected %v, got %v", expected, polymerCountsDifference)
	}
}
