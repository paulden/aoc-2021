package day25

import "fmt"

func CountStepsBeforeImmobilisation(input []string) int {
	height, width := len(input), len(input[0])
	seaCucumbersMap := parseSeaCucumbersMap(input)
	moves := 1
	steps := 0

	for moves > 0 {
		seaCucumbersMap, moves = getNextStep(seaCucumbersMap, height, width)
		steps++
	}

	return steps
}

func getNextStep(currentState map[seaCucumbersCoordinates]int, height, width int) (map[seaCucumbersCoordinates]int, int) {
	intermediateState := make(map[seaCucumbersCoordinates]int)
	finalState := make(map[seaCucumbersCoordinates]int)

	moves := 0

	for currentSpot, orientation := range currentState {
		if orientation == 1 {
			nextSpot := seaCucumbersCoordinates{currentSpot.x, (currentSpot.y + 1) % width}
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
			nextSpot := seaCucumbersCoordinates{(currentSpot.x + 1) % height, currentSpot.y}
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

func parseSeaCucumbersMap(input []string) map[seaCucumbersCoordinates]int {
	seaCucumbersMap := make(map[seaCucumbersCoordinates]int)
	for x, line := range input {
		for y, coord := range line {
			switch string(coord) {
			case ">":
				seaCucumbersMap[seaCucumbersCoordinates{x, y}] = 1
			case "v":
				seaCucumbersMap[seaCucumbersCoordinates{x, y}] = 2
			}
		}
	}
	return seaCucumbersMap
}

func prettyPrintMap(state map[seaCucumbersCoordinates]int, height, width int) {
	intToOrientation := map[int]string{1: ">", 2: "v"}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if orientation, isOccupied := state[seaCucumbersCoordinates{x, y}]; isOccupied {
				fmt.Printf("%v", intToOrientation[orientation])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

type seaCucumbersCoordinates struct {
	x, y int
}
