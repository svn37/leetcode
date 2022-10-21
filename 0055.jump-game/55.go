/**
 * Given an array of non-negative integers, you are initially positioned at the first index of the array.
 * Each element in the array represents your maximum jump length at that position.
 * Determine if you are able to reach the last index.
 *
 * Example 1:
 *
 * Input: nums = [2,3,1,1,4]
 * Output: true
 * Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
 *
 * Example 2:
 *
 * Input: nums = [3,2,1,0,4]
 * Output: false
 * Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
 *
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 3 * 10^4
 * 	0 <= nums[i][j] <= 10^5
 *
 */

package leetcode

// complicated solution with O(n) space
func canjump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	memo := make([]bool, len(nums))
	memo[0] = true
	for i := range nums {
		if memo[i] {
			jump := nums[i]
			for jump != 0 {
				if i+jump < len(nums) {
					memo[i+jump] = true
				}
				if i+jump == len(nums)-1 {
					return true
				}
				jump--
			}
		}
	}
	return memo[len(memo)-1]
}

// simple solution with one variable
func canJump(nums []int) bool {
	reach := 0
	for i := reach; i < len(nums) && i <= reach; i++ {
		if reach < i+nums[i] {
			reach = i + nums[i]
		}
	}
	return reach >= len(nums)-1
}
