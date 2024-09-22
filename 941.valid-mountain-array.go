package main

func ValidMountainArray(arr []int) bool {
	i := 0
	n := len(arr)

	for i+1 < n && arr[i] < arr[i+1] {
		i++
	}

	if i == 0 || i == n-1 {
		return false
	}

	for i+1 < n && arr[i] > arr[i+1] {
		i++
	}

	return i == n-1
}
