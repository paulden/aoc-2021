package day18

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Part 1

func GetMagnitudeAfterAddition(input []string) int {
	result, _ := snailfishAddition(input)
	return result.computeMagnitude()
}

func snailfishAddition(input []string) (*snailfishNumber, error) {
	result := parseSnailfishNumber(input[0])
	var err error

	for _, stringNumber := range input[1:] {
		number := parseSnailfishNumber(stringNumber)
		result, err = result.add(number)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (n *snailfishNumber) computeMagnitude() int {
	if n.isRegular() {
		return n.value
	}
	return 3*n.left.computeMagnitude() + 2*n.right.computeMagnitude()
}

// Part 2

func GetMaximumMagnitude(input []string) int {
	maxMagnitude := 0

	for i, number1 := range input {
		for j, number2 := range input {
			if j != i {
				snailfishAddition, _ := snailfishAddition([]string{number1, number2})
				magnitude := snailfishAddition.computeMagnitude()
				if magnitude > maxMagnitude {
					maxMagnitude = magnitude
				}
			}
		}
	}

	return maxMagnitude
}

// Data structure to be used

type snailfishNumber struct {
	value               int
	left, right, parent *snailfishNumber
}

// Parsing string to data structure

func parseSnailfishNumber(input string) *snailfishNumber {
	if strings.Count(input, ",") == 1 {
		// input is such as "[a,b]"
		split := strings.Split(input[1:len(input)-1], ",")
		leftInt, _ := strconv.ParseInt(split[0], 10, 64)
		rightInt, _ := strconv.ParseInt(split[1], 10, 64)
		number := newSnailfishNumber(int(leftInt), int(rightInt))
		return &number
	}
	if strings.Count(input, ",") == 0 {
		// input is such as "a"
		numberValue, _ := strconv.ParseInt(input, 10, 64)
		number := snailfishNumber{value: int(numberValue)}
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

	number1 := parseSnailfishNumber(firstPart)
	number2 := parseSnailfishNumber(secondPart)
	result := snailfishNumber{-1, number1, number2, nil}

	// Updating parent pointers
	// I don't know (yet) why this is needed but the debugger brought me here, and I want to get out
	if !number1.isRegular() {
		number1.left.parent = number1
		number1.right.parent = number1
	}
	if !number2.isRegular() {
		number2.left.parent = number2
		number2.right.parent = number2
	}
	number1.parent = &result
	number2.parent = &result
	return &result
}

func newSnailfishNumber(n1, n2 int) snailfishNumber {
	regularNumber1 := snailfishNumber{value: n1}
	regularNumber2 := snailfishNumber{value: n2}
	snailfishNumber := snailfishNumber{value: -1, left: &regularNumber1, right: &regularNumber2}
	regularNumber1.parent = &snailfishNumber
	regularNumber2.parent = &snailfishNumber
	return snailfishNumber
}

func (n *snailfishNumber) add(secondNumber *snailfishNumber) (*snailfishNumber, error) {
	result := addition(n, secondNumber)
	err := result.reduce()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func addition(n1, n2 *snailfishNumber) *snailfishNumber {
	additionNumber := snailfishNumber{value: -1, left: n1, right: n2}
	// Updating pointers to parents
	n1.left.parent = n1
	n2.left.parent = n2
	n1.right.parent = n1
	n2.right.parent = n2
	n1.parent = &additionNumber
	n2.parent = &additionNumber
	return &additionNumber
}

// reduce

func (n *snailfishNumber) reduce() error {
	shouldCheckExplode := true
	shouldCheckSplit := true

	for shouldCheckExplode || shouldCheckSplit {
		explodeErr := n.explode()
		if explodeErr != nil {
			shouldCheckExplode = false
			splitErr := n.split()
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

func (n *snailfishNumber) explode() error {
	pairThatShouldExplode, err := n.pairThatShouldExplode(0)
	if err != nil {
		return err
	}
	//fmt.Printf("Exploding pair %v with left %v and right %v!\n", pairThatShouldExplode, pairThatShouldExplode.left.value, pairThatShouldExplode.right.value)
	explodedNumber := snailfishNumber{0, nil, nil, pairThatShouldExplode.parent}

	rightNeighbour, rightErr := pairThatShouldExplode.findRightRegularNeighbour()
	if rightErr == nil {
		rightNeighbour.value += pairThatShouldExplode.right.value
	}
	leftNeighbour, leftErr := pairThatShouldExplode.findLeftRegularNeighbour()
	if leftErr == nil {
		leftNeighbour.value += pairThatShouldExplode.left.value
	}

	if pairThatShouldExplode.isRight() {
		pairThatShouldExplode.parent.right = &explodedNumber
	} else if pairThatShouldExplode.isLeft() {
		pairThatShouldExplode.parent.left = &explodedNumber
	}

	return nil
}

func (n *snailfishNumber) pairThatShouldExplode(currentDepth int) (*snailfishNumber, error) {
	if n.left.isRegular() && n.right.isRegular() {
		if currentDepth >= 4 {
			return n, nil
		} else {
			return nil, errors.New("no pair to explode")
		}
	}
	currentDepth++
	if !n.left.isRegular() {
		potentialLeftPairToExplode, errLeft := n.left.pairThatShouldExplode(currentDepth)
		if errLeft == nil {
			return potentialLeftPairToExplode, nil
		}
	}
	if !n.right.isRegular() {
		potentialRightPairToExplode, errRight := n.right.pairThatShouldExplode(currentDepth)
		if errRight == nil {
			return potentialRightPairToExplode, nil
		}
	}
	return nil, errors.New("no pair to explode")
}

// split part

func (n *snailfishNumber) split() error {
	numberToSplit, err := n.numberThatShouldSplit()

	if err != nil {
		return err
	}

	valueToSplit := numberToSplit.value

	leftNumber := snailfishNumber{valueToSplit / 2, nil, nil, numberToSplit}
	rightNumber := snailfishNumber{(valueToSplit + 1) / 2, nil, nil, numberToSplit}

	numberToSplit.left = &leftNumber
	numberToSplit.right = &rightNumber
	numberToSplit.value = -1

	return nil
}

func (n *snailfishNumber) numberThatShouldSplit() (*snailfishNumber, error) {
	if n.isRegular() {
		if n.value >= 10 {
			return n, nil
		} else {
			return nil, errors.New("no number to split")
		}
	}
	potentialLeftSplit, leftError := n.left.numberThatShouldSplit()
	if leftError == nil {
		return potentialLeftSplit, nil
	}
	potentialRightSplit, rightError := n.right.numberThatShouldSplit()
	if rightError == nil {
		return potentialRightSplit, nil
	}
	return nil, errors.New("no number to split")
}

// Navigate through "tree" number

func (n *snailfishNumber) findLeftmostRegularNumber() (*snailfishNumber, error) {
	if n.isRegular() {
		return nil, errors.New("already a regular number")
	}
	if n.left.isRegular() {
		return n.left, nil
	}
	return n.left.findLeftmostRegularNumber()
}

func (n *snailfishNumber) findRightmostRegularNumber() (*snailfishNumber, error) {
	if n.isRegular() {
		return nil, errors.New("already a regular number")
	}
	if n.right.isRegular() {
		return n.right, nil
	}
	return n.right.findRightmostRegularNumber()
}

func (n *snailfishNumber) findLeftRegularNeighbour() (*snailfishNumber, error) {
	if n.isRight() {
		if n.parent.left.isRegular() {
			return n.parent.left, nil
		}
		return n.parent.left.findRightmostRegularNumber()
	}
	if n.isRoot() {
		return nil, errors.New("already leftmost regular number")
	}
	return n.parent.findLeftRegularNeighbour()
}

func (n *snailfishNumber) findRightRegularNeighbour() (*snailfishNumber, error) {
	if n.isLeft() {
		if n.parent.right.isRegular() {
			return n.parent.right, nil
		}
		return n.parent.right.findLeftmostRegularNumber()
	}
	if n.isRoot() {
		return nil, errors.New("already rightmost regular number")
	}
	return n.parent.findRightRegularNeighbour()
}

// Basic utils

func (n *snailfishNumber) isRegular() bool {
	return n.value != -1
}

func (n *snailfishNumber) isRoot() bool {
	return n.parent == nil
}

func (n *snailfishNumber) isLeft() bool {
	if n.isRoot() {
		return false
	}
	return *n.parent.left == *n
}

func (n *snailfishNumber) isRight() bool {
	if n.isRoot() {
		return false
	}
	return *n.parent.right == *n
}

// Printing

func (n *snailfishNumber) Sprint() string {
	if n.isRegular() {
		return fmt.Sprintf("%v", n.value)
	} else {
		return fmt.Sprintf("[%v,%v]", n.left.Sprint(), n.right.Sprint())
	}
}
