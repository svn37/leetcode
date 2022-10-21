/**
 * Given an array of integers, return the maximum sum for a non-empty subarray (contiguous elements) with at most one element deletion. In other words, you want to choose a subarray and optionally delete one element from it so that there is still at least one element left and the sum of the remaining elements is maximum possible.
 *
 * Note that the subarray needs to be non-empty after deleting one element.
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [1,-2,0,3]
 * Output: 4
 * Explanation: Because we can choose [1, -2, 0, 3] and drop -2, thus the subarray [1, 0, 3] becomes the maximum value.
 *
 * Example 2:
 *
 *
 * Input: arr = [1,-2,-2,3]
 * Output: 3
 * Explanation: We just choose [3] and it's the maximum sum.
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [-1,-1,-1,-1]
 * Output: -1
 * Explanation: The final subarray needs to be non-empty. You can't choose [-1] and delete -1 from it, then get an empty subarray to make the sum equals to 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 	1 <= arr.length <= 10^5
 * 	-10^4 <= arr[i] <= 10^4
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ignored, notignored, maxSum, maxVal := 0, 0, nums[0], nums[0]

	for i := range nums {
		if nums[i] >= 0 {
			ignored += nums[i]
		} else {
			ignored = max(ignored+nums[i], notignored)
		}
		notignored += nums[i]

		if ignored < 0 {
			ignored = 0
		}
		if notignored < 0 {
			notignored = 0
		}

		maxSum = max(maxSum, max(ignored, notignored))
		maxVal = max(maxVal, nums[i])
	}
	if maxVal < 0 {
		return maxVal
	}
	return maxSum
}
