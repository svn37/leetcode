/**
 * You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
 * You may assume that you have an infinite number of each kind of coin.
 *
 * Example 1:
 *
 * Input: coins = [1,2,5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 *
 * Example 2:
 *
 * Input: coins = [2], amount = 3
 * Output: -1
 *
 * Example 3:
 *
 * Input: coins = [1], amount = 0
 * Output: 0
 *
 * Example 4:
 *
 * Input: coins = [1], amount = 1
 * Output: 1
 *
 * Example 5:
 *
 * Input: coins = [1], amount = 2
 * Output: 2
 *
 *
 * Constraints:
 *
 * 	1 <= coins.length <= 12
 * 	1 <= coins[i] <= 2^31 - 1
 * 	0 <= amount <= 10^4
 *
 */

package leetcode

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func helper(coins []int, amount int, memo []int) int {
	if amount == 0 {
		return 0
	} // completed
	if amount < 0 {
		return -1
	} // invalid

	if memo[amount] != math.MaxInt64 {
		return memo[amount]
	}

	minVal := math.MaxInt64
	for _, coin := range coins {
		num := helper(coins, amount-coin, memo)
		if num != -1 {
			minVal = min(minVal, num)
		}
	}

	if minVal == math.MaxInt64 {
		memo[amount] = -1
	} else {
		memo[amount] = minVal + 1
	}
	return memo[amount]
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = math.MaxInt64
	}
	return helper(coins, amount, memo)
}

// Bottom-up approach
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for curamount := 1; curamount <= amount; curamount++ {
		for _, coin := range coins {
			if coin <= curamount {
				dp[curamount] = min(dp[curamount], dp[curamount-coin]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
