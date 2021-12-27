package day22

import (
	"aoc-2021/internal/pkg/utils"
	"errors"
	"regexp"
	"strings"
)

// Part 1

func CountCubesOnNaive(input []string) int {
	steps := parseReactorInstructions(input)
	cubesState := make(map[cubeCoordinates]bool)

	for _, step := range steps {
		for x := utils.Max(-50, step.cuboid.xFrom); x <= utils.Min(50, step.cuboid.xTo); x++ {
			for y := utils.Max(-50, step.cuboid.yFrom); y <= utils.Min(50, step.cuboid.yTo); y++ {
				for z := utils.Max(-50, step.cuboid.zFrom); z <= utils.Min(50, step.cuboid.zTo); z++ {
					cubesState[cubeCoordinates{x, y, z}] = step.on
				}
			}
		}
	}

	for cube, isOn := range cubesState {
		if !isOn {
			delete(cubesState, cube)
		}
	}

	return len(cubesState)
}

// Part 2

func CountCubesOnOptimized(input []string) int {
	instructions := parseReactorInstructions(input)
	decomposedInstructions := make([]reactorInstruction, 0)

	for _, instruction := range instructions {
		newInstructions := decomposeInstruction(instruction, decomposedInstructions)
		decomposedInstructions = append(decomposedInstructions, newInstructions...)
	}

	cubesOn := getCubesOnFromAllInstructions(decomposedInstructions)

	return cubesOn
}

func decomposeInstruction(
	instruction reactorInstruction,
	decomposedInstructions []reactorInstruction,
) []reactorInstruction {
	newInstructions := make([]reactorInstruction, 0)

	for _, state := range decomposedInstructions {
		intersection, err := instruction.cuboid.getIntersectionWith(state.cuboid)
		if err == nil {
			// We need to compensate intersections found with previous individual instructions
			if instruction.on {
				if state.on {
					// We already turned on this portion and don't want to count it twice: compensate by adding an "off" instruction on it!
					newInstructions = append(newInstructions, reactorInstruction{false, *intersection})
				}
				if !state.on {
					// A portion that was turned off needs to be turned back on!
					newInstructions = append(newInstructions, reactorInstruction{true, *intersection})
				}
			}
			if !instruction.on {
				if state.on {
					// A portion that was turned on needs to be turned off!
					newInstructions = append(newInstructions, reactorInstruction{false, *intersection})
				}
				if !state.on {
					// We already turned off this portion and don't want to count it twice: compensate by adding an "on" instruction on it!
					newInstructions = append(newInstructions, reactorInstruction{true, *intersection})
				}
			}
		}
	}

	if instruction.on {
		// Finally, turn on the portion to count portions that were never "visited" before
		// If we turn off some portion that was never "visited", it does not matter since cubes stay off
		newInstructions = append(newInstructions, reactorInstruction{instruction.on, instruction.cuboid})
	}
	return newInstructions
}

func getCubesOnFromAllInstructions(instructions []reactorInstruction) int {
	cubesOn := 0

	for _, instruction := range instructions {
		if instruction.on {
			cubesOn += instruction.cuboid.getVolume()
		} else {
			cubesOn -= instruction.cuboid.getVolume()
		}
	}

	return cubesOn
}

type reactorInstruction struct {
	on     bool
	cuboid cuboid
}

type cubeCoordinates struct {
	x, y, z int
}

type cuboid struct {
	xFrom, xTo int
	yFrom, yTo int
	zFrom, zTo int
}

func (c *cuboid) getVolume() int {
	return (c.xTo - c.xFrom + 1) * (c.yTo - c.yFrom + 1) * (c.zTo - c.zFrom + 1)
}

func (c *cuboid) getIntersectionWith(c2 cuboid) (*cuboid, error) {
	xFrom, xTo := utils.Max(c.xFrom, c2.xFrom), utils.Min(c.xTo, c2.xTo)
	yFrom, yTo := utils.Max(c.yFrom, c2.yFrom), utils.Min(c.yTo, c2.yTo)
	zFrom, zTo := utils.Max(c.zFrom, c2.zFrom), utils.Min(c.zTo, c2.zTo)
	if xTo < xFrom || yTo < yFrom || zTo < zFrom {
		return nil, errors.New("no intersection found")
	}
	intersection := cuboid{xFrom, xTo, yFrom, yTo, zFrom, zTo}
	return &intersection, nil
}

// Parsing

func parseReactorInstructions(input []string) (instructions []reactorInstruction) {
	re := regexp.MustCompile("-?[0-9]+")
	for _, step := range input {
		numbers := utils.StringsToIntegers(re.FindAllString(step, 6))
		xFrom, xTo := numbers[0], numbers[1]
		yFrom, yTo := numbers[2], numbers[3]
		zFrom, zTo := numbers[4], numbers[5]
		cuboid := cuboid{xFrom, xTo, yFrom, yTo, zFrom, zTo}
		instruction := reactorInstruction{true, cuboid}

		if strings.Contains(step, "off") {
			instruction.on = false
		}

		instructions = append(instructions, instruction)
	}
	return
}
