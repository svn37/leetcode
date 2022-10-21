/**
 * Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
 * Each element in the array represents your maximum jump length at that position.
 * Your goal is to reach the last index in the minimum number of jumps.
 * You can assume that you can always reach the last index.
 *
 * Example 1:
 *
 * Input: nums = [2,3,1,1,4]
 * Output: 2
 * Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
 *
 * Example 2:
 *
 * Input: nums = [2,3,0,1,4]
 * Output: 2
 *
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 3 * 10^4
 * 	0 <= nums[i] <= 10^5
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Method 1. DP (slow)
func jump(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = math.MaxInt64
	}
	memo[0] = 0

	for i := range nums {
		j := nums[i]
		for j != 0 {
			if memo[i] != math.MaxInt64 && i+j < len(memo) {
				memo[i+j] = min(memo[i+j], memo[i]+1)
			}
			j--
		}
	}

	return memo[len(nums)-1]
}

// Method 2. greedy algorithm
func jump2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	jumps := 0
	curEnd, curFarthest := 0, 0
	for i := 0; i <= curEnd && i < len(nums); i++ {
		curFarthest = max(curFarthest, i+nums[i])
		if i == curEnd && i != len(nums)-1 {
			jumps++
			curEnd = curFarthest
		}
	}
	return jumps
}
