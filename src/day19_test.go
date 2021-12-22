package main

import (
	"testing"
)

func TestDay19ParseScannersAndBeacons(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day19_example.txt")
	expectedScannersNumbers := 5
	expectedBeaconsInScanner0 := 25
	expectedFirstCoordinates := Coordinates3D{404, -588, -901}

	// When
	result := ParseScannersAndBeacons(input)

	// Then
	if len(result) != expectedScannersNumbers {
		t.Errorf("Day 19 : scanners and beacons parsing invalid, expected %v scanners, got %v", expectedScannersNumbers, len(result))
	}
	if len(result[0]) != expectedBeaconsInScanner0 {
		t.Errorf("Day 19 : scanners and beacons parsing invalid, expected %v beacons for scanner 1, got %v", expectedBeaconsInScanner0, len(result[0]))
	}
	if result[0][0] != expectedFirstCoordinates {
		t.Errorf("Day 19 : scanners and beacons parsing invalid, expected first coordinates to be %v, got %v", expectedFirstCoordinates, result[0][0])
	}
}

func TestDay19DPart1Example(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day19_example.txt")
	expectedBeacons := 79

	// When
	beacons := CountBeacons(input)

	// Then
	if beacons != expectedBeacons {
		t.Errorf("Day 19 - Part 1 example, expected %v scanners, got %v", expectedBeacons, beacons)
	}
}

func TestDay19DPart1Real(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day19.txt")
	expectedBeacons := 326

	// When
	beacons := CountBeacons(input)

	// Then
	if beacons != expectedBeacons {
		t.Errorf("Day 19 - Part 1 real sample, expected %v scanners, got %v", expectedBeacons, beacons)
	}
}

func TestDay19DPart2Example(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day19_example.txt")
	expectedMaximumDistance := 3621

	// When
	manhattanDistance := GetMaximumManhattanDistance(input)

	// Then
	if manhattanDistance != expectedMaximumDistance {
		t.Errorf("Day 19 - Part 2 example, expected %v scanners, got %v", expectedMaximumDistance, manhattanDistance)
	}
}

func TestDay19DPart2Real(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day19.txt")
	expectedMaximumDistance := 10630

	// When
	manhattanDistance := GetMaximumManhattanDistance(input)

	// Then
	if manhattanDistance != expectedMaximumDistance {
		t.Errorf("Day 19 - Part 2 real sample, expected %v scanners, got %v", expectedMaximumDistance, manhattanDistance)
	}
}
