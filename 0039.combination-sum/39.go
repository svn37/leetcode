/**
 * Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.
 * The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.
 * It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
 *
 * Example 1:
 *
 * Input: candidates = [2,3,6,7], target = 7
 * Output: [[2,2,3],[7]]
 * Explanation:
 * 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
 * 7 is a candidate, and 7 = 7.
 * These are the only two combinations.
 *
 * Example 2:
 *
 * Input: candidates = [2,3,5], target = 8
 * Output: [[2,2,2,2],[2,3,3],[3,5]]
 *
 * Example 3:
 *
 * Input: candidates = [2], target = 1
 * Output: []
 *
 * Example 4:
 *
 * Input: candidates = [1], target = 1
 * Output: [[1]]
 *
 * Example 5:
 *
 * Input: candidates = [1], target = 2
 * Output: [[1,1]]
 *
 *
 * Constraints:
 *
 * 	1 <= candidates.length <= 30
 * 	1 <= candidates[i] <= 200
 * 	All elements of candidates are distinct.
 * 	1 <= target <= 500
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

func findCombinations(candidates []int, comb []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, comb)
		return
	}

	for i := 0; i < len(candidates) && target >= candidates[i]; i++ {
		newComb := make([]int, len(comb))
		copy(newComb, comb)
		newComb = append(newComb, candidates[i])

		// the trick here is we pass the array with the same index, because we need duplicates
		findCombinations(candidates[i:], newComb, target-candidates[i], res)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}

	qsort(candidates)

	findCombinations(candidates, []int{}, target, &res)
	return res
}
