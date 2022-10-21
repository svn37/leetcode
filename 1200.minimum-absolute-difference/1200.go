/**
 * Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements.
 * Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows
 *
 * 	a, b are from arr
 * 	a < b
 * 	b - a equals to the minimum absolute difference of any two elements in arr
 *
 *
 * Example 1:
 *
 * Input: arr = [4,2,1,3]
 * Output: [[1,2],[2,3],[3,4]]
 * Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.
 * Example 2:
 *
 * Input: arr = [1,3,6,10,15]
 * Output: [[1,3]]
 *
 * Example 3:
 *
 * Input: arr = [3,8,-10,23,19,-4,-14,27]
 * Output: [[-14,-10],[19,23],[23,27]]
 *
 *
 * Constraints:
 *
 * 	2 <= arr.length <= 10^5
 * 	-10^6 <= arr[i] <= 10^6
 *
 */

package leetcode

import (
	"math"
	"math/rand"
)

func qsort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivot := rand.Intn(len(a))

	a[right], a[pivot] = a[pivot], a[right]

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	qsort(a[:left])
	qsort(a[left+1:])
}

func minimumAbsDifference(arr []int) [][]int {
	qsort(arr)

	minDiff := math.MaxInt64
	res := [][]int{}
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < minDiff {
			minDiff = arr[i] - arr[i-1]
			res = res[:0]
			res = append(res, []int{arr[i-1], arr[i]})
		} else if arr[i]-arr[i-1] == minDiff {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
