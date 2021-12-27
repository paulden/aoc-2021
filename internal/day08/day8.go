package day08

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

func SumOutputDisplays(report []string) int {
	var sum int
	allPossibleMappings := generateAllPossibleMappings()
	for _, pattern := range report {
		sum += parseSevenSegmentDisplay(pattern, allPossibleMappings)
	}
	return sum
}

// Data structure

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

func (d *sevenSegmentDigit) isEqual(digit sevenSegmentDigit) bool {
	for _, segment := range "abcdefg" {
		if d.digits[string(segment)] != digit.digits[string(segment)] {
			return false
		}
	}
	return true
}

func parseSevenSegmentDisplay(display string, allPossibleMappings []map[string]string) int {
	split := strings.Split(display, " | ")
	fourDigits := strings.Split(split[1], " ")

	mapping := getSevenSegmentsMapping(split[0], allPossibleMappings)

	parsedDisplay := ""
	for _, digit := range fourDigits {
		parsedDisplay += getDigitFromMapping(digit, mapping)
	}

	parseInt, _ := strconv.ParseInt(parsedDisplay, 10, 64)
	return int(parseInt)
}

func getDigitFromMapping(value string, mapping map[string]string) string {
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
	if actualDigit.isEqual(zero) {
		return "0"
	}
	if actualDigit.isEqual(one) {
		return "1"
	}
	if actualDigit.isEqual(two) {
		return "2"
	}
	if actualDigit.isEqual(three) {
		return "3"
	}
	if actualDigit.isEqual(four) {
		return "4"
	}
	if actualDigit.isEqual(five) {
		return "5"
	}
	if actualDigit.isEqual(six) {
		return "6"
	}
	if actualDigit.isEqual(seven) {
		return "7"
	}
	if actualDigit.isEqual(eight) {
		return "8"
	}
	if actualDigit.isEqual(nine) {
		return "9"
	}
	return "ERR"
}

func getSevenSegmentsMapping(inputPattern string, allPossibleMappings []map[string]string) map[string]string {
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

	for _, m := range allPossibleMappings {
		mappedOne := newSevenSegmentDigit(m["c"] + m["f"])
		if !mappedOne.isEqual(one) {
			continue
		}
		mappedSeven := newSevenSegmentDigit(m["a"] + m["c"] + m["f"])
		if !mappedSeven.isEqual(seven) {
			continue
		}
		mappedFour := newSevenSegmentDigit(m["b"] + m["c"] + m["d"] + m["f"])
		if !mappedFour.isEqual(four) {
			continue
		}
		mappedTwo := newSevenSegmentDigit(m["a"] + m["c"] + m["d"] + m["e"] + m["g"])
		mappedThree := newSevenSegmentDigit(m["a"] + m["c"] + m["d"] + m["f"] + m["g"])
		mappedFive := newSevenSegmentDigit(m["a"] + m["b"] + m["d"] + m["f"] + m["g"])
		if !(mappedTwo.isEqual(possibleTwo1) || mappedTwo.isEqual(possibleTwo2) || mappedTwo.isEqual(possibleTwo3)) {
			continue
		}
		if !(mappedThree.isEqual(possibleThree1) || mappedThree.isEqual(possibleThree2) || mappedThree.isEqual(possibleThree3)) {
			continue
		}
		if !(mappedFive.isEqual(possibleFive1) || mappedFive.isEqual(possibleFive2) || mappedFive.isEqual(possibleFive3)) {
			continue
		}

		mappedZero := newSevenSegmentDigit(m["a"] + m["b"] + m["c"] + m["e"] + m["f"] + m["g"])
		mappedSix := newSevenSegmentDigit(m["a"] + m["b"] + m["d"] + m["e"] + m["f"] + m["g"])
		mappedNine := newSevenSegmentDigit(m["a"] + m["b"] + m["c"] + m["d"] + m["f"] + m["g"])
		if !(mappedZero.isEqual(possibleZero1) || mappedZero.isEqual(possibleZero2) || mappedZero.isEqual(possibleZero3)) {
			continue
		}
		if !(mappedSix.isEqual(possibleSix1) || mappedSix.isEqual(possibleSix2) || mappedSix.isEqual(possibleSix3)) {
			continue
		}
		if !(mappedNine.isEqual(possibleNine1) || mappedNine.isEqual(possibleNine2) || mappedNine.isEqual(possibleNine3)) {
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

func generateAllPossibleMappings() []map[string]string {
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
