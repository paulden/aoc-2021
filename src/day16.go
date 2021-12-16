package main

import (
	"math"
)

// Part 1

func SumVersionNumbers(input string) int {
	bitsSequence := HexadecimalToBits(input)
	_, versionsSum, _ := ParsePacket(bitsSequence, 0, 0)
	return versionsSum
}

// Part 2

func EvaluateBITSExpression(input string) int {
	bitsSequence := HexadecimalToBits(input)
	_, _, value := ParsePacket(bitsSequence, 0, 0)
	return value
}

func ParsePacket(input []int, originalPointer int, originalVersionsSum int) (int, int, int) {
	pointer := originalPointer
	versionsSum := originalVersionsSum
	var value int
	packetTypeID := BitsToDecimal(input[pointer+3 : pointer+6])

	if packetTypeID == 4 {
		version, literalValue, pointerAfterParsing := ParseLiteralPacket(input, pointer)
		pointer = pointerAfterParsing
		value = literalValue
		versionsSum += version
	} else {
		packetVersion := BitsToDecimal(input[pointer : pointer+3])
		versionsSum += packetVersion
		pointer += 6
		lengthTypeID := input[pointer]
		pointer++
		subValues := make([]int, 0)

		if lengthTypeID == 0 {
			subPacketsLength := BitsToDecimal(input[pointer : pointer+15])
			pointer += 15
			endPointer := pointer + subPacketsLength - 4

			for pointer < endPointer {
				var currentValue int
				pointer, versionsSum, currentValue = ParsePacket(input, pointer, versionsSum)
				subValues = append(subValues, currentValue)
			}
		} else if lengthTypeID == 1 {
			subPacketsNumber := BitsToDecimal(input[pointer : pointer+11])
			pointer += 11
			subPacketsParsed := 0

			for subPacketsParsed < subPacketsNumber {
				var currentValue int
				pointer, versionsSum, currentValue = ParsePacket(input, pointer, versionsSum)
				subValues = append(subValues, currentValue)
				subPacketsParsed++
			}
		}

		value = ComputeOperation(packetTypeID, subValues)
	}

	return pointer, versionsSum, value
}

func ComputeOperation(packetTypeID int, subValues []int) (value int) {
	switch packetTypeID {
	case 0:
		value = Sum(subValues)
	case 1:
		value = Product(subValues)
	case 2:
		value = Minimum(subValues)
	case 3:
		value = Maximum(subValues)
	case 5:
		value = GreaterThan(subValues[0], subValues[1])
	case 6:
		value = LessThan(subValues[0], subValues[1])
	case 7:
		value = EqualTo(subValues[0], subValues[1])
	}
	return
}

func ParseLiteralPacket(input []int, originalPointer int) (packetVersion int, literalValue int, pointer int) {
	packetVersion = BitsToDecimal(input[originalPointer : originalPointer+3])

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
	literalValue = BitsToDecimal(literalBinaryValue)

	return
}

func BitsToDecimal(bits []int) int {
	decimal := 0

	for i := range bits {
		decimal += bits[len(bits)-i-1] * Power(2, i)
	}

	return decimal
}

func HexadecimalToBits(input string) []int {
	hexDecimalToBits := map[string][]int{
		"0": {0, 0, 0, 0},
		"1": {0, 0, 0, 1},
		"2": {0, 0, 1, 0},
		"3": {0, 0, 1, 1},
		"4": {0, 1, 0, 0},
		"5": {0, 1, 0, 1},
		"6": {0, 1, 1, 0},
		"7": {0, 1, 1, 1},
		"8": {1, 0, 0, 0},
		"9": {1, 0, 0, 1},
		"A": {1, 0, 1, 0},
		"B": {1, 0, 1, 1},
		"C": {1, 1, 0, 0},
		"D": {1, 1, 0, 1},
		"E": {1, 1, 1, 0},
		"F": {1, 1, 1, 1},
	}

	bits := make([]int, 0)

	for _, character := range input {
		bits = append(bits, hexDecimalToBits[string(character)]...)
	}

	return bits
}

func Power(value, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= value
	}
	return result
}

func Sum(values []int) int {
	var result int
	for _, value := range values {
		result += value
	}
	return result
}

func Product(values []int) int {
	result := 1
	for _, value := range values {
		result *= value
	}
	return result
}

func Minimum(values []int) int {
	result := math.MaxInt
	for _, value := range values {
		if value < result {
			result = value
		}
	}
	return result
}

func Maximum(values []int) int {
	result := math.MinInt
	for _, value := range values {
		if value > result {
			result = value
		}
	}
	return result
}

func GreaterThan(a, b int) int {
	if a > b {
		return 1
	}
	return 0
}

func LessThan(a, b int) int {
	if a < b {
		return 1
	}
	return 0
}

func EqualTo(a, b int) int {
	if a == b {
		return 1
	}
	return 0
}
