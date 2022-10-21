/**
 * Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
 *
 * Example 1:
 *
 * Input: nums = [-1,2,1,-4], target = 1
 * Output: 2
 * Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 * Constraints:
 *
 * 	3 <= nums.length <= 10^3
 * 	-10^3 <= nums[i] <= 10^3
 * 	-10^4 <= target <= 10^4
 *
 */

package leetcode

import "math/rand"

func qsort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivot := rand.Intn(len(arr))

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	qsort(arr[:left])
	qsort(arr[left+1:])
}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	qsort(nums)
	res := nums[0] + nums[1] + nums[len(nums)-1]

	n := len(nums)
	// similar to the previous problem
	// no check for nums[i] <= 0
	for i := 0; i < n-2; i++ {
		first := nums[i]
		start := i + 1
		end := n - 1
		for start < end {
			sum := first + nums[start] + nums[end]
			if sum < target {
				start++
			} else if sum > target {
				end--
			} else {
				return target
			}

			if abs(sum-target) < abs(res-target) {
				res = sum
			}
		}
	}

	return res
}
