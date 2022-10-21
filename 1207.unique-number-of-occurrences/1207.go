/**
 * Given an array of integers arr, write a function that returns true if and only if the number of occurrences of each value in the array is unique.
 *
 * Example 1:
 *
 * Input: arr = [1,2,2,1,1,3]
 * Output: true
 * Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
 * Example 2:
 *
 * Input: arr = [1,2]
 * Output: false
 *
 * Example 3:
 *
 * Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
 * Output: true
 *
 *
 * Constraints:
 *
 * 	1 <= arr.length <= 1000
 * 	-1000 <= arr[i] <= 1000
 *
 */

package leetcode

import "math/rand"

func qsort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivot := rand.Intn(len(a))

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
}

func uniqueOccurrences(arr []int) bool {
	qsort(arr)

	set := map[int]struct{}{}
	count := 1
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i-1] != arr[i] {
			if _, ok := set[count]; ok {
				return false
			}
			set[count] = struct{}{}
			count = 1
		} else {
			count++
		}
	}
	return true
}
