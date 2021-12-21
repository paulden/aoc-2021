package main

import (
	"testing"
)

func TestDay20ParseEnhancementAlgorithm(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day20_example.txt")
	expectedLength := 512
	expectedStart := []int{0, 0, 1, 0, 1, 0, 0, 1, 1, 1}

	// When
	algorithm := ParseEnhancementAlgorithm(input[0])

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

func Test20Part1Example(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day20_example.txt")
	expectedLitPixels := 35

	// When
	litPixels := ProcessEnhancement(input, 2)

	// Then
	if litPixels != expectedLitPixels {
		t.Errorf("Day 20 - Part 1 example: expected %v, got %v", expectedLitPixels, litPixels)
	}
}

func Test20Part1Real(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day20.txt")
	expectedLitPixels := 5680

	// When
	litPixels := ProcessEnhancement(input, 2)

	// Then
	if litPixels != expectedLitPixels {
		t.Errorf("Day 20 - Part 1 real sample: expected %v, got %v", expectedLitPixels, litPixels)
	}
}

func Test20Part2Example(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day20_example.txt")
	expectedLitPixels := 3351

	// When
	litPixels := ProcessEnhancement(input, 50)

	// Then
	if litPixels != expectedLitPixels {
		t.Errorf("Day 20 - Part 1 example: expected %v, got %v", expectedLitPixels, litPixels)
	}
}


func Test20Part2Real(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day20.txt")
	expectedLitPixels := 19766

	// When
	litPixels := ProcessEnhancement(input, 50)

	// Then
	if litPixels != expectedLitPixels {
		t.Errorf("Day 20 - Part 1 real sample: expected %v, got %v", expectedLitPixels, litPixels)
	}
}
