/**
 * Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.
 *
 * Example 1:
 * Input: nums = [1,1,1], k = 2
 * Output: 2
 * Example 2:
 * Input: nums = [1,2,3], k = 3
 * Output: 2
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 2 * 10^4
 * 	-1000 <= nums[i] <= 1000
 * 	-10^7 <= k <= 10^7
 *
 */

package leetcode

func subarraySum(nums []int, k int) int {
	var count, sum int
	m := make(map[int]int)
	m[0] = 1

	for _, num := range nums {
		sum += num
		if c, ok := m[sum-k]; ok {
			count += c
		}
		m[sum]++
	}
	return count
}
