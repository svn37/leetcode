/**
 * Given an array, rotate the array to the right by k steps, where k is non-negative.
 * Follow up:
 *
 * 	Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
 * 	Could you do it in-place with O(1) extra space?
 *
 *
 * Example 1:
 *
 * Input: nums = [1,2,3,4,5,6,7], k = 3
 * Output: [5,6,7,1,2,3,4]
 * Explanation:
 * rotate 1 steps to the right: [7,1,2,3,4,5,6]
 * rotate 2 steps to the right: [6,7,1,2,3,4,5]
 * rotate 3 steps to the right: [5,6,7,1,2,3,4]
 *
 * Example 2:
 *
 * Input: nums = [-1,-100,3,99], k = 2
 * Output: [3,99,-1,-100]
 * Explanation:
 * rotate 1 steps to the right: [99,-1,-100,3]
 * rotate 2 steps to the right: [3,99,-1,-100]
 *
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 2 * 10^4
 * 	-2^31 <= nums[i] <= 2^31 - 1
 * 	0 <= k <= 10^5
 *
 */

package leetcode

// Method 1. Create a nums+nums array and copy the right interval.
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	extra := append(nums, nums...)
	copy(nums, extra[n-k:len(extra)-k])
}

// Method 2. Reverse.
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n

	// reverse the first n-k numbers
	reverse(nums[:n-k])

	// reverse the last k numbers
	reverse(nums[n-k:])

	// reverse all the numbers
	reverse(nums)
}

// Method 3. Cyclic replacements.
// [1,2,3,4,5,6,7]
// 3

// First outer for loop
// [5 6 7 4 1 2 3]
// Second outer for loop
// [1 2 3 4]

func rotate3(nums []int, k int) {
	for n := len(nums); n > 0 && k%n > 0; n = len(nums) {
		k %= n

		for i := 0; i < k; i++ {
			nums[i], nums[n-k+i] = nums[n-k+i], nums[i]
		}

		nums = nums[k:]
	}
}
