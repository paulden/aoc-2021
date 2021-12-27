package utils

import "math"

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func Power(value, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= value
	}
	return result
}

func Sum(values []int) int {
	var result int
	for _, value := range values {
		result += value
	}
	return result
}

func Product(values []int) int {
	result := 1
	for _, value := range values {
		result *= value
	}
	return result
}

func Minimum(values []int) int {
	result := math.MaxInt
	for _, value := range values {
		if value < result {
			result = value
		}
	}
	return result
}

func Maximum(values []int) int {
	result := math.MinInt
	for _, value := range values {
		if value > result {
			result = value
		}
	}
	return result
}

func GreaterThan(a, b int) int {
	if a > b {
		return 1
	}
	return 0
}

func LessThan(a, b int) int {
	if a < b {
		return 1
	}
	return 0
}

func EqualTo(a, b int) int {
	if a == b {
		return 1
	}
	return 0
}
