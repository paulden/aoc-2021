package main

import (
	"math"
	"strconv"
)

// Part 1

func GetLowestTotalRiskPath(input []string) int {
	caveRisks := ParseCaveRisks(input)
	return GetShortestPath(caveRisks)
}

// Part 2

func GetLowestTotalRiskPathInRealMap(input []string) int {
	caveRisks := TranscribeToRealMap(input)
	return GetShortestPath(caveRisks)
}

func GetShortestPath(caveRisks [][]int) int {
	// Source: https://fr.wikipedia.org/wiki/Algorithme_de_Dijkstra#Impl%C3%A9mentation_de_l'algorithme
	pathRisk := make(map[CaveCoordinates]int)
	allPositionsLeft := InitializeAllPositionsLeft(caveRisks)
	nextPositionsToVisit := map[CaveCoordinates]int{CaveCoordinates{0, 0}: 0}

	for len(allPositionsLeft) > 0 {
		currentPosition := GetNextPositionToVisit(nextPositionsToVisit)

		delete(allPositionsLeft, currentPosition)
		delete(nextPositionsToVisit, currentPosition)

		visitableNodes := GetVisitablePositions(currentPosition, caveRisks)

		for _, position := range visitableNodes {
			pathCost := pathRisk[currentPosition] + caveRisks[position.x][position.y]
			if pathCost < allPositionsLeft[position] {
				pathRisk[position] = pathCost
				allPositionsLeft[CaveCoordinates{position.x, position.y}] = pathCost
				nextPositionsToVisit[CaveCoordinates{position.x, position.y}] = pathCost
			}
		}
	}

	return pathRisk[CaveCoordinates{len(caveRisks) - 1, len(caveRisks[0]) - 1}]
}

func TranscribeToRealMap(input []string) [][]int {
	caveRisks := ParseCaveRisks(input)
	caveHeight := len(caveRisks)
	caveWidth := len(caveRisks[0])

	largerCave := make([][]int, caveHeight*5)
	for i := range largerCave {
		largerCave[i] = make([]int, caveWidth*5)
	}

	for d := 0; d < 5; d++ {
		for i := range caveRisks {
			for j := range caveRisks[0] {
				largerCave[i+d*caveHeight][j] = ResetIfAboveNine(caveRisks[i][j] + d)
			}
		}
	}

	for i := 0; i < caveHeight*5; i++ {
		for j := 0; j < caveWidth; j++ {
			for d := 1; d < 5; d++ {
				largerCave[i][j+d*caveWidth] = ResetIfAboveNine(largerCave[i][j] + d)
			}
		}
	}

	return largerCave
}

func ResetIfAboveNine(value int) int {
	if value > 9 {
		return value % 9
	}
	return value
}

func GetNextPositionToVisit(positionsToVisit map[CaveCoordinates]int) CaveCoordinates {
	min := math.MaxInt
	var nodeToVisit CaveCoordinates
	for node, cost := range positionsToVisit {
		if cost < min {
			min = cost
			nodeToVisit = node
		}
	}
	return nodeToVisit
}

func GetVisitablePositions(node CaveCoordinates, cave [][]int) []CaveCoordinates {
	visitableNodes := make([]CaveCoordinates, 0)
	topNode := CaveCoordinates{node.x - 1, node.y}
	leftNode := CaveCoordinates{node.x, node.y - 1}
	bottomNode := CaveCoordinates{node.x + 1, node.y}
	rightNode := CaveCoordinates{node.x, node.y + 1}

	if node.x > 0 {
		visitableNodes = append(visitableNodes, topNode)
	}
	if node.y > 0 {
		visitableNodes = append(visitableNodes, leftNode)
	}
	if node.x < len(cave)-1 {
		visitableNodes = append(visitableNodes, bottomNode)
	}
	if node.y < len(cave[0])-1 {
		visitableNodes = append(visitableNodes, rightNode)
	}

	return visitableNodes
}

type CaveCoordinates struct {
	x, y int
}

func ParseCaveRisks(input []string) [][]int {
	parsedCaveWalls := make([][]int, 0)

	for _, line := range input {
		parsedLine := make([]int, 0)
		for _, char := range line {
			parsedChar, _ := strconv.ParseInt(string(char), 10, 32)
			parsedLine = append(parsedLine, int(parsedChar))
		}
		parsedCaveWalls = append(parsedCaveWalls, parsedLine)
	}

	return parsedCaveWalls
}

func InitializeAllPositionsLeft(caveRisks [][]int) map[CaveCoordinates]int {
	nodesToVisit := make(map[CaveCoordinates]int)

	for i := range caveRisks {
		for j := range caveRisks[0] {
			nodesToVisit[CaveCoordinates{i, j}] = math.MaxInt
		}
	}

	nodesToVisit[CaveCoordinates{0, 0}] = 0

	return nodesToVisit
}
