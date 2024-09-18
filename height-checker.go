package main

import (
	"sort"
)

func HeightChecker(heights []int) int {
	expected := make([]int, len(heights))

	copy(expected, heights)

	sort.Ints(expected)

	cnt := 0
	for i := 0; i < len(heights); i++ {
		if expected[i] != heights[i] {
			cnt++
		}
	}

	return cnt
}
