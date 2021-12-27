package day19

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay19ParseScannersAndBeacons(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expectedScannersNumbers := 5
	expectedBeaconsInScanner0 := 25
	expectedFirstCoordinates := coordinates3D{404, -588, -901}

	// When
	result := parseScannersAndBeacons(input)

	// Then
	if len(result) != expectedScannersNumbers {
		t.Errorf("Day 19 - Scanners and beacons parsing invalid, expected %v scanners, got %v", expectedScannersNumbers, len(result))
	}
	if len(result[0]) != expectedBeaconsInScanner0 {
		t.Errorf("Day 19 - Scanners and beacons parsing invalid, expected %v beacons for scanner 1, got %v", expectedBeaconsInScanner0, len(result[0]))
	}
	if result[0][0] != expectedFirstCoordinates {
		t.Errorf("Day 19 - Scanners and beacons parsing invalid, expected first coordinates to be %v, got %v", expectedFirstCoordinates, result[0][0])
	}
}

func TestDay19Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expectedBeacons := 79

	// When
	beacons := CountBeacons(input)

	// Then
	if beacons != expectedBeacons {
		t.Errorf("Day 19 - Expected %v scanners, got %v", expectedBeacons, beacons)
	}
}

func TestDay19Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expectedBeacons := 326

	// When
	beacons := CountBeacons(input)

	// Then
	if beacons != expectedBeacons {
		t.Errorf("Day 19 - Expected %v scanners, got %v", expectedBeacons, beacons)
	}
}

func TestDay19Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expectedMaximumDistance := 3621

	// When
	manhattanDistance := GetMaximumManhattanDistance(input)

	// Then
	if manhattanDistance != expectedMaximumDistance {
		t.Errorf("Day 19 - Expected %v scanners, got %v", expectedMaximumDistance, manhattanDistance)
	}
}

func TestDay19Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expectedMaximumDistance := 10630

	// When
	manhattanDistance := GetMaximumManhattanDistance(input)

	// Then
	if manhattanDistance != expectedMaximumDistance {
		t.Errorf("Day 19 - Expected %v scanners, got %v", expectedMaximumDistance, manhattanDistance)
	}
}
