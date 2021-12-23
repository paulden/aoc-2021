package main

import (
	"fmt"
)

// Part 1

func OrderAmphipodsBurrow(
	burrow Burrow,
	currentCost int,
	allCosts []int,
) ([]int) {
	if burrow.IsOrdered() {
		fmt.Printf("Found a solution with cost %v!\n", currentCost)
		allCosts = append(allCosts, currentCost)
		return allCosts
	} else {
		minimalCost := Minimum(allCosts)
		burrowsWithCosts := burrow.GetNextPossibleBurrows()
		if len(burrowsWithCosts) == 0 {
			return allCosts
		}
		for nextBurrow, cost := range burrowsWithCosts {
			//time.Sleep(1000*time.Millisecond)
			//fmt.Printf("Exploring the following burrow, which would cost %v (current minimal is %v)\n", currentCost + cost, minimalCost)
			//nextBurrow.PrettyPrint()
			if currentCost + cost >= minimalCost {
				// Do not pursue any further
				return allCosts
			}
			if currentCost + cost < minimalCost {
				allCosts = OrderAmphipodsBurrow(*nextBurrow, currentCost + cost, allCosts)
			}
		}
	}
	return allCosts
}

func OrderAmphipodsBurrow2(
	burrow Burrow,
	currentCost int,
	allCosts []int,
) ([]int) {
	if burrow.IsOrdered() {
		fmt.Printf("Found a solution with cost %v!\n", currentCost)
		allCosts = append(allCosts, currentCost)
		return allCosts
	} else {
		minimalCost := Minimum(allCosts)
		burrowsWithCosts := burrow.GetNextPossibleBurrows2()
		if len(burrowsWithCosts) == 0 {
			return allCosts
		}
		for nextBurrow, cost := range burrowsWithCosts {
			//time.Sleep(1000*time.Millisecond)
			//fmt.Printf("Exploring the following burrow, which would cost %v (current minimal is %v)\n", currentCost + cost, minimalCost)
			//nextBurrow.PrettyPrint()
			if currentCost + cost >= minimalCost {
				// Do not pursue any further
				return allCosts
			}
			if currentCost + cost < minimalCost {
				allCosts = OrderAmphipodsBurrow2(*nextBurrow, currentCost + cost, allCosts)
			}
		}
	}
	return allCosts
}

type Burrow struct {
	state [][]string
}

func (b *Burrow) GetNextPossibleBurrows() map[*Burrow]int {
	possibleBurrowsWithCost := make(map[*Burrow]int)
	roomIndices := map[string]int{
		"A": 2,
		"B": 4,
		"C": 6,
		"D": 8,
	}
	costs := map[string]int{
		"A": 1,
		"B": 10,
		"C": 100,
		"D": 1000,
	}
	hallwayPossibleLocations := []int{0, 1, 3, 5, 7, 9, 10}

	for i, column := range b.state {
		if IsEmpty(column) {
			// Nothing to do
			continue
		}
		if IsComplete(column, i) {
			// This side room is complete, do not touch it!
			continue
		}
		if IsHalfComplete(column, i) {
			// This side room has the right amphipod at the bottom, do not move it!
			continue
		}
		if !IsSideRoom(column) && column[0] != "" {
			// An amphipod is in the hallway, it can only go to its specific room (if possible!)
			amphipod := column[0]
			newState := b.DuplicateState()
			amphipodRoom := roomIndices[amphipod]

			if b.state[amphipodRoom][2] == "" && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is there room at the bottom of the side room?
				newState[amphipodRoom][2] = amphipod
				newState[i][0] = ""
				cost := (Abs(amphipodRoom-i) + 2) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
			if b.state[amphipodRoom][1] == "" && b.state[amphipodRoom][2] == amphipod && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is the room already occupied by the same type and there is room for one more?
				newState[amphipodRoom][1] = amphipod
				newState[i][0] = ""
				cost := (Abs(amphipodRoom-i) + 1) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
		}
		if IsSideRoom(column) {
			// An amphipod is in a side room, list possible places where it can go!
			amphipod, depth := GetFirstAmphipod(column)

			// Listing possible locations in hallway
			for _, location := range hallwayPossibleLocations {
				if b.state[location][0] == ""  && b.CanVisitDestFromSource(location, i) {
					newState := b.DuplicateState()
					newState[location][0] = amphipod
					newState[i][depth] = ""
					cost := (depth + Abs(i-location)) * costs[amphipod]
					newBurrow := Burrow{newState}
					possibleBurrowsWithCost[&newBurrow] = cost
				}
			}

			// Listing possible locations in side rooms
			amphipodRoom := roomIndices[amphipod]
			if b.state[amphipodRoom][2] == ""  && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is there room at the bottom of the specific side room?
				newState := b.DuplicateState()
				newState[amphipodRoom][2] = amphipod
				newState[i][depth] = ""
				cost := (Abs(amphipodRoom-i) + 2 + depth) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
			if b.state[amphipodRoom][1] == "" && b.state[amphipodRoom][2] == amphipod && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is the room already occupied by the same type and there is room for one more?
				newState := b.DuplicateState()
				newState[amphipodRoom][1] = amphipod
				newState[i][depth] = ""
				cost := (Abs(amphipodRoom-i) + 1 + depth) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
		}
	}

	return possibleBurrowsWithCost
}


func (b *Burrow) GetNextPossibleBurrows2() map[*Burrow]int {
	possibleBurrowsWithCost := make(map[*Burrow]int)
	roomIndices := map[string]int{
		"A": 2,
		"B": 4,
		"C": 6,
		"D": 8,
	}
	costs := map[string]int{
		"A": 1,
		"B": 10,
		"C": 100,
		"D": 1000,
	}
	hallwayPossibleLocations := []int{0, 1, 3, 5, 7, 9, 10}

	for i, column := range b.state {
		if IsEmpty(column) {
			// Nothing to do
			continue
		}
		if IsComplete2(column, i) {
			// This side room is complete, do not touch it!
			continue
		}
		if (i == 2 || i == 4 || i == 6 || i == 8) && IsHalfComplete2(column, i) {
			// This side room has the right amphipod at the bottom, do not move it!
			continue
		}
		if !IsSideRoom(column) && column[0] != "" {
			// An amphipod is in the hallway, it can only go to its specific room (if possible!)
			amphipod := column[0]
			newState := b.DuplicateState()
			amphipodRoom := roomIndices[amphipod]

			if b.state[amphipodRoom][4] == "" && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is there room at the bottom of the side room?
				newState[amphipodRoom][4] = amphipod
				newState[i][0] = ""
				cost := (Abs(amphipodRoom-i) + 4) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
			if b.state[amphipodRoom][3] == "" && b.state[amphipodRoom][4] == amphipod && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is the room already occupied by the same type and there is room for one more?
				newState[amphipodRoom][3] = amphipod
				newState[i][0] = ""
				cost := (Abs(amphipodRoom-i) + 3) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
			if b.state[amphipodRoom][2] == "" && b.state[amphipodRoom][3] == amphipod && b.state[amphipodRoom][4] == amphipod && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is the room already occupied by the same type and there is room for one more?
				newState[amphipodRoom][2] = amphipod
				newState[i][0] = ""
				cost := (Abs(amphipodRoom-i) + 2) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
			if b.state[amphipodRoom][1] == "" && b.state[amphipodRoom][2] == amphipod && b.state[amphipodRoom][3] == amphipod && b.state[amphipodRoom][4] == amphipod && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is the room already occupied by the same type and there is room for one more?
				newState[amphipodRoom][1] = amphipod
				newState[i][0] = ""
				cost := (Abs(amphipodRoom-i) + 1) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
		}
		if IsSideRoom2(column) {
			// An amphipod is in a side room, list possible places where it can go!
			amphipod, depth := GetFirstAmphipod(column)

			// Listing possible locations in hallway
			for _, location := range hallwayPossibleLocations {
				if b.state[location][0] == ""  && b.CanVisitDestFromSource(location, i) {
					newState := b.DuplicateState()
					newState[location][0] = amphipod
					newState[i][depth] = ""
					cost := (depth + Abs(i-location)) * costs[amphipod]
					newBurrow := Burrow{newState}
					possibleBurrowsWithCost[&newBurrow] = cost
				}
			}

			// Listing possible locations in side rooms
			amphipodRoom := roomIndices[amphipod]
			if b.state[amphipodRoom][4] == "" && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is there room at the bottom of the specific side room?
				newState := b.DuplicateState()
				newState[amphipodRoom][4] = amphipod
				newState[i][depth] = ""
				cost := (Abs(amphipodRoom-i) + 4 + depth) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
			if b.state[amphipodRoom][3] == "" && b.state[amphipodRoom][4] == amphipod && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is the room already occupied by the same type and there is room for one more?
				newState := b.DuplicateState()
				newState[amphipodRoom][3] = amphipod
				newState[i][depth] = ""
				cost := (Abs(amphipodRoom-i) + 3 + depth) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
			if b.state[amphipodRoom][2] == "" && b.state[amphipodRoom][3] == amphipod && b.state[amphipodRoom][4] == amphipod && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is the room already occupied by the same type and there is room for one more?
				newState := b.DuplicateState()
				newState[amphipodRoom][2] = amphipod
				newState[i][depth] = ""
				cost := (Abs(amphipodRoom-i) + 2 + depth) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
			if b.state[amphipodRoom][1] == "" && b.state[amphipodRoom][2] == amphipod && b.state[amphipodRoom][3] == amphipod && b.state[amphipodRoom][4] == amphipod && b.CanVisitDestFromSource(amphipodRoom, i) {
				// Is the room already occupied by the same type and there is room for one more?
				newState := b.DuplicateState()
				newState[amphipodRoom][1] = amphipod
				newState[i][depth] = ""
				cost := (Abs(amphipodRoom-i) + 1 + depth) * costs[amphipod]
				newBurrow := Burrow{newState}
				possibleBurrowsWithCost[&newBurrow] = cost
			}
		}
	}

	return possibleBurrowsWithCost
}

func (b *Burrow) IsOrdered() bool {
	orderedA := b.state[2][1] == "A" && b.state[2][2] == "A"
	orderedB := b.state[4][1] == "B" && b.state[4][2] == "B"
	orderedC := b.state[6][1] == "C" && b.state[6][2] == "C"
	orderedD := b.state[8][1] == "D" && b.state[8][2] == "D"
	return orderedA && orderedB && orderedC && orderedD
}

func (b *Burrow) IsOrdered2() bool {
	orderedA := b.state[2][1] == "A" && b.state[2][2] == "A" && b.state[2][6] == "A" && b.state[2][4] == "A"
	orderedB := b.state[4][1] == "B" && b.state[4][2] == "B" && b.state[4][6] == "B" && b.state[4][4] == "B"
	orderedC := b.state[6][1] == "C" && b.state[6][2] == "C" && b.state[6][6] == "C" && b.state[6][4] == "C"
	orderedD := b.state[8][1] == "D" && b.state[8][2] == "D" && b.state[8][6] == "D" && b.state[8][4] == "D"
	return orderedA && orderedB && orderedC && orderedD
}

func (b *Burrow) DuplicateState() [][]string {
	newState := make([][]string, len(b.state))
	for i, _ := range b.state {
		newState[i] = make([]string, len(b.state[i]))
		copy(newState[i], b.state[i])
	}
	return newState
}

func (b *Burrow) CanVisitDestFromSource(destination, source int) bool {
	if destination > source {
		for path := source + 1; path <= destination; path++ {
			if b.state[path][0] != "" {
				return false
			}
		}
	} else {
		for path := destination; path < source; path++ {
			if b.state[path][0] != "" {
				return false
			}
		}
	}
	return true
}

func (b *Burrow) PrettyPrint() {
	fmt.Println("#############")

	fmt.Print("#")
	for i := 0; i < 11; i++ {
		if b.state[i][0] == "" {
			fmt.Print(".")
		} else {
			fmt.Print(b.state[i][0])
		}
	}
	fmt.Print("#\n")

	fmt.Print("###")
	for i := 2; i <= 8; i++ {
		if len(b.state[i]) == 1 {
			fmt.Print("#")
		} else if b.state[i][1] == "" {
			fmt.Print(".")
		} else {
			fmt.Print(b.state[i][1])
		}
	}
	fmt.Print("###\n")

	fmt.Print("  #")
	for i := 2; i <= 8; i++ {
		if len(b.state[i]) == 1 {
			fmt.Print("#")
		} else if b.state[i][2] == "" {
			fmt.Print(".")
		} else {
			fmt.Print(b.state[i][2])
		}
	}
	fmt.Print("#  \n")

	fmt.Println("  #########  ")
}


func (b *Burrow) PrettyPrint2() {
	fmt.Println("#############")

	fmt.Print("#")
	for i := 0; i < 11; i++ {
		if b.state[i][0] == "" {
			fmt.Print(".")
		} else {
			fmt.Print(b.state[i][0])
		}
	}
	fmt.Print("#\n")

	fmt.Print("###")
	for i := 2; i <= 8; i++ {
		if len(b.state[i]) == 1 {
			fmt.Print("#")
		} else if b.state[i][1] == "" {
			fmt.Print(".")
		} else {
			fmt.Print(b.state[i][1])
		}
	}
	fmt.Print("###\n")

	fmt.Print("  #")
	for i := 2; i <= 8; i++ {
		if len(b.state[i]) == 1 {
			fmt.Print("#")
		} else if b.state[i][2] == "" {
			fmt.Print(".")
		} else {
			fmt.Print(b.state[i][2])
		}
	}
	fmt.Print("#  \n")

	fmt.Print("  #")
	for i := 2; i <= 8; i++ {
		if len(b.state[i]) == 1 {
			fmt.Print("#")
		} else if b.state[i][3] == "" {
			fmt.Print(".")
		} else {
			fmt.Print(b.state[i][3])
		}
	}
	fmt.Print("#  \n")

	fmt.Print("  #")
	for i := 2; i <= 8; i++ {
		if len(b.state[i]) == 1 {
			fmt.Print("#")
		} else if b.state[i][4] == "" {
			fmt.Print(".")
		} else {
			fmt.Print(b.state[i][4])
		}
	}
	fmt.Print("#  \n")

	fmt.Println("  #########  ")
}

func IsEmpty(column []string) bool {
	for _, spot := range column {
		if spot != "" {
			return false
		}
	}
	return true
}


func IsSideRoom(column []string) bool {
	return len(column) == 3
}

func IsSideRoom2(column []string) bool {
	return len(column) == 5
}

func IsComplete(column []string, columnIndex int) bool {
	completeA := columnIndex == 2 && column[1] == "A" && column[2] == "A"
	completeB := columnIndex == 4 && column[1] == "B" && column[2] == "B"
	completeC := columnIndex == 6 && column[1] == "C" && column[2] == "C"
	completeD := columnIndex == 8 && column[1] == "D" && column[2] == "D"
	return completeA || completeB || completeC || completeD
}

func IsComplete2(column []string, columnIndex int) bool {
	completeA := columnIndex == 2 && column[1] == "A" && column[2] == "A" && column[3] == "A" && column[4] == "A"
	completeB := columnIndex == 4 && column[1] == "B" && column[2] == "B" && column[3] == "B" && column[4] == "B"
	completeC := columnIndex == 6 && column[1] == "C" && column[2] == "C" && column[3] == "C" && column[4] == "C"
	completeD := columnIndex == 8 && column[1] == "D" && column[2] == "D" && column[3] == "D" && column[4] == "D"
	return completeA || completeB || completeC || completeD
}

func IsHalfComplete(column []string, columnIndex int) bool {
	halfCompleteA := columnIndex == 2 && column[1] == "" && column[2] == "A"
	halfCompleteB := columnIndex == 4 && column[1] == "" && column[2] == "B"
	halfCompleteC := columnIndex == 6 && column[1] == "" && column[2] == "C"
	halfCompleteD := columnIndex == 8 && column[1] == "" && column[2] == "D"
	return halfCompleteA || halfCompleteB || halfCompleteC || halfCompleteD
}

func IsHalfComplete2(column []string, columnIndex int) bool {
	var depth int
	if column[3] == "" {
		depth = 4
	} else if column[2] == "" {
		depth = 3
	} else if column[1] == "" {
		depth = 2
	}
	halfCompleteA := columnIndex == 2 && column[1] == "" && column[depth] == "A" && column[Min(depth + 1, 4)] == "A" && column[Min(depth + 2, 4)] == "A"
	halfCompleteB := columnIndex == 4 && column[1] == "" && column[depth] == "B" && column[Min(depth + 1, 4)] == "B" && column[Min(depth + 2, 4)] == "B"
	halfCompleteC := columnIndex == 6 && column[1] == "" && column[depth] == "C" && column[Min(depth + 1, 4)] == "C" && column[Min(depth + 2, 4)] == "C"
	halfCompleteD := columnIndex == 8 && column[1] == "" && column[depth] == "D" && column[Min(depth + 1, 4)] == "D" && column[Min(depth + 2, 4)] == "D"
	return halfCompleteA || halfCompleteB || halfCompleteC || halfCompleteD
}

func GetFirstAmphipod(column []string) (string, int) {
	for depth, amphipod := range column {
		if amphipod != "" {
			return amphipod, depth
		}
	}
	return "", 0
}

func ParseAmphipodsBurrow(input []string) Burrow {
	burrowState := make([][]string, 11)

	for i, _ := range burrowState {
		spot := string(input[1][i+1])
		if spot != "." {
			burrowState[i] = []string{spot}
		} else {
			burrowState[i] = []string{""}
		}
	}

	sideRoomColumns := []int{2, 4, 6, 8}

	for _, sideRoomColumn := range sideRoomColumns {
		first := string(input[2][sideRoomColumn+1])
		second := string(input[3][sideRoomColumn+1])

		if first != "." {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], first)
		} else {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], "")
		}
		if second != "." {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], second)
		} else {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], "")
		}
	}

	return Burrow{burrowState}
}

func ParseAmphipodsBurrow2(input []string) Burrow {
	burrowState := make([][]string, 11)

	for i, _ := range burrowState {
		spot := string(input[1][i+1])
		if spot != "." {
			burrowState[i] = []string{spot}
		} else {
			burrowState[i] = []string{""}
		}
	}

	sideRoomColumns := []int{2, 4, 6, 8}

	for _, sideRoomColumn := range sideRoomColumns {
		first := string(input[2][sideRoomColumn+1])
		second := string(input[3][sideRoomColumn+1])
		third := string(input[4][sideRoomColumn+1])
		fourth := string(input[5][sideRoomColumn+1])

		if first != "." {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], first)
		} else {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], "")
		}

		if second != "." {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], second)
		} else {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], "")
		}

		if third != "." {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], third)
		} else {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], "")
		}

		if fourth != "." {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], fourth)
		} else {
			burrowState[sideRoomColumn] = append(burrowState[sideRoomColumn], "")
		}
	}

	return Burrow{burrowState}
}
