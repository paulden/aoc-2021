package day24

import (
	"aoc-2021/internal/pkg/utils"
	"errors"
	"strconv"
	"strings"
)

func FindLargestModelNumber(program []string) int {
	sequences := parseSequences(program)
	pushCombinations := getDifferentPushCombinations()

	for i := len(pushCombinations) - 1; i >= 0; i-- {
		finalCombination, err := isValidCombination(pushCombinations[i], sequences)
		if err == nil {
			return utils.SliceToInteger(finalCombination)
		}
	}
	return 0
}

func FindSmallestModelNumber(program []string) int {
	sequences := parseSequences(program)
	pushCombinations := getDifferentPushCombinations()

	for _, combination := range pushCombinations {
		finalCombination, err := isValidCombination(combination, sequences)
		if err == nil {
			return utils.SliceToInteger(finalCombination)
		}
	}
	return 0
}

func parseSequences(program []string) [][]string {
	sequences := make([][]string, 0)
	sequenceNumber := 14
	sequenceLength := 18

	for i := 0; i < sequenceNumber; i++ {
		sequence := make([]string, sequenceLength)
		for j := 0; j < sequenceLength; j++ {
			sequence[j] = program[i*sequenceLength+j]
		}
		sequences = append(sequences, sequence)
	}
	return sequences
}

func isValidCombination(combination []int, sequences [][]string) ([]int, error) {
	finalCombination := make([]int, 0)
	currentInputPointer := 0
	r := registry{0, 0, 0, 0}

	for _, sequence := range sequences {
		if isPushSequence(sequence) {
			finalCombination = append(finalCombination, combination[currentInputPointer])
			r = *runALUSequence(sequence, combination[currentInputPointer], &r)
			currentInputPointer++
		} else {
			input := computePopInput(sequence, r)
			if input > 9 || input < 1 {
				return finalCombination, errors.New("required input is out of bounds")
			}
			finalCombination = append(finalCombination, input)
			r = *runALUSequence(sequence, input, &r)
		}
	}

	if r.z == 0 {
		return finalCombination, nil
	}
	return finalCombination, errors.New("could not get z equal to zero")
}
func isPushSequence(sequence []string) bool {
	return sequence[4][len(sequence[4])-1:] == "1"
}

func computePopInput(sequence []string, r registry) int {
	split := strings.Split(sequence[5], " ")
	value, _ := strconv.ParseInt(split[2], 10, 64)
	input := (r.z + int(value)) % 26
	return input
}

func getDifferentPushCombinations() [][]int {
	combinations := make([][]int, 0)

	// I am sorry about that
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			for k := 1; k < 10; k++ {
				for l := 1; l < 10; l++ {
					for m := 1; m < 10; m++ {
						for n := 1; n < 10; n++ {
							for o := 1; o < 10; o++ {
								combinations = append(combinations, []int{i, j, k, l, m, n, o})
							}
						}
					}
				}
			}
		}
	}

	return combinations
}

type registry struct {
	w, x, y, z int
}

func runALUSequence(program []string, input int, r *registry) *registry {
	pointer := 0

	for _, instruction := range program {
		args := strings.Split(instruction, " ")

		var a *int
		var b *int

		switch args[1] {
		case "w":
			a = &r.w
		case "x":
			a = &r.x
		case "y":
			a = &r.y
		case "z":
			a = &r.z
		}

		if args[0] == "inp" {
			*a = input
			pointer++
			continue
		}

		switch args[2] {
		case "w":
			b = &r.w
		case "x":
			b = &r.x
		case "y":
			b = &r.y
		case "z":
			b = &r.z
		}

		var secondTerm int
		i, err := strconv.ParseInt(args[2], 10, 64)
		if err == nil {
			secondTerm = int(i)
		} else {
			secondTerm = *b
		}

		if args[0] == "add" {
			*a += secondTerm
		}
		if args[0] == "mul" {
			*a *= secondTerm
		}
		if args[0] == "div" {
			*a = *a / secondTerm
		}
		if args[0] == "mod" {
			*a = *a % secondTerm
		}
		if args[0] == "eql" {
			if *a == secondTerm {
				*a = 1
			} else {
				*a = 0
			}
		}
	}

	return r
}
