package main

import (
	"fmt"
	"testing"
)

func TestDay18SimpleAddition(t *testing.T) {
	// Given
	firstNumber := NewSnailfishNumber(1, 2)
	secondNumber := NewSnailfishNumber(3, 4)

	// When
	result := Add(&firstNumber, &secondNumber)

	// Then
	if result != firstNumber.parent || *result.left != firstNumber {
		t.Errorf("Day 18 : invalid parent / left association, expected %v, got %v", *firstNumber.parent, result)
	}
	if result != secondNumber.parent || *result.right != secondNumber {
		t.Errorf("Day 18 : invalid parent / left association, expected %v, got %v", *secondNumber.parent, result)
	}
}

func TestDay18IsLeftOrRight(t *testing.T) {
	// Given
	firstNumber := NewSnailfishNumber(1, 2)
	secondNumber := NewSnailfishNumber(3, 4)

	// When
	Add(&firstNumber, &secondNumber)

	// Then
	if !firstNumber.IsLeft() || firstNumber.IsRight() {
		t.Errorf("Day 18 : firstNumber should be left!")
	}
	if !secondNumber.IsRight() || secondNumber.IsLeft() {
		t.Errorf("Day 18 : secondNumber should be right!")
	}
}

func TestDay18ShouldNotBeExploded(t *testing.T) {
	// Given
	firstNumber := NewSnailfishNumber(1, 2)
	secondNumber := NewSnailfishNumber(3, 4)

	// When
	result := Add(&firstNumber, &secondNumber)
	_, err := result.PairThatShouldExplode(0)

	// Then
	if err == nil {
		t.Errorf("Day 18 : result should not be exploded but got no error to indicate so!")
	}
}

func TestDay18PairThatShouldExplode(t *testing.T) {
	// Given
	input := "[[[[[1,2],[3,4]],[5,6]],[7,8]],[9,0]]"
	result := ParseSnailfishNumber(input)

	// When
	pairThatShouldExplode, err := result.PairThatShouldExplode(0)

	// Then
	if err != nil && pairThatShouldExplode.left.value != 1 && pairThatShouldExplode.right.value != 2 {
		t.Errorf("Day 18 : result should be exploded but got an error or did not point the right pair to explode!")
	}
}

func TestDay18FindLeftmostAndRightmostRegularNumber(t *testing.T) {
	// Given
	input := "[[[[[1,2],[3,4]],[5,6]],[7,8]],[9,0]]"
	result := ParseSnailfishNumber(input)

	// When
	leftmostRegularNumber, _ := result.FindLeftmostRegularNumber()
	rightmostRegularNumber, _ := result.FindRightmostRegularNumber()

	// Then
	if leftmostRegularNumber.value != 1 {
		t.Errorf("Day 18 : wrong leftmost regular number, expected %v, got %v!", 1, leftmostRegularNumber.value)
	}
	if rightmostRegularNumber.value != 0 {
		t.Errorf("Day 18 : wrong rightmost regular number, expected %v, got %v!", 0, rightmostRegularNumber.value)
	}
}

func Test18FindRegularNeighbours(t *testing.T) {
	// Given
	input := "[[[[[1,2],[3,4]],[5,6]],[7,8]],[9,0]]"
	result := ParseSnailfishNumber(input)
	thirdNumber := result.left.left.right.right // is [5,6]

	// When
	leftNeighbour, leftErr := thirdNumber.FindLeftRegularNeighbour()
	rightNeighbour, rightErr := thirdNumber.FindRightRegularNeighbour()

	// Then
	if leftNeighbour.value != 5 || leftErr != nil {
		t.Errorf("Day 18 : wrong left regular neighbour, expected %v, got %v!", 5, leftNeighbour.value)
	}
	if rightNeighbour.value != 7 || rightErr != nil {
		t.Errorf("Day 18 : wrong right regular neighbour, expected %v, got %v!", 7, rightNeighbour.value)
	}
}

func Test18FindLeftRegularNeighbour(t *testing.T) {
	// Given
	input := "[[[[[1,2],[3,4]],[5,6]],[7,8]],[9,0]]"
	result := ParseSnailfishNumber(input)
	firstNumber := result.left.left.left // left is [1,2]
	fifthNumber := result // right is [9,0]

	// When
	fmt.Println(result.Sprint())

	leftInvalidNeighbour, leftErr := firstNumber.left.FindLeftRegularNeighbour()
	rightInvalidNeighbour, rightErr := fifthNumber.right.FindRightRegularNeighbour()

	// Then
	if leftInvalidNeighbour != nil || leftErr == nil {
		t.Errorf("Day 18 : should not find left neighbour and raise an error since it is already the leftmost element!")
	}
	if rightInvalidNeighbour != nil || rightErr == nil {
		t.Errorf("Day 18 : should not find right neighbour and raise an error since it is already the rightmost element!")
	}
}

func Test18ParseNumber(t *testing.T) {
	// Given
	input1 := "[1,2]"
	input2 := "[[1,2],3]"
	input3 := "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"

	// When
	number1 := ParseSnailfishNumber(input1)
	number2 := ParseSnailfishNumber(input2)
	number3 := ParseSnailfishNumber(input3)

	// Then
	if number1.Sprint() != input1 {
		t.Errorf("Day 18 : number parsing, expected %v, got %v", input1, number1.Sprint())
	}
	if number2.Sprint() != input2 {
		t.Errorf("Day 18 : number parsing, expected %v, got %v", input2, number2.Sprint())
	}
	if number3.Sprint() != input3 {
		t.Errorf("Day 18 : number parsing, expected %v, got %v", input3, number3)
	}
}

func Test18ExplodeNumber1(t *testing.T) {
	// Given
	input := "[[[[[9,8],1],2],3],4]"
	expected := "[[[[0,9],2],3],4]"

	// When
	snailfishNumber := ParseSnailfishNumber(input)
	err := snailfishNumber.Explode()

	// Then
	if err != nil {
		t.Errorf("Day 18 : expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18ExplodeNumber2(t *testing.T) {
	// Given
	input := "[7,[6,[5,[4,[3,2]]]]]"
	expected := "[7,[6,[5,[7,0]]]]"

	// When
	snailfishNumber := ParseSnailfishNumber(input)
	err := snailfishNumber.Explode()

	// Then
	if err != nil {
		t.Errorf("Day 18 : expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18ExplodeNumber3(t *testing.T) {
	// Given
	input := "[[6,[5,[4,[3,2]]]],1]"
	expected := "[[6,[5,[7,0]]],3]"

	// When
	snailfishNumber := ParseSnailfishNumber(input)
	err := snailfishNumber.Explode()

	// Then
	if err != nil {
		t.Errorf("Day 18 : expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18ExplodeNumber4(t *testing.T) {
	// Given
	input := "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"
	expected := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"

	// When
	snailfishNumber := ParseSnailfishNumber(input)
	err := snailfishNumber.Explode()

	// Then
	if err != nil {
		t.Errorf("Day 18 : expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18ExplodeNumber5(t *testing.T) {
	// Given
	input := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"
	expected := "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"

	// When
	snailfishNumber := ParseSnailfishNumber(input)
	err := snailfishNumber.Explode()

	// Then
	if err != nil {
		t.Errorf("Day 18 : expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18SplitAndExplodeNumber(t *testing.T) {
	// Given
	input := "[[[[4,0],[5,4]],[[7,0],[15,5]]],[10,[[11,9],[11,0]]]]"
	expected := "[[[[4,0],[5,4]],[[7,7],[0,13]]],[10,[[11,9],[11,0]]]]"

	// When
	snailfishNumber := ParseSnailfishNumber(input)
	err2 := snailfishNumber.Split()
	err1 := snailfishNumber.Explode()

	// Then
	if err1 != nil && err2 != nil {
		t.Errorf("Day 18 : expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18SplitNumber1(t *testing.T) {
	// Given
	input := "[10,1]"
	expected := "[[5,5],1]"

	// When
	snailfishNumber := ParseSnailfishNumber(input)
	err := snailfishNumber.Split()

	// Then
	if err != nil {
		t.Errorf("Day 18 : expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18SplitNumber2(t *testing.T) {
	// Given
	input := "[11,[1,3]]"
	expected := "[[5,6],[1,3]]"

	// When
	snailfishNumber := ParseSnailfishNumber(input)
	err := snailfishNumber.Split()

	// Then
	if err != nil {
		t.Errorf("Day 18 : expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18ReduceNumber1(t *testing.T) {
	// Given
	input := "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"
	expected := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"

	// When
	snailfishNumber := ParseSnailfishNumber(input)
	err := snailfishNumber.Reduce()

	// Then
	if err != nil {
		t.Errorf("Day 18 : expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18ReduceNumber2(t *testing.T) {
	// Given
	input := "[[[[4,0],[5,4]],[[7,0],[15,5]]],[10,[[11,9],[11,0]]]]"
	expected := "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"

	// When
	snailfishNumber := ParseSnailfishNumber(input)
	err := snailfishNumber.Reduce()

	// Then
	if err != nil {
		t.Errorf("Day 18 : expected number to be reducable but got error!")
	}
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected reduced number to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18SnailfishAddition1(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day18_example1.txt")
	expected := "[[[[1,1],[2,2]],[3,3]],[4,4]]"

	// When
	snailfishNumber, _ := SnailfishAddition(input)

	// Then
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected addition to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18SnailfishAddition2(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day18_example2.txt")
	expected := "[[[[5,0],[7,4]],[5,5]],[6,6]]"

	// When
	snailfishNumber, _ := SnailfishAddition(input)

	// Then
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected addition to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18SnailfishAddition3(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day18_example3.txt")
	expected := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"

	// When
	snailfishNumber, _ := SnailfishAddition(input)

	// Then
	if snailfishNumber.Sprint() != expected {
		t.Errorf("Day 18 : expected addition to be %v but got %v", expected, snailfishNumber.Sprint())
	}
}

func Test18ComputeMagnitude1(t *testing.T) {
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
	number1 := ParseSnailfishNumber(input1)
	number2 := ParseSnailfishNumber(input2)
	number3 := ParseSnailfishNumber(input3)
	number4 := ParseSnailfishNumber(input4)
	number5 := ParseSnailfishNumber(input5)
	number6 := ParseSnailfishNumber(input6)
	number7 := ParseSnailfishNumber(input7)

	// When
	magnitude1 := number1.ComputeMagnitude()
	magnitude2 := number2.ComputeMagnitude()
	magnitude3 := number3.ComputeMagnitude()
	magnitude4 := number4.ComputeMagnitude()
	magnitude5 := number5.ComputeMagnitude()
	magnitude6 := number6.ComputeMagnitude()
	magnitude7 := number7.ComputeMagnitude()

	// Then
	if magnitude1 != expected1 {
		t.Errorf("Day 18 : expected magnitude to be %v but got %v", expected1, magnitude1)
	}
	if magnitude2 != expected2 {
		t.Errorf("Day 18 : expected magnitude to be %v but got %v", expected2, magnitude2)
	}
	if magnitude3 != expected3 {
		t.Errorf("Day 18 : expected magnitude to be %v but got %v", expected3, magnitude3)
	}
	if magnitude4 != expected4 {
		t.Errorf("Day 18 : expected magnitude to be %v but got %v", expected4, magnitude4)
	}
	if magnitude5 != expected5 {
		t.Errorf("Day 18 : expected magnitude to be %v but got %v", expected5, magnitude5)
	}
	if magnitude6 != expected6 {
		t.Errorf("Day 18 : expected magnitude to be %v but got %v", expected6, magnitude6)
	}
	if magnitude7 != expected7 {
		t.Errorf("Day 18 : expected magnitude to be %v but got %v", expected7, magnitude7)
	}
}

func Test18SnailfishAdditionAndMagnitude(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day18_example4.txt")
	expected := 4140

	// When
	snailfishNumber, _ := SnailfishAddition(input)
	magnitude := snailfishNumber.ComputeMagnitude()

	// Then
	if magnitude != expected {
		t.Errorf("Day 18 : expected magnitude to be %v but got %v", expected, magnitude)
	}
}

func Test18GetMaximumMagnitude(t *testing.T) {
	// Given
	input := readStringsInFile("../data/day18_example4.txt")
	expected := 3993

	// When
	maxMagnitude := GetMaximumMagnitude(input)

	// Then
	if maxMagnitude != expected {
		t.Errorf("Day 18 : expected max magnitude to be %v but got %v", expected, maxMagnitude)
	}
}
