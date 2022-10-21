/**
 * Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
 * There is only one duplicate number in nums, return this duplicate number.
 * Follow-ups:
 * <ol>
 * 	How can we prove that at least one duplicate number must exist in nums?
 * 	Can you solve the problem without modifying the array nums?
 * 	Can you solve the problem using only constant, O(1) extra space?
 * 	Can you solve the problem with runtime complexity less than O(n^2)?
 * </ol>
 *
 * Example 1:
 * Input: nums = [1,3,4,2,2]
 * Output: 2
 * Example 2:
 * Input: nums = [3,1,3,4,2]
 * Output: 3
 * Example 3:
 * Input: nums = [1,1]
 * Output: 1
 * Example 4:
 * Input: nums = [1,1,2]
 * Output: 1
 *
 * Constraints:
 *
 * 	2 <= n <= 3 * 10^4
 * 	nums.length == n + 1
 * 	1 <= nums[i] <= n
 * 	All the integers in nums appear only once except for precisely one integer which appears two or more times.
 *
 */

package leetcode

// Floyd's Tortoise and Hare (Cycle Detection)
func findDuplicate(nums []int) int {
	tortoise := nums[0]
	hare := nums[0]
	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}

	ptr1 := nums[0]
	ptr2 := tortoise
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}
	return ptr1
}

// pigeonhole principle
func findDuplicate2(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		count := 0
		for i := range nums {
			if nums[i] <= mid {
				count++
			}
		}
		if count <= mid {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}
