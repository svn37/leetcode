/**
 * Given a non-empty integer array, find the minimum number of moves required to make all array elements equal, where a move is incrementing a selected element by 1 or decrementing a selected element by 1.
 *
 * You may assume the array's length is at most 10,000.
 *
 * Example:
 *
 * Input:
 * [1,2,3]
 *
 * Output:
 * 2
 *
 * Explanation:
 * Only two moves are needed (remember each move increments or decrements one element):
 *
 * [1,2,3]  =>  [2,2,3]  =>  [2,2,2]
 *
 *
 */

package leetcode

import "math/rand"

func quickSelect(nums []int, k int) int {
	left, right := 0, len(nums)-1
	pivot := rand.Intn(len(nums))

	nums[pivot], nums[right] = nums[right], nums[pivot]

	for i := range nums {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	if left == k {
		return nums[left]
	} else if k < left {
		return quickSelect(nums[:left], k)
	} else {
		return quickSelect(nums[left+1:], k-left-1)
	}
}

// This solution relies on the fact that if we increment/decrement each element to the median of
// all the elements, the optimal number of moves is necessary. The median of all elements can be
// found in expected O(n) time using QuickSelect (or deterministic O(n) time using Median of Medians).
func minMoves2(nums []int) int {
	median := quickSelect(nums, len(nums)/2)

	minNum := 0
	for i := range nums {
		if nums[i] > median {
			minNum += nums[i] - median
		} else {
			minNum += median - nums[i]
		}
	}
	return minNum
}
