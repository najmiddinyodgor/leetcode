package main

import (
	"strconv"
)

func FizzBuzz(n int) []string {
	r := make([]string, n)

	for i := 1; i <= n; i++ {
		var str string

		if i%15 == 0 {
			str += "FizzBuzz"
		} else if i%3 == 0 {
			str += "Fizz"
		} else if i%5 == 0 {
			str += "Buzz"
		} else {
			str += strconv.Itoa(i)
		}

		r[i-1] = str
	}

	return r
}
