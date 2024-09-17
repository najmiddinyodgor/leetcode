package main

func MinBitFlips(start int, goal int) int {
	c, xor := 0, start^goal

	for xor != 0 {
		xor &= (xor - 1)

		c++
	}

	return c
}
