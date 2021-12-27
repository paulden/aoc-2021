package day04

import (
	"aoc-2021/internal/pkg/utils"
	"reflect"
	"testing"
)

func TestDay04ParseBingo(t *testing.T) {
	firstBingoCard := bingoCard{[][]int{{1, 2, 3}, {3, 4, 5}}, createEmptyScoreCard(2, 3)}
	secondBingoCard := bingoCard{[][]int{{6, 7, 8}, {9, 10, 11}}, createEmptyScoreCard(2, 3)}

	type args struct {
		file       []string
		cardHeight int
		cardLength int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []bingoCard
	}{
		{"Simple parsing", args{[]string{"12,1", "", " 1 2 3", "3 4 5"}, 2, 3}, []int{12, 1}, []bingoCard{firstBingoCard}},
		{"More complex parsing", args{[]string{"12,1,5,4", "", " 1 2 3", " 3 4 5", "", " 6 7 8", " 9 10 11", ""}, 2, 3}, []int{12, 1, 5, 4}, []bingoCard{firstBingoCard, secondBingoCard}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseBingo(tt.args.file, tt.args.cardHeight, tt.args.cardLength)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBingo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseBingo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDay04_bingoCard_markDrawnNumber(t *testing.T) {
	firstBingoCard := bingoCard{[][]int{{1, 2, 3}, {3, 4, 5}}, createEmptyScoreCard(2, 3)}
	secondBingoCard := bingoCard{[][]int{{6, 7, 8}, {9, 10, 11}}, createEmptyScoreCard(2, 3)}

	firstBingoCard.markDrawnNumber(2)
	secondBingoCard.markDrawnNumber(11)

	if firstBingoCard.drawnNumbers[0][1] != 1 {
		t.Errorf("Expected %v, got %v", 1, firstBingoCard.drawnNumbers[0][1])
	}
	if secondBingoCard.drawnNumbers[1][2] != 1 {
		t.Errorf("Expected %v, got %v", 1, secondBingoCard.drawnNumbers[0][1])
	}
}

func TestDay04_bingoCard_isWinningCard(t *testing.T) {
	firstBingoCard := bingoCard{[][]int{{1, 2, 3}, {3, 4, 5}}, createEmptyScoreCard(2, 3)}
	secondBingoCard := bingoCard{[][]int{{6, 7, 8}, {9, 10, 11}}, createEmptyScoreCard(2, 3)}

	firstBingoCard.markDrawnNumber(2)
	secondBingoCard.markDrawnNumber(2)

	if firstBingoCard.isWinningCard() || secondBingoCard.isWinningCard() {
		t.Errorf("Expected not to be winning, won")
	}

	firstBingoCard.markDrawnNumber(5)
	secondBingoCard.markDrawnNumber(5)

	if firstBingoCard.isWinningCard() || secondBingoCard.isWinningCard() {
		t.Errorf("Expected not to be winning, won")
	}

	firstBingoCard.markDrawnNumber(4)
	secondBingoCard.markDrawnNumber(4)

	if secondBingoCard.isWinningCard() {
		t.Errorf("Expected not to be winning, won")
	}
	if !firstBingoCard.isWinningCard() {
		t.Errorf("Expected to be winning, did not win")
	}
}

func TestDay04Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 4512

	// When
	result := GetWinningBingoCardScore(input)

	// Then
	if result != expected {
		t.Errorf("Day 4 - Expected %v, got %v", expected, result)
	}
}

func TestDay04Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expected := 1924

	// When
	result := GetLosingBingoCardScore(input)

	// Then
	if result != expected {
		t.Errorf("Day 4 - Expected %v, got %v", expected, result)
	}
}

func TestDay04Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 44736

	// When
	result := GetWinningBingoCardScore(input)

	// Then
	if result != expected {
		t.Errorf("Day 4 - Expected %v, got %v", expected, result)
	}
}

func TestDay04Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 1827

	// When
	result := GetLosingBingoCardScore(input)

	// Then
	if result != expected {
		t.Errorf("Day 4 - Expected %v, got %v", expected, result)
	}
}
