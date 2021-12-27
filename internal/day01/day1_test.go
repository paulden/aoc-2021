package day01

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay01CountSimpleIncreases(t *testing.T) {
	// Given
	data := []int{1, 2}
	expected := 1

	// When
	result := CountIncreases(data)

	// Then
	if result != expected {
		t.Errorf("Day 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay01CountMultipleIncreases(t *testing.T) {
	// Given
	data := []int{1, 2, 3, 4, 5}
	expected := 4

	// When
	result := CountIncreases(data)

	// Then
	if result != expected {
		t.Errorf("Day 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay01CountMultipleIncreasesWithDecreases(t *testing.T) {
	// Given
	data := []int{4, 5, 1, 4, 1, 4}
	expected := 3

	// When
	result := CountIncreases(data)

	// Then
	if result != expected {
		t.Errorf("Day 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay01CountThreeMeasurementsIncreases(t *testing.T) {
	// Given
	data := []int{1, 2, 3, 4}
	expected := 1

	// When
	result := CountThreeMeasurementsIncreases(data)

	// Then
	if result != expected {
		t.Errorf("Day 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay01Part1(t *testing.T) {
	// Given
	input := utils.ReadIntegersInFile("testdata/input.txt")
	expected := 1766

	// When
	result := CountIncreases(input)

	// Then
	if result != expected {
		t.Errorf("Day 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay01Part2(t *testing.T) {
	// Given
	input := utils.ReadIntegersInFile("testdata/input.txt")
	expected := 1797

	// When
	result := CountThreeMeasurementsIncreases(input)

	// Then
	if result != expected {
		t.Errorf("Day 1 - Expected %v, got %v", expected, result)
	}
}
