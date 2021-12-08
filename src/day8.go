package main

import (
	"strconv"
	"strings"
)

// Part 1

func CountUniqueSegmentsDigits(report []string) int {
	uniqueSegmentsDigits := 0

	for _, pattern := range report {
		split := strings.Split(pattern, " | ")
		fourDigits := strings.Split(split[1], " ")
		for _, digit := range fourDigits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				uniqueSegmentsDigits++
			}
		}
	}

	return uniqueSegmentsDigits
}

// Part 2

// 0:      1:      2:      3:      4:
// aaaa    ....    aaaa    aaaa    ....
//b    c  .    c  .    c  .    c  b    c
//b    c  .    c  .    c  .    c  b    c
// ....    ....    dddd    dddd    dddd
//e    f  .    f  e    .  .    f  .    f
//e    f  .    f  e    .  .    f  .    f
// gggg    ....    gggg    gggg    ....
//
// 5:      6:      7:      8:      9:
// aaaa    aaaa    aaaa    aaaa    aaaa
//b    .  b    .  .    c  b    c  b    c
//b    .  b    .  .    c  b    c  b    c
// dddd    dddd    ....    dddd    dddd
//.    f  e    f  .    f  e    f  .    f
//.    f  e    f  .    f  e    f  .    f
// gggg    gggg    ....    gggg    gggg

type sevenSegmentDigit struct {
	digits map[string]bool
}

func newSevenSegmentDigit(segments string) sevenSegmentDigit {
	digit := make(map[string]bool)
	for _, segment := range segments {
		digit[string(segment)] = true
	}
	segmentDigit := sevenSegmentDigit{digit}
	return segmentDigit
}

func isEqual(digit1, digit2 sevenSegmentDigit) bool {
	for _, segment := range "abcdefg" {
		if digit1.digits[string(segment)] != digit2.digits[string(segment)] {
			return false
		}
	}
	return true
}

func SumOutputDisplays(report []string) int {
	sum := 0

	for _, pattern := range report {
		sum += ParseSevenSegmentDisplay(pattern)
	}

	return sum
}

func ParseSevenSegmentDisplay(display string) int {
	split := strings.Split(display, " | ")
	fourDigits := strings.Split(split[1], " ")

	mapping := GetSevenSegmentsMapping(split[0])

	parsedDisplay := ""
	for _, digit := range fourDigits {
		parsedDisplay += GetDigitFromMapping(digit, mapping)
	}

	parseInt, _ := strconv.ParseInt(parsedDisplay, 10, 64)
	return int(parseInt)
}

func GetDigitFromMapping(value string, mapping map[string]string) string {
	zero := newSevenSegmentDigit("abcefg")
	one := newSevenSegmentDigit("cf")
	two := newSevenSegmentDigit("acdeg")
	three := newSevenSegmentDigit("acdfg")
	four := newSevenSegmentDigit("bcdf")
	five := newSevenSegmentDigit("abdfg")
	six := newSevenSegmentDigit("abdefg")
	seven := newSevenSegmentDigit("acf")
	eight := newSevenSegmentDigit("abcdefg")
	nine := newSevenSegmentDigit("abcdfg")

	translatedValue := ""
	for _, segment := range value {
		translatedValue += mapping[string(segment)]
	}
	actualDigit := newSevenSegmentDigit(translatedValue)
	if isEqual(actualDigit, zero) {
		return "0"
	}
	if isEqual(actualDigit, one) {
		return "1"
	}
	if isEqual(actualDigit, two) {
		return "2"
	}
	if isEqual(actualDigit, three) {
		return "3"
	}
	if isEqual(actualDigit, four) {
		return "4"
	}
	if isEqual(actualDigit, five) {
		return "5"
	}
	if isEqual(actualDigit, six) {
		return "6"
	}
	if isEqual(actualDigit, seven) {
		return "7"
	}
	if isEqual(actualDigit, eight) {
		return "8"
	}
	if isEqual(actualDigit, nine) {
		return "9"
	}
	return "ERR"
}

func GetSevenSegmentsMapping(inputPattern string) map[string]string {
	inputDigits := strings.Split(inputPattern, " ")

	var possibleZero1 sevenSegmentDigit
	var possibleZero2 sevenSegmentDigit
	var possibleZero3 sevenSegmentDigit
	var one sevenSegmentDigit
	var possibleTwo1 sevenSegmentDigit
	var possibleTwo2 sevenSegmentDigit
	var possibleTwo3 sevenSegmentDigit
	var possibleThree1 sevenSegmentDigit
	var possibleThree2 sevenSegmentDigit
	var possibleThree3 sevenSegmentDigit
	var four sevenSegmentDigit
	var possibleFive1 sevenSegmentDigit
	var possibleFive2 sevenSegmentDigit
	var possibleFive3 sevenSegmentDigit
	var possibleSix1 sevenSegmentDigit
	var possibleSix2 sevenSegmentDigit
	var possibleSix3 sevenSegmentDigit
	var seven sevenSegmentDigit
	var possibleNine1 sevenSegmentDigit
	var possibleNine2 sevenSegmentDigit
	var possibleNine3 sevenSegmentDigit

	fiveSegments := 0
	sixSegments := 0

	for _, inputDigit := range inputDigits {
		if len(inputDigit) == 2 {
			one = newSevenSegmentDigit(inputDigit)
		} else if len(inputDigit) == 3 {
			seven = newSevenSegmentDigit(inputDigit)
		} else if len(inputDigit) == 4 {
			four = newSevenSegmentDigit(inputDigit)
		} else if len(inputDigit) == 5 && fiveSegments == 0 {
			possibleTwo1 = newSevenSegmentDigit(inputDigit)
			possibleThree1 = newSevenSegmentDigit(inputDigit)
			possibleFive1 = newSevenSegmentDigit(inputDigit)
			fiveSegments++
		} else if len(inputDigit) == 5 && fiveSegments == 1 {
			possibleTwo2 = newSevenSegmentDigit(inputDigit)
			possibleThree2 = newSevenSegmentDigit(inputDigit)
			possibleFive2 = newSevenSegmentDigit(inputDigit)
			fiveSegments++
		} else if len(inputDigit) == 5 && fiveSegments == 2 {
			possibleTwo3 = newSevenSegmentDigit(inputDigit)
			possibleThree3 = newSevenSegmentDigit(inputDigit)
			possibleFive3 = newSevenSegmentDigit(inputDigit)
		} else if len(inputDigit) == 6 && sixSegments == 0 {
			possibleZero1 = newSevenSegmentDigit(inputDigit)
			possibleSix1 = newSevenSegmentDigit(inputDigit)
			possibleNine1 = newSevenSegmentDigit(inputDigit)
			sixSegments++
		} else if len(inputDigit) == 6 && sixSegments == 1 {
			possibleZero2 = newSevenSegmentDigit(inputDigit)
			possibleSix2 = newSevenSegmentDigit(inputDigit)
			possibleNine2 = newSevenSegmentDigit(inputDigit)
			sixSegments++
		} else if len(inputDigit) == 6 && sixSegments == 2 {
			possibleZero3 = newSevenSegmentDigit(inputDigit)
			possibleSix3 = newSevenSegmentDigit(inputDigit)
			possibleNine3 = newSevenSegmentDigit(inputDigit)
		}
	}

	possibleMappings := make([]map[string]string, 0)

	allPossibleMappings := GenerateAllPossibleMappings()

	for _, m := range allPossibleMappings {
		mappedOne := newSevenSegmentDigit(m["c"] + m["f"])
		if !isEqual(mappedOne, one) {
			continue
		}
		mappedSeven := newSevenSegmentDigit(m["a"] + m["c"] + m["f"])
		if !isEqual(mappedSeven, seven) {
			continue
		}
		mappedFour := newSevenSegmentDigit(m["b"] + m["c"] + m["d"] + m["f"])
		if !isEqual(mappedFour, four) {
			continue
		}
		mappedTwo := newSevenSegmentDigit(m["a"] + m["c"] + m["d"] + m["e"] + m["g"])
		mappedThree := newSevenSegmentDigit(m["a"] + m["c"] + m["d"] + m["f"] + m["g"])
		mappedFive := newSevenSegmentDigit(m["a"] + m["b"] + m["d"] + m["f"] + m["g"])
		if !(isEqual(mappedTwo, possibleTwo1) || isEqual(mappedTwo, possibleTwo2) || isEqual(mappedTwo, possibleTwo3)) {
			continue
		}
		if !(isEqual(mappedThree, possibleThree1) || isEqual(mappedThree, possibleThree2) || isEqual(mappedThree, possibleThree3)) {
			continue
		}
		if !(isEqual(mappedFive, possibleFive1) || isEqual(mappedFive, possibleFive2) || isEqual(mappedFive, possibleFive3)) {
			continue
		}

		mappedZero := newSevenSegmentDigit(m["a"] + m["b"] + m["c"] + m["e"] + m["f"] + m["g"])
		mappedSix := newSevenSegmentDigit(m["a"] + m["b"] + m["d"] + m["e"] + m["f"] + m["g"])
		mappedNine := newSevenSegmentDigit(m["a"] + m["b"] + m["c"] + m["d"] + m["f"] + m["g"])
		if !(isEqual(mappedZero, possibleZero1) || isEqual(mappedZero, possibleZero2) || isEqual(mappedZero, possibleZero3)) {
			continue
		}
		if !(isEqual(mappedSix, possibleSix1) || isEqual(mappedSix, possibleSix2) || isEqual(mappedSix, possibleSix3)) {
			continue
		}
		if !(isEqual(mappedNine, possibleNine1) || isEqual(mappedNine, possibleNine2) || isEqual(mappedNine, possibleNine3)) {
			continue
		}

		possibleMappings = append(possibleMappings, m)
	}

	switchedMap := make(map[string]string)
	for k, v := range possibleMappings[0] {
		switchedMap[v] = k
	}
	return switchedMap
}

func GenerateAllPossibleMappings() []map[string]string {
	allPossibleMappings := make([]map[string]string, 0)
	permutations := generatePermutations([]rune("abcdefg"), 0, 6, make([]string, 0))
	for _, permutation := range permutations {
		possibleMapping := map[string]string{
			string(permutation[0]): "a",
			string(permutation[1]): "b",
			string(permutation[2]): "c",
			string(permutation[3]): "d",
			string(permutation[4]): "e",
			string(permutation[5]): "f",
			string(permutation[6]): "g",
		}
		allPossibleMappings = append(allPossibleMappings, possibleMapping)
	}
	return allPossibleMappings
}

// Adapted from https://golangbyexample.com/all-permutations-string-golang/
func generatePermutations(sampleRune []rune, left, right int, permutations []string) []string {
	if left == right {
		permutations = append(permutations, string(sampleRune))
	} else {
		for i := left; i <= right; i++ {
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
			permutations = generatePermutations(sampleRune, left+1, right, permutations)
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
		}
	}
	return permutations
}
