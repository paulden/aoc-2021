package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestDay23ParseALUInstructions(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day24_example1.txt")
	expectedInstruction0 := ALUInstruction{"inp", "z", ""}
	expectedInstruction1 := ALUInstruction{"inp", "x", ""}
	expectedInstruction2 := ALUInstruction{"mul", "z", "3"}
	expectedInstruction3 := ALUInstruction{"eql", "z", "x"}

	// When
	instructions := ParseALUInstructions(input)

	// Then
	if instructions[0] != expectedInstruction0 {
		t.Errorf("Day 24 - Parsing instructions: expected %v, got %v", expectedInstruction0, instructions[0])
	}
	if instructions[1] != expectedInstruction1 {
		t.Errorf("Day 24 - Parsing instructions: expected %v, got %v", expectedInstruction1, instructions[1])
	}
	if instructions[2] != expectedInstruction2 {
		t.Errorf("Day 24 - Parsing instructions: expected %v, got %v", expectedInstruction2, instructions[2])
	}
	if instructions[3] != expectedInstruction3 {
		t.Errorf("Day 24 - Parsing instructions: expected %v, got %v", expectedInstruction3, instructions[3])
	}
}

//func TestDay23RunALUProgram1(t *testing.T) {
//	// Given
//	inputProgram := readStringsInFile("../data/day24_example1.txt")
//	inputSequence1 := []int{2, 4}
//	inputSequence2 := []int{2, 6}
//
//	// When
//	_, x1, _, z1 := RunALUProgram(inputProgram, inputSequence1)
//	_, x2, _, z2 := RunALUProgram(inputProgram, inputSequence2)
//
//	if x1 != 4 {
//		t.Errorf("Day 24 - example1: expected x1 to be %v, got %v", 4, x1)
//	}
//	if z1 != 0 {
//		t.Errorf("Day 24 - example1: expected z1 to be %v, got %v", 0, z1)
//	}
//	if x2 != 6 {
//		t.Errorf("Day 24 - example1: expected x2 to be %v, got %v", 4, x2)
//	}
//	if z2 != 1 {
//		t.Errorf("Day 24 - example1: expected z2 to be %v, got %v", 0, z2)
//	}
//}

//func TestDay23RunALUProgram2(t *testing.T) {
//	// Given
//	inputProgram := readStringsInFile("../data/day24_example2.txt")
//	inputSequence1 := []int{7}
//	inputSequence2 := []int{11}
//
//	// When
//	w1, x1, y1, z1 := RunALUProgram(inputProgram, inputSequence1)
//	w2, x2, y2, z2 := RunALUProgram(inputProgram, inputSequence2)
//
//	if w1 != 0 && x1 != 1 && y1 != 1 && z1 != 1 {
//		t.Errorf("Day 24 - example2: error converting first decimal input to bits")
//	}
//	if w2 != 1 && x2 != 0 && y2 != 1 && z2 != 0 {
//		t.Errorf("Day 24 - example2: error converting second decimal input to bits")
//	}
//}

//func TestDay23RunALUProgramReal(t *testing.T) {
//	// Given
//	inputProgram := readStringsInFile("../data/day24.txt")
//	inputSequence := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 9, 9, 9, 9, 9}
//
//	//TryStuff(inputProgram)
//	//fmt.Println(numbers[0])
//
//	// When
//	now := time.Now()
//	w, x, y, z := RunALUProgram(inputProgram, inputSequence)
//	fmt.Printf("Elapsed time: %v\n", time.Since(now))
//
//	fmt.Printf("w: %v, x: %v, y: %v, z: %v\n", w, x, y, z)
//}

func TestDay23RunALUPart14(t *testing.T) {
	// Given
	inputProgram := readStringsInFile("../data/day24.txt")
	//inputSequence := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	// Just try every combination of this "div z 1" (so 9^7) since "div z 26" should be fixed to "z(i-1) + (add x <value>) % 26" if possible
	// I'll clean it out after Christmas!
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%v%v\n", i, j)
			for k := 1; k < 10; k++ {
				for l := 1; l < 10; l++ {
					for m := 1; m < 10; m++ {
						for n := 1; n < 10; n++ {
							for o := 1; o < 10; o++ {
								FindStuff(inputProgram, []int{i, j, k, l, m, n, o})
								//_, _, _, z := RunALUProgram(inputProgram, []int{i, j, k, l})
								//fmt.Printf("z: %v\n", z)
							}
						}
					}
				}
			}
		}
	}

	// When
	now := time.Now()
	//w, x, y, z := RunALUProgram(inputProgram, inputSequence)
	fmt.Printf("Elapsed time: %v\n", time.Since(now))

	//fmt.Printf("w: %v, x: %v, y: %v, z: %v\n", w, x, y, z)
}

func FindStuff(program []string, sequence []int) bool {
	s := make([]int, 0)
	w, x, y, z := 0, 0, 0, 0
	currentInputOne := 0

	for i := 1; i <= 14; i++ {
		path := "../data/day24_" + strconv.Itoa(i) + ".txt"
		inputProgram := readStringsInFile(path)
		instructions := ParseALUInstructions(inputProgram)
		if instructions[4].b == "1" {
			s = append(s, sequence[currentInputOne])
			w, x, y, z = RunALUProgram(instructions, []int{sequence[currentInputOne]}, w, x, y, z)
			currentInputOne++
		}  else {
			t, _ := strconv.ParseInt(instructions[5].b, 10, 64)
			inputToPass := (z + int(t)) % 26
			if inputToPass > 9 || inputToPass < 1 {
				return false
			}
			s = append(s, inputToPass)
			w, x, y, z = RunALUProgram(instructions, []int{inputToPass}, w, x, y, z)
		}
	}

	if z == 0 {
		fmt.Println(s)
		return true
	}
	return false
}
