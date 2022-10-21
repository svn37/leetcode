/**
 * Given an unsorted integer array nums, find the smallest missing positive integer.
 * Follow up: Could you implement an algorithm that runs in O(n) time and uses constant extra space.?
 *
 * Example 1:
 * Input: nums = [1,2,0]
 * Output: 3
 * Example 2:
 * Input: nums = [3,4,-1,1]
 * Output: 2
 * Example 3:
 * Input: nums = [7,8,9,11,12]
 * Output: 1
 *
 * Constraints:
 *
 * 	0 <= nums.length <= 300
 * 	-2^31 <= nums[i] <= 2^31 - 1
 *
 */

package leetcode

// cyclic swapping
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
