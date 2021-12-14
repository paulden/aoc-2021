package main

import (
	"testing"
)

func TestDay14Part1Example(t *testing.T) {
	// Given
	day14 := readStringsInFile("../data/day14_example.txt")
	expected := 1588

	// When
	polymerCountsDifference := CountPolymerCountsDifferenceNaive(day14, 10)

	// Then
	if polymerCountsDifference != expected {
		t.Errorf("Day 14 example - Part 1 example: expected %v, got %v", expected, polymerCountsDifference)
	}
}

func TestDay14Part1ExampleOptimized(t *testing.T) {
	// Given
	day14 := readStringsInFile("../data/day14_example.txt")
	expected := 1588

	// When
	polymerCountsDifference := CountPolymerCountsDifferenceOptimized(day14, 10)

	// Then
	if polymerCountsDifference != expected {
		t.Errorf("Day 14 example - Part 1 example: expected %v, got %v", expected, polymerCountsDifference)
	}
}

func TestDay14Part2Example(t *testing.T) {
	// Given
	day14 := readStringsInFile("../data/day14_example.txt")
	expected := 2188189693529

	// When
	polymerCountsDifference := CountPolymerCountsDifferenceOptimized(day14, 40)

	// Then
	if polymerCountsDifference != expected {
		t.Errorf("Day 14 example - Part 2 example: expected %v, got %v", expected, polymerCountsDifference)
	}
}
