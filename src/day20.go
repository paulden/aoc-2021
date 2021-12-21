package main

import "fmt"

func ProcessEnhancement(input []string, times int) int {
	algorithm := ParseEnhancementAlgorithm(input[0])
	image := ParseInputImage(input[2:])

	// Making images "infinite enough" for enhancement
	enhancedImage := PadImage(image, 3*times)

	for i := 0; i < times; i++ {
		enhancedImage = Enhance(algorithm, enhancedImage)
	}

	return CountLitPixels(enhancedImage)
}

func Enhance(algorithm []int, image [][]int) [][]int {
	enhancedImage := make([][]int, len(image)-2)

	for i := 1; i < len(image)-2; i++ {
		enhancedImage[i-1] = make([]int, len(image[0])-2)
		for j := 1; j < len(image[0])-2; j++ {
			bits := []int{
				image[i-1][j-1],
				image[i-1][j],
				image[i-1][j+1],
				image[i][j-1],
				image[i][j],
				image[i][j+1],
				image[i+1][j-1],
				image[i+1][j],
				image[i+1][j+1],
			}
			index := BinaryToDecimal(bits)
			enhancedImage[i-1][j-1] = algorithm[index]
		}
	}

	return enhancedImage[1:len(enhancedImage)]
}

func CountLitPixels(image [][]int) int {
	litPixels := 0

	for i := range image {
		for j := range image[i] {
			litPixels += image[i][j]
		}
	}

	return litPixels
}

func PadImage(image [][]int, padding int) [][]int {
	paddedImage := make([][]int, len(image)+padding*2)

	for i := 0; i < padding; i++ {
		paddedImage[i] = make([]int, len(image[0])+padding*2)
	}
	for i := len(paddedImage) - padding - 1; i < len(paddedImage); i++ {
		paddedImage[i] = make([]int, len(image[0])+padding*2)
	}

	for i := range image {
		paddedLine := make([]int, len(image[0])+padding*2)
		for j := range image[i] {
			paddedLine[j+padding] = image[i][j]
		}
		paddedImage[i+padding] = paddedLine
	}

	return paddedImage
}

// Parsing

func ParseEnhancementAlgorithm(input string) []int {
	charToBit := map[string]int{
		".": 0,
		"#": 1,
	}
	enhancementAlgorithm := make([]int, len(input))

	for i, char := range input {
		enhancementAlgorithm[i] = charToBit[string(char)]
	}
	return enhancementAlgorithm
}

func ParseInputImage(input []string) [][]int {
	image := make([][]int, len(input))

	for i, line := range input {
		parsedLine := make([]int, len(input[0]))
		for j, char := range line {
			if string(char) == "#" {
				parsedLine[j] = 1
			} else if string(char) == "." {
				parsedLine[j] = 0
			}
		}
		image[i] = parsedLine
	}

	return image
}

// Utils

func BinaryToDecimal(bits []int) int {
	decimal := 0

	for i := range bits {
		decimal += bits[len(bits)-i-1] * Power(2, i)
	}

	return decimal
}

// Utils for debugging

func PrettyPrintImage(image [][]int) {
	intToChar := map[int]string{
		0: ".",
		1: "█",
	}
	for i := range image {
		for j := range image[i] {
			fmt.Printf("%v", intToChar[image[i][j]])
		}
		fmt.Printf("\n")
	}
}
