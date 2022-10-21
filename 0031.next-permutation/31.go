/**
 * Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
 * If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).
 * The replacement must be <a href="http://en.wikipedia.org/wiki/In-place_algorithm" target="_blank">in place</a> and use only constant extra memory.
 *
 * Example 1:
 * Input: nums = [1,2,3]
 * Output: [1,3,2]
 * Example 2:
 * Input: nums = [3,2,1]
 * Output: [1,2,3]
 * Example 3:
 * Input: nums = [1,1,5]
 * Output: [1,5,1]
 * Example 4:
 * Input: nums = [1]
 * Output: [1]
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 100
 * 	0 <= nums[i] <= 100
 *
 */

package leetcode

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int) {
	// last peak element index
	l := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			l = i
		}
	}

	if l != 0 {
		// find index of the next element larger than nums[l-1]
		k := 0
		for i := l; i < len(nums); i++ {
			if nums[i] > nums[l-1] {
				k = i
			}
		}
		nums[k], nums[l-1] = nums[l-1], nums[k]
	}
	reverse(nums[l:])
}
