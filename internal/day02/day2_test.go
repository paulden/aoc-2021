package day02

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay02ParseInstruction(t *testing.T) {
	type args struct {
		instruction string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{"forward 5", args{"forward 5"}, "forward", 5},
		{"forward 3", args{"forward 3"}, "forward", 3},
		{"down 3", args{"down 3"}, "down", 3},
		{"up 8", args{"up 8"}, "up", 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInstruction(tt.args.instruction)
			if got != tt.want {
				t.Errorf("parseInstruction() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseInstruction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDay02DeterminePosition(t *testing.T) {
	type args struct {
		instructions []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Simple test case", args{[]string{"forward 5"}}, 0},
		{"Second test case", args{[]string{"forward 5", "forward 2"}}, 0},
		{"Third test case", args{[]string{"forward 5", "down 2"}}, 10},
		{"Fourth test case", args{[]string{"forward 5", "down 2", "up 1"}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeterminePosition(tt.args.instructions)
			if got != tt.want {
				t.Errorf("DeterminePosition() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay02DetermineAimPosition(t *testing.T) {
	type args struct {
		instructions []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Simple test case", args{[]string{"forward 5"}}, 0},
		{"Second test case", args{[]string{"forward 5", "down 2"}}, 0},
		{"Third test case", args{[]string{"forward 5", "down 5", "forward 8"}}, 520},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeterminePositionWithAim(tt.args.instructions)
			if got != tt.want {
				t.Errorf("DeterminePositionWithAim() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay02Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 2036120

	// When
	result := DeterminePosition(input)

	// Then
	if result != expected {
		t.Errorf("Day 2 - Expected %v, got %v", expected, result)
	}
}

func TestDay02Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 2015547716

	// When
	result := DeterminePositionWithAim(input)

	// Then
	if result != expected {
		t.Errorf("Day 2 - Expected %v, got %v", expected, result)
	}
}
