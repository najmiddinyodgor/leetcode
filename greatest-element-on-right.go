package main

func ReplaceElements(arr []int) []int {
	temp, max := 0, -1

	for i := len(arr) - 1; i >= 0; i-- {
		temp = arr[i]

		arr[i] = max

		if temp > max {
			max = temp
		}
	}

	return arr
}
