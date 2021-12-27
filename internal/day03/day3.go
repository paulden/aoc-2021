package day03

import (
	"strconv"
	"strings"
)

// Part 1

func GetPowerConsumption(input []string) int {
	return computeGamma(input) * computeEpsilon(input)
}

func computeGamma(report []string) int {
	mostCommonBits, _ := getCommonBits(report)

	gammaRate, _ := strconv.ParseInt(strings.Join(mostCommonBits, ""), 2, 64)
	return int(gammaRate)
}

func computeEpsilon(report []string) int {
	_, leastCommonBits := getCommonBits(report)

	epsilonRate, _ := strconv.ParseInt(strings.Join(leastCommonBits, ""), 2, 64)
	return int(epsilonRate)
}

// Part 2

func GetLifeSupportRating(input []string) int {
	return computeOxygenGeneratorRating(input) * computeCO2ScrubberRating(input)
}

func computeOxygenGeneratorRating(report []string) int {
	remainingSequences := report

	for i := 0; i < len(report[0]); i++ {
		mostCommonBits, _ := getCommonBits(remainingSequences)
		newRemainingReports := make([]string, 0)

		for _, sequence := range remainingSequences {
			currentBit := string(sequence[i])
			if currentBit == mostCommonBits[i] {
				newRemainingReports = append(newRemainingReports, sequence)
			}
		}

		remainingSequences = newRemainingReports

		if len(remainingSequences) <= 1 {
			break
		}
	}
	oxygen, _ := strconv.ParseInt(remainingSequences[0], 2, 64)
	return int(oxygen)
}

func computeCO2ScrubberRating(report []string) int {
	remainingReports := report

	for i := 0; i < len(report[0]); i++ {
		_, leastCommonBits := getCommonBits(remainingReports)
		newRemainingReports := make([]string, 0)

		for _, sequence := range remainingReports {
			currentBit := string(sequence[i])
			if currentBit == leastCommonBits[i] {
				newRemainingReports = append(newRemainingReports, sequence)
			}
		}

		remainingReports = newRemainingReports

		if len(remainingReports) <= 1 {
			break
		}
	}
	oxygen, _ := strconv.ParseInt(remainingReports[0], 2, 64)
	return int(oxygen)
}

// Utils

func getCommonBits(report []string) ([]string, []string) {
	bitsLength := len(report[0])
	mostCommonBits := make([]string, bitsLength)
	leastCommonBits := make([]string, bitsLength)

	for i := 0; i < bitsLength; i++ {
		zeros := 0
		ones := 0

		for _, sequence := range report {
			if string(sequence[i]) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if zeros > ones {
			mostCommonBits[i], leastCommonBits[i] = "0", "1"
		} else {
			mostCommonBits[i], leastCommonBits[i] = "1", "0"
		}
	}
	return mostCommonBits, leastCommonBits
}
