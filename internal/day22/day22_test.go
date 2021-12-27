package day22

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay22ParseRebootSteps(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example1.txt")
	expectedStepsNumber := 4
	firstCuboid := cuboid{10, 12, 10, 12, 10, 12}
	thirdCuboid := cuboid{9, 11, 9, 11, 9, 11}
	expectedFirstReactorInstruction := reactorInstruction{true, firstCuboid}
	expectedThirdReactorInstruction := reactorInstruction{false, thirdCuboid}

	// When
	rebootSteps := parseReactorInstructions(input)

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
	input := utils.ReadStringsInFile("testdata/example2.txt")
	expectedStepsNumber := 22
	firstCuboid := cuboid{-20, 26, -36, 17, -47, 7}
	secondCuboid := cuboid{-20, 33, -21, 23, -26, 28}
	expectedFirstReactorInstruction := reactorInstruction{true, firstCuboid}
	expectedSecondReactorInstruction := reactorInstruction{true, secondCuboid}

	// When
	rebootSteps := parseReactorInstructions(input)

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

func TestDay22Part1SimpleExample(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example1.txt")
	expectedCubesOn := 39

	// When
	cubesOn := CountCubesOnNaive(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 - Expected %v, got %v", expectedCubesOn, cubesOn)
	}
}

func TestDay22Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example2.txt")
	expectedCubesOn := 590784

	// When
	cubesOn := CountCubesOnNaive(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 - Expected %v, got %v", expectedCubesOn, cubesOn)
	}
}

func TestDay22IntersectionWhenNoOverlap(t *testing.T) {
	// Given
	c1 := cuboid{0, 2, 0, 2, 0, 2}
	c2 := cuboid{3, 5, 3, 5, 3, 5}

	// When
	_, err := c1.getIntersectionWith(c2)

	// Then
	if err == nil {
		t.Errorf("Day 22 - cuboid intersection, expected error but got nil")
	}
}

func TestDay22IntersectionWhenOverlap(t *testing.T) {
	// Given
	c1 := cuboid{0, 2, 0, 2, 0, 2}
	c2 := cuboid{1, 5, 1, 5, 1, 5}
	expectedIntersection := cuboid{1, 2, 1, 2, 1, 2}

	// When
	intersection, err := c1.getIntersectionWith(c2)

	// Then
	if err != nil {
		t.Errorf("Day 22 - cuboid intersection, expected no error but got one")
	}
	if *intersection != expectedIntersection {
		t.Errorf("Day 22 - cuboid intersection, expected %v, got %v", expectedIntersection, intersection)
	}
}

func TestDay22Part2SimpleExample(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example1.txt")
	expectedCubesOn := 39

	// When
	cubesOn := CountCubesOnOptimized(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 - Expected %v, got %v", expectedCubesOn, cubesOn)
	}
}

func TestDay22Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example3.txt")
	expectedCubesOn := 2758514936282235

	// When
	cubesOn := CountCubesOnOptimized(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 - Expected %v, got %v", expectedCubesOn, cubesOn)
	}
}

func TestDay22Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expectedCubesOn := 655005

	// When
	cubesOn := CountCubesOnNaive(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 - Expected %v, got %v", expectedCubesOn, cubesOn)
	}
}

func TestDay22Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expectedCubesOn := 1125649856443608

	// When
	cubesOn := CountCubesOnOptimized(input)

	// Then
	if cubesOn != expectedCubesOn {
		t.Errorf("Day 22 - Expected %v, got %v", expectedCubesOn, cubesOn)
	}
}
