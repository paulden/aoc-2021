package main

import (
	"strconv"
	"strings"
)

// Part 1

func GetDangerousZonesNumber(ventLines []string) int {
	oceanFloorMapping := MapOceanFloorVentLines(ventLines)
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

func MapOceanFloorVentLines(ventLines []string) [][]int {
	xLength, yLength := GetFieldSize(ventLines)
	oceanFloorMapping := CreateEmptyMapping(xLength, yLength)

	for _, ventLine := range ventLines {
		x1, y1, x2, y2 := ParseVentLine(ventLine)
		if x1 == x2 {
			for v := Min(y1, y2); v <= Max(y1, y2); v++ {
				oceanFloorMapping[x1][v]++
			}
		}
		if y1 == y2 {
			for v := Min(x1, x2); v <= Max(x1, x2); v++ {
				oceanFloorMapping[v][y1]++
			}
		}
	}

	return oceanFloorMapping
}

// Part 2

func GetDangerousZonesNumberWithDiagonals(ventLines []string) int {
	oceanFloorMapping := MapOceanFloorVentLinesWithDiagonals(ventLines)
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

func MapOceanFloorVentLinesWithDiagonals(ventLines []string) [][]int {
	xLength, yLength := GetFieldSize(ventLines)
	oceanFloorMapping := CreateEmptyMapping(xLength, yLength)

	for _, ventLine := range ventLines {
		x1, y1, x2, y2 := ParseVentLine(ventLine)
		if x1 == x2 {
			for v := Min(y1, y2); v <= Max(y1, y2); v++ {
				oceanFloorMapping[x1][v]++
			}
		} else if y1 == y2 {
			for v := Min(x1, x2); v <= Max(x1, x2); v++ {
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

func CreateEmptyMapping(maxX int, maxY int) [][]int {
	emptyMapping := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		emptyMapping[i] = make([]int, maxY)
	}
	return emptyMapping
}

func GetFieldSize(ventLines []string) (int, int) {
	maxX := 0
	maxY := 0

	for _, ventLine := range ventLines {
		x1, y1, x2, y2 := ParseVentLine(ventLine)
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

func ParseVentLine(ventLine string) (int, int, int, int) {
	startToEnd := strings.Split(ventLine, " -> ")
	start := strings.Split(startToEnd[0], ",")
	end := strings.Split(startToEnd[1], ",")

	x1, _ := strconv.ParseInt(start[0], 10, 32)
	y1, _ := strconv.ParseInt(start[1], 10, 32)
	x2, _ := strconv.ParseInt(end[0], 10, 32)
	y2, _ := strconv.ParseInt(end[1], 10, 32)

	return int(x1), int(y1), int(x2), int(y2)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
