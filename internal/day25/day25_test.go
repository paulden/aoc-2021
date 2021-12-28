package day25

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay25ParseSeaCucumbersMap(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example2.txt")
	seaCucumberEast1 := seaCucumbersCoordinates{0, 3}
	seaCucumberEast2 := seaCucumbersCoordinates{2, 6}
	seaCucumberEast3 := seaCucumbersCoordinates{3, 6}
	seaCucumberEast4 := seaCucumbersCoordinates{4, 6}
	seaCucumberSouth1 := seaCucumbersCoordinates{3, 0}
	seaCucumberSouth2 := seaCucumbersCoordinates{6, 2}
	seaCucumberSouth3 := seaCucumbersCoordinates{6, 3}
	seaCucumberSouth4 := seaCucumbersCoordinates{6, 4}

	// When
	seaCucumbersMap := parseSeaCucumbersMap(input)

	if seaCucumbersMap[seaCucumberEast1] != 1 && seaCucumbersMap[seaCucumberEast2] != 1 && seaCucumbersMap[seaCucumberEast3] != 1 && seaCucumbersMap[seaCucumberEast4] != 1 {
		t.Errorf("Day 25 - Expected sea cucumbers facing east at these locations but were not")
	}
	if seaCucumbersMap[seaCucumberSouth1] != 2 && seaCucumbersMap[seaCucumberSouth2] != 2 && seaCucumbersMap[seaCucumberSouth3] != 2 && seaCucumbersMap[seaCucumberSouth4] != 2 {
		t.Errorf("Day 25 - Expected sea cucumbers facing south at these locations but were not")
	}
}

func TestDay25Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example1.txt")
	expectedSteps := 58

	// When
	steps := CountStepsBeforeImmobilisation(input)

	if steps != expectedSteps {
		t.Errorf("Day 25 - Expected %v, got %v", expectedSteps, steps)
	}
}

func TestDay25Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expectedSteps := 380

	// When
	steps := CountStepsBeforeImmobilisation(input)

	if steps != expectedSteps {
		t.Errorf("Day 25 - Expected %v, got %v", expectedSteps, steps)
	}
}
