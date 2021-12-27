package day07

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay07Part1Example(t *testing.T) {
	// Given
	day7 := "16,1,2,0,4,2,7,1,2,14"
	expected := 37

	// When
	fuelConsumption := GetCheapestFuelConsumption(day7)

	// Then
	if fuelConsumption != expected {
		t.Errorf("Day 7 - Expected %v, got %v", expected, fuelConsumption)
	}
}

func TestDay07Part2Example(t *testing.T) {
	// Given
	day7 := "16,1,2,0,4,2,7,1,2,14"
	expected := 168

	// When
	fuelConsumption := GetCheapestFuelConsumptionUpdated(day7)

	// Then
	if fuelConsumption != expected {
		t.Errorf("Day 7 - Expected %v, got %v", expected, fuelConsumption)
	}
}

func TestDay07Part1(t *testing.T) {
	// Given
	day7 := utils.ReadStringsInFile("testdata/input.txt")
	expected := 352997

	// When
	fuelConsumption := GetCheapestFuelConsumption(day7[0])

	// Then
	if fuelConsumption != expected {
		t.Errorf("Day 7 - Expected %v, got %v", expected, fuelConsumption)
	}
}

func TestDay07Part2(t *testing.T) {
	// Given
	day7 := utils.ReadStringsInFile("testdata/input.txt")
	expected := 101571302

	// When
	fuelConsumption := GetCheapestFuelConsumptionUpdated(day7[0])

	// Then
	if fuelConsumption != expected {
		t.Errorf("Day 7 - Expected %v, got %v", expected, fuelConsumption)
	}
}
