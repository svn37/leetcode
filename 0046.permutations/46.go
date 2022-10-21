/**
 * Given a collection of distinct integers, return all possible permutations.
 *
 * Example:
 *
 *
 * Input: [1,2,3]
 * Output:
 * [
 *   [1,2,3],
 *   [1,3,2],
 *   [2,1,3],
 *   [2,3,1],
 *   [3,1,2],
 *   [3,2,1]
 * ]
 *
 *
 */

package leetcode

// Method 1. Backtracking
func permuteHelper(nums []int, begin int, res *[][]int) {
	if begin >= len(nums) {
		permutation := make([]int, len(nums))
		copy(permutation, nums)
		*res = append(*res, permutation)
		return
	}

	for i := begin; i < len(nums); i++ {
		nums[begin], nums[i] = nums[i], nums[begin]
		permuteHelper(nums, begin+1, res)
		// swap back
		nums[begin], nums[i] = nums[i], nums[begin]
	}
}

func permute(nums []int) [][]int {
	res := [][]int{}
	permuteHelper(nums, 0, &res)
	return res
}

// Method 2. Heap's algorithm
func permuteHelper2(nums []int, size int, res *[][]int) {
	if size == 1 {
		permutation := make([]int, len(nums))
		copy(permutation, nums)
		*res = append(*res, permutation)
		return
	}

	for i := 0; i < size; i++ {
		permuteHelper2(nums, size-1, res)

		// heap's algorithm
		if len(nums[:size])%2 == 1 {
			nums[0], nums[size-1] = nums[size-1], nums[0]
		} else {
			nums[i], nums[size-1] = nums[size-1], nums[i]
		}
	}
}

func permute2(nums []int) [][]int {
	res := [][]int{}
	permuteHelper2(nums, len(nums), &res)
	return res
}
