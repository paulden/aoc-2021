package main

import "testing"

func TestCountSimpleIncreases(t *testing.T) {
	// Given
	data := []int{1, 2}
	expected := 1

	// When
	result := CountIncreases(data)

	// Then
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCountMultipleIncreases(t *testing.T) {
	// Given
	data := []int{1, 2, 3, 4, 5}
	expected := 4

	// When
	result := CountIncreases(data)

	// Then
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCountMultipleIncreasesWithDecreases(t *testing.T) {
	// Given
	data := []int{4, 5, 1, 4, 1, 4}
	expected := 3

	// When
	result := CountIncreases(data)

	// Then
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCountThreeMeasurementsIncreases(t *testing.T) {
	// Given
	data := []int{1, 2, 3, 4}
	expected := 1

	// When
	result := CountThreeMeasurementsIncreases(data)

	// Then
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
