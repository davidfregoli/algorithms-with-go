package module01

func GCD(a, b int) int {
	gcd := 1
	for i := 2; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return gcd
}
