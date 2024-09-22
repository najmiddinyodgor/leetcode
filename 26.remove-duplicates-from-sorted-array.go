package main

func RemoveDuplicates(nums []int) int {
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[cnt] = nums[i]
			cnt++
		}
	}

	return cnt
}
