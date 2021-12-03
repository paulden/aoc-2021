package main

import "testing"

func TestParseInstruction(t *testing.T) {
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
			got, got1 := ParseInstruction(tt.args.instruction)
			if got != tt.want {
				t.Errorf("ParseInstruction() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseInstruction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_determinePosition(t *testing.T) {
	type args struct {
		instructions []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"Simple test case", args{ []string{"forward 5"} }, 5, 0},
		{"Second test case", args{ []string{"forward 5", "forward 2"} }, 7, 0},
		{"Third test case", args{ []string{"forward 5", "down 2"} }, 5, 2},
		{"Fourth test case", args{ []string{"forward 5", "down 2", "up 1"} }, 5, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := determinePosition(tt.args.instructions)
			if got != tt.want {
				t.Errorf("determinePosition() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("determinePosition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_determineAimPosition(t *testing.T) {
	type args struct {
		instructions []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"Simple test case", args{ []string{"forward 5"} }, 5, 0},
		{"Second test case", args{ []string{"forward 5", "down 2"} }, 5, 0},
		{"Third test case", args{ []string{"forward 5", "down 5", "forward 8"} }, 13, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := determinePositionWithAim(tt.args.instructions)
			if got != tt.want {
				t.Errorf("determinePositionWithAim() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("determinePositionWithAim() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}