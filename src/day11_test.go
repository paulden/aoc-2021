package main

import (
	"testing"
)

func TestDay11ExamplePart1(t *testing.T) {
	// Given
	day11 := readStringsInFile("../data/day11_example.txt")
	expected := 1656

	// When
	totalFlashes := CountFlashes(day11)

	// Then
	if totalFlashes != expected {
		t.Errorf("Day 11 example - Part 1: expected %v, got %v", expected, totalFlashes)
	}
}

func TestDay11ExamplePart2(t *testing.T) {
	// Given
	day11 := readStringsInFile("../data/day11_example.txt")
	expected := 195

	// When
	mingBlowingStep := FindMindBlowingStep(day11)

	// Then
	if mingBlowingStep != expected {
		t.Errorf("Day 11 example - Part 12: expected %v, got %v", expected, mingBlowingStep)
	}
}

func TestDay11SingleStep(t *testing.T) {
	// Given
	// 1 1 1 1 1
	// 1 9 9 9 1
	// 1 9 1 9 1
	// 1 9 9 9 1
	// 1 1 1 1 1
	day11 := readStringsInFile("../data/day11_simple_example.txt")
	parsedGrid := ParseOctopusesGrid(day11)
	expectedFlashes := 9
	expectedGrid := [][]int{
		{3, 4, 5, 4, 3},
		{4, 0, 0, 0, 4},
		{5, 0, 0, 0, 5},
		{4, 0, 0, 0, 4},
		{3, 4, 5, 4, 3},
	}

	// When
	grid, flashes := PlaySingleStep(parsedGrid)

	// Then
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != expectedGrid[i][j] {
				t.Errorf("Single step : expected %v, got %v at coordinates (%v, %v)", expectedGrid[i][j], grid[i][j], i, j)
			}
		}
	}
	if flashes != expectedFlashes {
		t.Errorf("Single step : expected %v flashes, got %v", expectedFlashes, flashes)
	}
}

func TestDay11SingleComplicatedStep(t *testing.T) {
	// Given
	// 6 5 9 4 2 5 4 3 3 4
	// 3 8 5 6 9 6 5 8 2 2
	// 6 3 7 5 6 6 7 2 8 4
	// 7 2 5 2 4 4 7 2 5 7
	// 7 4 6 8 4 9 6 5 8 9
	// 5 2 7 8 6 3 5 7 5 6
	// 3 2 8 7 9 5 2 8 3 2
	// 7 9 9 3 9 9 2 2 4 5
	// 5 9 5 7 9 5 9 6 6 5
	// 6 3 9 4 8 6 2 6 3 7

	day11 := readStringsInFile("../data/day11_second_example.txt")
	parsedGrid := ParseOctopusesGrid(day11)
	expectedFlashes := 35
	expectedGrid := [][]int{
		{8, 8, 0, 7, 4, 7, 6, 5, 5, 5},
		{5, 0, 8, 9, 0, 8, 7, 0, 5, 4},
		{8, 5, 9, 7, 8, 8, 9, 6, 0, 8},
		{8, 4, 8, 5, 7, 6, 9, 6, 0, 0},
		{8, 7, 0, 0, 9, 0, 8, 8, 0, 0},
		{6, 6, 0, 0, 0, 8, 8, 9, 8, 9},
		{6, 8, 0, 0, 0, 0, 5, 9, 4, 3},
		{0, 0, 0, 0, 0, 0, 7, 4, 5, 6},
		{9, 0, 0, 0, 0, 0, 0, 8, 7, 6},
		{8, 7, 0, 0, 0, 0, 6, 8, 4, 8},
	}

	// When
	grid, flashes := PlaySingleStep(parsedGrid)

	// Then
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != expectedGrid[i][j] {
				t.Errorf("Single step : expected %v, got %v at coordinates (%v, %v)", expectedGrid[i][j], grid[i][j], i, j)
			}
		}
	}
	if flashes != expectedFlashes {
		t.Errorf("Single step : expected %v flashes, got %v", expectedFlashes, flashes)
	}
}
