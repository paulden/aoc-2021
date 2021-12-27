package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntegersInFile(filePath string) []int {
	file, _ := os.Open(filePath)

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		integer, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, int(integer))
	}

	return lines
}

func ReadStringsInFile(filePath string) []string {
	file, _ := os.Open(filePath)

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
