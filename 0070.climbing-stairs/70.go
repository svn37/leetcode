/**
 * You are climbing a stair case. It takes n steps to reach to the top.
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
 * Example 1:
 *
 * Input: 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 *
 * Example 2:
 *
 * Input: 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 *
 *
 * Constraints:
 *
 * 	1 <= n <= 45
 *
 */

package leetcode

// Method 1. Recursive with memo.
func climbWithMap(n int, m map[int]int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	oneStepBack, ok := m[n-1]
	if !ok {
		oneStepBack = climbWithMap(n-1, m)
		m[n-1] = oneStepBack
	}

	twoStepsBack, ok := m[n-2]
	if !ok {
		twoStepsBack = climbWithMap(n-2, m)
		m[n-2] = twoStepsBack
	}

	return oneStepBack + twoStepsBack
}

func climbStairs(n int) int {
	return climbWithMap(n, make(map[int]int))
}

// Method 2. Dynamic programming.
func climbStairs2(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}

// Method 3. We don't need the dp array, just two variables.
func climbStairs3(n int) int {
	if n == 0 {
		return 0
	}

	oneStepBack, twoStepsBack := 1, 1
	for i := 2; i <= n; i++ {
		oneStepBack, twoStepsBack = oneStepBack+twoStepsBack, oneStepBack
	}

	return oneStepBack
}
