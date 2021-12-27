package day19

import (
	"aoc-2021/internal/pkg/utils"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Part 1

func CountBeacons(input []string) int {
	beaconsPerScanner := parseScannersAndBeacons(input)
	beacons, _ := locateBeaconsAndScanners(beaconsPerScanner)
	return len(beacons)
}

// Part 2

func GetMaximumManhattanDistance(input []string) int {
	maximumManhattanDistance := 0
	beaconsPerScanner := parseScannersAndBeacons(input)
	_, scanners := locateBeaconsAndScanners(beaconsPerScanner)

	for _, scanner1 := range scanners {
		for _, scanner2 := range scanners {
			manhattanDistance := scanner1.getManhattanDistanceTo(scanner2)
			if manhattanDistance > maximumManhattanDistance {
				maximumManhattanDistance = manhattanDistance
			}
		}
	}

	return maximumManhattanDistance
}

func locateBeaconsAndScanners(
	beaconsPerScanner map[int][]coordinates3D,
) (
	beacons []coordinates3D,
	scanners []coordinates3D,
) {
	allBeacons := make(map[coordinates3D]bool)
	scannersToVisit := listScannersToVisit(beaconsPerScanner)

	scanner0 := 0
	scanner0Coordinates := coordinates3D{0, 0, 0}

	// List beacons from scanner 0
	for _, beacon := range beaconsPerScanner[scanner0] {
		allBeacons[beacon] = true
	}
	delete(scannersToVisit, scanner0)

	// While not all scanners are visited, keep testing if there are overlapping regions and add beacons to the list!
	for len(scannersToVisit) > 0 {
		for scanner := range scannersToVisit {
			currentBeacons := beaconsPerScanner[scanner]
			vectorToScanner, rotationToScanner, err := isOverlappingRegion(currentBeacons, allBeacons)
			if err == nil {
				// Found an overlapping region!
				currentScannerCoordinates := scanner0Coordinates.translate(vectorToScanner)
				scanners = append(scanners, currentScannerCoordinates)

				transformedBeacons := rotateCoordinates(currentBeacons, rotationToScanner)
				for _, beacon := range transformedBeacons {
					beaconAfterTranslation := beacon.translate(vectorToScanner)
					allBeacons[beaconAfterTranslation] = true
				}
				delete(scannersToVisit, scanner)
			}
		}
	}

	for beacon := range allBeacons {
		beacons = append(beacons, beacon)
	}
	return
}

func listScannersToVisit(beaconsPerScanner map[int][]coordinates3D) map[int]bool {
	scannersToVisit := make(map[int]bool)
	for scanner := range beaconsPerScanner {
		scannersToVisit[scanner] = true
	}
	return scannersToVisit
}

func isOverlappingRegion(beacons []coordinates3D, allBeacons map[coordinates3D]bool) (vector, int, error) {
	var vectorToScanner vector
	var rotationToScanner int
	var err error

	for rotation := 1; rotation <= 24; rotation++ {
		vectorToScanner, rotationToScanner, err = isOverlappingRegionForRotation(beacons, allBeacons, rotation)
		if err == nil {
			break
		}
	}
	return vectorToScanner, rotationToScanner, err
}

func isOverlappingRegionForRotation(beacons []coordinates3D, allBeacons map[coordinates3D]bool, rotation int) (vector, int, error) {
	beaconsToTest := rotateCoordinates(beacons, rotation)
	distances := make(map[vector]int)

	for beacon0 := range allBeacons {
		for _, beaconScanner := range beaconsToTest {
			distances[beacon0.getVectorTo(beaconScanner)]++
		}
	}

	for distance, occurrences := range distances {
		if occurrences > 11 {
			return distance, rotation, nil
		}
	}

	return vector{0, 0, 0}, 0, errors.New("region is not overlapping from this rotation")
}

// Data structures

type coordinates3D struct {
	x, y, z int
}

type vector struct {
	x, y, z int
}

func (c *coordinates3D) getVectorTo(b coordinates3D) vector {
	return vector{c.x - b.x, c.y - b.y, c.z - b.z}
}

func (c *coordinates3D) translate(v vector) coordinates3D {
	return coordinates3D{c.x + v.x, c.y + v.y, c.z + v.z}
}

func (c *coordinates3D) getManhattanDistanceTo(b coordinates3D) int {
	return utils.Abs(c.x-b.x) + utils.Abs(c.y-b.y) + utils.Abs(c.z-b.z)
}

func (v *vector) exists() bool {
	return v.x != 0 && v.y != 0 && v.z != 0
}

// Rotations

//     .+------+
//   .'   y  .'|
//  +---+--+'  |
//  |      | x |
//  |   z  +   +
//  |      | .'
//  +------+'

func (c *coordinates3D) rotate(rotation int) coordinates3D {
	switch rotation {
	// Positive x
	case 1:
		return coordinates3D{c.x, c.y, c.z}
	case 2:
		return coordinates3D{c.x, c.z, -c.y}
	case 3:
		return coordinates3D{c.x, -c.y, -c.z}
	case 4:
		return coordinates3D{c.x, -c.z, c.y}
	// Positive y
	case 5:
		return coordinates3D{c.y, c.z, c.x}
	case 6:
		return coordinates3D{c.y, c.x, -c.z}
	case 7:
		return coordinates3D{c.y, -c.z, -c.x}
	case 8:
		return coordinates3D{c.y, -c.x, c.z}
	// Positive z
	case 9:
		return coordinates3D{c.z, c.x, c.y}
	case 10:
		return coordinates3D{c.z, c.y, -c.x}
	case 11:
		return coordinates3D{c.z, -c.x, -c.y}
	case 12:
		return coordinates3D{c.z, -c.y, c.x}
	// Negative x
	case 13:
		return coordinates3D{-c.x, -c.y, c.z}
	case 14:
		return coordinates3D{-c.x, c.z, c.y}
	case 15:
		return coordinates3D{-c.x, c.y, -c.z}
	case 16:
		return coordinates3D{-c.x, -c.z, -c.y}
	// Negative y
	case 17:
		return coordinates3D{-c.y, -c.z, c.x}
	case 18:
		return coordinates3D{-c.y, c.x, c.z}
	case 19:
		return coordinates3D{-c.y, c.z, -c.x}
	case 20:
		return coordinates3D{-c.y, -c.x, -c.z}
	// Negative z
	case 21:
		return coordinates3D{-c.z, -c.x, c.y}
	case 22:
		return coordinates3D{-c.z, c.y, c.x}
	case 23:
		return coordinates3D{-c.z, c.x, -c.y}
	case 24:
		return coordinates3D{-c.z, -c.y, -c.x}
	}
	return *c
}

func rotateCoordinates(beacons []coordinates3D, rotation int) []coordinates3D {
	transformedBeacons := make([]coordinates3D, len(beacons))

	for i, beacon := range beacons {
		transformedBeacons[i] = beacon.rotate(rotation)
	}

	return transformedBeacons
}

// Parsing

func parseScannersAndBeacons(input []string) map[int][]coordinates3D {
	var currentScanner int
	beaconsPerScanner := make(map[int][]coordinates3D)

	re := regexp.MustCompile("[0-9]+")
	for _, line := range input {
		if strings.Contains(line, "scanner") {
			scannerNumber := re.FindString(line)
			scannerInt, _ := strconv.ParseInt(scannerNumber, 10, 64)
			currentScanner = int(scannerInt)
		} else if strings.Contains(line, ",") {
			split := strings.Split(line, ",")
			parsedX, _ := strconv.ParseInt(split[0], 10, 64)
			parsedY, _ := strconv.ParseInt(split[1], 10, 64)
			parsedZ, _ := strconv.ParseInt(split[2], 10, 64)
			coordinates3d := coordinates3D{int(parsedX), int(parsedY), int(parsedZ)}
			beaconsPerScanner[currentScanner] = append(beaconsPerScanner[currentScanner], coordinates3d)
		}
	}
	return beaconsPerScanner
}
