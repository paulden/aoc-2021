package day15

import (
	"math"
	"strconv"
)

// Part 1

func GetLowestTotalRiskPath(input []string) int {
	caveRisks := parseCaveRisks(input)
	return getShortestPath(caveRisks)
}

// Part 2

func GetLowestTotalRiskPathInRealMap(input []string) int {
	caveRisks := transcribeToRealMap(input)
	return getShortestPath(caveRisks)
}

func getShortestPath(caveRisks [][]int) int {
	// Source: https://fr.wikipedia.org/wiki/Algorithme_de_Dijkstra#Impl%C3%A9mentation_de_l'algorithme
	pathRisk := make(map[caveCoordinates]int)
	allPositionsLeft := InitializeAllPositionsLeft(caveRisks)
	nextPositionsToVisit := map[caveCoordinates]int{caveCoordinates{0, 0}: 0}

	for len(allPositionsLeft) > 0 {
		currentPosition := getNextPositionToVisit(nextPositionsToVisit)

		delete(allPositionsLeft, currentPosition)
		delete(nextPositionsToVisit, currentPosition)

		visitableNodes := getVisitablePositions(currentPosition, caveRisks)

		for _, position := range visitableNodes {
			pathCost := pathRisk[currentPosition] + caveRisks[position.x][position.y]
			if pathCost < allPositionsLeft[position] {
				pathRisk[position] = pathCost
				allPositionsLeft[caveCoordinates{position.x, position.y}] = pathCost
				nextPositionsToVisit[caveCoordinates{position.x, position.y}] = pathCost
			}
		}
	}

	return pathRisk[caveCoordinates{len(caveRisks) - 1, len(caveRisks[0]) - 1}]
}

func transcribeToRealMap(input []string) [][]int {
	caveRisks := parseCaveRisks(input)
	caveHeight := len(caveRisks)
	caveWidth := len(caveRisks[0])

	largerCave := make([][]int, caveHeight*5)
	for i := range largerCave {
		largerCave[i] = make([]int, caveWidth*5)
	}

	for d := 0; d < 5; d++ {
		for i := range caveRisks {
			for j := range caveRisks[0] {
				largerCave[i+d*caveHeight][j] = resetIfAboveNine(caveRisks[i][j] + d)
			}
		}
	}

	for i := 0; i < caveHeight*5; i++ {
		for j := 0; j < caveWidth; j++ {
			for d := 1; d < 5; d++ {
				largerCave[i][j+d*caveWidth] = resetIfAboveNine(largerCave[i][j] + d)
			}
		}
	}

	return largerCave
}

func resetIfAboveNine(value int) int {
	if value > 9 {
		return value % 9
	}
	return value
}

func getNextPositionToVisit(positionsToVisit map[caveCoordinates]int) caveCoordinates {
	min := math.MaxInt
	var nodeToVisit caveCoordinates
	for node, cost := range positionsToVisit {
		if cost < min {
			min = cost
			nodeToVisit = node
		}
	}
	return nodeToVisit
}

func getVisitablePositions(node caveCoordinates, cave [][]int) []caveCoordinates {
	visitableNodes := make([]caveCoordinates, 0)
	topNode := caveCoordinates{node.x - 1, node.y}
	leftNode := caveCoordinates{node.x, node.y - 1}
	bottomNode := caveCoordinates{node.x + 1, node.y}
	rightNode := caveCoordinates{node.x, node.y + 1}

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

type caveCoordinates struct {
	x, y int
}

func parseCaveRisks(input []string) [][]int {
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

func InitializeAllPositionsLeft(caveRisks [][]int) map[caveCoordinates]int {
	nodesToVisit := make(map[caveCoordinates]int)

	for i := range caveRisks {
		for j := range caveRisks[0] {
			nodesToVisit[caveCoordinates{i, j}] = math.MaxInt
		}
	}

	nodesToVisit[caveCoordinates{0, 0}] = 0

	return nodesToVisit
}
