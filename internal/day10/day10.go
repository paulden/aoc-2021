package day10

import (
	"fmt"
	"sort"
)

// Part 1

func GetSyntaxErrorScore(report []string) int {
	syntaxErrorScore := 0
	pointsPerInvalidDelimiter := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}

	for _, line := range report {
		_, err := parseLineChunks(line)
		if syntaxError, ok := err.(*syntaxError); ok {
			syntaxErrorScore += pointsPerInvalidDelimiter[syntaxError.invalidCharacter]
		}
	}
	return syntaxErrorScore
}

// Part 2

func GetAutocompletionScore(report []string) int {
	autocompletionScores := make([]int, 0)
	pointsPerDelimiter := map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}

	for _, line := range report {
		openingDelimiters, err := parseLineChunks(line)
		if err == nil {
			score := 0
			for i := len(openingDelimiters) - 1; i >= 0; i-- {
				delimiter := openingDelimiters[i]
				score = score * 5
				score += pointsPerDelimiter[delimiter]
			}
			autocompletionScores = append(autocompletionScores, score)
		}
	}

	sort.Ints(autocompletionScores)
	medianIndex := (len(autocompletionScores) - 1) / 2
	return autocompletionScores[medianIndex]
}

// Utils

func parseLineChunks(line string) ([]string, error) {
	var openingDelimiters []string

	for _, char := range line {
		delimiter := string(char)
		var latestDelimiter string
		if len(openingDelimiters) > 0 {
			latestDelimiter = openingDelimiters[len(openingDelimiters)-1]
		}
		if delimiter == "{" || delimiter == "(" || delimiter == "[" || delimiter == "<" {
			openingDelimiters = append(openingDelimiters, delimiter)
		} else if delimiter == "}" && latestDelimiter != "{" {
			return openingDelimiters, &syntaxError{delimiter, "unexpected character"}
		} else if delimiter == ">" && latestDelimiter != "<" {
			return openingDelimiters, &syntaxError{delimiter, "unexpected character"}
		} else if delimiter == ")" && latestDelimiter != "(" {
			return openingDelimiters, &syntaxError{delimiter, "unexpected character"}
		} else if delimiter == "]" && latestDelimiter != "[" {
			return openingDelimiters, &syntaxError{delimiter, "unexpected character"}
		} else {
			openingDelimiters = openingDelimiters[:len(openingDelimiters)-1]
		}
	}

	return openingDelimiters, nil
}

type syntaxError struct {
	invalidCharacter string
	message          string
}

func (e *syntaxError) Error() string {
	return fmt.Sprintf("%s: %s", e.invalidCharacter, e.message)
}
