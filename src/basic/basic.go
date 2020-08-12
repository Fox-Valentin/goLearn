package basic

import "math"

func CalculateTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
func LengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, v := range []rune(s) {
		if lastI, ok := lastOccurred[v]; ok && lastI >= start {
			start = lastI + i
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[v] = i
	}
	return maxLength
}
