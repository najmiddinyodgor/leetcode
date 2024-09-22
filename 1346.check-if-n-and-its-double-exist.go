package main

func CheckIfExist(arr []int) bool {
	z := 0
	m := make(map[int]bool)

	for _, num := range arr {
		m[num] = true

		if num == 0 {
			z++
		}
	}

	for _, num := range arr {
		if num == 0 {
			if z > 1 {
				return true
			}
		} else if m[num*2] {
			return true
		}
	}

	return false
}
