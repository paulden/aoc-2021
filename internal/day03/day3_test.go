package day03

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay03ComputeGamma(t *testing.T) {
	type args struct {
		report []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"10110", args{[]string{"10110"}}, 22},
		{"01001", args{[]string{"01001"}}, 9},
		{"Three sequences", args{[]string{"10110", "01001", "11100"}}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeGamma(tt.args.report); got != tt.want {
				t.Errorf("computeGamma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay03ComputeEpsilon(t *testing.T) {
	type args struct {
		report []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"10110", args{[]string{"10110"}}, 9},
		{"01001", args{[]string{"01001"}}, 22},
		{"01001", args{[]string{"10110", "01001", "11100"}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeEpsilon(tt.args.report); got != tt.want {
				t.Errorf("computeGamma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay03ComputeOxygenGeneratorRating(t *testing.T) {
	type args struct {
		report []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Three sequences", args{[]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeOxygenGeneratorRating(tt.args.report); got != tt.want {
				t.Errorf("computeOxygenGeneratorRating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay03ComputeCO2ScrubberRating(t *testing.T) {
	type args struct {
		report []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Three sequences", args{[]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeCO2ScrubberRating(tt.args.report); got != tt.want {
				t.Errorf("computeOxygenGeneratorRating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay02Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 3882564

	// When
	result := GetPowerConsumption(input)

	// Then
	if result != expected {
		t.Errorf("Day 3 - Expected %v, got %v", expected, result)
	}
}

func TestDay02Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 3385170

	// When
	result := GetLifeSupportRating(input)

	// Then
	if result != expected {
		t.Errorf("Day 3 - Expected %v, got %v", expected, result)
	}
}
