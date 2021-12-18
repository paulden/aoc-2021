package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Part 1

func GetMaxHeight(input string) int {
	targetArea := ParseTargetArea(input)
	vX, vY := 1, 0
	maxHeight := 0

	allVelocities := make(map[probeVelocity]bool)

	for i := 0; i < 1000; i++ {
		inArea, x, y, maxY := ComputeTrajectory(targetArea, vX, vY)

		if !inArea {
			if x < targetArea.xLeft && y < targetArea.yDown {
				//fmt.Printf("End position is (%v, %v) - Falling too fast before reaching target!\n", x, y)
				vX++
			} else if x > targetArea.yTop && y > targetArea.yTop {
				//fmt.Printf("End position is (%v, %v) - Going too far after target area, and too high!\n", x, y)
				vX--
			} else if x > targetArea.yTop && y < targetArea.yDown {
				//fmt.Printf("End position is (%v, %v) - Going too far after target area, and not high enough!\n", x, y)
				vX--
				vY++
			} else if x > targetArea.yTop {
				//fmt.Printf("End position is (%v, %v) - Going too far but in the right range!\n", x, y)
				vX--
			} else if x >= targetArea.xLeft && x <= targetArea.xRight {
				//fmt.Printf("End position is (%v, %v) - Falling too fast but in the right range!\n", x, y)
			}
		} else {
			allVelocities[probeVelocity{vX, vY}] = true
			if maxY > maxHeight {
				maxHeight = maxY
			}
			vY++
		}

		//fmt.Printf("Adjusted initial velocity is (%v, %v)!\n", vX, vY)
	}

	fmt.Printf("All velocities : %v\n", len(allVelocities))

	return maxHeight
}

func GetAllVelocities(input string) int {
	targetArea := ParseTargetArea(input)
	allVelocities := make(map[probeVelocity]bool)

	//minVX :=

	for vX := 0; vX < 233; vX++ {
		for vY := -233; vY < 233; vY++ {
			inArea, _, _, _ := ComputeTrajectory(targetArea, vX, vY)
			if inArea {
				allVelocities[probeVelocity{vX, vY}] = true
			}
		}
	}

	return len(allVelocities)
}

//func GetVelocityXToReach(value int) int {
//	delta := 1 * 4 * value * 2
//
//}

func ComputeTrajectory(
	area TargetArea,
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

func ParseTargetArea(input string) TargetArea {
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

	return TargetArea{xTop, xDown, yTop, yDown}
}

type TargetArea struct {
	xRight, xLeft, yTop, yDown int
}

type probeVelocity struct {
	x, y int
}

func (area *TargetArea) isInArea(x, y int) bool {
	return x >= area.xLeft && x <= area.xRight && y >= area.yDown && y <= area.yTop
}
