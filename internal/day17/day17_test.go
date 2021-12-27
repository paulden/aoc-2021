package day17

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay17ParseTargetArea(t *testing.T) {
	// Given
	input := "target area: x=20..30, y=-10..-5"
	expected := targetArea{30, 20, -5, -10}

	// When
	result := parseTargetArea(input)

	// Then
	if result != expected {
		t.Errorf("Day 17 - Expected %v, got %v", expected, result)
	}
}

func TestDay17IsValidTrajectory(t *testing.T) {
	// Given
	input := "target area: x=20..30, y=-10..-5"
	area := parseTargetArea(input)
	xv1, yv1 := 7, 2
	xv2, yv2 := 6, 3
	xv3, yv3 := 9, 0
	xv4, yv4 := 17, -4

	// When
	result1, _, _, _ := computeTrajectory(area, xv1, yv1)
	result2, _, _, _ := computeTrajectory(area, xv2, yv2)
	result3, _, _, _ := computeTrajectory(area, xv3, yv3)
	result4, _, _, _ := computeTrajectory(area, xv4, yv4)

	// Then
	if !result1 {
		t.Errorf("Day 17 - computeTrajectory should be true for initial velocity (%v, %v)", xv1, yv1)
	}
	if !result2 {
		t.Errorf("Day 17 - computeTrajectory should be true for initial velocity (%v, %v)", xv2, yv2)
	}
	if !result3 {
		t.Errorf("Day 17 - computeTrajectory should be true for initial velocity (%v, %v)", xv3, yv3)
	}
	if result4 {
		t.Errorf("Day 17 - computeTrajectory should be false for initial velocity (%v, %v)", xv4, yv4)
	}
}

func TestDay17Part1Example(t *testing.T) {
	// Given
	input := "target area: x=20..30, y=-10..-5"
	expected := 45

	// When
	result := GetMaxHeight(input)

	// Then
	if result != expected {
		t.Errorf("Day 17 - Expected %v, got %v", expected, result)
	}
}

func TestDay17Part2Example(t *testing.T) {
	// Given
	input := "target area: x=20..30, y=-10..-5"
	expected := 112

	// When
	result := getAllVelocities(input)

	// Then
	if len(result) != expected {
		t.Errorf("Day 17 - Expected %v, got %v", expected, len(result))
	}
}

func TestDay17Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 7626

	// When
	result := GetMaxHeight(input[0])

	// Then
	if result != expected {
		t.Errorf("Day 17 - Expected %v, got %v", expected, result)
	}
}

func TestDay17Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 2032

	// When
	result := CountAllVelocities(input[0])

	// Then
	if result != expected {
		t.Errorf("Day 17 - Expected %v, got %v", expected, result)
	}
}
