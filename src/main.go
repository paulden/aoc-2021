package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	day1 := readIntegersInFile("data/day1.txt")
	increases := CountIncreases(day1)
	fmt.Printf("Day 1 - Part 1: %v\n", increases)
	threeMeasurementsIncreases := CountThreeMeasurementsIncreases(day1)
	fmt.Printf("Day 1 - Part 2: %v\n", threeMeasurementsIncreases)

	day2 := readStringsInFile("data/day2.txt")
	forwardPosition, depthPosition := determinePosition(day2)
	fmt.Printf("Day 2 - Part 1: %v\n", forwardPosition*depthPosition)
	forwardPositionWithAim, depthPositionWithAim := determinePositionWithAim(day2)
	fmt.Printf("Day 2 - Part 2: %v\n", forwardPositionWithAim*depthPositionWithAim)

	day3 := readStringsInFile("data/day3.txt")
	gammaRate := ComputeGamma(day3)
	epsilonRate := ComputeEpsilon(day3)
	fmt.Printf("Day 3 - Part 1: %v\n", gammaRate*epsilonRate)
	oxygenRating := ComputeOxygenGeneratorRating(day3)
	co2Rating := ComputeCO2ScrubberRating(day3)
	fmt.Printf("Day 3 - Part 2: %v\n", oxygenRating*co2Rating)

	day4 := readStringsInFile("data/day4.txt")
	winningScore := GetWinningBingoCardScore(day4)
	fmt.Printf("Day 4 - Part 1: %v\n", winningScore)
	losingScore := GetLosingBingoCardScore(day4)
	fmt.Printf("Day 4 - Part 2: %v\n", losingScore)

	day5 := readStringsInFile("data/day5.txt")
	dangerousZonesNumber := GetDangerousZonesNumber(day5)
	fmt.Printf("Day 5 - Part 1: %v\n", dangerousZonesNumber)
	dangerousZonesNumberWithDiagonals := GetDangerousZonesNumberWithDiagonals(day5)
	fmt.Printf("Day 5 - Part 2: %v\n", dangerousZonesNumberWithDiagonals)

	day6 := readStringsInFile("data/day6.txt")
	lanterfishesNumber := CountLanternfishesNaive(day6[0], 80)
	fmt.Printf("Day 6 - Part 1: %v\n", lanterfishesNumber)
	lanterfishesNumber256 := CountLanternfishesOptimized(day6[0], 256)
	fmt.Printf("Day 6 - Part 2: %v\n", lanterfishesNumber256)

	day7 := readStringsInFile("data/day7.txt")
	fuelConsumption := GetCheapestFuelConsumption(day7[0])
	fmt.Printf("Day 7 - Part 1: %v\n", fuelConsumption)
	fuelConsumptionUpdated := GetCheapestFuelConsumptionUpdated(day7[0])
	fmt.Printf("Day 7 - Part 2: %v\n", fuelConsumptionUpdated)

	day8 := readStringsInFile("data/day8.txt")
	uniqueSegmentsDigits := CountUniqueSegmentsDigits(day8)
	fmt.Printf("Day 8 - Part 1: %v\n", uniqueSegmentsDigits)
	sum := SumOutputDisplays(day8)
	fmt.Printf("Day 8 - Part 2: %v\n", sum)

	day9 := readStringsInFile("data/day9.txt")
	riskLevel := GetSmokeRiskLevel(day9)
	fmt.Printf("Day 9 - Part 1: %v\n", riskLevel)
	biggestBasinsSize := GetBiggestBasins(day9)
	fmt.Printf("Day 9 - Part 2: %v\n", biggestBasinsSize)

	day10 := readStringsInFile("data/day10.txt")
	syntaxScoreError := GetSyntaxErrorScore(day10)
	fmt.Printf("Day 10 - Part 1: %v\n", syntaxScoreError)
	autocompletionScore := GetAutocompletionScore(day10)
	fmt.Printf("Day 10 - Part 2: %v\n", autocompletionScore)

	day11 := readStringsInFile("data/day11.txt")
	flashes := CountFlashes(day11)
	fmt.Printf("Day 11 - Part 1: %v\n", flashes)
	mindBlowingStep := FindMindBlowingStep(day11)
	fmt.Printf("Day 11 - Part 2: %v\n", mindBlowingStep)

	day12 := readStringsInFile("data/day12.txt")
	paths := CountCavePathsPart1(day12)
	fmt.Printf("Day 12 - Part 1: %v\n", paths)
	pathsVisitingTwice := CountCavePathsPart2(day12)
	fmt.Printf("Day 12 - Part 2: %v\n", pathsVisitingTwice)

	day13 := readStringsInFile("data/day13.txt")
	paperFoldedOnce := FoldPaperOnce(day13)
	fmt.Printf("Day 13 - Part 1: %v\n", len(paperFoldedOnce))
	paperFoldedCompletely := FoldPaperCompletely(day13)
	fmt.Printf("Day 13 - Part 2:\n")
	PrettyPrintDots(paperFoldedCompletely)

	day14 := readStringsInFile("data/day14.txt")
	polymerCountsDifference := CountPolymerCountsDifferenceNaive(day14, 10)
	fmt.Printf("Day 14 - Part 1: %v\n", polymerCountsDifference)
	polymerCountsDifferenceAfter40 := CountPolymerCountsDifferenceOptimized(day14, 40)
	fmt.Printf("Day 14 - Part 2: %v\n", polymerCountsDifferenceAfter40)

	day15 := readStringsInFile("data/day15.txt")
	lowestTotalRiskPath := GetLowestTotalRiskPath(day15)
	fmt.Printf("Day 15 - Part 1: %v\n", lowestTotalRiskPath)
	lowestTotalRiskPathInRealMap := GetLowestTotalRiskPathInRealMap(day15)
	fmt.Printf("Day 15 - Part 2: %v\n", lowestTotalRiskPathInRealMap)

	day16 := readStringsInFile("data/day16.txt")
	versionsSum := SumVersionNumbers(day16[0])
	fmt.Printf("Day 16 - Part 1: %v\n", versionsSum)
	evaluatedBITSExpression := EvaluateBITSExpression(day16[0])
	fmt.Printf("Day 16 - Part 2: %v\n", evaluatedBITSExpression)

	day17 := readStringsInFile("data/day17.txt")
	maxHeight := GetMaxHeight(day17[0])
	fmt.Printf("Day 17 - Part 1: %v\n", maxHeight)
	allVelocities := GetAllVelocities(day17[0])
	fmt.Printf("Day 17 - Part 2: %v\n", len(allVelocities))

	day18 := readStringsInFile("data/day18.txt")
	snailfishAddition, _ := SnailfishAddition(day18)
	magnitude := snailfishAddition.ComputeMagnitude()
	fmt.Printf("Day 18 - Part 1: %v\n", magnitude)
	maximumMagnitude := GetMaximumMagnitude(day18)
	fmt.Printf("Day 18 - Part 2: %v\n", maximumMagnitude)
}

func readIntegersInFile(filePath string) []int {
	file, _ := os.Open(filePath)

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		integer, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, int(integer))
	}

	return lines
}

func readStringsInFile(filePath string) []string {
	file, _ := os.Open(filePath)

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
