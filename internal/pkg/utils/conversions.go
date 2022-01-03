package utils

import "strconv"

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

func StringsToIntegers(strings []string) []int {
	integers := make([]int, len(strings))
	for i := range strings {
		parseInt, _ := strconv.ParseInt(strings[i], 10, 64)
		integers[i] = int(parseInt)
	}
	return integers
}

func SliceToInteger(slice []int) (result int) {
	for i, integer := range slice {
		result += integer * Power(10, len(slice)-i-1)
	}

	return result
}
