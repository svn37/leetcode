/**
 * Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: [1,2,2]
 * Output:
 * [
 *   [2],
 *   [1],
 *   [1,2,2],
 *   [2,2],
 *   [1,2],
 *   []
 * ]
 *
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

	arr[right], arr[pivot] = arr[pivot], arr[right]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	qsort(arr[:left])
	qsort(arr[left+1:])
}

func subsetsWithDup(nums []int) [][]int {
	qsort(nums)

	res := make([][]int, 0)
	res = append(res, []int{})

	start := len(res)
	for i := range nums {
		if i > 0 && nums[i-1] != nums[i] {
			start = len(res)
		}
		newSet := make([][]int, 0)

		for j := len(res) - start; j < len(res); j++ {
			subset := make([]int, len(res[j]))
			copy(subset, res[j])

			subset = append(subset, nums[i])
			newSet = append(newSet, subset)
		}
		start = len(newSet)

		res = append(res, newSet...)
	}

	return res
}
