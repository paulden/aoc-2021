package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Part 1

func FoldPaperOnce(input []string) map[dotCoordinates]bool {
	paper, foldingInstructions := ParseInstructions(input)
	foldedPaper := Fold(foldingInstructions[0], paper)
	return foldedPaper
}

// Part 2

func FoldPaperCompletely(input []string) map[dotCoordinates]bool {
	paper, foldingInstructions := ParseInstructions(input)
	for _, instruction := range foldingInstructions {
		paper = Fold(instruction, paper)
	}
	return paper
}

func Fold(instruction foldingInstruction, originalCoordinates map[dotCoordinates]bool) map[dotCoordinates]bool {
	foldedCoordinates := make(map[dotCoordinates]bool)
	for coord, _ := range originalCoordinates {
		var newX, newY int
		if instruction.foldAlong == "y" && coord.y >= instruction.foldAt {
			newX = coord.x
			newY = 2*instruction.foldAt - coord.y
		} else if instruction.foldAlong == "x" && coord.x >= instruction.foldAt {
			newX = 2*instruction.foldAt - coord.x
			newY = coord.y
		} else {
			newX = coord.x
			newY = coord.y
		}
		foldedCoordinates[dotCoordinates{newX, newY}] = true
	}

	return foldedCoordinates
}

// Utils

type dotCoordinates struct {
	x, y int
}

type foldingInstruction struct {
	foldAlong string
	foldAt    int
}

func PrettyPrintDots(paper map[dotCoordinates]bool) {
	xMax := 0
	yMax := 0
	for coord, _ := range paper {
		if coord.x > xMax {
			xMax = coord.x
		}
		if coord.y > yMax {
			yMax = coord.y
		}
	}

	for j := 0; j <= yMax; j++ {
		for i := 0; i <= xMax; i++ {
			if paper[dotCoordinates{i, j}] {
				fmt.Printf("##")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}

func ParseInstructions(instructions []string) (map[dotCoordinates]bool, []foldingInstruction) {
	dotsCoordinates := make(map[dotCoordinates]bool)
	foldingInstructions := make([]foldingInstruction, 0)

	for _, line := range instructions {
		if strings.Contains(line, "fold along") {
			instruction := strings.Trim(line, "fold along ")
			split := strings.Split(instruction, "=")

			foldAt, _ := strconv.ParseInt(split[1], 10, 32)

			foldingInstructions = append(foldingInstructions, foldingInstruction{split[0], int(foldAt)})
		} else if strings.Contains(line, ",") {
			split := strings.Split(line, ",")
			dotX, _ := strconv.ParseInt(split[0], 10, 32)
			dotY, _ := strconv.ParseInt(split[1], 10, 32)

			dot := dotCoordinates{int(dotX), int(dotY)}
			dotsCoordinates[dot] = true
		}
	}
	return dotsCoordinates, foldingInstructions
}
