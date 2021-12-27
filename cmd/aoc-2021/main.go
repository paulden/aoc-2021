package main

import (
	"aoc-2021/internal/day01"
	"aoc-2021/internal/day02"
	"aoc-2021/internal/day03"
	"aoc-2021/internal/day04"
	"aoc-2021/internal/day05"
	"aoc-2021/internal/day06"
	"aoc-2021/internal/day07"
	"aoc-2021/internal/day08"
	"aoc-2021/internal/day09"
	"aoc-2021/internal/day10"
	"aoc-2021/internal/day11"
	"aoc-2021/internal/day12"
	"aoc-2021/internal/day13"
	"aoc-2021/internal/day14"
	"aoc-2021/internal/day15"
	"aoc-2021/internal/day16"
	"aoc-2021/internal/day17"
	"aoc-2021/internal/day18"
	"aoc-2021/internal/day19"
	"aoc-2021/internal/day20"
	"aoc-2021/internal/day21"
	"aoc-2021/internal/day22"
	"aoc-2021/internal/pkg/utils"
	"fmt"
)

func main() {
	input1 := utils.ReadIntegersInFile("assets/day1.txt")
	day1Part1 := day01.CountIncreases(input1)
	day1Part2 := day01.CountThreeMeasurementsIncreases(input1)
	fmt.Println("--- Day 1: Sonar Sweep ---")
	fmt.Printf("Part 1: %v\n", day1Part1)
	fmt.Printf("Part 2: %v\n", day1Part2)

	input2 := utils.ReadStringsInFile("assets/day2.txt")
	day2Part1 := day02.DeterminePosition(input2)
	day2Part2 := day02.DeterminePositionWithAim(input2)
	fmt.Println("--- Day 2: Dive! ---")
	fmt.Printf("Part 1: %v\n", day2Part1)
	fmt.Printf("Part 2: %v\n", day2Part2)

	input3 := utils.ReadStringsInFile("assets/day3.txt")
	day3Part1 := day03.GetPowerConsumption(input3)
	day3Part2 := day03.GetLifeSupportRating(input3)
	fmt.Println("--- Day 3: Binary Diagnostic ---")
	fmt.Printf("Part 1: %v\n", day3Part1)
	fmt.Printf("Part 2: %v\n", day3Part2)

	input4 := utils.ReadStringsInFile("assets/day4.txt")
	day4Part1 := day04.GetWinningBingoCardScore(input4)
	day4Part2 := day04.GetLosingBingoCardScore(input4)
	fmt.Println("--- Day 4: Giant Squid ---")
	fmt.Printf("Part 1: %v\n", day4Part1)
	fmt.Printf("Part 2: %v\n", day4Part2)

	input5 := utils.ReadStringsInFile("assets/day5.txt")
	day5Part1 := day05.GetDangerousZonesNumber(input5)
	day5Part2 := day05.GetDangerousZonesNumberWithDiagonals(input5)
	fmt.Println("--- Day 5: Hydrothermal Venture ---")
	fmt.Printf("Part 1: %v\n", day5Part1)
	fmt.Printf("Part 2: %v\n", day5Part2)

	input6 := utils.ReadStringsInFile("assets/day6.txt")
	day6Part1 := day06.CountLanternfishesOptimized(input6[0], 80)
	day6Part2 := day06.CountLanternfishesOptimized(input6[0], 256)
	fmt.Println("--- Day 6: Lanternfish ---")
	fmt.Printf("Part 1: %v\n", day6Part1)
	fmt.Printf("Part 2: %v\n", day6Part2)

	input7 := utils.ReadStringsInFile("assets/day7.txt")
	day7Part1 := day07.GetCheapestFuelConsumption(input7[0])
	day7Part2 := day07.GetCheapestFuelConsumption(input7[0])
	fmt.Println("--- Day 7: The Treachery of Whales ---")
	fmt.Printf("Part 1: %v\n", day7Part1)
	fmt.Printf("Part 2: %v\n", day7Part2)

	input8 := utils.ReadStringsInFile("assets/day8.txt")
	day8Part1 := day08.CountUniqueSegmentsDigits(input8)
	day8Part2 := day08.SumOutputDisplays(input8)
	fmt.Println("--- Day 8: Seven Segment Search ---")
	fmt.Printf("Part 1: %v\n", day8Part1)
	fmt.Printf("Part 2: %v\n", day8Part2)

	input9 := utils.ReadStringsInFile("assets/day9.txt")
	day9Part1 := day09.GetSmokeRiskLevel(input9)
	day9Part2 := day09.GetBiggestBasins(input9)
	fmt.Println("--- Day 9: Smoke Basin ---")
	fmt.Printf("Part 1: %v\n", day9Part1)
	fmt.Printf("Part 2: %v\n", day9Part2)

	input10 := utils.ReadStringsInFile("assets/day10.txt")
	day10Part1 := day10.GetSyntaxErrorScore(input10)
	day10Part2 := day10.GetAutocompletionScore(input10)
	fmt.Println("--- Day 10: Syntax Scoring ---")
	fmt.Printf("Part 1: %v\n", day10Part1)
	fmt.Printf("Part 2: %v\n", day10Part2)

	input11 := utils.ReadStringsInFile("assets/day11.txt")
	day11Part1 := day11.CountFlashes(input11)
	day11Part2 := day11.FindMindBlowingStep(input11)
	fmt.Println("--- Day 11: Dumbo Octopus ---")
	fmt.Printf("Part 1: %v\n", day11Part1)
	fmt.Printf("Part 2: %v\n", day11Part2)

	input12 := utils.ReadStringsInFile("assets/day12.txt")
	day12Part1 := day12.CountCavePathsPart1(input12)
	day12Part2 := day12.CountCavePathsPart2(input12)
	fmt.Println("--- Day 12: Passage Pathing ---")
	fmt.Printf("Part 1: %v\n", day12Part1)
	fmt.Printf("Part 2: %v\n", day12Part2)

	input13 := utils.ReadStringsInFile("assets/day13.txt")
	day13Part1 := day13.FoldPaperOnce(input13)
	day13Part2 := day13.FoldPaperCompletely(input13)
	fmt.Println("--- Day 13: Transparent Origami ---")
	fmt.Printf("Part 1: %v\n", day13Part1)
	fmt.Printf("Part 2:\n")
	day13.PrettyPrintDots(day13Part2)

	input14 := utils.ReadStringsInFile("assets/day14.txt")
	day14Part1 := day14.CountPolymerCountsDifferenceOptimized(input14, 10)
	day14Part2 := day14.CountPolymerCountsDifferenceOptimized(input14, 40)
	fmt.Println("--- Day 14: Extended Polymerization ---")
	fmt.Printf("Part 1: %v\n", day14Part1)
	fmt.Printf("Part 2: %v\n", day14Part2)

	input15 := utils.ReadStringsInFile("assets/day15.txt")
	day15Part1 := day15.GetLowestTotalRiskPath(input15)
	day15Part2 := day15.GetLowestTotalRiskPathInRealMap(input15)
	fmt.Println("--- Day 15: Chiton ---")
	fmt.Printf("Part 1: %v\n", day15Part1)
	fmt.Printf("Part 2: %v\n", day15Part2)

	input16 := utils.ReadStringsInFile("assets/day16.txt")
	day16Part1 := day16.SumVersionNumbers(input16[0])
	day16Part2 := day16.EvaluateBITSExpression(input16[0])
	fmt.Println("--- Day 16: Packet Decoder ---")
	fmt.Printf("Part 1: %v\n", day16Part1)
	fmt.Printf("Part 2: %v\n", day16Part2)

	input17 := utils.ReadStringsInFile("assets/day17.txt")
	day17Part1 := day17.GetMaxHeight(input17[0])
	day17Part2 := day17.CountAllVelocities(input17[0])
	fmt.Println("--- Day 17: Trick Shot ---")
	fmt.Printf("Part 1: %v\n", day17Part1)
	fmt.Printf("Part 2: %v\n", day17Part2)

	input18 := utils.ReadStringsInFile("assets/day18.txt")
	day18Part1 := day18.GetMagnitudeAfterAddition(input18)
	day18Part2 := day18.GetMaximumMagnitude(input18)
	fmt.Println("--- Day 18: Trick Shot ---")
	fmt.Printf("Part 1: %v\n", day18Part1)
	fmt.Printf("Part 2: %v\n", day18Part2)

	input19 := utils.ReadStringsInFile("assets/day19.txt")
	day19Part1 := day19.CountBeacons(input19)
	day19Part2 := day19.GetMaximumManhattanDistance(input19)
	fmt.Println("--- Day 19: Beacon Scanner ---")
	fmt.Printf("Part 1: %v\n", day19Part1)
	fmt.Printf("Part 2: %v\n", day19Part2)

	input20 := utils.ReadStringsInFile("assets/day20.txt")
	day20Part1 := day20.ProcessEnhancement(input20, 2)
	day20Part2 := day20.ProcessEnhancement(input20, 50)
	fmt.Println("--- Day 20: Beacon Scanner ---")
	fmt.Printf("Part 1: %v\n", day20Part1)
	fmt.Printf("Part 2: %v\n", day20Part2)

	input21 := utils.ReadStringsInFile("assets/day21.txt")
	day21Part1 := day21.PracticeDirac(input21)
	day21Part2 := day21.GetDiracWinnerVictories(input21)
	fmt.Println("--- Day 21: Dirac Dice ---")
	fmt.Printf("Part 1: %v\n", day21Part1)
	fmt.Printf("Part 2: %v\n", day21Part2)

	input22 := utils.ReadStringsInFile("assets/day22.txt")
	day22Part1 := day22.CountCubesOnNaive(input22)
	day22Part2 := day22.CountCubesOnOptimized(input22)
	fmt.Println("--- Day 22: Reactor Reboot ---")
	fmt.Printf("Part 1: %v\n", day22Part1)
	fmt.Printf("Part 2: %v\n", day22Part2)
}
