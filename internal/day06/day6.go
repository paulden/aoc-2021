package day06

import (
	"strconv"
	"strings"
)

func CountLanternfishesNaive(lanternfishesAgesString string, daysBeforeCheck int) int {
	lanternfishesAges := parseLanternfishesAgesAsSlice(lanternfishesAgesString)

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
	lanternfishesAges := parseLanternfishesAgesAsMap(lanternfishesAgesString)

	for i := 0; i < daysBeforeCheck; i++ {
		newLanternfishesAges := make(map[int]int)

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

func parseLanternfishesAgesAsMap(lanternfishesAgesString string) map[int]int {
	split := strings.Split(lanternfishesAgesString, ",")
	lanternfishesAges := make(map[int]int)

	for _, lanternfishAge := range split {
		parsedInt, _ := strconv.ParseInt(lanternfishAge, 10, 32)
		lanternfishesAges[int(parsedInt)]++
	}
	return lanternfishesAges
}

func parseLanternfishesAgesAsSlice(lanternfishesAgesString string) []int {
	split := strings.Split(lanternfishesAgesString, ",")
	lanternfishesAges := make([]int, len(split))

	for i, lanternfishAge := range split {
		parsedInt, _ := strconv.ParseInt(lanternfishAge, 10, 32)
		lanternfishesAges[i] = int(parsedInt)
	}
	return lanternfishesAges
}
