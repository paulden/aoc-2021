package day23

import (
	"aoc-2021/internal/pkg/utils"
	"fmt"
	"math"
)

func GetMinimalEnergyCostToOrderPart2(input []string) int {
	burrow := parseUnfoldedBurrow(input)
	minimumEnergyCost := math.MaxInt

	movementsTotalCost := map[unfoldedBurrow]int{burrow: 0}
	nextMovements := map[unfoldedBurrow]int{burrow: 0}

	for len(nextMovements) > 0 {
		currentBurrow := getNextMovementPart2(nextMovements)

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

func getNextMovementPart2(possibleMovements map[unfoldedBurrow]int) unfoldedBurrow {
	min := math.MaxInt
	var nextBurrow unfoldedBurrow
	for burrow, cost := range possibleMovements {
		if cost < min {
			min = cost
			nextBurrow = burrow
		}
	}
	return nextBurrow
}

func (b *unfoldedBurrow) getNextBurrowsWithCost() map[unfoldedBurrow]int {
	nextBurrows := make(map[unfoldedBurrow]int)
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
			_, depth := getTopSpotUnfolded(b.sideRoomA)
			newBurrow := b.copy()
			newBurrow.hallway[i] = ""
			newBurrow.sideRoomA[depth-1] = spot
			cost := (utils.Abs(sideRoomIndices[spot]-i) + depth) * costs[spot]
			nextBurrows[newBurrow] = cost
		}
		if spot == "B" && b.canVisit(i, sideRoomIndices[spot]) && (b.hasEmptySideRoomB() || b.hasHalfCompleteSideRoomB()) {
			_, depth := getTopSpotUnfolded(b.sideRoomB)
			newBurrow := b.copy()
			newBurrow.hallway[i] = ""
			newBurrow.sideRoomB[depth-1] = spot
			cost := (utils.Abs(sideRoomIndices[spot]-i) + depth) * costs[spot]
			nextBurrows[newBurrow] = cost
		}
		if spot == "C" && b.canVisit(i, sideRoomIndices[spot]) && (b.hasEmptySideRoomC() || b.hasHalfCompleteSideRoomC()) {
			_, depth := getTopSpotUnfolded(b.sideRoomC)
			newBurrow := b.copy()
			newBurrow.hallway[i] = ""
			newBurrow.sideRoomC[depth-1] = spot
			cost := (utils.Abs(sideRoomIndices[spot]-i) + depth) * costs[spot]
			nextBurrows[newBurrow] = cost
		}
		if spot == "D" && b.canVisit(i, sideRoomIndices[spot]) && (b.hasEmptySideRoomD() || b.hasHalfCompleteSideRoomD()) {
			_, depth := getTopSpotUnfolded(b.sideRoomD)
			newBurrow := b.copy()
			newBurrow.hallway[i] = ""
			newBurrow.sideRoomD[depth-1] = spot
			cost := (utils.Abs(sideRoomIndices[spot]-i) + depth) * costs[spot]
			nextBurrows[newBurrow] = cost
		}
	}

	topA, depthA := getTopSpotUnfolded(b.sideRoomA)
	topB, depthB := getTopSpotUnfolded(b.sideRoomB)
	topC, depthC := getTopSpotUnfolded(b.sideRoomC)
	topD, depthD := getTopSpotUnfolded(b.sideRoomD)

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

type unfoldedBurrow struct {
	hallway                                    [11]string
	sideRoomA, sideRoomB, sideRoomC, sideRoomD [4]string
}

func (b *unfoldedBurrow) isOrdered() bool {
	return b.hasCompleteSideRoomA() && b.hasCompleteSideRoomB() && b.hasCompleteSideRoomC() && b.hasCompleteSideRoomD()
}

func (b *unfoldedBurrow) hasCompleteSideRoomA() bool {
	return isUnfoldedSideRoomFilledWith(b.sideRoomA, "A")
}

func (b *unfoldedBurrow) hasCompleteSideRoomB() bool {
	return isUnfoldedSideRoomFilledWith(b.sideRoomB, "B")
}

func (b *unfoldedBurrow) hasCompleteSideRoomC() bool {
	return isUnfoldedSideRoomFilledWith(b.sideRoomC, "C")
}

func (b *unfoldedBurrow) hasCompleteSideRoomD() bool {
	return isUnfoldedSideRoomFilledWith(b.sideRoomD, "D")
}

func (b *unfoldedBurrow) hasHalfCompleteSideRoomA() bool {
	_, depth := getTopSpotUnfolded(b.sideRoomA)
	if depth == 0 {
		return false
	}
	for i := depth; i < len(b.sideRoomA); i++ {
		if b.sideRoomA[i] != "A" {
			return false
		}
	}
	return true
}

func (b *unfoldedBurrow) hasHalfCompleteSideRoomB() bool {
	_, depth := getTopSpotUnfolded(b.sideRoomB)
	if depth == 0 {
		return false
	}
	for i := depth; i < len(b.sideRoomB); i++ {
		if b.sideRoomB[i] != "B" {
			return false
		}
	}
	return true
}

func (b *unfoldedBurrow) hasHalfCompleteSideRoomC() bool {
	_, depth := getTopSpotUnfolded(b.sideRoomC)
	if depth == 0 {
		return false
	}
	for i := depth; i < len(b.sideRoomC); i++ {
		if b.sideRoomC[i] != "C" {
			return false
		}
	}
	return true
}

func (b *unfoldedBurrow) hasHalfCompleteSideRoomD() bool {
	_, depth := getTopSpotUnfolded(b.sideRoomD)
	if depth == 0 {
		return false
	}
	for i := depth; i < len(b.sideRoomD); i++ {
		if b.sideRoomD[i] != "D" {
			return false
		}
	}
	return true
}

func (b *unfoldedBurrow) hasEmptySideRoomA() bool {
	return isUnfoldedSideRoomFilledWith(b.sideRoomA, "")
}

func (b *unfoldedBurrow) hasEmptySideRoomB() bool {
	return isUnfoldedSideRoomFilledWith(b.sideRoomB, "")
}

func (b *unfoldedBurrow) hasEmptySideRoomC() bool {
	return isUnfoldedSideRoomFilledWith(b.sideRoomC, "")
}

func (b *unfoldedBurrow) hasEmptySideRoomD() bool {
	return isUnfoldedSideRoomFilledWith(b.sideRoomD, "")
}

func (b *unfoldedBurrow) canVisit(source, destination int) bool {
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

func (b *unfoldedBurrow) copy() unfoldedBurrow {
	return unfoldedBurrow{b.hallway, b.sideRoomA, b.sideRoomB, b.sideRoomC, b.sideRoomD}
}

func isUnfoldedSideRoomFilledWith(column [4]string, expected string) bool {
	for _, spot := range column {
		if spot != expected {
			return false
		}
	}
	return true
}

func getTopSpotUnfolded(column [4]string) (string, int) {
	for depth, spot := range column {
		if spot != "" {
			return spot, depth
		}
	}
	return "", len(column)
}

func parseUnfoldedBurrow(input []string) unfoldedBurrow {
	hallway := parseHallway(input)
	sideRoomA := parseUnfoldedSideRoom(input, 3)
	sideRoomB := parseUnfoldedSideRoom(input, 5)
	sideRoomC := parseUnfoldedSideRoom(input, 7)
	sideRoomD := parseUnfoldedSideRoom(input, 9)

	return unfoldedBurrow{hallway, sideRoomA, sideRoomB, sideRoomC, sideRoomD}
}

func parseUnfoldedSideRoom(input []string, sideRoomIndex int) (sideRoom [4]string) {
	topChar := string(input[2][sideRoomIndex])
	bottomChar := string(input[3][sideRoomIndex])
	if string(topChar) == "." {
		sideRoom[0] = ""
	} else {
		sideRoom[0] = topChar
	}
	if sideRoomIndex == 3 {
		sideRoom[1], sideRoom[2] = "D", "D"
	} else if sideRoomIndex == 5 {
		sideRoom[1], sideRoom[2] = "C", "B"
	}else if sideRoomIndex == 7 {
		sideRoom[1], sideRoom[2] = "B", "A"
	} else if sideRoomIndex == 9 {
		sideRoom[1], sideRoom[2] = "A", "C"
	}
	if string(bottomChar) == "." {
		sideRoom[3] = ""
	} else {
		sideRoom[3] = bottomChar
	}
	return sideRoom
}

func (b *unfoldedBurrow) prettyPrint() {
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

