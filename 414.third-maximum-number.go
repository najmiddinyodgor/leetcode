package main

import "math"

func ThirdMax(nums []int) int {
	var num float64
	first, second, third := math.Inf(-1), math.Inf(-1), math.Inf(-1)

	for i := range nums {
		num = float64(nums[i])
		if first == num || second == num || third == num {
			continue
		}

		if first < num {
			first, second, third = num, first, second
		} else if second < num {
			second, third = num, second
		} else if third < num {
			third = num
		}
	}

	if third == math.Inf(-1) {
		return int(first)
	}

	return int(third)
}
