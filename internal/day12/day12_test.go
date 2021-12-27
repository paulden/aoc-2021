package day12

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay12Part1StartToEnd(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/start_to_end.txt")
	expected := 1

	// When
	paths := CountCavePathsPart1(input)

	// Then
	if paths != expected {
		t.Errorf("Day 12 - Expected %v, got %v", expected, paths)
	}
}

func TestDay12Part1StartToAToEnd(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/start_to_A_to_end.txt")
	expected := 1

	// When
	paths := CountCavePathsPart1(input)

	// Then
	if paths != expected {
		t.Errorf("Day 12 - Expected %v, got %v", expected, paths)
	}
}

func TestDay12Part1TwoCaves(t *testing.T) {
	// Given
	//     start
	//    /
	//   A ----- b
	//    \     /
	//      end
	input := utils.ReadStringsInFile("testdata/two_caves.txt")
	// start - A - b - end
	// start - A - b - A - end
	// start - A - end
	expected := 3

	// When
	paths := CountCavePathsPart1(input)

	// Then
	if paths != expected {
		t.Errorf("Day 12 - Expected %v, got %v", expected, paths)
	}
}

func TestDay12Part1FirstExample(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/simple_example.txt")
	expected := 10

	// When
	paths := CountCavePathsPart1(input)

	// Then
	if paths != expected {
		t.Errorf("Day 12 - Expected %v, got %v", expected, paths)
	}
}

func TestDay12Part1SecondExample(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/second_example.txt")
	expected := 19

	// When
	paths := CountCavePathsPart1(input)

	// Then
	if paths != expected {
		t.Errorf("Day 12 - Expected %v, got %v", expected, paths)
	}
}

func TestDay12Part1ThirdExample(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/third_example.txt")
	expected := 226

	// When
	paths := CountCavePathsPart1(input)

	// Then
	if paths != expected {
		t.Errorf("Day 12 - Expected %v, got %v", expected, paths)
	}
}

func TestDay12Part2ThirdExample(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/third_example.txt")
	expected := 3509

	// When
	paths := CountCavePathsPart2(input)

	// Then
	if paths != expected {
		t.Errorf("Day 12 - Expected %v, got %v", expected, paths)
	}
}

func TestDay12Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 4573

	// When
	paths := CountCavePathsPart1(input)

	// Then
	if paths != expected {
		t.Errorf("Day 12 - Expected %v, got %v", expected, paths)
	}
}

func TestDay12Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 117509

	// When
	paths := CountCavePathsPart2(input)

	// Then
	if paths != expected {
		t.Errorf("Day 12 - Expected %v, got %v", expected, paths)
	}
}
