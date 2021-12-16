package main

import (
	"testing"
)

func TestDay16HexadecimalToBits(t *testing.T) {
	// Given
	input := "EE00D40C823060"
	expected := []int{1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0}

	// When
	result := HexadecimalToBits(input)

	// Then
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Day 16 example - Part 1 hexadecimal to bits: expected %v, got %v", expected[i], result[i])
		}
	}
}

func Test16BitsToDecimal(t *testing.T) {
	// Given
	input := []int{0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1}
	expected := 2021

	// When
	result := BitsToDecimal(input)

	// Then
	if result != expected {
		t.Errorf("Day 16 example - Part 1 first example: expected %v, got %v", expected, result)
	}
}

func Test16ParseLiteralPacket(t *testing.T) {
	// Given
	input := HexadecimalToBits("D2FE28")
	// 110100101111111000101000
	// VVVTTTAAAAABBBBBCCCCC
	expectedVersion := 6
	expectedLiteralValue := 2021
	expectedPointer := 21

	// When
	packetVersion, literalValue, pointer := ParseLiteralPacket(input, 0)

	// Then
	if packetVersion != expectedVersion {
		t.Errorf("Day 16 example - Part 1 first example: expected version %v, got %v", expectedVersion, packetVersion)
	}
	if literalValue != expectedLiteralValue {
		t.Errorf("Day 16 example - Part 1 first example: expected value %v, got %v", expectedLiteralValue, literalValue)
	}
	if pointer != expectedPointer {
		t.Errorf("Day 16 example - Part 1 first example: expected pointer %v, got %v", expectedPointer, pointer)
	}
}

func Test16ParseOperatorPacketType0(t *testing.T) {
	// Given
	input := HexadecimalToBits("38006F45291200")
	// 00111000000000000110111101000101001010010001001000000000
	// VVVTTTILLLLLLLLLLLLLLLAAAAAAAAAAABBBBBBBBBBBBBBBB
	expectedVersion := 1 + 6 + 2
	expectedPointer := 49

	// When
	pointer, packetVersions, _ := ParsePacket(input, 0, 0)

	// Then
	if packetVersions != expectedVersion {
		t.Errorf("Day 16 example - Part 1 first example: expected version %v, got %v", expectedVersion, packetVersions)
	}
	if pointer != expectedPointer {
		t.Errorf("Day 16 example - Part 1 first example: expected pointer %v, got %v", expectedPointer, pointer)
	}
}

func Test16ParseOperatorPacketType1(t *testing.T) {
	// Given
	input := HexadecimalToBits("EE00D40C823060")
	// 11101110000000001101010000001100100000100011000001100000
	// VVVTTTILLLLLLLLLLLAAAAAAAAAAABBBBBBBBBBBCCCCCCCCCCC
	expectedVersion := 7 + 2 + 4 + 1
	expectedPointer := 51

	// When
	pointer, packetVersions, _ := ParsePacket(input, 0, 0)

	// Then
	if packetVersions != expectedVersion {
		t.Errorf("Day 16 example - Part 1 first example: expected version %v, got %v", expectedVersion, packetVersions)
	}
	if pointer != expectedPointer {
		t.Errorf("Day 16 example - Part 1 first example: expected pointer %v, got %v", expectedPointer, pointer)
	}
}

func Test16ParsePacketSimple(t *testing.T) {
	// Given
	input := HexadecimalToBits("D2FE28")
	expectedVersion := 6
	expectedPointer := 21

	// When
	pointer, packetVersions, _ := ParsePacket(input, 0, 0)

	// Then
	if packetVersions != expectedVersion {
		t.Errorf("Day 16 example - Part 1 first example: expected version %v, got %v", expectedVersion, packetVersions)
	}
	if pointer != expectedPointer {
		t.Errorf("Day 16 example - Part 1 first example: expected pointer %v, got %v", expectedPointer, pointer)
	}
}

func Test16Part1FirstExample(t *testing.T) {
	// Given
	input := "8A004A801A8002F478"
	expected := 16

	// When
	result := SumVersionNumbers(input)

	// Then
	if result != expected {
		t.Errorf("Day 16 example - Part 1 first example: expected %v, got %v", expected, result)
	}
}

func Test16Part1SecondExample(t *testing.T) {
	// Given
	input := "620080001611562C8802118E34"
	expected := 12

	// When
	result := SumVersionNumbers(input)

	// Then
	if result != expected {
		t.Errorf("Day 16 example - Part 1 second example: expected %v, got %v", expected, result)
	}
}

func Test16Part1ThirdExample(t *testing.T) {
	// Given
	input := "C0015000016115A2E0802F182340"
	expected := 23

	// When
	result := SumVersionNumbers(input)

	// Then
	if result != expected {
		t.Errorf("Day 16 example - Part 1 third example: expected %v, got %v", expected, result)
	}
}

func Test16Part1FourthExample(t *testing.T) {
	// Given
	input := "A0016C880162017C3686B18A3D4780"
	expected := 31

	// When
	result := SumVersionNumbers(input)

	// Then
	if result != expected {
		t.Errorf("Day 16 example - Part 1 fourth example: expected %v, got %v", expected, result)
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