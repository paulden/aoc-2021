package day06

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay06Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	daysBeforeCheck := 80
	expected := 5934

	// When
	lanternfishNumber := CountLanternfishesNaive(input[0], daysBeforeCheck)

	// Then
	if lanternfishNumber != expected {
		t.Errorf("Day 6 - Expected %v, got %v", expected, lanternfishNumber)
	}
}

func TestDay06Part1ExampleOptimized(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	daysBeforeCheck := 80
	expected := 5934

	// When
	lanternfishNumber := CountLanternfishesOptimized(input[0], daysBeforeCheck)

	// Then
	if lanternfishNumber != expected {
		t.Errorf("Day 6 - Expected %v, got %v", expected, lanternfishNumber)
	}
}

func TestDay06Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	daysBeforeCheck := 256
	expected := 26984457539

	// When
	lanternfishNumber := CountLanternfishesOptimized(input[0], daysBeforeCheck)

	// Then
	if lanternfishNumber != expected {
		t.Errorf("Day 6 - Expected %v, got %v", expected, lanternfishNumber)
	}
}

func TestDay06Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	daysBeforeCheck := 80
	expected := 395627

	// When
	lanternfishNumber := CountLanternfishesNaive(input[0], daysBeforeCheck)

	// Then
	if lanternfishNumber != expected {
		t.Errorf("Day 6 - Expected %v, got %v", expected, lanternfishNumber)
	}
}

func TestDay06Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	daysBeforeCheck := 256
	expected := 1767323539209

	// When
	lanternfishNumber := CountLanternfishesOptimized(input[0], daysBeforeCheck)

	// Then
	if lanternfishNumber != expected {
		t.Errorf("Day 6 - Expected %v, got %v", expected, lanternfishNumber)
	}
}

func TestCountLanternfishesOptimized(t *testing.T) {
	type args struct {
		lanternfishesAgesString string
		daysBeforeCheck         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"1", 1}, 1},
		{"0,1", args{"0,1", 1}, 3},
		{"0,1,6", args{"0,1,6", 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLanternfishesOptimized(tt.args.lanternfishesAgesString, tt.args.daysBeforeCheck); got != tt.want {
				t.Errorf("CountLanternfishesOptimized() = %v, want %v", got, tt.want)
			}
		})
	}
}
