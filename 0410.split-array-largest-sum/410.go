/**
 * Given an array nums which consists of non-negative integers and an integer m, you can split the array into m non-empty continuous subarrays.
 * Write an algorithm to minimize the largest sum among these m subarrays.
 *
 * Example 1:
 *
 * Input: nums = [7,2,5,10,8], m = 2
 * Output: 18
 * Explanation:
 * There are four ways to split nums into two subarrays.
 * The best way is to split it into [7,2,5] and [10,8],
 * where the largest sum among the two subarrays is only 18.
 *
 * Example 2:
 *
 * Input: nums = [1,2,3,4,5], m = 2
 * Output: 9
 *
 * Example 3:
 *
 * Input: nums = [1,4,4], m = 3
 * Output: 4
 *
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 1000
 * 	0 <= nums[i] <= 10^6
 * 	1 <= m <= min(50, nums.length)
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cannotDivide(target int, nums []int, parts int) bool {
	count := 1
	total := 0
	for i := range nums {
		total += nums[i]

		if total > target {
			total = nums[i]
			count++
			if count > parts {
				return false
			}
		}
	}
	return true
}

func splitArray(nums []int, m int) int {
	maxVal, sum := 0, 0
	for i := range nums {
		maxVal = max(maxVal, nums[i])
		sum += nums[i]
	}

	if m == 1 {
		return sum
	}

	// the answer is between maximum value of input array numbers and sum of those numbers
	// use binary search to approach the correct answer
	lo, hi := maxVal, sum
	for lo <= hi {
		mid := (lo + hi) / 2
		// try to divide the array into more than m subarrays
		if cannotDivide(mid, nums, m) {
			// either we successfully divide the array into m parts
			// and the sum of each part is less than mid,
			// or we used up all numbers before we reach m
			hi = mid - 1
		} else {
			// we still have numbers left
			// so, it is impossible to cut the array into m parts
			// and make sure each parts is no larger than mid
			lo = mid + 1
		}
	}
	return lo
}
