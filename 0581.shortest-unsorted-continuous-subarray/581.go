/**
 * Given an integer array nums, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order.
 * Return the shortest such subarray and output its length.
 *
 * Example 1:
 *
 * Input: nums = [2,6,4,8,10,9,15]
 * Output: 5
 * Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
 *
 * Example 2:
 *
 * Input: nums = [1,2,3,4]
 * Output: 0
 *
 * Example 3:
 *
 * Input: nums = [1]
 * Output: 0
 *
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 10^4
 * 	-10^5 <= nums[i] <= 10^5
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	start, end := -1, -2
	maxVal, minVal := nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		maxVal = max(maxVal, nums[i])
		minVal = min(minVal, nums[n-1-i])
		if nums[i] < maxVal {
			end = i
		}
		if nums[n-1-i] > minVal {
			start = n - 1 - i
		}
	}
	return end - start + 1
}
