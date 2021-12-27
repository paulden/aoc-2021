package day25

import "fmt"

func GetNextStep(currentState map[SeaCucumbersCoordinates]int, height, width int) (map[SeaCucumbersCoordinates]int, int) {
	intermediateState := make(map[SeaCucumbersCoordinates]int)
	finalState := make(map[SeaCucumbersCoordinates]int)
	moves := 0

	for currentSpot, orientation := range currentState {
		if orientation == 1 {
			nextSpot := SeaCucumbersCoordinates{currentSpot.x, (currentSpot.y + 1) % width}
			if _, isOccupied := currentState[nextSpot]; !isOccupied {
				moves++
				intermediateState[nextSpot] = orientation
			} else {
				intermediateState[currentSpot] = orientation
			}
		} else {
			intermediateState[currentSpot] = orientation
		}
	}

	for currentSpot, orientation := range intermediateState {
		if orientation == 2 {
			nextSpot := SeaCucumbersCoordinates{(currentSpot.x + 1) % height, currentSpot.y}
			if _, isOccupied := intermediateState[nextSpot]; !isOccupied {
				finalState[nextSpot] = orientation
				moves++
			} else {
				finalState[currentSpot] = orientation
			}
		} else {
			finalState[currentSpot] = orientation
		}
	}

	return finalState, moves
}

func ParseSeaCucumbersMap(input []string) map[SeaCucumbersCoordinates]int {
	seaCucumbersMap := make(map[SeaCucumbersCoordinates]int)
	for x, line := range input {
		for y, coord := range line {
			switch string(coord) {
			case ">":
				seaCucumbersMap[SeaCucumbersCoordinates{x, y}] = 1
			case "v":
				seaCucumbersMap[SeaCucumbersCoordinates{x, y}] = 2
			}
		}
	}
	return seaCucumbersMap
}

func PrettyPrint(state map[SeaCucumbersCoordinates]int, height, width int) {
	intToOrientation := map[int]string{1: ">", 2: "v"}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if orientation, isOccupied := state[SeaCucumbersCoordinates{x, y}]; isOccupied {
				fmt.Printf("%v", intToOrientation[orientation])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

type SeaCucumbersCoordinates struct {
	x, y int
}
