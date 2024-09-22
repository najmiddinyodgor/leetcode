package main

func CountConsistentStrings(allowed string, words []string) int {
	c, m := 0, 0

	for _, w := range allowed {
		m |= 1 << (w - 'a')
	}

	for _, word := range words {
		consistent := true
		for _, w := range word {
			if (m>>(w-'a'))&1 == 0 {
				consistent = false
				break
			}
		}

		if consistent {
			c++
		}
	}

	return c
}
