package main

import (
	"strconv"
	"strings"
)

func determinePosition(instructions []string) (int, int) {
	forwardPosition := 0
	depthPosition := 0
	for _, instruction := range instructions {
		instructionType, value := ParseInstruction(instruction)

		if instructionType == "forward" {
			forwardPosition += value
		} else if instructionType == "down" {
			depthPosition += value
		} else if instructionType == "up" {
			depthPosition -= value
		}
	}
	return forwardPosition, depthPosition
}

func determinePositionWithAim(instructions []string) (int, int) {
	forwardPosition := 0
	depthPosition := 0
	aim := 0

	for _, instruction := range instructions {
		instructionType, value := ParseInstruction(instruction)

		if instructionType == "forward" {
			forwardPosition += value
			depthPosition += value * aim
		} else if instructionType == "down" {
			aim += value
		} else if instructionType == "up" {
			aim -= value
		}
	}
	return forwardPosition, depthPosition
}

func ParseInstruction(instruction string) (string, int) {
	s := strings.Split(instruction, " ")

	value, _ := strconv.ParseInt(s[1], 10, 64)
	return s[0], int(value)
}
