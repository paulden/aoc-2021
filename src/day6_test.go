package main

import (
	"testing"
)

func TestDay6Example(t *testing.T) {
	// Given
	day6 := "3,4,3,1,2"
	daysBeforeCheck := 80

	// When
	lanternfishNumber := CountLanternfishesNaive(day6, daysBeforeCheck)

	// Then
	if lanternfishNumber != 5934 {
		t.Errorf("Day 6 example: expected %v, got %v", 5934, lanternfishNumber)
	}
}

func TestDay6ExampleOptimized(t *testing.T) {
	// Given
	day6 := "3,4,3,1,2"
	daysBeforeCheck := 80

	// When
	lanternfishNumber := CountLanternfishesOptimized(day6, daysBeforeCheck)

	// Then
	if lanternfishNumber != 5934 {
		t.Errorf("Day 6 example: expected %v, got %v", 5934, lanternfishNumber)
	}
}

func TestDay6ExamplePart2(t *testing.T) {
	// Given
	day6 := "3,4,3,1,2"
	daysBeforeCheck := 256

	// When
	lanternfishNumber := CountLanternfishesOptimized(day6, daysBeforeCheck)

	// Then
	if lanternfishNumber != 26984457539 {
		t.Errorf("Day 6 example: expected %v, got %v", 5934, lanternfishNumber)
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