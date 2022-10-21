/**
 * Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
 * Example:
 *
 * Input:  [1,2,3,4]
 * Output: [24,12,8,6]
 *
 * Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array (including the whole array) fits in a 32 bit integer.
 * Note: Please solve it without division and in O(n).
 * Follow up:<br />
 * Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)
 *
 */

package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)

	// total product = left products * right products
	// use result array as space for left products
	res := make([]int, n)
	left := 1
	for i := 0; i < n; i++ {
		if i > 0 {
			left *= nums[i-1]
		}
		res[i] = left
	}

	// instead of using O(n) space and another array, right variable is enough
	right := 1
	for i := n - 1; i >= 0; i-- {
		if i < n-1 {
			right *= nums[i+1]
		}
		res[i] *= right
	}

	return res
}
