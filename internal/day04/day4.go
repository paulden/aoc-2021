package day04

import (
	"aoc-2021/internal/pkg/utils"
	"strconv"
	"strings"
)

// Part 1

func GetWinningBingoCardScore(file []string) int {
	sequence, bingoCards := parseBingo(file, 5, 5)

	for _, number := range sequence {
		for _, card := range bingoCards {
			card.markDrawnNumber(number)
			if card.isWinningCard() {
				return card.getRemainingNumbers() * number
			}
		}
	}
	return 0
}

// Part 2

func GetLosingBingoCardScore(file []string) int {
	sequence, bingoCards := parseBingo(file, 5, 5)
	cardsNumber := len(bingoCards)
	winningCardsNumbers := 0

	for _, number := range sequence {
		remainingBingoCards := make([]bingoCard, 0)
		for _, card := range bingoCards {
			card.markDrawnNumber(number)
			if card.isWinningCard() {
				winningCardsNumbers++
				if winningCardsNumbers == cardsNumber {
					return card.getRemainingNumbers() * number
				}
			} else {
				remainingBingoCards = append(remainingBingoCards, card)
			}
		}
		bingoCards = remainingBingoCards
	}
	return 0
}

// Data structure

type bingoCard struct {
	card         [][]int
	drawnNumbers [][]int
}

func (b *bingoCard) markDrawnNumber(drawnNumber int) {
	for i := 0; i < len(b.card); i++ {
		for j := 0; j < len(b.card[0]); j++ {
			if b.card[i][j] == drawnNumber {
				b.drawnNumbers[i][j] = 1
			}
		}
	}
}

func (b *bingoCard) isWinningCard() bool {
	for i := 0; i < len(b.card); i++ {
		rowScore := 0
		for j := 0; j < len(b.card[0]); j++ {
			rowScore += b.drawnNumbers[i][j]
		}
		if rowScore == len(b.card[0]) {
			return true
		}
	}

	for j := 0; j < len(b.card[0]); j++ {
		columnScore := 0
		for i := 0; i < len(b.card); i++ {
			columnScore += b.drawnNumbers[i][j]
		}
		if columnScore == len(b.card) {
			return true
		}
	}

	return false
}

func (b *bingoCard) getRemainingNumbers() int {
	remainingNumbers := 0
	for i := 0; i < len(b.card); i++ {
		for j := 0; j < len(b.card[0]); j++ {
			if b.drawnNumbers[i][j] == 0 {
				remainingNumbers += b.card[i][j]
			}
		}
	}
	return remainingNumbers
}

// Utils

func parseBingo(file []string, cardHeight int, cardLength int) ([]int, []bingoCard) {
	integerSequence := convertStringToIntSlice(strings.Split(file[0], ","))
	bingoCards := make([]bingoCard, 0)
	currentRowIndex := 0
	currentCard := make([][]int, 0)

	for _, line := range file[2:] {
		if line == "" {
			continue
		}

		row := convertStringToIntSlice(strings.Split(line, " "))
		currentRowIndex++
		currentCard = append(currentCard, row)

		if currentRowIndex == cardHeight {
			bingoCards = append(bingoCards, bingoCard{currentCard, createEmptyScoreCard(cardHeight, cardLength)})
			currentCard = make([][]int, 0)
			currentRowIndex = 0
		}

	}

	return integerSequence, bingoCards
}

func createEmptyScoreCard(cardHeight int, cardLength int) [][]int {
	return utils.CreateEmptyGrid(cardHeight, cardLength)
}

func convertStringToIntSlice(sequence []string) []int {
	integerSequence := make([]int, 0)
	for i, char := range sequence {
		if char == "" {
			continue
		}
		parsedCharacter, _ := strconv.Atoi(sequence[i])
		integerSequence = append(integerSequence, parsedCharacter)
	}
	return integerSequence
}
