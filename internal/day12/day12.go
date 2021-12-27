package day12

import (
	"strings"
	"unicode"
)

// Part 1

func CountCavePathsPart1(input []string) int {
	caveConnections := parseCaveConnections(input)
	allPaths := make([][]string, 0)
	alreadyVisitedCaves := make([]string, 0)

	allPaths = getAllPathsFromCave(
		"start",
		alreadyVisitedCaves,
		allPaths,
		caveConnections,
		isSmallCaveAlreadyVisited,
	)

	return len(allPaths)
}

// Part 2

func CountCavePathsPart2(input []string) int {
	caveConnections := parseCaveConnections(input)
	allPaths := make([][]string, 0)
	alreadyVisitedCaves := make([]string, 0)

	allPaths = getAllPathsFromCave(
		"start",
		alreadyVisitedCaves,
		allPaths,
		caveConnections,
		isSmallCaveWithVisitsQuotaExceeded,
	)

	return len(allPaths)
}

// Cave visits

type isInvalidCave func(string, []string) bool

func isSmallCaveAlreadyVisited(cave string, alreadyVisitedCaves []string) bool {
	return isSmallCave(cave) && isAlreadyVisited(cave, alreadyVisitedCaves)
}

func isSmallCaveWithVisitsQuotaExceeded(cave string, alreadyVisitedCaves []string) bool {
	return isSmallCaveAlreadyVisited(cave, alreadyVisitedCaves) && hasAlreadyVisitedASmallCaveTwice(alreadyVisitedCaves)
}

func getAllPathsFromCave(
	currentCave string,
	alreadyVisitedCaves []string,
	allPaths [][]string,
	caveConnections map[string][]string,
	isDeadend isInvalidCave,
) [][]string {
	alreadyVisitedCaves = append(alreadyVisitedCaves, currentCave)
	if currentCave == "end" {
		allPaths = append(allPaths, alreadyVisitedCaves)
		return allPaths
	}
	for _, adjacentCave := range caveConnections[currentCave] {
		if adjacentCave == "start" || isDeadend(adjacentCave, alreadyVisitedCaves) {
			continue
		}
		allPaths = getAllPathsFromCave(adjacentCave, alreadyVisitedCaves, allPaths, caveConnections, isDeadend)
	}
	return allPaths
}

func isSmallCave(cave string) bool {
	return unicode.IsLower([]rune(cave)[0])
}

func isAlreadyVisited(cave string, visitedCaves []string) bool {
	for _, visitedCave := range visitedCaves {
		if visitedCave == cave {
			return true
		}
	}
	return false
}

func hasAlreadyVisitedASmallCaveTwice(visitedCaves []string) bool {
	visitedSmallCaves := make(map[string]int)

	for _, visitedCave := range visitedCaves {
		if isSmallCave(visitedCave) {
			visitedSmallCaves[visitedCave]++
		}
		if visitedSmallCaves[visitedCave] > 1 {
			return true
		}
	}

	return false
}

// Utils

func parseCaveConnections(caveConnections []string) map[string][]string {
	cavePaths := make(map[string][]string)

	for _, connection := range caveConnections {
		split := strings.Split(connection, "-")
		cavePaths[split[0]] = append(cavePaths[split[0]], split[1])
		cavePaths[split[1]] = append(cavePaths[split[1]], split[0])
	}
	return cavePaths
}
