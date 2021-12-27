package day10

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay10Part1Example(t *testing.T) {
	// Given
	day10 := utils.ReadStringsInFile("testdata/example.txt")
	expected := 26397

	// When
	syntaxErrorScore := GetSyntaxErrorScore(day10)

	// Then
	if syntaxErrorScore != expected {
		t.Errorf("Day 10 - Expected %v, got %v", expected, syntaxErrorScore)
	}
}

func TestDay10Part2Example(t *testing.T) {
	// Given
	day9 := utils.ReadStringsInFile("testdata/example.txt")
	expected := 288957

	// When
	autocompletionScore := GetAutocompletionScore(day9)

	// Then
	if autocompletionScore != expected {
		t.Errorf("Day 10 - Expected %v, got %v", expected, autocompletionScore)
	}
}

func TestDay10Part1(t *testing.T) {
	// Given
	day10 := utils.ReadStringsInFile("testdata/input.txt")
	expected := 321237

	// When
	syntaxErrorScore := GetSyntaxErrorScore(day10)

	// Then
	if syntaxErrorScore != expected {
		t.Errorf("Day 10 - Expected %v, got %v", expected, syntaxErrorScore)
	}
}

func TestDay10Part2(t *testing.T) {
	// Given
	day9 := utils.ReadStringsInFile("testdata/input.txt")
	expected := 2360030859

	// When
	autocompletionScore := GetAutocompletionScore(day9)

	// Then
	if autocompletionScore != expected {
		t.Errorf("Day 10 - Expected %v, got %v", expected, autocompletionScore)
	}
}
