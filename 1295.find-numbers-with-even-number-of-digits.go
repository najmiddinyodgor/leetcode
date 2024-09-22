package main

func FindNumbers(nums []int) int {
	c := 0

	for _, num := range nums {
		t, m := 0, num/10

		for m != 0 {
			m /= 10
			t++
		}

		if t%2 == 0 {
			c++
		}
	}

	return c
}
