package day24

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay24RunALUSequenceDecimalToBits1(t *testing.T) {
	// Given
	program := utils.ReadStringsInFile("testdata/example1.txt")
	initialRegistry := registry{0, 0, 0, 0}
	expectedRegistry := registry{0, 1, 1, 1}
	input := 7

	// When
	r := runALUSequence(program, input, &initialRegistry)

	// Then
	if *r != expectedRegistry {
		t.Errorf("Day 24 - Error converting decimal input to bits, expected %v, got %v", expectedRegistry, *r)
	}
}

func TestDay24RunALUSequenceDecimalToBits2(t *testing.T) {
	// Given
	program := utils.ReadStringsInFile("testdata/example1.txt")
	initialRegistry := registry{0, 0, 0, 0}
	expectedRegistry := registry{1, 0, 1, 1}
	input := 11

	// When
	r := runALUSequence(program, input, &initialRegistry)

	// Then
	if *r != expectedRegistry {
		t.Errorf("Day 24 - Error converting decimal input to bits, expected %v, got %v", expectedRegistry, *r)
	}
}

func TestDay24Part1(t *testing.T) {
	// Given
	program := utils.ReadStringsInFile("testdata/day24.txt")
	expectedNumber := 92915979999498

	// When
	smallestModelNumber := FindLargestModelNumber(program)

	// Then
	if smallestModelNumber != expectedNumber {
		t.Errorf("Day 24 - Expected %v, got %v", expectedNumber, smallestModelNumber)
	}
}

func TestDay24Part2(t *testing.T) {
	// Given
	program := utils.ReadStringsInFile("testdata/day24.txt")
	expectedNumber := 21611513911181

	// When
	smallestModelNumber := FindSmallestModelNumber(program)

	// Then
	if smallestModelNumber != expectedNumber {
		t.Errorf("Day 24 - Expected %v, got %v", expectedNumber, smallestModelNumber)
	}
}
