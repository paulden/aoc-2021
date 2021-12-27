package day24

import (
	"strconv"
	"strings"
)

// TODO: Refacto this package!

func GenerateAllPossibleNumbers() (sequences [][]int) {
	for i := 11111111111111; i <= 99999999999999; i++ {
		sequence := make([]int, 0)
		for k := 1; k <= 14; k++ {
			sequence = append(sequence, (i/k)%10)
		}
		isValid := true
		for _, j := range sequence {
			if j == 0 {
				isValid = false
			}
		}
		if isValid {
			sequences = append(sequences, sequence)
		}
	}
	return
}

func RunALUProgram(instructions []ALUInstruction, sequence []int, w, x, y, z int) (int, int, int, int) {
	//z = 2
	pointer := 0

	for _, instruction := range instructions {
		var a *int
		var b *int

		switch instruction.a {
		case "w":
			a = &w
		case "x":
			a = &x
		case "y":
			a = &y
		case "z":
			a = &z
		}
		switch instruction.b {
		case "w":
			b = &w
		case "x":
			b = &x
		case "y":
			b = &y
		case "z":
			b = &z
		}

		if instruction.isInp() {
			//fmt.Printf("\nAssigning input %v: %v\n", sequence[pointer], instruction)
			*a = sequence[pointer]
			pointer++
			continue
		}

		var secondTerm int
		i, err := strconv.ParseInt(instruction.b, 10, 64)
		if err == nil {
			secondTerm = int(i)
		} else {
			secondTerm = *b
		}

		if instruction.isAdd() {
			*a += secondTerm
		}
		if instruction.isMul() {
			*a *= secondTerm
		}
		if instruction.isDiv() {
			*a = *a / secondTerm
		}
		if instruction.isMod() {
			*a = *a % secondTerm
		}
		if instruction.isEql() {
			if *a == secondTerm {
				*a = 1
			} else {
				*a = 0
			}
		}
		//fmt.Printf("%v - %v, %v, %v, %v\n", instruction, w, x, y, z)
	}

	return w, x, y, z
}

func ParseALUInstructions(input []string) (instructions []ALUInstruction) {
	for _, instruction := range input {
		split := strings.Split(instruction, " ")
		var newInstruction ALUInstruction
		if len(split) >= 3 {
			newInstruction = ALUInstruction{split[0], split[1], split[2]}
		} else {
			newInstruction = ALUInstruction{split[0], split[1], ""}
		}
		instructions = append(instructions, newInstruction)
	}
	return
}

type ALUInstruction struct {
	instruction string
	a, b        string
}

func (i *ALUInstruction) isInp() bool {
	return i.instruction == "inp"
}

func (i *ALUInstruction) isAdd() bool {
	return i.instruction == "add"
}

func (i *ALUInstruction) isMul() bool {
	return i.instruction == "mul"
}

func (i *ALUInstruction) isDiv() bool {
	return i.instruction == "div"
}

func (i *ALUInstruction) isMod() bool {
	return i.instruction == "mod"
}

func (i *ALUInstruction) isEql() bool {
	return i.instruction == "eql"
}
