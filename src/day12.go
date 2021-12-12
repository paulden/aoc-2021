package main

import (
	"strings"
	"unicode"
)

// Part 1

func CountCavePathsPart1(input []string) int {
	caveConnections := ParseCaveConnections(input)
	allPaths := make([][]string, 0)
	alreadyVisitedCaves := make([]string, 0)

	allPaths = GetAllPathsFromCave(
		"start",
		alreadyVisitedCaves,
		allPaths,
		caveConnections,
		IsSmallCaveAlreadyVisited,
	)

	return len(allPaths)
}

// Part 2

func CountCavePathsPart2(input []string) int {
	caveConnections := ParseCaveConnections(input)
	allPaths := make([][]string, 0)
	alreadyVisitedCaves := make([]string, 0)

	allPaths = GetAllPathsFromCave(
		"start",
		alreadyVisitedCaves,
		allPaths,
		caveConnections,
		IsSmallCaveWithVisitsQuotaExceeded,
	)

	return len(allPaths)
}

// Cave visits

type isInvalidCave func(string, []string) bool

func IsSmallCaveAlreadyVisited(cave string, alreadyVisitedCaves []string) bool {
	return IsSmallCave(cave) && IsAlreadyVisited(cave, alreadyVisitedCaves)
}

func IsSmallCaveWithVisitsQuotaExceeded(cave string, alreadyVisitedCaves []string) bool {
	return IsSmallCaveAlreadyVisited(cave, alreadyVisitedCaves) && HasAlreadyVisitedASmallCaveTwice(alreadyVisitedCaves)
}

func GetAllPathsFromCave(
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
		allPaths = GetAllPathsFromCave(adjacentCave, alreadyVisitedCaves, allPaths, caveConnections, isDeadend)
	}
	return allPaths
}

func IsSmallCave(cave string) bool {
	return unicode.IsLower([]rune(cave)[0])
}

func IsAlreadyVisited(cave string, visitedCaves []string) bool {
	for _, visitedCave := range visitedCaves {
		if visitedCave == cave {
			return true
		}
	}
	return false
}

func HasAlreadyVisitedASmallCaveTwice(visitedCaves []string) bool {
	visitedSmallCaves := make(map[string]int)

	for _, visitedCave := range visitedCaves {
		if IsSmallCave(visitedCave) {
			visitedSmallCaves[visitedCave]++
		}
		if visitedSmallCaves[visitedCave] > 1 {
			return true
		}
	}

	return false
}

// Utils

func ParseCaveConnections(caveConnections []string) map[string][]string {
	cavePaths := make(map[string][]string)

	for _, connection := range caveConnections {
		split := strings.Split(connection, "-")
		cavePaths[split[0]] = append(cavePaths[split[0]], split[1])
		cavePaths[split[1]] = append(cavePaths[split[1]], split[0])
	}
	return cavePaths
}
