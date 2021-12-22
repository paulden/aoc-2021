package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Part 1

func CountBeacons(input []string) int {
	beaconsPerScanner := ParseScannersAndBeacons(input)
	beacons, _ := LocateBeaconsAndScanners(beaconsPerScanner)
	return len(beacons)
}

// Part 2

func GetMaximumManhattanDistance(input []string) int {
	maximumManhattanDistance := 0
	beaconsPerScanner := ParseScannersAndBeacons(input)
	_, scanners := LocateBeaconsAndScanners(beaconsPerScanner)

	for _, scanner1 := range scanners {
		for _, scanner2 := range scanners {
			manhattanDistance := scanner1.GetManhattanDistanceTo(scanner2)
			if manhattanDistance > maximumManhattanDistance {
				maximumManhattanDistance = manhattanDistance
			}
		}
	}

	return maximumManhattanDistance
}

func LocateBeaconsAndScanners(beaconsPerScanner map[int][]Coordinates3D) (beacons []Coordinates3D, scanners []Coordinates3D) {
	allBeacons := make(map[Coordinates3D]bool)
	scannersToVisit := ListScannersToVisit(beaconsPerScanner)

	scanner0 := 0
	scanner0Coordinates := Coordinates3D{0, 0, 0}

	// List beacons from scanner 0
	for _, beacon := range beaconsPerScanner[scanner0] {
		allBeacons[beacon] = true
	}
	delete(scannersToVisit, scanner0)

	// While not all scanners are visited, keep testing if there are overlapping regions and add beacons to the list!
	for len(scannersToVisit) > 0 {
		for scanner, _ := range scannersToVisit {
			currentBeacons := beaconsPerScanner[scanner]
			vectorToScanner, rotationToScanner, err := TestIfOverlappingRegion(currentBeacons, allBeacons)
			if err == nil {
				// Found an overlapping region!
				currentScannerCoordinates := scanner0Coordinates.Translate(vectorToScanner)
				scanners = append(scanners, currentScannerCoordinates)

				transformedBeacons := RotateCoordinates(currentBeacons, rotationToScanner)
				for _, beacon := range transformedBeacons {
					beaconAfterTranslation := beacon.Translate(vectorToScanner)
					allBeacons[beaconAfterTranslation] = true
				}
				delete(scannersToVisit, scanner)
			}
		}
	}

	for beacon, _ := range allBeacons {
		beacons = append(beacons, beacon)
	}
	return
}

func ListScannersToVisit(beaconsPerScanner map[int][]Coordinates3D) map[int]bool {
	scannersToVisit := make(map[int]bool)
	for scanner, _ := range beaconsPerScanner {
		scannersToVisit[scanner] = true
	}
	return scannersToVisit
}

func TestIfOverlappingRegion(beacons []Coordinates3D, allBeacons map[Coordinates3D]bool) (Vector, int, error) {
	var vectorToScanner Vector
	var rotationToScanner int
	var err error

	for rotation := 1; rotation <= 24; rotation++ {
		vectorToScanner, rotationToScanner, err = TestRegionAfterRotation(beacons, allBeacons, rotation)
		if err == nil {
			break
		}
	}
	return vectorToScanner, rotationToScanner, err
}

func TestRegionAfterRotation(beacons []Coordinates3D, allBeacons map[Coordinates3D]bool, rotation int) (Vector, int, error) {
	beaconsToTest := RotateCoordinates(beacons, rotation)
	distances := make(map[Vector]int)

	for beacon0, _ := range allBeacons {
		for _, beaconScanner := range beaconsToTest {
			distances[beacon0.GetVectorTo(beaconScanner)]++
		}
	}

	for distance, occurrences := range distances {
		if occurrences > 11 {
			return distance, rotation, nil
		}
	}

	return Vector{0, 0, 0}, 0, errors.New("region is not overlapping from this rotation")
}

// Data structures

type Coordinates3D struct {
	x, y, z int
}

type Vector struct {
	x, y, z int
}

func (c *Coordinates3D) GetVectorTo(b Coordinates3D) Vector {
	return Vector{c.x - b.x, c.y - b.y, c.z - b.z}
}

func (c *Coordinates3D) Translate(v Vector) Coordinates3D {
	return Coordinates3D{c.x + v.x, c.y + v.y, c.z + v.z}
}

func (c *Coordinates3D) GetManhattanDistanceTo(b Coordinates3D) int {
	return Abs(c.x-b.x) + Abs(c.y-b.y) + Abs(c.z-b.z)
}

func (v *Vector) exists() bool {
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

func (c *Coordinates3D) Rotate(rotation int) Coordinates3D {
	switch rotation {
	// Positive x
	case 1:
		return Coordinates3D{c.x, c.y, c.z}
	case 2:
		return Coordinates3D{c.x, c.z, -c.y}
	case 3:
		return Coordinates3D{c.x, -c.y, -c.z}
	case 4:
		return Coordinates3D{c.x, -c.z, c.y}
	// Positive y
	case 5:
		return Coordinates3D{c.y, c.z, c.x}
	case 6:
		return Coordinates3D{c.y, c.x, -c.z}
	case 7:
		return Coordinates3D{c.y, -c.z, -c.x}
	case 8:
		return Coordinates3D{c.y, -c.x, c.z}
	// Positive z
	case 9:
		return Coordinates3D{c.z, c.x, c.y}
	case 10:
		return Coordinates3D{c.z, c.y, -c.x}
	case 11:
		return Coordinates3D{c.z, -c.x, -c.y}
	case 12:
		return Coordinates3D{c.z, -c.y, c.x}
	// Negative x
	case 13:
		return Coordinates3D{-c.x, -c.y, c.z}
	case 14:
		return Coordinates3D{-c.x, c.z, c.y}
	case 15:
		return Coordinates3D{-c.x, c.y, -c.z}
	case 16:
		return Coordinates3D{-c.x, -c.z, -c.y}
	// Negative y
	case 17:
		return Coordinates3D{-c.y, -c.z, c.x}
	case 18:
		return Coordinates3D{-c.y, c.x, c.z}
	case 19:
		return Coordinates3D{-c.y, c.z, -c.x}
	case 20:
		return Coordinates3D{-c.y, -c.x, -c.z}
	// Negative z
	case 21:
		return Coordinates3D{-c.z, -c.x, c.y}
	case 22:
		return Coordinates3D{-c.z, c.y, c.x}
	case 23:
		return Coordinates3D{-c.z, c.x, -c.y}
	case 24:
		return Coordinates3D{-c.z, -c.y, -c.x}
	}
	return *c
}

func RotateCoordinates(beacons []Coordinates3D, rotation int) []Coordinates3D {
	transformedBeacons := make([]Coordinates3D, len(beacons))

	for i, beacon := range beacons {
		transformedBeacons[i] = beacon.Rotate(rotation)
	}

	return transformedBeacons
}

// Parsing

func ParseScannersAndBeacons(input []string) map[int][]Coordinates3D {
	var currentScanner int
	beaconsPerScanner := make(map[int][]Coordinates3D)

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
			coordinates3d := Coordinates3D{int(parsedX), int(parsedY), int(parsedZ)}
			beaconsPerScanner[currentScanner] = append(beaconsPerScanner[currentScanner], coordinates3d)
		}
	}
	return beaconsPerScanner
}
