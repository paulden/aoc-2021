package main

import (
	"fmt"
	"math"
	"testing"
)

func TestDay23ParseAmphipodsBurrow(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day23_ordered.txt")
	burrow := ParseAmphipodsBurrow(input)

	// When
	isOrdered := burrow.IsOrdered()

	// Then
	if !isOrdered {
		t.Errorf("Day 23 - Expected burrow to be ordered but was not")
	}
}

func TestDay23ParseAmphipodsBurrow2(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day23_unfolded.txt")
	burrow := ParseAmphipodsBurrow2(input)

	// When
	burrow.PrettyPrint2()
}

func TestDay23Part1(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day23_second_example.txt")
	burrow := ParseAmphipodsBurrow(input)

	//burrow.PrettyPrint()

	// When
	burrow.GetNextPossibleBurrows()
}

func TestDay23CanVisitDestFromSource(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day23_second_example.txt")
	burrow := ParseAmphipodsBurrow(input)

	// When
	bToB := burrow.CanVisitDestFromSource(5, 1)
	bToA := burrow.CanVisitDestFromSource(2, 6)
	aToA := burrow.CanVisitDestFromSource(2, 10)
	dToD := burrow.CanVisitDestFromSource(8, 4)

	if !bToB {
		t.Errorf("Day 23 : expected to be possible to go from B to side room B")
	}
	if !bToA {
		t.Errorf("Day 23 : expected to be possible to go from B to side room A")
	}
	if aToA {
		t.Errorf("Day 23 : expected not to be possible to go from A to side room A")
	}
	if dToD {
		t.Errorf("Day 23 : expected not to be possible to go from D to side room D")
	}
}

func TestDay23Part1Example(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day23_example.txt")
	burrow := ParseAmphipodsBurrow(input)

	// When
	allCosts := []int{math.MaxInt}

	ints := OrderAmphipodsBurrow(burrow, 0, allCosts)
	fmt.Println(ints)
}

func TestDay23Part2Example(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day23_unfolded.txt")
	burrow := ParseAmphipodsBurrow2(input)

	// When
	allCosts := []int{math.MaxInt}

	ints := OrderAmphipodsBurrow2(burrow, 0, allCosts)
	fmt.Println(ints)
}
