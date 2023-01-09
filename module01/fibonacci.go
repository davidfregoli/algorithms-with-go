package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//	Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//	Fibonacci(0) => 0
//	Fibonacci(1) => 1
//
// Examples:
//
//	Fibonacci(0) => 0
//	Fibonacci(1) => 1
//	Fibonacci(2) => 1
//	Fibonacci(3) => 2
//	Fibonacci(4) => 3
//	Fibonacci(5) => 5
//	Fibonacci(6) => 8
//	Fibonacci(7) => 13
//	Fibonacci(14) => 377

var FibonacciMem = map[int]int{}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	val, cached := FibonacciMem[n]
	if !cached {
		val = Fibonacci(n-1) + Fibonacci(n-2)
		FibonacciMem[n] = val
	}
	return val
}

func FibonacciIter(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var next int
	seq := []int{0, 1}
	for i := 2; i <= n; i++ {
		next = seq[i-2] + seq[i-1]
		seq = append(seq, next)
	}
	return next
}

func FibonacciRec(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibonacciRec(n-1) + FibonacciRec(n-2)
}
