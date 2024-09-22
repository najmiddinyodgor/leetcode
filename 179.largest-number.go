package main

import (
	"sort"
	"strconv"
	"strings"
)

func LargestNumber(nums []int) string {
	sNums := make([]string, len(nums))

	for i, num := range nums {
		sNums[i] = strconv.Itoa(num)
	}

	sort.Slice(sNums, func(i, j int) bool {
		return sNums[i]+sNums[j] > sNums[j]+sNums[i]
	})

	if sNums[0] == "0" {
		return "0"
	}

	return strings.Join(sNums, "")
}
