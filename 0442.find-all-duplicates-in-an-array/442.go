/**
 * Given an array of integers, 1 &le; a[i] &le; n (n = size of array), some elements appear twice and others appear once.
 *
 * Find all the elements that appear twice in this array.
 *
 * Could you do it without extra space and in O(n) runtime?
 *
 * Example:<br/>
 *
 * Input:
 * [4,3,2,7,8,2,3,1]
 *
 * Output:
 * [2,3]
 *
 */

package leetcode

func abs(a int) int {
	if a < 0 {
		return a * (-1)
	}
	return a
}

func findDuplicates(nums []int) []int {
	res := []int{}
	for i := range nums {
		absVal := abs(nums[i])
		if nums[absVal-1] < 0 {
			res = append(res, absVal)
		} else {
			nums[absVal-1] *= -1
		}
	}
	return res
}
