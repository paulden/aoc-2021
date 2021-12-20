package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Part 1

func SnailfishAddition(input []string) (*Number, error) {
	result := ParseSnailfishNumber(input[0])
	var err error

	for _, stringNumber := range input[1:] {
		number := ParseSnailfishNumber(stringNumber)
		result, err = result.Add(number)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (n *Number) ComputeMagnitude() int {
	if n.IsRegular() {
		return n.value
	}
	return 3*n.left.ComputeMagnitude() + 2*n.right.ComputeMagnitude()
}

// Part 2

func GetMaximumMagnitude(input []string) int {
	maxMagnitude := 0

	for i, number1 := range input {
		for j, number2 := range input {
			if j != i {
				snailfishAddition, _ := SnailfishAddition([]string{number1, number2})
				magnitude := snailfishAddition.ComputeMagnitude()
				if magnitude > maxMagnitude {
					maxMagnitude = magnitude
				}
			}
		}
	}

	return maxMagnitude
}

// Data structure to be used

type Number struct {
	value               int
	left, right, parent *Number
}

// Parsing string to data structure

func ParseSnailfishNumber(input string) *Number {
	if strings.Count(input, ",") == 1 {
		// input is such as "[a,b]"
		split := strings.Split(input[1:len(input)-1], ",")
		leftInt, _ := strconv.ParseInt(split[0], 10, 64)
		rightInt, _ := strconv.ParseInt(split[1], 10, 64)
		number := NewSnailfishNumber(int(leftInt), int(rightInt))
		return &number
	}
	if strings.Count(input, ",") == 0 {
		// input is such as "a"
		numberValue, _ := strconv.ParseInt(input, 10, 64)
		number := Number{value: int(numberValue)}
		return &number
	}
	// input has nested pairs inside, such as "[[a,b],c]
	inPairDepth := 0
	var commaIndex int
	for i, char := range input[1 : len(input)-1] {
		element := string(char)
		if element == "[" {
			inPairDepth++
		}
		if element == "]" {
			inPairDepth--
		}
		if inPairDepth == 0 && element == "," {
			commaIndex = i
		}
	}
	firstPart := input[1 : commaIndex+1]
	secondPart := input[commaIndex+2 : len(input)-1]

	number1 := ParseSnailfishNumber(firstPart)
	number2 := ParseSnailfishNumber(secondPart)
	result := Number{-1, number1, number2, nil}

	// Updating parent pointers
	// I don't know (yet) why this is needed but the debugger brought me here, and I want to get out
	if !number1.IsRegular() {
		number1.left.parent = number1
		number1.right.parent = number1
	}
	if !number2.IsRegular() {
		number2.left.parent = number2
		number2.right.parent = number2
	}
	number1.parent = &result
	number2.parent = &result
	return &result
}

func NewSnailfishNumber(n1, n2 int) Number {
	regularNumber1 := Number{value: n1}
	regularNumber2 := Number{value: n2}
	snailfishNumber := Number{value: -1, left: &regularNumber1, right: &regularNumber2}
	regularNumber1.parent = &snailfishNumber
	regularNumber2.parent = &snailfishNumber
	return snailfishNumber
}

func (n *Number) Add(secondNumber *Number) (*Number, error) {
	result := Add(n, secondNumber)
	err := result.Reduce()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Add(n1, n2 *Number) *Number {
	additionNumber := Number{value: -1, left: n1, right: n2}
	// Updating pointers to parents
	n1.left.parent = n1
	n2.left.parent = n2
	n1.right.parent = n1
	n2.right.parent = n2
	n1.parent = &additionNumber
	n2.parent = &additionNumber
	return &additionNumber
}

// Reduce

func (n *Number) Reduce() error {
	shouldCheckExplode := true
	shouldCheckSplit := true

	for shouldCheckExplode || shouldCheckSplit {
		explodeErr := n.Explode()
		if explodeErr != nil {
			shouldCheckExplode = false
			splitErr := n.Split()
			if splitErr != nil {
				shouldCheckSplit = false
			} else {
				shouldCheckExplode, shouldCheckSplit = true, true
			}
		}
	}
	return nil
}

// Explosion part

func (n *Number) Explode() error {
	pairThatShouldExplode, err := n.PairThatShouldExplode(0)
	if err != nil {
		return err
	}
	//fmt.Printf("Exploding pair %v with left %v and right %v!\n", pairThatShouldExplode, pairThatShouldExplode.left.value, pairThatShouldExplode.right.value)
	explodedNumber := Number{0, nil, nil, pairThatShouldExplode.parent}

	rightNeighbour, rightErr := pairThatShouldExplode.FindRightRegularNeighbour()
	if rightErr == nil {
		rightNeighbour.value += pairThatShouldExplode.right.value
	}
	leftNeighbour, leftErr := pairThatShouldExplode.FindLeftRegularNeighbour()
	if leftErr == nil {
		leftNeighbour.value += pairThatShouldExplode.left.value
	}

	if pairThatShouldExplode.IsRight() {
		pairThatShouldExplode.parent.right = &explodedNumber
	} else if pairThatShouldExplode.IsLeft() {
		pairThatShouldExplode.parent.left = &explodedNumber
	}

	return nil
}

func (n *Number) PairThatShouldExplode(currentDepth int) (*Number, error) {
	if n.left.IsRegular() && n.right.IsRegular() {
		if currentDepth >= 4 {
			return n, nil
		} else {
			return nil, errors.New("no pair to explode")
		}
	}
	currentDepth++
	if !n.left.IsRegular() {
		potentialLeftPairToExplode, errLeft := n.left.PairThatShouldExplode(currentDepth)
		if errLeft == nil {
			return potentialLeftPairToExplode, nil
		}
	}
	if !n.right.IsRegular() {
		potentialRightPairToExplode, errRight := n.right.PairThatShouldExplode(currentDepth)
		if errRight == nil {
			return potentialRightPairToExplode, nil
		}
	}
	return nil, errors.New("no pair to explode")
}

// Split part

func (n *Number) Split() error {
	numberToSplit, err := n.NumberThatShouldSplit()

	if err != nil {
		return err
	}

	valueToSplit := numberToSplit.value

	leftNumber := Number{valueToSplit / 2, nil, nil, numberToSplit}
	rightNumber := Number{(valueToSplit + 1) / 2, nil, nil, numberToSplit}

	numberToSplit.left = &leftNumber
	numberToSplit.right = &rightNumber
	numberToSplit.value = -1

	return nil
}

func (n *Number) NumberThatShouldSplit() (*Number, error) {
	if n.IsRegular() {
		if n.value >= 10 {
			return n, nil
		} else {
			return nil, errors.New("no number to split")
		}
	}
	potentialLeftSplit, leftError := n.left.NumberThatShouldSplit()
	if leftError == nil {
		return potentialLeftSplit, nil
	}
	potentialRightSplit, rightError := n.right.NumberThatShouldSplit()
	if rightError == nil {
		return potentialRightSplit, nil
	}
	return nil, errors.New("no number to split")
}

// Navigate through "tree" number

func (n *Number) FindLeftmostRegularNumber() (*Number, error) {
	if n.IsRegular() {
		return nil, errors.New("already a regular number")
	}
	if n.left.IsRegular() {
		return n.left, nil
	}
	return n.left.FindLeftmostRegularNumber()
}

func (n *Number) FindRightmostRegularNumber() (*Number, error) {
	if n.IsRegular() {
		return nil, errors.New("already a regular number")
	}
	if n.right.IsRegular() {
		return n.right, nil
	}
	return n.right.FindRightmostRegularNumber()
}

func (n *Number) FindLeftRegularNeighbour() (*Number, error) {
	if n.IsRight() {
		if n.parent.left.IsRegular() {
			return n.parent.left, nil
		}
		return n.parent.left.FindRightmostRegularNumber()
	}
	if n.IsRoot() {
		return nil, errors.New("already leftmost regular number")
	}
	return n.parent.FindLeftRegularNeighbour()
}

func (n *Number) FindRightRegularNeighbour() (*Number, error) {
	if n.IsLeft() {
		if n.parent.right.IsRegular() {
			return n.parent.right, nil
		}
		return n.parent.right.FindLeftmostRegularNumber()
	}
	if n.IsRoot() {
		return nil, errors.New("already rightmost regular number")
	}
	return n.parent.FindRightRegularNeighbour()
}

// Basic utils

func (n *Number) IsRegular() bool {
	return n.value != -1
}

func (n *Number) IsRoot() bool {
	return n.parent == nil
}

func (n *Number) IsLeft() bool {
	if n.IsRoot() {
		return false
	}
	return *n.parent.left == *n
}

func (n *Number) IsRight() bool {
	if n.IsRoot() {
		return false
	}
	return *n.parent.right == *n
}

// Printing

func (n *Number) Sprint() string {
	if n.IsRegular() {
		return fmt.Sprintf("%v", n.value)
	} else {
		return fmt.Sprintf("[%v,%v]", n.left.Sprint(), n.right.Sprint())
	}
}
