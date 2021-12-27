package day14

import (
	"math"
	"strings"
)

// Part 1

func CountPolymerCountsDifferenceNaive(input []string, steps int) int {
	polymerTemplate, pairInsertionRules := parsePolymerInput(input)

	for step := 0; step < steps; step++ {
		polymerTemplate = polymerize(polymerTemplate, pairInsertionRules)
	}

	occurrencesPerElement := make(map[string]int)

	for _, element := range polymerTemplate {
		occurrencesPerElement[element]++
	}

	maxOccurrences, minOccurrences := getMaxAndMinOccurrences(occurrencesPerElement)

	return maxOccurrences - minOccurrences
}

func polymerize(polymerTemplate []string, pairInsertionRules map[string]string) []string {
	newPolymerTemplate := make([]string, len(polymerTemplate))
	copy(newPolymerTemplate, polymerTemplate)
	j := 0

	for i := 0; i < len(polymerTemplate)-1; i, j = i+1, j+1 {
		pair := polymerTemplate[i] + polymerTemplate[i+1]

		if newElement, ok := pairInsertionRules[pair]; ok {
			newPolymerTemplate = append(newPolymerTemplate[:j+1], newPolymerTemplate[j:]...)
			newPolymerTemplate[j+1] = newElement
			j++
		}
	}
	return newPolymerTemplate
}

// Part 2

func CountPolymerCountsDifferenceOptimized(input []string, steps int) int {
	polymerTemplate, pairInsertionRules := parsePolymerInput(input)
	occurrencesPerPair := getOccurrencesPerPair(polymerTemplate)
	lastElement := polymerTemplate[len(polymerTemplate)-1]

	for step := 0; step < steps; step++ {
		newOccurrencesPerPair := make(map[string]int)
		for pair, occurrences := range occurrencesPerPair {
			polymerizedElement := pairInsertionRules[pair]
			firstPair := string(pair[0]) + polymerizedElement
			secondPair := polymerizedElement + string(pair[1])

			newOccurrencesPerPair[firstPair] += occurrences
			newOccurrencesPerPair[secondPair] += occurrences
		}
		occurrencesPerPair = newOccurrencesPerPair
	}

	occurrencesPerElement := countOccurrencesPerElement(occurrencesPerPair)
	occurrencesPerElement[lastElement]++
	maxOccurrences, minOccurrences := getMaxAndMinOccurrences(occurrencesPerElement)

	return maxOccurrences - minOccurrences
}

func countOccurrencesPerElement(occurrencesPerPair map[string]int) map[string]int {
	occurrencesPerElement := make(map[string]int)
	for pair, occurrences := range occurrencesPerPair {
		occurrencesPerElement[string(pair[0])] += occurrences
	}
	return occurrencesPerElement
}

func getOccurrencesPerPair(polymerTemplate []string) map[string]int {
	occurrencesPerPair := make(map[string]int)

	for i := 0; i < len(polymerTemplate)-1; i++ {
		pair := polymerTemplate[i] + polymerTemplate[i+1]
		occurrencesPerPair[pair]++
	}
	return occurrencesPerPair
}

func getMaxAndMinOccurrences(occurrencesPerElement map[string]int) (int, int) {
	maxOccurrences := 0
	minOccurrences := math.MaxInt

	for _, occurrences := range occurrencesPerElement {
		if occurrences < minOccurrences {
			minOccurrences = occurrences
		}
		if occurrences > maxOccurrences {
			maxOccurrences = occurrences
		}
	}
	return maxOccurrences, minOccurrences
}

func parsePolymerInput(input []string) ([]string, map[string]string) {
	polymerTemplate := strings.Split(input[0], "")

	pairInsertionRules := make(map[string]string)
	for _, rule := range input {
		if strings.Contains(rule, " -> ") {
			split := strings.Split(rule, " -> ")
			pairInsertionRules[split[0]] = split[1]
		}
	}

	return polymerTemplate, pairInsertionRules
}
