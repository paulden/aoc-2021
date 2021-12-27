package day17

import (
	"strconv"
	"strings"
)

// Part 1

func GetMaxHeight(input string) int {
	area := parseTargetArea(input)
	maxHeight := 0

	allVelocities := getAllVelocities(input)

	for velocity, _ := range allVelocities {
		_, _, _, maxY := computeTrajectory(area, velocity.x, velocity.y)
		if maxY > maxHeight {
			maxHeight = maxY
		}
	}

	return maxHeight
}

func CountAllVelocities(input string) int {
	return len(getAllVelocities(input))
}

func getAllVelocities(input string) map[probeVelocity]bool {
	area := parseTargetArea(input)
	allVelocities := make(map[probeVelocity]bool)

	for vX := 0; vX < area.xRight+1; vX++ {
		for vY := -500; vY < 500; vY++ {
			inArea, _, _, _ := computeTrajectory(area, vX, vY)
			if inArea {
				allVelocities[probeVelocity{vX, vY}] = true
			}
		}
	}

	return allVelocities
}

func computeTrajectory(
	area targetArea,
	xVelocity int,
	yVelocity int,
) (isInArea bool, currentX int, currentY int, maxY int) {
	currentX, currentY = 0, 0

	for currentX <= area.xRight && currentY >= area.yDown {
		currentX += xVelocity
		currentY += yVelocity

		if currentY > maxY {
			maxY = currentY
		}

		if xVelocity > 0 {
			xVelocity--
		} else if xVelocity < 0 {
			xVelocity++
		}

		yVelocity--

		if area.isInArea(currentX, currentY) {
			isInArea = true
			return
		}
	}

	return
}

func parseTargetArea(input string) targetArea {
	var xTop, xDown, yTop, yDown int

	xIndex := strings.Index(input, "x=")
	xIndexEnd := strings.Index(input, ", y=")

	xSplit := strings.Split(input[xIndex+2:xIndexEnd], "..")

	x0, _ := strconv.ParseInt(xSplit[0], 10, 64)
	x1, _ := strconv.ParseInt(xSplit[1], 10, 64)

	if x0 > x1 {
		xTop = int(x0)
		xDown = int(x1)
	} else {
		xTop = int(x1)
		xDown = int(x0)
	}

	yIndex := strings.Index(input, "y=")
	ySplit := strings.Split(input[yIndex+2:], "..")

	y0, _ := strconv.ParseInt(ySplit[0], 10, 64)
	y1, _ := strconv.ParseInt(ySplit[1], 10, 64)

	if y0 > y1 {
		yTop = int(y0)
		yDown = int(y1)
	} else {
		yTop = int(y1)
		yDown = int(y0)
	}

	return targetArea{xTop, xDown, yTop, yDown}
}

type targetArea struct {
	xRight, xLeft, yTop, yDown int
}

type probeVelocity struct {
	x, y int
}

func (area *targetArea) isInArea(x, y int) bool {
	return x >= area.xLeft && x <= area.xRight && y >= area.yDown && y <= area.yTop
}
