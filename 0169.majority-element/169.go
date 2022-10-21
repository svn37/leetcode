/**
 * Given an array of size n, find the majority element. The majority element is the element that appears more than &lfloor; n/2 &rfloor; times.
 *
 * You may assume that the array is non-empty and the majority element always exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */

package leetcode

import "math/rand"

// Method 1. Boyer-Moore voting algorithm
// You may assume that the array is non-empty
// and the majority element always exists in the array.
func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}

		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// Method 2. Sorting
func qsort(arr []int) {
	if len(arr) < 2 {
		return
	}

	lo, hi := 0, len(arr)-1
	pivot := rand.Intn(len(arr))
	arr[pivot], arr[hi] = arr[hi], arr[pivot]

	for i := range arr {
		if arr[i] < arr[hi] {
			arr[lo], arr[i] = arr[i], arr[lo]
			lo++
		}
	}
	arr[lo], arr[hi] = arr[hi], arr[lo]

	qsort(arr[:lo])
	qsort(arr[lo+1:])
}

func majorityElement2(nums []int) int {
	qsort(nums)
	return nums[(len(nums)-1)/2]
}

// Method 3. Hashmap
func majorityElement3(nums []int) int {
	majority := len(nums) / 2

	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num] > majority {
			return num
		}
	}
	return 0
}
