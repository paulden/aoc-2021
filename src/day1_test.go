package main

import "testing"

func TestCountSimpleIncreases(t *testing.T) {
	// Given
	data := []int{1, 2}

	// When
	result := CountIncreases(data)

	// Then
	if result != 1 {
		t.Errorf("Expected %v, got %v", 1, result)
	}
}

func TestCountMultipleIncreases(t *testing.T) {
	// Given
	data := []int{1, 2, 3, 4, 5}

	// When
	result := CountIncreases(data)

	// Then
	if result != 4 {
		t.Errorf("Expected %v, got %v", 4, result)
	}
}

func TestCountMultipleIncreasesWithDecreases(t *testing.T) {
	// Given
	data := []int{4, 5, 1, 4, 1, 4}

	// When
	result := CountIncreases(data)

	// Then
	if result != 3 {
		t.Errorf("Expected %v, got %v", 3, result)
	}
}

func TestCountThreeMeasurementsIncreases(t *testing.T) {
	// Given
	data := []int{1, 2, 3, 4}

	// When
	result := CountThreeMeasurementsIncreases(data)

	// Then
	if result != 1 {
		t.Errorf("Expected %v, got %v", 1, result)
	}
}