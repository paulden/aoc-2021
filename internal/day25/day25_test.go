package day25

import (
	"aoc-2021/internal/pkg/utils"
	"fmt"
	"testing"
)

func TestDay25ParseSeaCucumbersMap(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example2.txt")
	seaCucumberEast1 := SeaCucumbersCoordinates{0, 3}
	seaCucumberEast2 := SeaCucumbersCoordinates{2, 6}
	seaCucumberEast3 := SeaCucumbersCoordinates{3, 6}
	seaCucumberEast4 := SeaCucumbersCoordinates{4, 6}
	seaCucumberSouth1 := SeaCucumbersCoordinates{3, 0}
	seaCucumberSouth2 := SeaCucumbersCoordinates{6, 2}
	seaCucumberSouth3 := SeaCucumbersCoordinates{6, 3}
	seaCucumberSouth4 := SeaCucumbersCoordinates{6, 4}

	// When
	seaCucumbersMap := ParseSeaCucumbersMap(input)

	if seaCucumbersMap[seaCucumberEast1] != 1 && seaCucumbersMap[seaCucumberEast2] != 1 && seaCucumbersMap[seaCucumberEast3] != 1 && seaCucumbersMap[seaCucumberEast4] != 1 {
		t.Errorf("Expected sea cucumbers facing east at these locations but were not")
	}
	if seaCucumbersMap[seaCucumberSouth1] != 2 && seaCucumbersMap[seaCucumberSouth2] != 2 && seaCucumbersMap[seaCucumberSouth3] != 2 && seaCucumbersMap[seaCucumberSouth4] != 2 {
		t.Errorf("Expected sea cucumbers facing south at these locations but were not")
	}
}

func TestDay25Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example2.txt")
	h, w := len(input), len(input[0])

	// When
	seaCucumbersMap := ParseSeaCucumbersMap(input)
	PrettyPrint(seaCucumbersMap, h, w)
	fmt.Printf("\n")
	seaCucumbersMap, _ = GetNextStep(seaCucumbersMap, h, w)
	PrettyPrint(seaCucumbersMap, h, w)
	fmt.Printf("\n")
	seaCucumbersMap, _ = GetNextStep(seaCucumbersMap, h, w)
	PrettyPrint(seaCucumbersMap, h, w)
	fmt.Printf("\n")
	seaCucumbersMap, _ = GetNextStep(seaCucumbersMap, h, w)
	PrettyPrint(seaCucumbersMap, h, w)
}


func TestDay25Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	h, w := len(input), len(input[0])

	// When
	seaCucumbersMap := ParseSeaCucumbersMap(input)
	moves := 1
	steps := 0

	for moves > 0 {
		seaCucumbersMap, moves = GetNextStep(seaCucumbersMap, h, w)
		steps++
	}

	fmt.Println(steps)
}
