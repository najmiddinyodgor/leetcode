package main

func MaximumWealth(accounts [][]int) int {
	max := 0
	acc_count := len(accounts)

	for i := 0; i < acc_count; i++ {
		sum := 0
		money_count := len(accounts[i])

		for j := 0; j < money_count; j++ {
			sum += accounts[i][j]
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
