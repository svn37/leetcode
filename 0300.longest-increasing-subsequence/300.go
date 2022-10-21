/**
 * Given an unsorted array of integers, find the length of longest increasing subsequence.
 *
 * Example:
 *
 *
 * Input: [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
 *
 * Note:
 *
 *
 * 	There may be more than one LIS combination, it is only necessary for you to return the length.
 * 	Your algorithm should run in O(n^2) complexity.
 *
 *
 * Follow up: Could you improve it to O(n log n) time complexity?
 *
 */

package leetcode

func lengthOfLIS(nums []int) int {
	tails := make([]int, len(nums))
	size := 0
	for _, num := range nums {
		i, j := 0, size
		for i != j {
			m := (i + j) / 2
			if tails[m] < num {
				i = m + 1
			} else {
				j = m
			}
		}
		tails[i] = num
		if i == size {
			size++
		}
	}
	return size
}
