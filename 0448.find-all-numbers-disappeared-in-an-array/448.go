/**
 * Given an array of integers where 1 &le; a[i] &le; n (n = size of array), some elements appear twice and others appear once.
 *
 * Find all the elements of [1, n] inclusive that do not appear in this array.
 *
 * Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
 *
 * Example:
 *
 * Input:
 * [4,3,2,7,8,2,3,1]
 *
 * Output:
 * [5,6]
 *
 *
 */

package leetcode

// Method 1. Cyclic swapping
func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	res := []int{}
	for i := range nums {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

// Method 2. Making use of the fact that numbers appear once or twice.
func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func findDisappearedNumbers2(nums []int) []int {
	for i := range nums {
		val := abs(nums[i]) - 1
		if nums[val] > 0 {
			nums[val] *= -1
		}
	}
	res := []int{}
	for i := range nums {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
