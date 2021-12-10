package main

import (
	"testing"
)

func TestDay10ExamplePart1(t *testing.T) {
	// Given
	day10 := readStringsInFile("../data/day10_example.txt")
	expected := 26397

	// When
	syntaxErrorScore := GetSyntaxErrorScore(day10)

	// Then
	if syntaxErrorScore != expected {
		t.Errorf("Day 10 example - Part 1: expected %v, got %v", expected, syntaxErrorScore)
	}
}

func TestDay10ExamplePart2(t *testing.T) {
	// Given
	day9 := readStringsInFile("../data/day10_example.txt")
	expected := 288957

	// When
	autocompletionScore := GetAutocompletionScore(day9)

	// Then
	if autocompletionScore != expected {
		t.Errorf("Day 10 example - Part 2: expected %v, got %v", expected, autocompletionScore)
	}
}
