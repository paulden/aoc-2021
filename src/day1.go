package main

func CountIncreases(data []int) int {
	increases := 0
	previousDepth := data[0]

	for _, depth := range data[1:] {
		if depth > previousDepth {
			increases++
		}
		previousDepth = depth
	}
	return increases
}

func CountThreeMeasurementsIncreases(data []int) int {
	increases := 0
	averagePreviousDepth := data[0] + data[1] + data[2]

	for i, _ := range data[1:len(data) - 2] {
		averageDepth := data[i + 1] + data[i + 2] + data[i + 3]
		if averageDepth > averagePreviousDepth {
			increases++
		}
		averagePreviousDepth = averageDepth
	}
	return increases
}
