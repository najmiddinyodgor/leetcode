package main

func XorQueries(arr []int, queries [][]int) []int {
	for i := 1; i < len(arr); i++ {
		arr[i] ^= arr[i-1]
	}

	r := make([]int, len(queries))

	for i, q := range queries {
		if q[0] > 0 {
			r[i] = arr[q[0]-1] ^ arr[q[1]]
		} else {
			r[i] = arr[q[1]]
		}
	}

	return r
}
