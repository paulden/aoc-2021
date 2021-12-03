package main

import "testing"

func TestComputeGamma(t *testing.T) {
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
			if got := ComputeGamma(tt.args.report); got != tt.want {
				t.Errorf("ComputeGamma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeEpsilon(t *testing.T) {
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
		{"01001", args{ []string{"10110", "01001", "11100"} }, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeEpsilon(tt.args.report); got != tt.want {
				t.Errorf("ComputeGamma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeOxygenGeneratorRating(t *testing.T) {
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
			if got := ComputeOxygenGeneratorRating(tt.args.report); got != tt.want {
				t.Errorf("ComputeOxygenGeneratorRating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeCO2ScrubberRating(t *testing.T) {
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
			if got := ComputeCO2ScrubberRating(tt.args.report); got != tt.want {
				t.Errorf("ComputeOxygenGeneratorRating() = %v, want %v", got, tt.want)
			}
		})
	}
}
