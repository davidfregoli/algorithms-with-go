package module01

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	res := []string{}
	for i := 1; i <= n; i++ {
		by3 := i%3 == 0
		by5 := i%5 == 0
		el := strconv.Itoa(i)
		if by3 && by5 {
			el = "Fizz Buzz"
		} else if by3 {
			el = "Fizz"
		} else if by5 {
			el = "Buzz"
		}
		res = append(res, el)
	}
	fmt.Println(strings.Join(res, ", "))
}
