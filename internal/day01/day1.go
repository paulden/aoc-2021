package day01

// Part 1

func CountIncreases(input []int) int {
	increases := 0
	previousDepth := input[0]

	for _, depth := range input[1:] {
		if depth > previousDepth {
			increases++
		}
		previousDepth = depth
	}
	return increases
}

// Part 2

func CountThreeMeasurementsIncreases(data []int) int {
	increases := 0
	averagePreviousDepth := data[0] + data[1] + data[2]

	for i, _ := range data[1 : len(data)-2] {
		averageDepth := data[i+1] + data[i+2] + data[i+3]
		if averageDepth > averagePreviousDepth {
			increases++
		}
		averagePreviousDepth = averageDepth
	}
	return increases
}
