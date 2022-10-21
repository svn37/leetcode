/**
 * Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.
 * Each number in candidates may only be used once in the combination.
 * Note: The solution set must not contain duplicate combinations.
 *
 * Example 1:
 *
 * Input: candidates = [10,1,2,7,6,1,5], target = 8
 * Output:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * Example 2:
 *
 * Input: candidates = [2,5,2,1,2], target = 5
 * Output:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 * Constraints:
 *
 * 	1 <= candidates.length <= 100
 * 	1 <= candidates[i] <= 50
 * 	1 <= target <= 30
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
		// skip duplicates
		// in [1, 2, 2], comb is [1, 2] and we start with [2], so this combination will be included.
		if i > 0 && candidates[i-1] == candidates[i] {
			continue
		}

		newComb := make([]int, len(comb))
		copy(newComb, comb)
		newComb = append(newComb, candidates[i])

		// the difference from the previous problem is i + 1, because we don't need duplicates.
		findCombinations(candidates[i+1:], newComb, target-candidates[i], res)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}

	qsort(candidates)

	findCombinations(candidates, []int{}, target, &res)
	return res
}
