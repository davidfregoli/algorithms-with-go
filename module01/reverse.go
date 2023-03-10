package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	s := ""
	for _, l := range word {
		s = string(l) + s
	}
	return s
}
