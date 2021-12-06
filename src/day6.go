package main

import (
	"strconv"
	"strings"
)

func CountLanternfishesNaive(lanternfishesAgesString string, daysBeforeCheck int) int {
	lanternfishesAges := ParseLanternfishesAgesAsSlice(lanternfishesAgesString)

	for i := 0; i < daysBeforeCheck; i++ {
		newLanternfishesAges := lanternfishesAges

		for i, lanternfishAge := range lanternfishesAges {
			if lanternfishAge == 0 {
				newLanternfishesAges[i] = 6
				newLanternfishesAges = append(newLanternfishesAges, 8)
			} else {
				newLanternfishesAges[i]--
			}
		}

		lanternfishesAges = newLanternfishesAges
	}

	return len(lanternfishesAges)
}

func CountLanternfishesOptimized(lanternfishesAgesString string, daysBeforeCheck int) int {
	lanternfishesAges := ParseLanternfishesAgesAsMap(lanternfishesAgesString)

	for i := 0; i < daysBeforeCheck; i++ {
		newLanternfishesAges := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

		for daysToSpawn, lanternfishesNumber := range lanternfishesAges {

			if daysToSpawn == 0 {
				newLanternfishesAges[6] += lanternfishesNumber
				newLanternfishesAges[8] = lanternfishesNumber
			} else {
				newLanternfishesAges[daysToSpawn-1] += lanternfishesNumber
			}
		}

		lanternfishesAges = newLanternfishesAges
	}

	totalLanternfishes := 0

	for _, number := range lanternfishesAges {
		totalLanternfishes += number
	}

	return totalLanternfishes
}

func ParseLanternfishesAgesAsMap(lanternfishesAgesString string) map[int]int {
	split := strings.Split(lanternfishesAgesString, ",")
	lanternfishesAges := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

	for _, lanternfishAge := range split {
		parsedInt, _ := strconv.ParseInt(lanternfishAge, 10, 32)
		lanternfishesAges[int(parsedInt)]++
	}
	return lanternfishesAges
}

func ParseLanternfishesAgesAsSlice(lanternfishesAgesString string) []int {
	split := strings.Split(lanternfishesAgesString, ",")
	lanternfishesAges := make([]int, len(split))

	for i, lanternfishAge := range split {
		parsedInt, _ := strconv.ParseInt(lanternfishAge, 10, 32)
		lanternfishesAges[i] = int(parsedInt)
	}
	return lanternfishesAges
}
