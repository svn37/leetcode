/**
 * Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.
 *
 * Example 1:
 *
 * Input: nums = [1,1,2]
 * Output:
 * [[1,1,2],
 *  [1,2,1],
 *  [2,1,1]]
 *
 * Example 2:
 *
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 8
 * 	-10 <= nums[i] <= 10
 *
 */

package leetcode

func backtrack(nums []int, begin int, res *[][]int) {
	if begin == len(nums) {
		permutation := make([]int, len(nums))
		copy(permutation, nums)
		*res = append(*res, permutation)
		return
	}

	visited := make(map[int]struct{})

	for i := begin; i < len(nums); i++ {
		if _, ok := visited[nums[i]]; ok {
			continue
		}
		nums[begin], nums[i] = nums[i], nums[begin]
		backtrack(nums, begin+1, res)
		nums[begin], nums[i] = nums[i], nums[begin]

		visited[nums[i]] = struct{}{}
	}
}

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	backtrack(nums, 0, &res)
	return res
}
