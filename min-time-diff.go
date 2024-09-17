package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func FindMinDifference(timePoints []string) int {
	converted := make([]int, 0)

	for i := range timePoints {
		parsed := strings.Split(timePoints[i], ":")

		h, _ := strconv.Atoi(parsed[0])
		m, _ := strconv.Atoi(parsed[1])

		converted = append(converted, h*60+m)

		if h == 0 {
			converted = append(converted, 24*60)
		}
	}

	min := -1
	for i := 1; i <= len(timePoints); i++ {
		fmt.Println(converted)
		diff := int(math.Abs(float64(converted[i-1] - converted[i])))

		if min < 0 || diff < min {
			min = diff
		}
	}

	return min
}
