package utils

import "math"

func ApplyPrecision(value float64, precision int) float64 {
	n := math.Pow(10, float64(precision))
	return math.Round(value*n) / n
}
