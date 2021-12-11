package main

import (
	"strconv"
)

// Part 1

func CountFlashes(grid []string) int {
	parsedGrid := ParseOctopusesGrid(grid)
	totalFlashes := 0

	for i := 0; i < 100; i++ {
		newGrid, flashes := PlaySingleStep(parsedGrid)
		parsedGrid = newGrid
		totalFlashes += flashes
	}

	return totalFlashes
}

// Part 2

func FindMindBlowingStep(grid []string) int {
	parsedGrid := ParseOctopusesGrid(grid)
	currentStep := 0
	currentFlashes := 0

	for currentFlashes != 100 {
		parsedGrid, currentFlashes = PlaySingleStep(parsedGrid)
		currentStep++
	}

	return currentStep
}

func PlaySingleStep(grid [][]int) ([][]int, int) {
	newGrid := grid
	flashesGrid := CreateEmptyGrid(len(grid), len(grid[0]))
	flashes := 0
	iMax := len(grid) - 1
	jMax := len(grid[0]) - 1

	for i, line := range newGrid {
		for j, _ := range line {
			newGrid[i][j]++
		}
	}

	for true {
		currentFlashes := 0
		for i, line := range newGrid {
			for j, octopus := range line {
				if octopus >= 10 && flashesGrid[i][j] == 0 {
					if j < jMax {
						newGrid[i][j+1]++
					}
					if j > 0 {
						newGrid[i][j-1]++
					}
					if i < iMax {
						newGrid[i+1][j]++
					}
					if i > 0 {
						newGrid[i-1][j]++
					}
					if i < iMax && j < jMax {
						newGrid[i+1][j+1]++
					}
					if i > 0 && j > 0 {
						newGrid[i-1][j-1]++
					}
					if i < iMax && j > 0 {
						newGrid[i+1][j-1]++
					}
					if i > 0 && j < jMax {
						newGrid[i-1][j+1]++
					}

					flashesGrid[i][j] = 1
					currentFlashes++


				}
			}
		}
		flashes += currentFlashes
		if currentFlashes == 0 {
			break
		}
	}

	for i, line := range newGrid {
		for j, octopus := range line {
			if octopus >= 10 {
				newGrid[i][j] = 0
			}
		}
	}

	return newGrid, flashes
}

// Utils

func ParseOctopusesGrid(grid []string) [][]int {
	parsedGrid := make([][]int, 0)

	for _, line := range grid {
		parsedLine := make([]int, 0)
		for _, char := range line {
			parsedChar, _ := strconv.ParseInt(string(char), 10, 32)
			parsedLine = append(parsedLine, int(parsedChar))
		}
		parsedGrid = append(parsedGrid, parsedLine)
	}

	return parsedGrid
}

func CreateEmptyGrid(height int, length int) [][]int {
	emptyGrid := make([][]int, height)
	for i := 0; i < height; i++ {
		emptyGrid[i] = make([]int, length)
	}
	return emptyGrid
}
