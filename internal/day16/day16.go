package day16

import (
	"aoc-2021/internal/pkg/utils"
)

// Part 1

func SumVersionNumbers(input string) int {
	bitsSequence := utils.HexadecimalToBits(input)
	_, versionsSum, _ := parsePacket(bitsSequence, 0, 0)
	return versionsSum
}

// Part 2

func EvaluateBITSExpression(input string) int {
	bitsSequence := utils.HexadecimalToBits(input)
	_, _, value := parsePacket(bitsSequence, 0, 0)
	return value
}

func parsePacket(input []int, originalPointer int, originalVersionsSum int) (int, int, int) {
	pointer := originalPointer
	versionsSum := originalVersionsSum
	var value int
	packetTypeID := utils.BitsToDecimal(input[pointer+3 : pointer+6])

	if packetTypeID == 4 {
		version, literalValue, pointerAfterParsing := parseLiteralPacket(input, pointer)
		pointer = pointerAfterParsing
		value = literalValue
		versionsSum += version
	} else {
		packetVersion := utils.BitsToDecimal(input[pointer : pointer+3])
		versionsSum += packetVersion
		pointer += 6
		lengthTypeID := input[pointer]
		pointer++
		subValues := make([]int, 0)

		if lengthTypeID == 0 {
			subPacketsLength := utils.BitsToDecimal(input[pointer : pointer+15])
			pointer += 15
			endPointer := pointer + subPacketsLength - 4

			for pointer < endPointer {
				var currentValue int
				pointer, versionsSum, currentValue = parsePacket(input, pointer, versionsSum)
				subValues = append(subValues, currentValue)
			}
		} else if lengthTypeID == 1 {
			subPacketsNumber := utils.BitsToDecimal(input[pointer : pointer+11])
			pointer += 11
			subPacketsParsed := 0

			for subPacketsParsed < subPacketsNumber {
				var currentValue int
				pointer, versionsSum, currentValue = parsePacket(input, pointer, versionsSum)
				subValues = append(subValues, currentValue)
				subPacketsParsed++
			}
		}

		value = computeOperation(packetTypeID, subValues)
	}

	return pointer, versionsSum, value
}

func computeOperation(packetTypeID int, subValues []int) (value int) {
	switch packetTypeID {
	case 0:
		value = utils.Sum(subValues)
	case 1:
		value = utils.Product(subValues)
	case 2:
		value = utils.Minimum(subValues)
	case 3:
		value = utils.Maximum(subValues)
	case 5:
		value = utils.GreaterThan(subValues[0], subValues[1])
	case 6:
		value = utils.LessThan(subValues[0], subValues[1])
	case 7:
		value = utils.EqualTo(subValues[0], subValues[1])
	}
	return
}

func parseLiteralPacket(input []int, originalPointer int) (packetVersion int, literalValue int, pointer int) {
	packetVersion = utils.BitsToDecimal(input[originalPointer : originalPointer+3])

	pointer = originalPointer + 6
	shouldKeepReading := input[pointer]
	literalBinaryValue := make([]int, 0)
	for shouldKeepReading == 1 {
		literalBinaryValue = append(literalBinaryValue, input[pointer+1:pointer+5]...)
		pointer += 5
		shouldKeepReading = input[pointer]
	}
	literalBinaryValue = append(literalBinaryValue, input[pointer+1:pointer+5]...)
	pointer += 5
	literalValue = utils.BitsToDecimal(literalBinaryValue)

	return
}
