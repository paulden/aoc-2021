package day23

import (
	"aoc-2021/internal/pkg/utils"
	"fmt"
	"math"
)

func GetMinimalEnergyCostToOrder(input []string) int {
	burrow := parseBurrow(input)
	minimumEnergyCost := math.MaxInt

	movementsTotalCost := map[foldedBurrow]int{burrow: 0}
	nextMovements := map[foldedBurrow]int{burrow: 0}

	for len(nextMovements) > 0 {
		currentBurrow := getNextMovement(nextMovements)

		delete(nextMovements, currentBurrow)

		nextBurrows := currentBurrow.getNextBurrowsWithCost()

		for potentialBurrow, cost := range nextBurrows {
			currentTotalCost := movementsTotalCost[currentBurrow] + cost
			if pastTotalCost, ok := movementsTotalCost[potentialBurrow]; (!ok || currentTotalCost < pastTotalCost) && currentTotalCost < minimumEnergyCost {
				movementsTotalCost[potentialBurrow] = currentTotalCost
				nextMovements[potentialBurrow] = currentTotalCost
			}
			if potentialBurrow.isOrdered() && currentTotalCost < minimumEnergyCost {
				minimumEnergyCost = currentTotalCost
			}
		}
	}

	return minimumEnergyCost
}

func getNextMovement(possibleMovements map[foldedBurrow]int) foldedBurrow {
	min := math.MaxInt
	var nextBurrow foldedBurrow
	for burrow, cost := range possibleMovements {
		if cost < min {
			min = cost
			nextBurrow = burrow
		}
	}
	return nextBurrow
}

func (b *foldedBurrow) getNextBurrowsWithCost() map[foldedBurrow]int {
	nextBurrows := make(map[foldedBurrow]int)
	hallwayPossibleLocations := []int{0, 1, 3, 5, 7, 9, 10}
	sideRoomIndices := map[string]int{
		"A": 2,
		"B": 4,
		"C": 6,
		"D": 8,
	}
	costs := map[string]int{
		"A": 1,
		"B": 10,
		"C": 100,
		"D": 1000,
	}

	for i, spot := range b.hallway {
		if spot == "A" && b.canVisit(i, sideRoomIndices[spot]) && (b.hasEmptySideRoomA() || b.hasHalfCompleteSideRoomA()) {
			_, depth := getTopSpot(b.sideRoomA)
			newBurrow := b.copy()
			newBurrow.hallway[i] = ""
			newBurrow.sideRoomA[depth-1] = spot
			cost := (utils.Abs(sideRoomIndices[spot]-i) + depth) * costs[spot]
			nextBurrows[newBurrow] = cost
		}
		if spot == "B" && b.canVisit(i, sideRoomIndices[spot]) && (b.hasEmptySideRoomB() || b.hasHalfCompleteSideRoomB()) {
			_, depth := getTopSpot(b.sideRoomB)
			newBurrow := b.copy()
			newBurrow.hallway[i] = ""
			newBurrow.sideRoomB[depth-1] = spot
			cost := (utils.Abs(sideRoomIndices[spot]-i) + depth) * costs[spot]
			nextBurrows[newBurrow] = cost
		}
		if spot == "C" && b.canVisit(i, sideRoomIndices[spot]) && (b.hasEmptySideRoomC() || b.hasHalfCompleteSideRoomC()) {
			_, depth := getTopSpot(b.sideRoomC)
			newBurrow := b.copy()
			newBurrow.hallway[i] = ""
			newBurrow.sideRoomC[depth-1] = spot
			cost := (utils.Abs(sideRoomIndices[spot]-i) + depth) * costs[spot]
			nextBurrows[newBurrow] = cost
		}
		if spot == "D" && b.canVisit(i, sideRoomIndices[spot]) && (b.hasEmptySideRoomD() || b.hasHalfCompleteSideRoomD()) {
			_, depth := getTopSpot(b.sideRoomD)
			newBurrow := b.copy()
			newBurrow.hallway[i] = ""
			newBurrow.sideRoomD[depth-1] = spot
			cost := (utils.Abs(sideRoomIndices[spot]-i) + depth) * costs[spot]
			nextBurrows[newBurrow] = cost
		}
	}

	topA, depthA := getTopSpot(b.sideRoomA)
	topB, depthB := getTopSpot(b.sideRoomB)
	topC, depthC := getTopSpot(b.sideRoomC)
	topD, depthD := getTopSpot(b.sideRoomD)

	if topA != "" && (topA != "A" || !b.hasHalfCompleteSideRoomA()) {
		for _, i := range hallwayPossibleLocations {
			if b.canVisit(sideRoomIndices["A"], i) {
				newBurrow := b.copy()
				newBurrow.hallway[i] = topA
				newBurrow.sideRoomA[depthA] = ""
				cost := (utils.Abs(sideRoomIndices["A"]-i) + (depthA + 1)) * costs[topA]
				nextBurrows[newBurrow] = cost
			}
		}
	}

	if topB != "" && (topB != "B" || !b.hasHalfCompleteSideRoomB()) {
		for _, i := range hallwayPossibleLocations {
			if b.canVisit(sideRoomIndices["B"], i) {
				newBurrow := b.copy()
				newBurrow.hallway[i] = topB
				newBurrow.sideRoomB[depthB] = ""
				cost := (utils.Abs(sideRoomIndices["B"]-i) + (depthB + 1)) * costs[topB]
				nextBurrows[newBurrow] = cost
			}
		}
	}

	if topC != "" && (topC != "C" || !b.hasHalfCompleteSideRoomC()) {
		for _, i := range hallwayPossibleLocations {
			if b.canVisit(sideRoomIndices["C"], i) {
				newBurrow := b.copy()
				newBurrow.hallway[i] = topC
				newBurrow.sideRoomC[depthC] = ""
				cost := (utils.Abs(sideRoomIndices["C"]-i) + (depthC + 1)) * costs[topC]
				nextBurrows[newBurrow] = cost
			}
		}
	}

	if topD != "" && (topD != "D" || !b.hasHalfCompleteSideRoomD()) {
		for _, i := range hallwayPossibleLocations {
			if b.canVisit(sideRoomIndices["D"], i) {
				newBurrow := b.copy()
				newBurrow.hallway[i] = topD
				newBurrow.sideRoomD[depthD] = ""
				cost := (utils.Abs(sideRoomIndices["D"]-i) + (depthD + 1)) * costs[topD]
				nextBurrows[newBurrow] = cost
			}
		}
	}

	return nextBurrows
}

type foldedBurrow struct {
	hallway                                    [11]string
	sideRoomA, sideRoomB, sideRoomC, sideRoomD [2]string
}

func (b *foldedBurrow) isOrdered() bool {
	return b.hasCompleteSideRoomA() && b.hasCompleteSideRoomB() && b.hasCompleteSideRoomC() && b.hasCompleteSideRoomD()
}

func (b *foldedBurrow) hasCompleteSideRoomA() bool {
	return isSideRoomFilledWith(b.sideRoomA, "A")
}

func (b *foldedBurrow) hasCompleteSideRoomB() bool {
	return isSideRoomFilledWith(b.sideRoomB, "B")
}

func (b *foldedBurrow) hasCompleteSideRoomC() bool {
	return isSideRoomFilledWith(b.sideRoomC, "C")
}

func (b *foldedBurrow) hasCompleteSideRoomD() bool {
	return isSideRoomFilledWith(b.sideRoomD, "D")
}

func (b *foldedBurrow) hasHalfCompleteSideRoomA() bool {
	return b.sideRoomA[0] == "" && b.sideRoomA[1] == "A"
}

func (b *foldedBurrow) hasHalfCompleteSideRoomB() bool {
	return b.sideRoomB[0] == "" && b.sideRoomB[1] == "B"
}

func (b *foldedBurrow) hasHalfCompleteSideRoomC() bool {
	return b.sideRoomC[0] == "" && b.sideRoomC[1] == "C"
}

func (b *foldedBurrow) hasHalfCompleteSideRoomD() bool {
	return b.sideRoomD[0] == "" && b.sideRoomD[1] == "D"
}

func (b *foldedBurrow) hasEmptySideRoomA() bool {
	return isSideRoomFilledWith(b.sideRoomA, "")
}

func (b *foldedBurrow) hasEmptySideRoomB() bool {
	return isSideRoomFilledWith(b.sideRoomB, "")
}

func (b *foldedBurrow) hasEmptySideRoomC() bool {
	return isSideRoomFilledWith(b.sideRoomC, "")
}

func (b *foldedBurrow) hasEmptySideRoomD() bool {
	return isSideRoomFilledWith(b.sideRoomD, "")
}

func (b *foldedBurrow) canVisit(source, destination int) bool {
	if destination > source {
		for i := source + 1; i <= destination; i++ {
			if b.hallway[i] != "" {
				return false
			}
		}
	} else {
		for i := destination; i < source; i++ {
			if b.hallway[i] != "" {
				return false
			}
		}
	}
	return true
}

func (b *foldedBurrow) copy() foldedBurrow {
	return foldedBurrow{b.hallway, b.sideRoomA, b.sideRoomB, b.sideRoomC, b.sideRoomD}
}

func isSideRoomFilledWith(column [2]string, expected string) bool {
	for _, spot := range column {
		if spot != expected {
			return false
		}
	}
	return true
}

func getTopSpot(column [2]string) (string, int) {
	for depth, spot := range column {
		if spot != "" {
			return spot, depth
		}
	}
	return "", len(column)
}

func parseBurrow(input []string) foldedBurrow {
	hallway := parseHallway(input)
	sideRoomA := parseSideRoom(input, 3)
	sideRoomB := parseSideRoom(input, 5)
	sideRoomC := parseSideRoom(input, 7)
	sideRoomD := parseSideRoom(input, 9)

	return foldedBurrow{hallway, sideRoomA, sideRoomB, sideRoomC, sideRoomD}
}

func parseHallway(input []string) (hallway [11]string) {
	for i := 0; i < len(hallway); i++ {
		char := string(input[1][i+1])
		if string(char) == "." {
			hallway[i] = ""
		} else {
			hallway[i] = char
		}
	}
	return hallway
}

func parseSideRoom(input []string, sideRoomIndex int) (sideRoom [2]string) {
	for i := 0; i < len(sideRoom); i++ {
		char := string(input[i+2][sideRoomIndex])
		if string(char) == "." {
			sideRoom[i] = ""
		} else {
			sideRoom[i] = char
		}
	}
	return sideRoom
}

func (b *foldedBurrow) prettyPrint() {
	fmt.Println("#############")

	fmt.Printf("#")
	for _, spot := range b.hallway {
		if spot == "" {
			fmt.Printf(".")
		} else {
			fmt.Printf(spot)
		}
	}
	fmt.Printf("#\n")

	for stage := range b.sideRoomA {
		fmt.Printf("###")
		fmt.Printf("%v#", getAmphipodForPrinting(b.sideRoomA[stage]))
		fmt.Printf("%v#", getAmphipodForPrinting(b.sideRoomB[stage]))
		fmt.Printf("%v#", getAmphipodForPrinting(b.sideRoomC[stage]))
		fmt.Printf("%v#", getAmphipodForPrinting(b.sideRoomD[stage]))
		fmt.Printf("##\n")
	}
	fmt.Println("  #########  ")
}

func getAmphipodForPrinting(amphipod string) string {
	if amphipod == "" {
		return "."
	}
	return amphipod
}
