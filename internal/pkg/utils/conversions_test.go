package utils

import "testing"

func TestDay16HexadecimalToBits(t *testing.T) {
	// Given
	input := "EE00D40C823060"
	expected := []int{1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0}

	// When
	result := HexadecimalToBits(input)

	// Then
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Day 16 - Expected %v, got %v", expected[i], result[i])
		}
	}
}

func TestDay16BitsToDecimal(t *testing.T) {
	// Given
	input := []int{0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1}
	expected := 2021

	// When
	result := BitsToDecimal(input)

	// Then
	if result != expected {
		t.Errorf("Day 16 - Expected %v, got %v", expected, result)
	}
}
