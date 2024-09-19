package main

import (
	"strconv"
	"unicode"
)

func DiffWaysToCompute(expression string) []int {
	if len(expression) == 0 {
		return []int{}
	}

	if len(expression) <= 2 && unicode.IsDigit(rune(expression[0])) {
		num, _ := strconv.Atoi(expression)

		return []int{num}
	}

	result := []int{}

	for i, elem := range expression {
		if unicode.IsDigit(elem) {
			continue
		}

		left_values := DiffWaysToCompute(expression[:i])
		right_values := DiffWaysToCompute(expression[i+1:])

		for _, left_value := range left_values {
			for _, right_value := range right_values {
				switch elem {
				case '+':
					{
						result = append(result, left_value+right_value)
					}
				case '-':
					{
						result = append(result, left_value-right_value)
					}
				case '*':
					{
						result = append(result, left_value*right_value)
					}
				}
			}
		}
	}

	return result
}
