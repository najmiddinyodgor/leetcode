package main

func RemoveElement(nums []int, val int) int {
	i, t, r := 0, 0, len(nums)-1

	for i <= r {
		if nums[i] == val {
			if nums[r] != val {
				t = nums[i]
				nums[i] = nums[r]
				nums[r] = t
				i++
			}

			r--
		} else {
			i++
		}
	}

	return i
}
