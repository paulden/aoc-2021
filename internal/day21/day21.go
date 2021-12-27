package day21

import (
	"aoc-2021/internal/pkg/utils"
	"strconv"
	"strings"
)

// Part 1

func PracticeDirac(input []string) int {
	dice := 1
	rolls := 0
	player1Score, player2Score := 0, 0
	playerTurn := 1

	player1Position, player2Position := parsePlayerPositions(input)

	for player1Score < 1000 && player2Score < 1000 {
		movement := dice*3 + 3
		dice += 3
		rolls += 3

		if playerTurn == 1 {
			player1Position = (player1Position+movement-1)%10 + 1
			player1Score += player1Position
			playerTurn = 2
		} else if playerTurn == 2 {
			player2Position = (player2Position+movement-1)%10 + 1
			player2Score += player2Position
			playerTurn = 1
		}
	}
	if player1Score >= 1000 {
		return rolls * player2Score
	} else {
		return rolls * player1Score
	}
}

// Part 2

func GetDiracWinnerVictories(input []string) int {
	return utils.Max(countDiracVictories(input))
}

func countDiracVictories(input []string) (int, int) {
	victories := playerVictories{0, 0}
	p1InitialPosition, p2InitialPosition := parsePlayerPositions(input)
	initialUniverse := universe{p1InitialPosition, p2InitialPosition, 0, 0, 1}
	outcomePerUniverse := make(map[universe]playerVictories)

	totalVictories, _ := getVictoriesForUniverse(initialUniverse, victories, outcomePerUniverse)
	return totalVictories.p1Victories, totalVictories.p2Victories
}

func getVictoriesForUniverse(
	universe universe,
	currentVictories playerVictories,
	outcomePerUniverse map[universe]playerVictories,
) (playerVictories, map[universe]playerVictories) {
	if outcome, ok := outcomePerUniverse[universe]; ok {
		// universe has already been visited, better use some cache!
		return outcome, outcomePerUniverse
	}

	if universe.p1Score >= 21 {
		return playerVictories{1, 0}, outcomePerUniverse
	}
	if universe.p2Score >= 21 {
		return playerVictories{0, 1}, outcomePerUniverse
	}

	branchVictories := playerVictories{0, 0}
	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				totalRoll := r1 + r2 + r3
				nextUniverse := getNextAlternativeUniverse(universe, totalRoll)
				var resultVictories playerVictories
				resultVictories, outcomePerUniverse = getVictoriesForUniverse(nextUniverse, currentVictories, outcomePerUniverse)
				branchVictories.UpdateWith(resultVictories)
			}
		}
	}
	currentVictories.UpdateWith(branchVictories)

	// Saving outcome of this universe
	outcomePerUniverse[universe] = branchVictories

	return currentVictories, outcomePerUniverse
}

func getNextAlternativeUniverse(universe universe, roll int) universe {
	nextUniverse := universe
	if nextUniverse.playerTurn == 1 {
		nextUniverse.p1Position = (nextUniverse.p1Position+roll-1)%10 + 1
		nextUniverse.p1Score += nextUniverse.p1Position
		nextUniverse.playerTurn = 2
	} else if nextUniverse.playerTurn == 2 {
		nextUniverse.p2Position = (nextUniverse.p2Position+roll-1)%10 + 1
		nextUniverse.p2Score += nextUniverse.p2Position
		nextUniverse.playerTurn = 1
	}
	return nextUniverse
}

// universe State of all the Dirac game elements in a universe
type universe struct {
	p1Position, p2Position, p1Score, p2Score int
	playerTurn                               int
}

// playerVictories Possible outcomes of a game
type playerVictories struct {
	p1Victories, p2Victories int
}

func (victories *playerVictories) UpdateWith(additionalVictories playerVictories) *playerVictories {
	victories.p1Victories += additionalVictories.p1Victories
	victories.p2Victories += additionalVictories.p2Victories
	return victories
}

func parsePlayerPositions(input []string) (int, int) {
	split0 := strings.Split(input[0], ": ")
	split1 := strings.Split(input[1], ": ")

	p1, _ := strconv.ParseInt(split0[1], 10, 64)
	p2, _ := strconv.ParseInt(split1[1], 10, 64)

	return int(p1), int(p2)
}
