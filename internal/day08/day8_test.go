package day08

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay08Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 26

	// When
	uniqueSegmentsDigits := CountUniqueSegmentsDigits(input)

	// Then
	if uniqueSegmentsDigits != expected {
		t.Errorf("Day 8 - Expected %v, got %v", expected, uniqueSegmentsDigits)
	}
}

func TestDay08Part2ExampleSingleLine(t *testing.T) {
	// Given
	inputDigits := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	allPossibleMappings := generateAllPossibleMappings()
	expected := 5353

	// When
	result := parseSevenSegmentDisplay(inputDigits, allPossibleMappings)

	// Then
	if result != expected {
		t.Errorf("Day 8 - Expected %v, got %v", expected, result)
	}
}

func TestDay08Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 61229

	// When
	sum := SumOutputDisplays(input)

	// Then
	if sum != expected {
		t.Errorf("Day 8 - Expected %v, got %v", expected, sum)
	}
}

func TestDay08Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 530

	// When
	uniqueSegmentsDigits := CountUniqueSegmentsDigits(input)

	// Then
	if uniqueSegmentsDigits != expected {
		t.Errorf("Day 8 - Expected %v, got %v", expected, uniqueSegmentsDigits)
	}
}

func TestDay08Part2(t *testing.T) {
	// Given
	day8 := utils.ReadStringsInFile("testdata/input.txt")
	expected := 1051087

	// When
	sum := SumOutputDisplays(day8)

	// Then
	if sum != expected {
		t.Errorf("Day 8 - Expected %v, got %v", expected, sum)
	}
}

func TestDay08Part2Mapping(t *testing.T) {
	// Given
	inputDigits := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"
	allPossibleMappings := generateAllPossibleMappings()

	// When
	mapping := getSevenSegmentsMapping(inputDigits, allPossibleMappings)

	// True
	//  aaaa
	// b    c
	// b    c
	//  dddd
	// e    f
	// e    f
	//  gggg

	// Altered
	// dddd
	//e    a
	//e    a
	// ffff
	//g    b
	//g    b
	// cccc

	// Then
	if mapping["d"] != "a" || mapping["e"] != "b" || mapping["a"] != "c" || mapping["f"] != "d" || mapping["g"] != "e" || mapping["b"] != "f" || mapping["c"] != "g" {
		t.Errorf("Day 8 mapping is invalid!")
	}
}

func Test_newSevenSegmentDigit(t *testing.T) {
	// Given
	onePattern := "cf"

	// When
	oneDigit := newSevenSegmentDigit(onePattern)

	// Then
	if !oneDigit.digits["c"] || !oneDigit.digits["f"] {
		t.Errorf("Segments of one digit should be up, were not")
	}
	if oneDigit.digits["a"] || oneDigit.digits["b"] || oneDigit.digits["d"] || oneDigit.digits["e"] || oneDigit.digits["g"] {
		t.Errorf("Segments of one digit should be up, were not")
	}
}

func Test_isEqual(t *testing.T) {
	type args struct {
		digit1 sevenSegmentDigit
		digit2 sevenSegmentDigit
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Passing", args{newSevenSegmentDigit("ab"), newSevenSegmentDigit("ab")}, true},
		{"Failing", args{newSevenSegmentDigit("ab"), newSevenSegmentDigit("fab")}, false},
		{"Passing order", args{newSevenSegmentDigit("abcde"), newSevenSegmentDigit("decba")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.digit1.isEqual(tt.args.digit2); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
