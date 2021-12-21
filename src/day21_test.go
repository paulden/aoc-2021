package main

import (
	"testing"
)

func TestDay20ParsePlayerPositions(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day21_example.txt")
	expectedPlayer1StartingPosition := 4
	expectedPlayer2StartingPosition := 8

	// When
	p1, p2 := ParsePlayerPositions(input)

	// Then
	if p1 != expectedPlayer1StartingPosition {
		t.Errorf("Day 21 - parsing first player position: expected %v, got %v", expectedPlayer1StartingPosition, p1)
	}
	if p2 != expectedPlayer2StartingPosition {
		t.Errorf("Day 21 - parsing first player position: expected %v, got %v", expectedPlayer2StartingPosition, p2)
	}
}

func TestDay20Part1Example(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day21_example.txt")
	expectedResult := 739785

	// When
	result := PracticeDirac(input)

	// Then
	if result != expectedResult {
		t.Errorf("Day 21 - Part 1 example: expected %v, got %v", expectedResult, result)
	}
}

func TestDay20Part1Real(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day21.txt")
	expectedResult := 757770

	// When
	result := PracticeDirac(input)

	// Then
	if result != expectedResult {
		t.Errorf("Day 21 - Part 1 real sample: expected %v, got %v", expectedResult, result)
	}
}

func TestDay20Part2Example(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day21_example.txt")
	expectedPlayer1Victories := 444356092776315
	expectedPlayer2Victories := 341960390180808

	// When
	player1Victories, player2Victories := CountDiracVictories(input)

	// Then
	if player1Victories != expectedPlayer1Victories {
		t.Errorf("Day 21 - Part 2 example: expected player 1 to win %v times, got %v", expectedPlayer1Victories, player1Victories)
	}
	if player2Victories != expectedPlayer2Victories {
		t.Errorf("Day 21 - Part 2 example: expected player 2 to win %v times, got %v", expectedPlayer2Victories, player2Victories)
	}
}

func TestDay20Part2Real(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day21.txt")
	expectedPlayer1Victories := 712381680443927
	expectedPlayer2Victories := 540941920252956

	// When
	player1Victories, player2Victories := CountDiracVictories(input)

	// Then
	if player1Victories != expectedPlayer1Victories {
		t.Errorf("Day 21 - Part 2 real sample: expected player 1 to win %v times, got %v", expectedPlayer1Victories, player1Victories)
	}
	if player2Victories != expectedPlayer2Victories {
		t.Errorf("Day 21 - Part 2 real sample: expected player 2 to win %v times, got %v", expectedPlayer2Victories, player2Victories)
	}
}
