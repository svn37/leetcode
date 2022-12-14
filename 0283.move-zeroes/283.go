/**
 * Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
 *
 * Example:
 *
 *
 * Input: [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 *
 * Note:
 *
 * <ol>
 * 	You must do this in-place without making a copy of the array.
 * 	Minimize the total number of operations.
 * </ol>
 */

package leetcode

func moveZeroes(nums []int) {
	var lastNonZero int
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}
	}
}
