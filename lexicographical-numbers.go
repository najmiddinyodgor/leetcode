package main

func LexicalOrder(n int) []int {
	result := []int{}

	for i := 1; i < 10; i++ {
		result = append(result, dfs(i, n)...)
	}

	return result
}

func dfs(from, until int) []int {
	if from > until {
		return []int{}
	}

	var next int
	result := []int{from}

	for i := range 10 {
		next = from*10 + i
		if next <= until {
			result = append(result, dfs(next, until)...)
		} else {
			break
		}
	}

	return result
}
