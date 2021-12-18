package main

import (
	"testing"
)

func TestDay17ParseTargetArea(t *testing.T) {
	// Given
	input := "target area: x=20..30, y=-10..-5"
	expected := TargetArea{30, 20, -5, -10}

	// When
	result := ParseTargetArea(input)

	// Then
	if result != expected {
		t.Errorf("Day 17 - parse target area: expected %v, got %v", expected, result)
	}
}

func TestDay17IsValidTrajectory(t *testing.T) {
	// Given
	input := "target area: x=20..30, y=-10..-5"
	targetArea := ParseTargetArea(input)
	xv1, yv1 := 7, 2
	xv2, yv2 := 6, 3
	xv3, yv3 := 9, 0
	xv4, yv4 := 17, -4

	// When
	result1, _, _, _ := ComputeTrajectory(targetArea, xv1, yv1)
	result2, _, _, _ := ComputeTrajectory(targetArea, xv2, yv2)
	result3, _, _, _ := ComputeTrajectory(targetArea, xv3, yv3)
	result4, _, _, _ := ComputeTrajectory(targetArea, xv4, yv4)

	// Then
	if !result1 {
		t.Errorf("Day 17 - ComputeTrajectory should be true for initial velocity (%v, %v)", xv1, yv1)
	}
	if !result2 {
		t.Errorf("Day 17 - ComputeTrajectory should be true for initial velocity (%v, %v)", xv2, yv2)
	}
	if !result3 {
		t.Errorf("Day 17 - ComputeTrajectory should be true for initial velocity (%v, %v)", xv3, yv3)
	}
	if result4 {
		t.Errorf("Day 17 - ComputeTrajectory should be false for initial velocity (%v, %v)", xv4, yv4)
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
		t.Errorf("Day 17 - example: expected %v, got %v", expected, result)
	}
}

func TestDay17Part2Example(t *testing.T) {
	// Given
	input := "target area: x=20..30, y=-10..-5"
	expected := 112

	// When
	result := GetAllVelocities(input)

	// Then
	if result != expected {
		t.Errorf("Day 17 Part 2- example: expected %v, got %v", expected, result)
	}
}

func TestDay17Part2Real(t *testing.T) {
	// Given
	input := "target area: x=211..232, y=-124..-69"
	expected := 2032

	// When
	result := GetAllVelocities(input)

	// Then
	if result != expected {
		t.Errorf("Day 17 Part 2- example: expected %v, got %v", expected, result)
	}
}

func TestDay17Part1Real(t *testing.T) {
	// Given
	input := "target area: x=211..232, y=-124..-69"
	expected := 7626

	// When
	result := GetMaxHeight(input)

	// Then
	if result != expected {
		t.Errorf("Day 17 - example: expected %v, got %v", expected, result)
	}
}
