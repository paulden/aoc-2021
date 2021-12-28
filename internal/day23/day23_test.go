package day23

import (
	"aoc-2021/internal/pkg/utils"
	"fmt"
	"math"
	"testing"
)

func TestDay23ParseAmphipodsBurrow(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example_ordered.txt")
	burrow := parseAmphipodsBurrow(input)

	// When
	isOrdered := burrow.IsOrdered()

	// Then
	if !isOrdered {
		t.Errorf("Day 23 - Expected burrow to be ordered but was not")
	}
}


func TestDay23ParseAmphipodsBurrow2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input_unfolded.txt")
	burrow := parseAmphipodsBurrow2(input)

	// When
	burrow.PrettyPrint2()
}

func TestDay23CanVisitDestFromSource(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example_ongoing.txt")
	burrow := parseAmphipodsBurrow(input)

	// When
	bToB := burrow.CanVisitDestFromSource(5, 1)
	bToA := burrow.CanVisitDestFromSource(2, 6)
	aToA := burrow.CanVisitDestFromSource(2, 10)
	dToD := burrow.CanVisitDestFromSource(8, 4)

	if !bToB {
		t.Errorf("Day 23 - Expected to be possible to go from B to side room B")
	}
	if !bToA {
		t.Errorf("Day 23 - Expected to be possible to go from B to side room A")
	}
	if aToA {
		t.Errorf("Day 23 - Expected not to be possible to go from A to side room A")
	}
	if dToD {
		t.Errorf("Day 23 - Expected not to be possible to go from D to side room D")
	}
}

func TestDay23Part1Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/example.txt")
	expectedCost := 12521

	// When
	result := GetMinimumEnergyCost(input)

	if result != expectedCost {
		t.Errorf("Day 23 - Expected %v, got %v", expectedCost, result)
	}
}

func TestDay23Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expectedCost := 14460

	// When
	result := GetMinimumEnergyCost(input)

	if result != expectedCost {
		t.Errorf("Day 23 - Expected %v, got %v", expectedCost, result)
	}
}

func TestDay23Part2Example(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input_unfolded.txt")
	burrow := parseAmphipodsBurrow2(input)

	// When
	allCosts := []int{math.MaxInt}

	ints := OrderAmphipodsBurrow2(burrow, 0, allCosts)
	fmt.Println(ints)
}

func TestIsComplete(t *testing.T) {
	type args struct {
		column      []string
		columnIndex int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"False", args{[]string{"", "A", "B"}, 2}, false},
		{"True A", args{[]string{"", "A", "A"}, 2}, true},
		{"True B", args{[]string{"", "B", "B"}, 4}, true},
		{"True C", args{[]string{"", "C", "C"}, 6}, true},
		{"True D", args{[]string{"", "D", "D"}, 8}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isComplete(tt.args.column, tt.args.columnIndex); got != tt.want {
				t.Errorf("isComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}