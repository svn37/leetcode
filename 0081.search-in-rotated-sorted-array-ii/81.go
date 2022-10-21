/**
 * Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
 *
 * (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
 *
 * You are given a target value to search. If found in the array return true, otherwise return false.
 *
 * Example 1:
 *
 *
 * Input: nums = [2,5,6,0,0,1,2], target = 0
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,5,6,0,0,1,2], target = 3
 * Output: false
 *
 * Follow up:
 *
 *
 * 	This is a follow up problem to <a href="/problems/search-in-rotated-sorted-array/description/">Search in Rotated Sorted Array</a>, where nums may contain duplicates.
 * 	Would this affect the run-time complexity? How and why?
 *
 *
 */

package leetcode

func search(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := (lo + hi) / 2

		if nums[mid] == target {
			return true
		}

		if nums[lo] == nums[mid] && nums[hi] == nums[mid] {
			// duplicate case - [1,3,1,1,1], target 3
			lo++
			hi--
		} else if nums[lo] <= nums[mid] {
			// left part is sorted
			if nums[lo] <= target && nums[mid] > target {
				hi = mid - 1
			} else {
				// target is still in the rotated part
				lo = mid + 1
			}
		} else {
			// right part is sorted
			if nums[hi] >= target && nums[mid] < target {
				lo = mid + 1
			} else {
				// target is still in the rotated part
				hi = mid - 1
			}
		}
	}

	return false
}
