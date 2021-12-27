package day05

import (
	"aoc-2021/internal/pkg/utils"
	"strconv"
	"strings"
)

// Part 1

func GetDangerousZonesNumber(ventLines []string) int {
	oceanFloorMapping := mapOceanFloorVentLines(ventLines)
	dangerousZonesNumber := 0

	for i := 0; i < len(oceanFloorMapping); i++ {
		for j := 0; j < len(oceanFloorMapping[0]); j++ {
			if oceanFloorMapping[i][j] > 1 {
				dangerousZonesNumber++
			}
		}
	}

	return dangerousZonesNumber
}

func mapOceanFloorVentLines(ventLines []string) [][]int {
	xLength, yLength := getFieldSize(ventLines)
	oceanFloorMapping := createEmptyMapping(xLength, yLength)

	for _, ventLine := range ventLines {
		x1, y1, x2, y2 := parseVentLine(ventLine)
		if x1 == x2 {
			for v := utils.Min(y1, y2); v <= utils.Max(y1, y2); v++ {
				oceanFloorMapping[x1][v]++
			}
		}
		if y1 == y2 {
			for v := utils.Min(x1, x2); v <= utils.Max(x1, x2); v++ {
				oceanFloorMapping[v][y1]++
			}
		}
	}

	return oceanFloorMapping
}

// Part 2

func GetDangerousZonesNumberWithDiagonals(ventLines []string) int {
	oceanFloorMapping := mapOceanFloorVentLinesWithDiagonals(ventLines)
	dangerousZonesNumber := 0

	for i := 0; i < len(oceanFloorMapping); i++ {
		for j := 0; j < len(oceanFloorMapping[0]); j++ {
			if oceanFloorMapping[i][j] > 1 {
				dangerousZonesNumber++
			}
		}
	}

	return dangerousZonesNumber
}

func mapOceanFloorVentLinesWithDiagonals(ventLines []string) [][]int {
	xLength, yLength := getFieldSize(ventLines)
	oceanFloorMapping := createEmptyMapping(xLength, yLength)

	for _, ventLine := range ventLines {
		x1, y1, x2, y2 := parseVentLine(ventLine)
		if x1 == x2 {
			for v := utils.Min(y1, y2); v <= utils.Max(y1, y2); v++ {
				oceanFloorMapping[x1][v]++
			}
		} else if y1 == y2 {
			for v := utils.Min(x1, x2); v <= utils.Max(x1, x2); v++ {
				oceanFloorMapping[v][y1]++
			}
		} else {
			if x1 < x2 && y1 < y2 {
				for i, j := x1, y1; i <= x2; i, j = i+1, j+1 {
					oceanFloorMapping[i][j]++
				}
			} else if x1 > x2 && y1 > y2 {
				for i, j := x1, y1; i >= x2; i, j = i-1, j-1 {
					oceanFloorMapping[i][j]++
				}
			} else if x1 > x2 && y1 < y2 {
				for i, j := x1, y1; i >= x2; i, j = i-1, j+1 {
					oceanFloorMapping[i][j]++
				}
			} else if x1 < x2 && y1 > y2 {
				for i, j := x1, y1; i <= x2; i, j = i+1, j-1 {
					oceanFloorMapping[i][j]++
				}
			}
		}
	}

	return oceanFloorMapping
}

// Utils

func createEmptyMapping(maxX int, maxY int) [][]int {
	return utils.CreateEmptyGrid(maxX, maxY)
}

func getFieldSize(ventLines []string) (int, int) {
	maxX := 0
	maxY := 0

	for _, ventLine := range ventLines {
		x1, y1, x2, y2 := parseVentLine(ventLine)
		if x1 > maxX {
			maxX = x1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y1 > maxY {
			maxY = y1
		}
		if y2 > maxY {
			maxY = y2
		}
	}

	return maxX + 1, maxY + 1
}

func parseVentLine(ventLine string) (int, int, int, int) {
	startToEnd := strings.Split(ventLine, " -> ")
	start := strings.Split(startToEnd[0], ",")
	end := strings.Split(startToEnd[1], ",")

	x1, _ := strconv.ParseInt(start[0], 10, 32)
	y1, _ := strconv.ParseInt(start[1], 10, 32)
	x2, _ := strconv.ParseInt(end[0], 10, 32)
	y2, _ := strconv.ParseInt(end[1], 10, 32)

	return int(x1), int(y1), int(x2), int(y2)
}
