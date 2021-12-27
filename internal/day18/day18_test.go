package day18

import (
	"aoc-2021/internal/pkg/utils"
	"fmt"
	"testing"
)

func TestDay18SimpleAddition(t *testing.T) {
	// Given
	firstNumber := newSnailfishNumber(1, 2)
	secondNumber := newSnailfishNumber(3, 4)

	// When
	result := addition(&firstNumber, &secondNumber)

	// Then
	if result != firstNumber.parent || *result.left != firstNumber {
		t.Errorf("Day 18 - invalid parent / left association, expected %v, got %v", *firstNumber.parent, result)
	}
	if result != secondNumber.parent || *result.right != secondNumber {
		t.Errorf("Day 18 - invalid parent / left association, expected %v, got %v", *secondNumber.parent, result)
	}
}

func TestDay18IsLeftOrRight(t *testing.T) {
	// Given
	firstNumber := newSnailfishNumber(1, 2)
	secondNumber := newSnailfishNumber(3, 4)

	// When
	addition(&firstNumber, &secondNumber)

	// Then
	if !firstNumber.isLeft() || firstNumber.isRight() {
		t.Errorf("Day 18 - firstNumber should be left!")
	}
	if !secondNumber.isRight() || secondNumber.isLeft() {
		t.Errorf("Day 18 - secondNumber should be right!")
	}
}

func TestDay18ShouldNotBeExploded(t *testing.T) {
	// Given
	firstNumber := newSnailfishNumber(1, 2)
	secondNumber := newSnailfishNumber(3, 4)

	// When
	result := addition(&firstNumber, &secondNumber)
	_, err := result.pairThatShouldExplode(0)

	// Then
	if err == nil {
		t.Errorf("Day 18 - result should not be exploded but got no error to indicate so!")
	}
}

func TestDay18PairThatShouldExplode(t *testing.T) {
	// Given
	input := "[[[[[1,2],[3,4]],[5,6]],[7,8]],[9,0]]"
	result := parseSnailfishNumber(input)

	// When
	pairThatShouldExplode, err := result.pairThatShouldExplode(0)

	// Then
	if err != nil && pairThatShouldExplode.left.value != 1 && pairThatShouldExplode.right.value != 2 {
		t.Errorf("Day 18 - result should be exploded but got an error or did not point the right pair to explode!")
	}
}

func TestDay18FindLeftmostAndRightmostRegularNumber(t *testing.T) {
	// Given
	input := "[[[[[1,2],[3,4]],[5,6]],[7,8]],[9,0]]"
	result := parseSnailfishNumber(input)

	// When
	leftmostRegularNumber, _ := result.findLeftmostRegularNumber()
	rightmostRegularNumber, _ := result.findRightmostRegularNumber()

	// Then
	if leftmostRegularNumber.value != 1 {
		t.Errorf("Day 18 - wrong leftmost regular number, expected %v, got %v!", 1, leftmostRegularNumber.value)
	}
	if rightmostRegularNumber.value != 0 {
		t.Errorf("Day 18 - wrong rightmost regular number, expected %v, got %v!", 0, rightmostRegularNumber.value)
	}
}

func TestDay18FindRegularNeighbours(t *testing.T) {
	// Given
	input := "[[[[[1,2],[3,4]],[5,6]],[7,8]],[9,0]]"
	result := parseSnailfishNumber(input)
	thirdNumber := result.left.left.right.right // is [5,6]

	// When
	leftNeighbour, leftErr := thirdNumber.findLeftRegularNeighbour()
	rightNeighbour, rightErr := thirdNumber.findRightRegularNeighbour()

	// Then
	if leftNeighbour.value != 5 || leftErr != nil {
		t.Errorf("Day 18 - wrong left regular neighbour, expected %v, got %v!", 5, leftNeighbour.value)
	}
	if rightNeighbour.value != 7 || rightErr != nil {
		t.Errorf("Day 18 - wrong right regular neighbour, expected %v, got %v!", 7, rightNeighbour.value)
	}
}

func TestDay18FindLeftRegularNeighbour(t *testing.T) {
	// Given
	input := "[[[[[1,2],[3,4]],[5,6]],[7,8]],[9,0]]"
	result := parseSnailfishNumber(input)
	firstNumber := result.left.left.left // left is [1,2]
	fifthNumber := result                // right is [9,0]

	// When
	fmt.Println(result.Sprint())

	leftInvalidNeighbour, leftErr := firstNumber.left.findLeftRegularNeighbour()
	rightInvalidNeighbour, rightErr := fifthNumber.right.findRightRegularNeighbour()

	// Then
	if leftInvalidNeighbour != nil || leftErr == nil {
		t.Errorf("Day 18 - should not find left neighbour and raise an error since it is already the leftmost element!")
	}
	if rightInvalidNeighbour != nil || rightErr == nil {
		t.Errorf("Day 18 - should not find right neighbour and raise an error since it is already the rightmost element!")
	}
}

func TestDay18ParseNumber(t *testing.T) {
	// Given
	input1 := "[1,2]"
	input2 := "[[1,2],3]"
	input3 := "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"

	// When
	number1 := parseSnailfishNumber(input1)
	number2 := parseSnailfishNumber(input2)
	number3 := parseSnailfishNumber(input3)

	// Then
	if number1.Sprint() != input1 {
		t.Errorf("Day 18 - number parsing, expected %v, got %v", input1, number1.Sprint())
	}
	if number2.Sprint() != input2 {
		t.Errorf("Day 18 - number parsing, expected %v, got %v", input2, number2.Sprint())
	}
	if number3.Sprint() != input3 {
		t.Errorf("Day 18 - number parsing, expected %v, got %v", input3, number3)
	}
}

func TestDay18ExplodeNumber1(t *testing.T) {
	// Given
	input := "[[[[[9,8],1],2],3],4]"
	expected := "[[[[0,9],2],3],4]"

	// When
	snailfishNumber := parseSnailfishNumber(input)
	err := snailfishNumber.explode()

	// Then
	if err != nil {
		t.Errorf("Day 18 - expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18ExplodeNumber2(t *testing.T) {
	// Given
	input := "[7,[6,[5,[4,[3,2]]]]]"
	expected := "[7,[6,[5,[7,0]]]]"

	// When
	snailfishNumber := parseSnailfishNumber(input)
	err := snailfishNumber.explode()

	// Then
	if err != nil {
		t.Errorf("Day 18 - expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18ExplodeNumber3(t *testing.T) {
	// Given
	input := "[[6,[5,[4,[3,2]]]],1]"
	expected := "[[6,[5,[7,0]]],3]"

	// When
	snailfishNumber := parseSnailfishNumber(input)
	err := snailfishNumber.explode()

	// Then
	if err != nil {
		t.Errorf("Day 18 - expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18ExplodeNumber4(t *testing.T) {
	// Given
	input := "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"
	expected := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"

	// When
	snailfishNumber := parseSnailfishNumber(input)
	err := snailfishNumber.explode()

	// Then
	if err != nil {
		t.Errorf("Day 18 - expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18ExplodeNumber5(t *testing.T) {
	// Given
	input := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"
	expected := "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"

	// When
	snailfishNumber := parseSnailfishNumber(input)
	err := snailfishNumber.explode()

	// Then
	if err != nil {
		t.Errorf("Day 18 - expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18SplitAndExplodeNumber(t *testing.T) {
	// Given
	input := "[[[[4,0],[5,4]],[[7,0],[15,5]]],[10,[[11,9],[11,0]]]]"
	expected := "[[[[4,0],[5,4]],[[7,7],[0,13]]],[10,[[11,9],[11,0]]]]"

	// When
	snailfishNumber := parseSnailfishNumber(input)
	err2 := snailfishNumber.split()
	err1 := snailfishNumber.explode()

	// Then
	if err1 != nil && err2 != nil {
		t.Errorf("Day 18 - expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18SplitNumber1(t *testing.T) {
	// Given
	input := "[10,1]"
	expected := "[[5,5],1]"

	// When
	snailfishNumber := parseSnailfishNumber(input)
	err := snailfishNumber.split()

	// Then
	if err != nil {
		t.Errorf("Day 18 - expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18SplitNumber2(t *testing.T) {
	// Given
	input := "[11,[1,3]]"
	expected := "[[5,6],[1,3]]"

	// When
	snailfishNumber := parseSnailfishNumber(input)
	err := snailfishNumber.split()

	// Then
	if err != nil {
		t.Errorf("Day 18 - expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18ReduceNumber1(t *testing.T) {
	// Given
	input := "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"
	expected := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"

	// When
	snailfishNumber := parseSnailfishNumber(input)
	err := snailfishNumber.reduce()

	// Then
	if err != nil {
		t.Errorf("Day 18 - expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18ReduceNumber2(t *testing.T) {
	// Given
	input := "[[[[4,0],[5,4]],[[7,0],[15,5]]],[10,[[11,9],[11,0]]]]"
	expected := "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"

	// When
	snailfishNumber := parseSnailfishNumber(input)
	err := snailfishNumber.reduce()

	// Then
	if err != nil {
		t.Errorf("Day 18 - expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18SnailfishAddition1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example1.txt")
	expected := "[[[[1,1],[2,2]],[3,3]],[4,4]]"

	// When
	snailfishNumber, _ := snailfishAddition(input)

	// Then
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected addition to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18SnailfishAddition2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example2.txt")
	expected := "[[[[5,0],[7,4]],[5,5]],[6,6]]"

	// When
	snailfishNumber, _ := snailfishAddition(input)

	// Then
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected addition to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18SnailfishAddition3(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example3.txt")
	expected := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"

	// When
	snailfishNumber, _ := snailfishAddition(input)

	// Then
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 - expected addition to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func TestDay18ComputeMagnitude1(t *testing.T) {
	// Given
	input1 := "[9,1]"
	input2 := "[[1,2],[[3,4],5]]"
	input3 := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	input4 := "[[[[1,1],[2,2]],[3,3]],[4,4]]"
	input5 := "[[[[3,0],[5,3]],[4,4]],[5,5]]"
	input6 := "[[[[5,0],[7,4]],[5,5]],[6,6]]"
	input7 := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"
	expected1 := 29
	expected2 := 143
	expected3 := 1384
	expected4 := 445
	expected5 := 791
	expected6 := 1137
	expected7 := 3488
	number1 := parseSnailfishNumber(input1)
	number2 := parseSnailfishNumber(input2)
	number3 := parseSnailfishNumber(input3)
	number4 := parseSnailfishNumber(input4)
	number5 := parseSnailfishNumber(input5)
	number6 := parseSnailfishNumber(input6)
	number7 := parseSnailfishNumber(input7)

	// When
	magnitude1 := number1.computeMagnitude()
	magnitude2 := number2.computeMagnitude()
	magnitude3 := number3.computeMagnitude()
	magnitude4 := number4.computeMagnitude()
	magnitude5 := number5.computeMagnitude()
	magnitude6 := number6.computeMagnitude()
	magnitude7 := number7.computeMagnitude()

	// Then
	if magnitude1 != expected1 {
		t.Errorf("Day 18 - expected magnitude to be %v but got %v", expected1, magnitude1)
	}
	if magnitude2 != expected2 {
		t.Errorf("Day 18 - expected magnitude to be %v but got %v", expected2, magnitude2)
	}
	if magnitude3 != expected3 {
		t.Errorf("Day 18 - expected magnitude to be %v but got %v", expected3, magnitude3)
	}
	if magnitude4 != expected4 {
		t.Errorf("Day 18 - expected magnitude to be %v but got %v", expected4, magnitude4)
	}
	if magnitude5 != expected5 {
		t.Errorf("Day 18 - expected magnitude to be %v but got %v", expected5, magnitude5)
	}
	if magnitude6 != expected6 {
		t.Errorf("Day 18 - expected magnitude to be %v but got %v", expected6, magnitude6)
	}
	if magnitude7 != expected7 {
		t.Errorf("Day 18 - expected magnitude to be %v but got %v", expected7, magnitude7)
	}
}

func TestDay18SnailfishAdditionAndMagnitude(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example4.txt")
	expected := 4140

	// When
	snailfishNumber, _ := snailfishAddition(input)
	magnitude := snailfishNumber.computeMagnitude()

	// Then
	if magnitude != expected {
		t.Errorf("Day 18 - expected magnitude to be %v but got %v", expected, magnitude)
	}
}

func TestDay18GetMaximumMagnitude(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example4.txt")
	expected := 3993

	// When
	maxMagnitude := GetMaximumMagnitude(input)

	// Then
	if maxMagnitude != expected {
		t.Errorf("Day 18 - expected max magnitude to be %v but got %v", expected, maxMagnitude)
	}
}

func TestDay18Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 4088

	// When
	maxMagnitude := GetMagnitudeAfterAddition(input)

	// Then
	if maxMagnitude != expected {
		t.Errorf("Day 18 - expected magnitude to be %v but got %v", expected, maxMagnitude)
	}
}

func TestDay18Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 4536

	// When
	maxMagnitude := GetMaximumMagnitude(input)

	// Then
	if maxMagnitude != expected {
		t.Errorf("Day 18 - expected max magnitude to be %v but got %v", expected, maxMagnitude)
	}
}
