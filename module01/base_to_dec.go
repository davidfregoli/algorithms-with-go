package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	charset := "0123456789ABCDEF"
	sum := 0
	s := Reverse(value)
	for i, n := range s {
		sum += int(math.Pow(float64(base), float64(i))) * strings.IndexRune(charset, n)
	}
	return sum
}
