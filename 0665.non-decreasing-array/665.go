/**
 * Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.
 * We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).
 *
 * Example 1:
 *
 * Input: nums = [4,2,3]
 * Output: true
 * Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
 *
 * Example 2:
 *
 * Input: nums = [4,2,1]
 * Output: false
 * Explanation: You can't get a non-decreasing array by modify at most one element.
 *
 *
 * Constraints:
 *
 * 	1 <= n <= 10 ^ 4
 * 	- 10 ^ 5 <= nums[i] <= 10 ^ 5
 *
 */

package leetcode

// This problem is like a greedy problem. When you find nums[i-1] > nums[i] for some i,
// you will prefer to change nums[i-1]'s value, since a larger nums[i] will give you more
// risks that you get inversion errors after position i. But, if you also find nums[i-2] > nums[i],
// then you have to change nums[i]'s value instead, or else you need to change both of nums[i-2]'s and nums[i-1]'s values
func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums) && count <= 1; i++ {
		if nums[i-1] > nums[i] {
			count++
			// Correction can be made in either of two ways:
			// 1) Make the previous number smaller or equal to current number
			// 2) Make the current number equal to previous number
			if i-2 < 0 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}
	return count <= 1
}
