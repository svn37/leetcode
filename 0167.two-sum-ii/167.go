/**
 * Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
 * The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.
 * Note:
 *
 * 	Your returned answers (both index1 and index2) are not zero-based.
 * 	You may assume that each input would have exactly one solution and you may not use the same element twice.
 *
 *
 * Example 1:
 *
 * Input: numbers = [2,7,11,15], target = 9
 * Output: [1,2]
 * Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 *
 * Example 2:
 *
 * Input: numbers = [2,3,4], target = 6
 * Output: [1,3]
 *
 * Example 3:
 *
 * Input: numbers = [-1,0], target = -1
 * Output: [1,2]
 *
 *
 * Constraints:
 *
 * 	2 <= nums.length <= 3 * 10^4
 * 	-1000 <= nums[i] <= 1000
 * 	nums is sorted in increasing order.
 * 	-1000 <= target <= 1000
 *
 */

package leetcode

func bsearch(arr []int, target int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return -1
}

// Your returned answers (both index1 and index2) are not zero-based.
// You may assume that each input would have exactly one solution
// and you may not use the same element twice
func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		num := target - numbers[i]
		if num < numbers[i] {
			break
		}
		idx := bsearch(numbers[i+1:], num)
		if idx >= 0 {
			return []int{i + 1, i + idx + 2}
		}
	}
	// no solution
	return []int{}
}
