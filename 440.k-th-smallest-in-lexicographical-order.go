package main

func FindKthNumber(n int, k int) int {
	current, currentIndex := 1, 1

	for i := 0; i < n; i++ {
		if currentIndex == k {
			return current
		}

		if current*10 <= n {
			current *= 10
		} else {
			for current%10 == 9 || current >= n {
				current /= 10
			}
			current++
		}

		currentIndex++
	}

	return 0
}
