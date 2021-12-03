package main

import (
	"strconv"
	"strings"
)

func ComputeGamma(report []string) int {
	mostCommonBits, _ := GetCommonBits(report)

	gammaRate, _ := strconv.ParseInt(strings.Join(mostCommonBits, ""), 2, 64)
	return int(gammaRate)
}

func ComputeEpsilon(report []string) int {
	_, leastCommonBits := GetCommonBits(report)

	epsilonRate, _ := strconv.ParseInt(strings.Join(leastCommonBits, ""), 2, 64)
	return int(epsilonRate)
}

func ComputeOxygenGeneratorRating(report []string) int {
	remainingSequences := report

	for i := 0; i < len(report[0]); i++ {
		mostCommonBits, _ := GetCommonBits(remainingSequences)
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

func ComputeCO2ScrubberRating(report []string) int {
	remainingReports := report

	for i := 0; i < len(report[0]); i++ {
		_, leastCommonBits := GetCommonBits(remainingReports)
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

func GetCommonBits(report []string) ([]string, []string) {
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