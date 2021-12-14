package main

import (
	"testing"
)

func TestDay13Part1Example(t *testing.T) {
	// Given
	day13 := readStringsInFile("../data/day13_example.txt")
	expected := 17

	// When
	visibleDotsAfterOneFold := FoldPaperOnce(day13)

	// Then
	if len(visibleDotsAfterOneFold) != expected {
		t.Errorf("Day 13 example - Part 1 example: expected %v, got %v", expected, visibleDotsAfterOneFold)
	}
}
