/**
 * Given a binary array, find the maximum number of consecutive 1s in this array.
 *
 * Example 1:<br />
 *
 * Input: [1,1,0,1,1,1]
 * Output: 3
 * Explanation: The first two digits or the last three digits are consecutive 1s.
 *     The maximum number of consecutive 1s is 3.
 *
 *
 *
 * Note:
 *
 * The input array will only contain 0 and 1.
 * The length of input array is a positive integer and will not exceed 10,000
 *
 *
 */

package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	maxVal := 0
	cur := 0
	for i := range nums {
		if nums[i] == 1 {
			cur++
		} else {
			cur = 0
		}
		if maxVal < cur {
			maxVal = cur
		}
	}
	return maxVal
}
