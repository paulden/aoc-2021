package day23

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay23ParseBurrow(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expectedHallway := [11]string{"", "", "", "", "", "", "", "", "", "", ""}
	expectedSideRoomA := [2]string{"B", "A"}
	expectedSideRoomB := [2]string{"C", "D"}
	expectedSideRoomC := [2]string{"B", "C"}
	expectedSideRoomD := [2]string{"D", "A"}
	expectedBurrow := foldedBurrow{expectedHallway, expectedSideRoomA, expectedSideRoomB, expectedSideRoomC, expectedSideRoomD}

	// When
	burrow := parseBurrow(input)

	// Then
	if expectedBurrow != burrow {
		t.Errorf("Day 23 - Parsing burrow, expected %v, got %v", expectedBurrow, burrow)
	}
}

func TestDay23ParseOngoingBurrow(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example_ongoing.txt")
	expectedHallway := [11]string{"", "B", "", "", "", "D", "", "D", "", "A", ""}
	expectedSideRoomA := [2]string{"", "A"}
	expectedSideRoomB := [2]string{"", ""}
	expectedSideRoomC := [2]string{"B", "C"}
	expectedSideRoomD := [2]string{"", "C"}
	expectedBurrow := foldedBurrow{expectedHallway, expectedSideRoomA, expectedSideRoomB, expectedSideRoomC, expectedSideRoomD}

	// When
	burrow := parseBurrow(input)

	// Then
	if expectedBurrow != burrow {
		t.Errorf("Day 23 - Parsing burrow, expected %v, got %v", expectedBurrow, burrow)
	}
}

func TestDay23ParseAmphipodsBurrow(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example_ordered.txt")
	burrow := parseBurrow(input)

	// When
	isOrdered := burrow.isOrdered()

	// Then
	if !isOrdered {
		t.Errorf("Day 23 - Expected burrow to be ordered but was not")
	}
}
func TestDay23ParseUnfoldedBurrow(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expectedHallway := [11]string{"", "", "", "", "", "", "", "", "", "", ""}
	expectedSideRoomA := [4]string{"B", "D", "D", "A"}
	expectedSideRoomB := [4]string{"C", "C", "B", "D"}
	expectedSideRoomC := [4]string{"B", "B", "A", "C"}
	expectedSideRoomD := [4]string{"D", "A", "C", "A"}
	expectedBurrow := unfoldedBurrow{expectedHallway, expectedSideRoomA, expectedSideRoomB, expectedSideRoomC, expectedSideRoomD}

	// When
	burrow := parseUnfoldedBurrow(input)

	// Then
	if expectedBurrow != burrow {
		t.Errorf("Day 23 - Parsing burrow, expected %v, got %v", expectedBurrow, burrow)
	}
}

func TestDay23Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 12521

	// When
	energyCost := GetMinimalEnergyCostToOrder(input)

	// Then
	if energyCost != expected {
		t.Errorf("Day 23 - Expected %v, got %v", expected, energyCost)
	}
}


func TestDay23Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 44169

	// When
	energyCost := GetMinimalEnergyCostToOrderPart2(input)

	// Then
	if energyCost != expected {
		t.Errorf("Day 23 - Expected %v, got %v", expected, energyCost)
	}
}

func TestDay23Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 14460

	// When
	energyCost := GetMinimalEnergyCostToOrder(input)

	// Then
	if energyCost != expected {
		t.Errorf("Day 23 - Expected %v, got %v", expected, energyCost)
	}
}

func TestDay23Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 41366

	// When
	energyCost := GetMinimalEnergyCostToOrderPart2(input)

	// Then
	if energyCost != expected {
		t.Errorf("Day 23 - Expected %v, got %v", expected, energyCost)
	}
}
