package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Part 1

func CountCubesOnNaive(input []string) int {
	steps := ParseReactorInstructions(input)
	cubesState := make(map[CubeCoordinate]bool)

	for _, step := range steps {
		for x := Max(-50, step.cuboid.xFrom); x <= Min(50, step.cuboid.xTo); x++ {
			for y := Max(-50, step.cuboid.yFrom); y <= Min(50, step.cuboid.yTo); y++ {
				for z := Max(-50, step.cuboid.zFrom); z <= Min(50, step.cuboid.zTo); z++ {
					cubesState[CubeCoordinate{x, y, z}] = step.on
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
	instructions := ParseReactorInstructions(input)
	decomposedInstructions := make([]ReactorInstruction, 0)

	for _, instruction := range instructions {
		newInstructions := decomposeInstruction(instruction, decomposedInstructions)
		decomposedInstructions = append(decomposedInstructions, newInstructions...)
	}

	cubesOn := GetCubesOnFromAllInstructions(decomposedInstructions)

	return cubesOn
}

func decomposeInstruction(
	instruction ReactorInstruction,
	decomposedInstructions []ReactorInstruction,
) []ReactorInstruction {
	newInstructions := make([]ReactorInstruction, 0)

	for _, state := range decomposedInstructions {
		intersection, err := instruction.cuboid.Intersection(state.cuboid)
		if err == nil {
			// We need to compensate intersections found with previous individual instructions
			if instruction.on {
				if state.on {
					// We already turned on this portion and don't want to count it twice: compensate by adding an "off" instruction on it!
					newInstructions = append(newInstructions, ReactorInstruction{false, *intersection})
				}
				if !state.on {
					// A portion that was turned off needs to be turned back on!
					newInstructions = append(newInstructions, ReactorInstruction{true, *intersection})
				}
			}
			if !instruction.on {
				if state.on {
					// A portion that was turned on needs to be turned off!
					newInstructions = append(newInstructions, ReactorInstruction{false, *intersection})
				}
				if !state.on {
					// We already turned off this portion and don't want to count it twice: compensate by adding an "on" instruction on it!
					newInstructions = append(newInstructions, ReactorInstruction{true, *intersection})
				}
			}
		}
	}

	if instruction.on {
		// Finally, turn on the portion to count portions that were never "visited" before
		// If we turn off some portion that was never "visited", it does not matter since cubes stay off
		newInstructions = append(newInstructions, ReactorInstruction{instruction.on, instruction.cuboid})
	}
	return newInstructions
}

func GetCubesOnFromAllInstructions(instructions []ReactorInstruction) int {
	cubesOn := 0

	for _, instruction := range instructions {
		if instruction.on {
			cubesOn += instruction.cuboid.Volume()
		} else {
			cubesOn -= instruction.cuboid.Volume()
		}
	}

	return cubesOn
}

type ReactorInstruction struct {
	on     bool
	cuboid Cuboid
}

type CubeCoordinate struct {
	x, y, z int
}

type Cuboid struct {
	xFrom, xTo int
	yFrom, yTo int
	zFrom, zTo int
}

func (c *Cuboid) Volume() int {
	return (c.xTo - c.xFrom + 1) * (c.yTo - c.yFrom + 1) * (c.zTo - c.zFrom + 1)
}

func (c *Cuboid) Intersection(c2 Cuboid) (*Cuboid, error) {
	xFrom, xTo := Max(c.xFrom, c2.xFrom), Min(c.xTo, c2.xTo)
	yFrom, yTo := Max(c.yFrom, c2.yFrom), Min(c.yTo, c2.yTo)
	zFrom, zTo := Max(c.zFrom, c2.zFrom), Min(c.zTo, c2.zTo)
	if xTo < xFrom || yTo < yFrom || zTo < zFrom {
		return nil, errors.New("no intersection found")
	}
	intersection := Cuboid{xFrom, xTo, yFrom, yTo, zFrom, zTo}
	return &intersection, nil
}

// Parsing

func ParseReactorInstructions(input []string) (instructions []ReactorInstruction) {
	re := regexp.MustCompile("-?[0-9]+")
	for _, step := range input {
		numbers := StringsToIntegers(re.FindAllString(step, 6))
		xFrom, xTo := numbers[0], numbers[1]
		yFrom, yTo := numbers[2], numbers[3]
		zFrom, zTo := numbers[4], numbers[5]
		cuboid := Cuboid{xFrom, xTo, yFrom, yTo, zFrom, zTo}
		instruction := ReactorInstruction{true, cuboid}

		if strings.Contains(step, "off") {
			instruction.on = false
		}

		instructions = append(instructions, instruction)
	}
	return
}

func StringsToIntegers(strings []string) []int {
	integers := make([]int, len(strings))
	for i := range strings {
		parseInt, _ := strconv.ParseInt(strings[i], 10, 64)
		integers[i] = int(parseInt)
	}
	return integers
}
