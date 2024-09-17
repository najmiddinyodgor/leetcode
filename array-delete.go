package main

func RemoveElement(nums []int, val int) int {
	cnt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[cnt] = nums[cnt], nums[i]
			cnt++
		}
	}

	return cnt
}
