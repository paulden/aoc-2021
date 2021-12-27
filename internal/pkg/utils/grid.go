package utils

func CreateEmptyGrid(height int, width int) [][]int {
	emptyGrid := make([][]int, height)
	for i := 0; i < height; i++ {
		emptyGrid[i] = make([]int, width)
	}
	return emptyGrid
}
