package main

import (
	"testing"
)

func TestDay7Example(t *testing.T) {
	// Given
	day7 := "16,1,2,0,4,2,7,1,2,14"
	expected := 37

	// When
	fuelConsumption := GetCheapestFuelConsumption(day7)

	// Then
	if fuelConsumption != expected {
		t.Errorf("Day 7 example - Part 1: expected %v, got %v", expected, fuelConsumption)
	}
}

func TestDay7ExamplePart6(t *testing.T) {
	// Given
	day7 := "16,1,2,0,4,2,7,1,2,14"
	expected := 168

	// When
	fuelConsumption := GetCheapestFuelConsumptionUpdated(day7)

	// Then
	if fuelConsumption != expected {
		t.Errorf("Day 7 example - Part 2: expected %v, got %v", expected, fuelConsumption)
	}
}
