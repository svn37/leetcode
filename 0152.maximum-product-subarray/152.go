/**
 * Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.
 *
 * Example 1:
 *
 *
 * Input: [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 *
 *
 * Example 2:
 *
 *
 * Input: [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 */

package leetcode

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

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]

	// imax/imin stores the max/min product of
	// subarray that ends with the current number nums[i]
	for i, imin, imax := 1, res, res; i < len(nums); i++ {
		// multiplied by a negative makes big number smaller, small number bigger
		// so we swap imin and imax
		if nums[i] < 0 {
			imin, imax = imax, imin
		}

		// max/min product for the current number is either the current number itself
		// or the max/min by the previous number times the current one
		imin = min(nums[i], nums[i]*imin)
		imax = max(nums[i], nums[i]*imax)

		// the newly computed max value is a candidate for our global result
		res = max(res, imax)
	}
	return res
}
