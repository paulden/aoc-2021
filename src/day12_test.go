package main

import (
	"testing"
)

func TestDay12Part1StartToEnd(t *testing.T) {
	// Given
	day12 := readStringsInFile("../data/day12_start_to_end.txt")
	expected := 1

	// When
	paths := CountCavePathsPart1(day12)

	// Then
	if paths != expected {
		t.Errorf("Day 12 example - Part 1 start to end: expected %v, got %v", expected, paths)
	}
}

func TestDay12Part1StartToAToEnd(t *testing.T) {
	// Given
	day12 := readStringsInFile("../data/day12_start_to_A_to_end.txt")
	expected := 1

	// When
	paths := CountCavePathsPart1(day12)

	// Then
	if paths != expected {
		t.Errorf("Day 12 example - Part 1 start to A to end: expected %v, got %v", expected, paths)
	}
}

func TestDay12MPart1TwoCaves(t *testing.T) {
	// Given
	//     start
	//    /
	//   A ----- b
	//    \     /
	//      end
	day12 := readStringsInFile("../data/day12_two_caves.txt")
	// start - A - b - end
	// start - A - b - A - end
	// start - A - end
	expected := 3

	// When
	paths := CountCavePathsPart1(day12)

	// Then
	if paths != expected {
		t.Errorf("Day 12 example - Part 1 two caves: expected %v, got %v", expected, paths)
	}
}

func TestDay12Part1FirstExample(t *testing.T) {
	// Given
	day12 := readStringsInFile("../data/day12_simple_example.txt")
	expected := 10

	// When
	paths := CountCavePathsPart1(day12)

	// Then
	if paths != expected {
		t.Errorf("Day 12 example - Part 1 simple example: expected %v, got %v", expected, paths)
	}
}

func TestDay12Part1SecondExample(t *testing.T) {
	// Given
	day12 := readStringsInFile("../data/day12_second_example.txt")
	expected := 19

	// When
	paths := CountCavePathsPart1(day12)

	// Then
	if paths != expected {
		t.Errorf("Day 12 example - Part 1 second example: expected %v, got %v", expected, paths)
	}
}

func TestDay12Part1ThirdExample(t *testing.T) {
	// Given
	day12 := readStringsInFile("../data/day12_third_example.txt")
	expected := 226

	// When
	paths := CountCavePathsPart1(day12)

	// Then
	if paths != expected {
		t.Errorf("Day 12 example - Part 1 third example: expected %v, got %v", expected, paths)
	}
}

func TestDay12Part2ThirdExample(t *testing.T) {
	// Given
	day12 := readStringsInFile("../data/day12_third_example.txt")
	expected := 3509

	// When
	paths := CountCavePathsPart2(day12)

	// Then
	if paths != expected {
		t.Errorf("Day 12 example - Part 2 third example: expected %v, got %v", expected, paths)
	}
}
