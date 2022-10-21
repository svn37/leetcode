/**
 * Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
 * Notice that the solution set must not contain duplicate triplets.
 *
 * Example 1:
 * Input: nums = [-1,0,1,2,-1,-4]
 * Output: [[-1,-1,2],[-1,0,1]]
 * Example 2:
 * Input: nums = []
 * Output: []
 * Example 3:
 * Input: nums = [0]
 * Output: []
 *
 * Constraints:
 *
 * 	0 <= nums.length <= 3000
 * 	-10^5 <= nums[i] <= 10^5
 *
 */

package leetcode

import (
	"math/rand"
)

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

func threeSum(nums []int) [][]int {
	qsort(nums)

	res := [][]int{}
	// if nums[i] > 0, then sum can't be zero, stop
	for i := 0; i < len(nums)-1 && nums[i] <= 0; i++ {
		// nums[i] is the first element
		// two pointers search from both ends for the second two elements
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo, hi := i+1, len(nums)-1
			for lo < hi {
				candidate := nums[i] + nums[lo] + nums[hi]
				if candidate == 0 {
					res = append(res, []int{nums[i], nums[lo], nums[hi]})
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					hi--
					lo++
				} else if candidate > 0 {
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					hi--
				} else {
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					lo++
				}
			}
		}
	}
	return res
}
