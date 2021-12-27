package day16

import (
	"aoc-2021/internal/pkg/utils"
	"testing"
)

func TestDay16ParseLiteralPacket(t *testing.T) {
	// Given
	input := utils.HexadecimalToBits("D2FE28")
	// 110100101111111000101000
	// VVVTTTAAAAABBBBBCCCCC
	expectedVersion := 6
	expectedLiteralValue := 2021
	expectedPointer := 21

	// When
	packetVersion, literalValue, pointer := parseLiteralPacket(input, 0)

	// Then
	if packetVersion != expectedVersion {
		t.Errorf("Day 16 - Expected version %v, got %v", expectedVersion, packetVersion)
	}
	if literalValue != expectedLiteralValue {
		t.Errorf("Day 16 - Expected value %v, got %v", expectedLiteralValue, literalValue)
	}
	if pointer != expectedPointer {
		t.Errorf("Day 16 - Expected pointer %v, got %v", expectedPointer, pointer)
	}
}

func TestDay16ParseOperatorPacketType0(t *testing.T) {
	// Given
	input := utils.HexadecimalToBits("38006F45291200")
	// 00111000000000000110111101000101001010010001001000000000
	// VVVTTTILLLLLLLLLLLLLLLAAAAAAAAAAABBBBBBBBBBBBBBBB
	expectedVersion := 1 + 6 + 2
	expectedPointer := 49

	// When
	pointer, packetVersions, _ := parsePacket(input, 0, 0)

	// Then
	if packetVersions != expectedVersion {
		t.Errorf("Day 16 - Expected version %v, got %v", expectedVersion, packetVersions)
	}
	if pointer != expectedPointer {
		t.Errorf("Day 16 - Expected pointer %v, got %v", expectedPointer, pointer)
	}
}

func TestDay16ParseOperatorPacketType1(t *testing.T) {
	// Given
	input := utils.HexadecimalToBits("EE00D40C823060")
	// 11101110000000001101010000001100100000100011000001100000
	// VVVTTTILLLLLLLLLLLAAAAAAAAAAABBBBBBBBBBBCCCCCCCCCCC
	expectedVersion := 7 + 2 + 4 + 1
	expectedPointer := 51

	// When
	pointer, packetVersions, _ := parsePacket(input, 0, 0)

	// Then
	if packetVersions != expectedVersion {
		t.Errorf("Day 16 - Expected version %v, got %v", expectedVersion, packetVersions)
	}
	if pointer != expectedPointer {
		t.Errorf("Day 16 - Expected pointer %v, got %v", expectedPointer, pointer)
	}
}

func TestDay16ParsePacketSimple(t *testing.T) {
	// Given
	input := utils.HexadecimalToBits("D2FE28")
	expectedVersion := 6
	expectedPointer := 21

	// When
	pointer, packetVersions, _ := parsePacket(input, 0, 0)

	// Then
	if packetVersions != expectedVersion {
		t.Errorf("Day 16 - Expected version %v, got %v", expectedVersion, packetVersions)
	}
	if pointer != expectedPointer {
		t.Errorf("Day 16 - Expected pointer %v, got %v", expectedPointer, pointer)
	}
}

func TestDay16Part1FirstExample(t *testing.T) {
	// Given
	input := "8A004A801A8002F478"
	expected := 16

	// When
	result := SumVersionNumbers(input)

	// Then
	if result != expected {
		t.Errorf("Day 16 - Expected %v, got %v", expected, result)
	}
}

func TestDay16Part1SecondExample(t *testing.T) {
	// Given
	input := "620080001611562C8802118E34"
	expected := 12

	// When
	result := SumVersionNumbers(input)

	// Then
	if result != expected {
		t.Errorf("Day 16 - Expected %v, got %v", expected, result)
	}
}

func TestDay16Part1ThirdExample(t *testing.T) {
	// Given
	input := "C0015000016115A2E0802F182340"
	expected := 23

	// When
	result := SumVersionNumbers(input)

	// Then
	if result != expected {
		t.Errorf("Day 16 - Expected %v, got %v", expected, result)
	}
}

func TestDay16Part1FourthExample(t *testing.T) {
	// Given
	input := "A0016C880162017C3686B18A3D4780"
	expected := 31

	// When
	result := SumVersionNumbers(input)

	// Then
	if result != expected {
		t.Errorf("Day 16 - Expected %v, got %v", expected, result)
	}
}

func TestEvaluateBITSExpression(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"C200B40A82", args{"C200B40A82"}, 3},
		{"04005AC33890", args{"04005AC33890"}, 54},
		{"880086C3E88112", args{"880086C3E88112"}, 7},
		{"CE00C43D881120", args{"CE00C43D881120"}, 9},
		{"D8005AC2A8F0", args{"D8005AC2A8F0"}, 1},
		{"F600BC2D8F", args{"F600BC2D8F"}, 0},
		{"9C005AC2F8F0", args{"9C005AC2F8F0"}, 0},
		{"9C0141080250320F1802104A08", args{"9C0141080250320F1802104A08"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvaluateBITSExpression(tt.args.input); got != tt.want {
				t.Errorf("EvaluateBITSExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay16Part1(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 843

	// When
	result := SumVersionNumbers(input[0])

	// Then
	if result != expected {
		t.Errorf("Day 16 - Expected %v, got %v", expected, result)
	}
}

func TestDay16Part2(t *testing.T) {
	// Given
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 5390807940351

	// When
	result := EvaluateBITSExpression(input[0])

	// Then
	if result != expected {
		t.Errorf("Day 16 - Expected %v, got %v", expected, result)
	}
}
