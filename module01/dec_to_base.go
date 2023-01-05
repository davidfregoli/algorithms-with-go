package module01

import "strconv"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {
	s, d := "", ""
	for n := dec; n > 0; n = n / base {
		rest := n % base
		if rest > 9 {
			d = string(rune(55 + rest))
		} else {
			d = strconv.Itoa(rest)
		}
		s = d + s
	}
	return s
}
