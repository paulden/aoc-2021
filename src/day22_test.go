package main

import (
	"testing"
)

func TestDay22ParseRebootSteps(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day22_simple_example.txt")
	expectedStepsNumber := 4
	firstCuboid := Cuboid{10, 12, 10, 12, 10, 12}
	thirdCuboid := Cuboid{9, 11, 9, 11, 9, 11}
	expectedFirstReactorInstruction := ReactorInstruction{true, firstCuboid}
	expectedThirdReactorInstruction := ReactorInstruction{false, thirdCuboid}

	// When
	rebootSteps := ParseReactorInstructions(input)

	// Then
	if len(rebootSteps) != expectedStepsNumber {
		t.Errorf("Day 22 - parsing reboot steps: expected %v, got %v", expectedStepsNumber, len(rebootSteps))
	}
	if rebootSteps[0] != expectedFirstReactorInstruction {
		t.Errorf("Day 22 - parsing first reboot step: expected %v, got %v", expectedFirstReactorInstruction, rebootSteps[0])
	}
	if rebootSteps[2] != expectedThirdReactorInstruction {
		t.Errorf("Day 22 - parsing third reboot step: expected %v, got %v", expectedThirdReactorInstruction, rebootSteps[2])
	}
}

func TestDay22ParseRebootStepsNegative(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day22_example.txt")
	expectedStepsNumber := 22
	firstCuboid := Cuboid{-20, 26, -36, 17, -47, 7}
	secondCuboid := Cuboid{-20, 33, -21, 23, -26, 28}
	expectedFirstReactorInstruction := ReactorInstruction{true, firstCuboid}
	expectedSecondReactorInstruction := ReactorInstruction{true, secondCuboid}

	// When
	rebootSteps := ParseReactorInstructions(input)

	// Then
	if len(rebootSteps) != expectedStepsNumber {
		t.Errorf("Day 22 - parsing reboot steps: expected %v, got %v", expectedStepsNumber, len(rebootSteps))
	}
	if rebootSteps[0] != expectedFirstReactorInstruction {
		t.Errorf("Day 22 - parsing first reboot step: expected %v, got %v", expectedFirstReactorInstruction, rebootSteps[0])
	}
	if rebootSteps[1] != expectedSecondReactorInstruction {
		t.Errorf("Day 22 - parsing second reboot step: expected %v, got %v", expectedSecondReactorInstruction, rebootSteps[2])
	}
}

func Test22Part1SimpleExample(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day22_simple_example.txt")
	expectedCubesOn := 39

	// When
	cubesOn := CountCubesOnNaive(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 Part 1 - example: expected %v, got %v", expectedCubesOn, cubesOn)
	}
}

func Test22Part1Example(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day22_example.txt")
	expectedCubesOn := 590784

	// When
	cubesOn := CountCubesOnNaive(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 Part 1 - example: expected %v, got %v", expectedCubesOn, cubesOn)
	}
}

func Test22Part1Real(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day22.txt")
	expectedCubesOn := 655005

	// When
	cubesOn := CountCubesOnNaive(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 Part 1 - real sample: expected %v, got %v", expectedCubesOn, cubesOn)
	}
}

func TestDay22IntersectionWhenNoOverlap(t *testing.T) {
	// Given
	c1 := Cuboid{0, 2, 0, 2, 0, 2}
	c2 := Cuboid{3, 5, 3, 5, 3, 5}

	// When
	_, err := c1.Intersection(c2)

	// Then
	if err == nil {
		t.Errorf("Day 22 - Cuboid intersection, expected error but got nil")
	}
}

func TestDay22IntersectionWhenOverlap(t *testing.T) {
	// Given
	c1 := Cuboid{0, 2, 0, 2, 0, 2}
	c2 := Cuboid{1, 5, 1, 5, 1, 5}
	expectedIntersection := Cuboid{1, 2, 1, 2, 1, 2}

	// When
	intersection, err := c1.Intersection(c2)

	// Then
	if err != nil {
		t.Errorf("Day 22 - Cuboid intersection, expected no error but got one")
	}
	if *intersection != expectedIntersection {
		t.Errorf("Day 22 - Cuboid intersection, expected %v, got %v", expectedIntersection, intersection)
	}
}

func Test22Part2SimpleExample(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day22_simple_example.txt")
	expectedCubesOn := 39

	// When
	cubesOn := CountCubesOnOptimized(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 Part 1 - example: expected %v, got %v", expectedCubesOn, cubesOn)
	}
}

func Test22Part2Example(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day22_example_part2.txt")
	expectedCubesOn := 2758514936282235

	// When
	cubesOn := CountCubesOnOptimized(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 Part 2 - example: expected %v, got %v", expectedCubesOn, cubesOn)
	}
}

func Test22Part2Real(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day22.txt")
	expectedCubesOn := 1125649856443608

	// When
	cubesOn := CountCubesOnOptimized(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 Part 2 - real sample: expected %v, got %v", expectedCubesOn, cubesOn)
	}
}
