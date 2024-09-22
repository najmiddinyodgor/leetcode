package main

func DuplicateZeroes(arr []int) {
	s := make([]int, 0)
	for _, n := range arr {
		if n == 0 {
			s = append(s, 0, 0)
		} else {
			s = append(s, n)
		}

		if len(s) >= len(arr) {
			copy(arr, s[:len(arr)])
			break
		}
	}
}
