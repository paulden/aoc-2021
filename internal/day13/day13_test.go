package day13

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay13Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 17

	// When
	visibleDotsAfterOneFold := FoldPaperOnce(input)

	// Then
	if visibleDotsAfterOneFold != expected {
		t.Errorf("Day 13 - Expected %v, got %v", expected, visibleDotsAfterOneFold)
	}
}

func TestDay13Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 737

	// When
	visibleDotsAfterOneFold := FoldPaperOnce(input)

	// Then
	if visibleDotsAfterOneFold != expected {
		t.Errorf("Day 13 - Expected %v, got %v", expected, visibleDotsAfterOneFold)
	}
}
