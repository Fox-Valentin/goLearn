package testingFn

import "math"

func calcTriangle(a int, b int) int {
	return int(math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2)))
}
