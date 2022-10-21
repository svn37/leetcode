/**
 * Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
 * Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
 *
 * Example 1:
 *
 * Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 * Example 2:
 *
 * Input: nums = [1]
 * Output: 1
 *
 * Example 3:
 *
 * Input: nums = [0]
 * Output: 0
 *
 * Example 4:
 *
 * Input: nums = [-1]
 * Output: -1
 *
 * Example 5:
 *
 * Input: nums = [-2147483647]
 * Output: -2147483647
 *
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 2 * 10^4
 * 	-2^31 <= nums[i] <= 2^31 - 1
 *
 */

package leetcode

// Method 1. Kadane's algorithm.
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curSum, maxSum := nums[0], nums[0]
	for _, num := range nums[1:] {
		if curSum+num > num {
			curSum += num
		} else {
			curSum = num
		}

		if maxSum < curSum {
			maxSum = curSum
		}
	}

	return maxSum
}

// Method 2. Dynamic programming.
// maxSubArray(nums, i) = maxSubArray(nums, i - 1) > 0 ? maxSubArray(nums, i - 1) : 0 + nums[i];
func maxSubArray2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}

		if maxSum < dp[i] {
			maxSum = dp[i]
		}
	}

	return maxSum
}
