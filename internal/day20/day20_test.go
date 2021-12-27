package day20

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay20ParseEnhancementAlgorithm(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expectedLength := 512
	expectedStart := []int{0, 0, 1, 0, 1, 0, 0, 1, 1, 1}

	// When
	algorithm := parseEnhancementAlgorithm(input[0])

	// Then
	if len(algorithm) != expectedLength {
		t.Errorf("Day 20 - parse enhancement algorithm: expectedStart length %v, got %v", len(expectedStart), expectedLength)
	}
	for i := range expectedStart {
		if algorithm[i] != expectedStart[i] {
			t.Errorf("Day 20 - parse enhancement algorithm: expectedStart %v, got %v", expectedStart[i], algorithm[i])
		}
	}
}

func TestDay20Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expectedLitPixels := 35

	// When
	litPixels := ProcessEnhancement(input, 2)

	// Then
	if litPixels != expectedLitPixels {
		t.Errorf("Day 20 - Expected %v, got %v", expectedLitPixels, litPixels)
	}
}

func TestDay20Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expectedLitPixels := 3351

	// When
	litPixels := ProcessEnhancement(input, 50)

	// Then
	if litPixels != expectedLitPixels {
		t.Errorf("Day 20 - Expected %v, got %v", expectedLitPixels, litPixels)
	}
}

func TestDay20Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expectedLitPixels := 5680

	// When
	litPixels := ProcessEnhancement(input, 2)

	// Then
	if litPixels != expectedLitPixels {
		t.Errorf("Day 20 - Expected %v, got %v", expectedLitPixels, litPixels)
	}
}

func TestDay20Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expectedLitPixels := 19766

	// When
	litPixels := ProcessEnhancement(input, 50)

	// Then
	if litPixels != expectedLitPixels {
		t.Errorf("Day 20 - Expected %v, got %v", expectedLitPixels, litPixels)
	}
}
