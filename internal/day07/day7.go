package day07

import (
	"aoc-2021/internal/pkg/utils"
	"math"
	"strconv"
	"strings"
)

func GetCheapestFuelConsumption(crabPositions string) int {
	split := strings.Split(crabPositions, ",")
	positions := make([]int, len(split))
	maxPosition := 0

	for i, crabPosition := range split {
		parsedInt, _ := strconv.ParseInt(crabPosition, 10, 32)
		positions[i] = int(parsedInt)
		if positions[i] > maxPosition {
			maxPosition = positions[i]
		}
	}

	fuelConsumption := math.MaxInt
	for i := 0; i < maxPosition; i++ {
		currentConsumption := 0
		for _, position := range positions {
			currentConsumption += utils.Abs(position - i)
		}
		if currentConsumption < fuelConsumption {
			fuelConsumption = currentConsumption
		}
	}
	return fuelConsumption
}

func GetCheapestFuelConsumptionUpdated(crabPositions string) int {
	split := strings.Split(crabPositions, ",")
	positions := make([]int, len(split))
	maxPosition := 0

	for i, crabPosition := range split {
		parsedInt, _ := strconv.ParseInt(crabPosition, 10, 32)
		positions[i] = int(parsedInt)
		if positions[i] > maxPosition {
			maxPosition = positions[i]
		}
	}

	fuelConsumption := math.MaxInt
	for i := 0; i < maxPosition; i++ {
		currentConsumption := 0
		for _, position := range positions {
			currentConsumption += getNthTriangleNumber(utils.Abs(position - i))
		}
		if currentConsumption < fuelConsumption {
			fuelConsumption = currentConsumption
		}
	}
	return fuelConsumption
}

func getNthTriangleNumber(value int) int {
	return (value*value + value) / 2
}
