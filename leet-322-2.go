package main

func coinChange2(coins []int, amount int) int {

	dp := make([]int, amount+1)

	dp[0] = 0

	// dp[i] denotes the optimal solution for the amount i
	// initialize to "infinity" which in this context can be amount + 1 or math.MaxInt64

	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	// iterate through every possible amount starting from 1

	for i := 1; i <= amount; i++ {

		// for given amount compute optimal (minimum) amount as per recurrence relation
		// filtering out denominations greater than the current amount
		// optimal solution will be stored in dp[i]

		for _, coin := range coins {

			if coin <= i {

				dp[i] = min(dp[i], dp[i-coin]+1)

			}
		}
	}

	// flag case where optimal solution does not exist
	// dp[amount] still equal to initial "infinity" amount + 1

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
