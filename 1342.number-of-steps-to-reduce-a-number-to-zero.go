package main

func NumberOfSteps(num int) int {
	i := 0

	for num != 0 {
		if num&1 == 0 {
			num >>= 1
		} else {
			num -= 1
		}

		i++
	}

	return i
}
