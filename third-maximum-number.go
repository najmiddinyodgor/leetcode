package main

func ThirdMax(nums []int) int {
	n := 3
	result := make([]int, 0)

	for _, num := range nums {
		for j := 0; j < n; j++ {
			if len(result) == j {
				result = append(result, num)
			} else if result[j] < num {
				result[j] = num
			}
		}
	}

	return result[2]
}
