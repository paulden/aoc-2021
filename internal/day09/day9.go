package day09

import (
	"strconv"
)

// Part 1

func GetSmokeRiskLevel(heightmap []string) int {
	maxX := len(heightmap) - 1
	maxY := len(heightmap[0]) - 1
	parsedHeightmap := parseHeightmap(heightmap)
	riskLevel := 0

	for x, line := range parsedHeightmap {
		for y, height := range line {
			hasHigherTop := x > 0 && height < parsedHeightmap[x-1][y] || x == 0
			hasHigherDown := x < maxX && height < parsedHeightmap[x+1][y] || x == maxX
			hasHigherLeft := y > 0 && height < line[y-1] || y == 0
			hasHigherRight := y < maxY && height < line[y+1] || y == maxY
			if hasHigherTop && hasHigherDown && hasHigherRight && hasHigherLeft {
				riskLevel += 1 + height
			}
		}
	}

	return riskLevel
}

// Part 2

func GetBiggestBasins(heightmap []string) int {
	maxX := len(heightmap) - 1
	maxY := len(heightmap[0]) - 1
	parsedHeightmap := parseHeightmap(heightmap)
	firstBiggestBasinSize := 0
	secondBiggestBasinSize := 0
	thirdBiggestBasinSize := 0

	for x, line := range parsedHeightmap {
		for y, height := range line {
			hasHigherTop := x > 0 && height < parsedHeightmap[x-1][y] || x == 0
			hasHigherDown := x < maxX && height < parsedHeightmap[x+1][y] || x == maxX
			hasHigherLeft := y > 0 && height < line[y-1] || y == 0
			hasHigherRight := y < maxY && height < line[y+1] || y == maxY
			if hasHigherTop && hasHigherDown && hasHigherRight && hasHigherLeft {
				basinSize := getBasinSize(x, y, parsedHeightmap)
				if basinSize > firstBiggestBasinSize {
					thirdBiggestBasinSize = secondBiggestBasinSize
					secondBiggestBasinSize = firstBiggestBasinSize
					firstBiggestBasinSize = basinSize
				} else if basinSize > secondBiggestBasinSize {
					thirdBiggestBasinSize = secondBiggestBasinSize
					secondBiggestBasinSize = basinSize
				} else if basinSize > thirdBiggestBasinSize {
					thirdBiggestBasinSize = basinSize
				}
			}
		}
	}

	return firstBiggestBasinSize * secondBiggestBasinSize * thirdBiggestBasinSize
}

func getBasinSize(x, y int, heightmap [][]int) int {
	basinSet := make(map[coordinates]bool)
	basinSet = getBasinSet(x, y, heightmap, basinSet)
	return len(basinSet)
}

func getBasinSet(x, y int, heightmap [][]int, basinSet map[coordinates]bool) map[coordinates]bool {
	maxX := len(heightmap) - 1
	maxY := len(heightmap[0]) - 1
	height := heightmap[x][y]

	basinSet[coordinates{x, y}] = true

	hasHigherTop := x > 0 && height < heightmap[x-1][y] && heightmap[x-1][y] < 9
	hasHigherDown := x < maxX && height < heightmap[x+1][y] && heightmap[x+1][y] < 9
	hasHigherLeft := y > 0 && height < heightmap[x][y-1] && heightmap[x][y-1] < 9
	hasHigherRight := y < maxY && height < heightmap[x][y+1] && heightmap[x][y+1] < 9

	if !hasHigherTop && !hasHigherDown && !hasHigherRight && !hasHigherLeft {
		return basinSet
	} else {
		if hasHigherTop {
			basinSet = getBasinSet(x-1, y, heightmap, basinSet)
		}
		if hasHigherDown {
			basinSet = getBasinSet(x+1, y, heightmap, basinSet)
		}
		if hasHigherLeft {
			basinSet = getBasinSet(x, y-1, heightmap, basinSet)
		}
		if hasHigherRight {
			basinSet = getBasinSet(x, y+1, heightmap, basinSet)
		}
		return basinSet
	}
}

func parseHeightmap(heightmap []string) [][]int {
	maxX := len(heightmap)
	maxY := len(heightmap[0])
	parsedHeightmap := make([][]int, maxX)

	for x, line := range heightmap {
		parsedLine := make([]int, maxY)
		for y, height := range line {
			parsedHeight, _ := strconv.ParseInt(string(height), 10, 64)
			parsedLine[y] = int(parsedHeight)
		}
		parsedHeightmap[x] = parsedLine
	}

	return parsedHeightmap
}

// Data structure

type coordinates struct {
	x int
	y int
}
