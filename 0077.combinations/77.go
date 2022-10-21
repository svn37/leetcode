/**
 * Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
 * You may return the answer in any order.
 *
 * Example 1:
 *
 * Input: n = 4, k = 2
 * Output:
 * [
 *   [2,4],
 *   [3,4],
 *   [2,3],
 *   [1,2],
 *   [1,3],
 *   [1,4],
 * ]
 *
 * Example 2:
 *
 * Input: n = 1, k = 1
 * Output: [[1]]
 *
 *
 * Constraints:
 *
 * 	1 <= n <= 20
 * 	1 <= k <= n
 *
 */

package leetcode

func findCombinations(begin, n, k int, comb []int, res *[][]int) {
	if len(comb) == k {
		temp := append([]int{}, comb...)
		*res = append(*res, temp)
		return
	}

	if begin > n {
		return
	}

	for i := begin; i <= n; i++ {
		comb = append(comb, i)
		findCombinations(i+1, n, k, comb, res)
		comb = comb[:len(comb)-1]
	}
}

func combine(n int, k int) [][]int {
	res := [][]int{}

	findCombinations(1, n, k, []int{}, &res)
	return res
}
