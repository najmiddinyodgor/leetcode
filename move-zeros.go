package main

func MoveZeroes(nums []int) {
	i, cnt := 0, 0

	for i < len(nums) {
		if nums[i] != 0 {
			nums[cnt] = nums[i]
			cnt++
		}
		i++
	}

	for cnt < len(nums) {
		nums[cnt] = 0
		cnt++
	}
}
