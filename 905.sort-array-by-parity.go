package main

func SortArrayByParity(nums []int) []int {
	even := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i], nums[even] = nums[even], nums[i]
			even++
		}
	}

	return nums
}
