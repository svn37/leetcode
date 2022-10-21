/**
 * Given a set of distinct integers, nums, return all possible subsets (the power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: nums = [1,2,3]
 * Output:
 * [
 *   [3],
 *   [1],
 *   [2],
 *   [1,2,3],
 *   [1,3],
 *   [2,3],
 *   [1,2],
 *   []
 * ]
 *
 */

package leetcode

import "math"

func pow(n int) int {
	return int(math.Pow(2, float64(n)))
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0, pow(len(nums)))
	res = append(res, []int{}) // empty set

	for i := range nums {
		for j := range res {
			subset := make([]int, len(res[j]))
			copy(subset, res[j])

			subset = append(subset, nums[i])
			res = append(res, subset)
		}
	}

	return res
}
