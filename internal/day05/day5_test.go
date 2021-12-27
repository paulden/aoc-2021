package day05

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay05ParseVentLine(t *testing.T) {
	type args struct {
		ventLine string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
		want3 int
	}{
		{"First vent line", args{"0,9 -> 5,9"}, 0, 9, 5, 9},
		{"Second vent line", args{"8,0 -> 0,8"}, 8, 0, 0, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := parseVentLine(tt.args.ventLine)
			if got != tt.want {
				t.Errorf("parseVentLine() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseVentLine() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("parseVentLine() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("parseVentLine() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

func TestDay05GetFieldSize(t *testing.T) {
	type args struct {
		ventLines []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"First field", args{[]string{"0,9 -> 5,9"}}, 6, 10},
		{"Second field", args{[]string{"0,9 -> 5,9", "0,9 -> 2,9", "8,0 -> 0,8"}}, 9, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getFieldSize(tt.args.ventLines)
			if got != tt.want {
				t.Errorf("getFieldSize() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getFieldSize() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDay05_MapOceanFloorVentLinesWithX(t *testing.T) {
	ventLines := []string{"0,1 -> 0,2"}

	oceanFloorMapping := mapOceanFloorVentLines(ventLines)

	if oceanFloorMapping[0][0] == 1 {
		t.Errorf("Expected 0, 0 to be empty, was not")
	}

	if oceanFloorMapping[0][1] != 1 && oceanFloorMapping[0][2] != 1 {
		t.Errorf("Expected 0, 1 and 0, 2 not to be empty, was empty")
	}
}

func TestDay05_MapOceanFloorVentLinesWithY(t *testing.T) {
	ventLines := []string{"0,1 -> 3,1"}

	oceanFloorMapping := mapOceanFloorVentLines(ventLines)

	if oceanFloorMapping[0][0] == 1 && oceanFloorMapping[1][0] == 1 && oceanFloorMapping[2][0] == 1 && oceanFloorMapping[3][0] == 1 {
		t.Errorf("Expected 0, 0 to be empty, was not")
	}

	if oceanFloorMapping[0][1] != 1 && oceanFloorMapping[1][1] != 1 && oceanFloorMapping[2][1] != 1 && oceanFloorMapping[3][1] != 1 {
		t.Errorf("Expected 0, 1 to 3, 1 not to be empty, was empty")
	}
}

func TestDay04Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 5

	// When
	result := GetDangerousZonesNumber(input)

	// Then
	if result != expected {
		t.Errorf("Day 4 - Expected %v, got %v", expected, result)
	}
}

func TestDay04Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 12

	// When
	result := GetDangerousZonesNumberWithDiagonals(input)

	// Then
	if result != expected {
		t.Errorf("Day 4 - Expected %v, got %v", expected, result)
	}
}

func TestDay04Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 4873

	// When
	result := GetDangerousZonesNumber(input)

	// Then
	if result != expected {
		t.Errorf("Day 4 - Expected %v, got %v", expected, result)
	}
}

func TestDay04Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 19472

	// When
	result := GetDangerousZonesNumberWithDiagonals(input)

	// Then
	if result != expected {
		t.Errorf("Day 4 - Expected %v, got %v", expected, result)
	}
}
