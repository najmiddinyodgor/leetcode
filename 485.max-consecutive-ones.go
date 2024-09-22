package main

func FindMaxConsecutiveOnes(nums []int) int {
	c, max := 0, 0

	for _, num := range nums {
		if num == 1 {
			c++

			if c > max {
				max = c
			}
		} else {
			c = 0
		}
	}

	return max
}
