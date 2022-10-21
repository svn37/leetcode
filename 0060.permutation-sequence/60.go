/**
 * The set [1, 2, 3, ..., n] contains a total of n! unique permutations.
 * By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
 * <ol>
 * 	"123"
 * 	"132"
 * 	"213"
 * 	"231"
 * 	"312"
 * 	"321"
 * </ol>
 * Given n and k, return the k^th permutation sequence.
 *
 * Example 1:
 * Input: n = 3, k = 3
 * Output: "213"
 * Example 2:
 * Input: n = 4, k = 9
 * Output: "2314"
 * Example 3:
 * Input: n = 3, k = 1
 * Output: "123"
 *
 * Constraints:
 *
 * 	1 <= n <= 9
 * 	1 <= k <= n!
 *
 */

package leetcode

import "bytes"

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	// calculate factorials
	factorial[0] = 1
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
		factorial[i] = sum
	}

	numbers := make([]int, 0)
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}

	k--
	var str bytes.Buffer

	for i := 1; i <= n; i++ {
		idx := k / factorial[n-i]
		k -= idx * factorial[n-i]

		str.WriteByte(byte(numbers[idx] + '0'))
		numbers = append(numbers[:idx], numbers[idx+1:]...)
	}

	return str.String()
}
