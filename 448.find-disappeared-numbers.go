package main

func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, num := range nums {
		i := (num - 1) % n
		nums[i] += n
	}

	result := []int{}
	for i, num := range nums {
		if num <= n {
			result = append(result, i+1)
		}
	}

	return result
}
