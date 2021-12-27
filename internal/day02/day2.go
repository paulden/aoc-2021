package day02

import (
	"strconv"
	"strings"
)

// Part 1

func DeterminePosition(input []string) int {
	forwardPosition := 0
	depthPosition := 0
	for _, instruction := range input {
		instructionType, value := parseInstruction(instruction)

		if instructionType == "forward" {
			forwardPosition += value
		} else if instructionType == "down" {
			depthPosition += value
		} else if instructionType == "up" {
			depthPosition -= value
		}
	}
	return forwardPosition * depthPosition
}

// Part 2

func DeterminePositionWithAim(instructions []string) int {
	forwardPosition := 0
	depthPosition := 0
	aim := 0

	for _, instruction := range instructions {
		instructionType, value := parseInstruction(instruction)

		if instructionType == "forward" {
			forwardPosition += value
			depthPosition += value * aim
		} else if instructionType == "down" {
			aim += value
		} else if instructionType == "up" {
			aim -= value
		}
	}
	return forwardPosition * depthPosition
}

func parseInstruction(instruction string) (string, int) {
	s := strings.Split(instruction, " ")

	value, _ := strconv.ParseInt(s[1], 10, 64)
	return s[0], int(value)
}
