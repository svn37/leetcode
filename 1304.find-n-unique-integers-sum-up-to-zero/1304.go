/**
 * Given an integer n, return any array containing n unique integers such that they add up to 0.
 *
 * Example 1:
 *
 * Input: n = 5
 * Output: [-7,-1,1,3,4]
 * Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].
 *
 * Example 2:
 *
 * Input: n = 3
 * Output: [-1,0,1]
 *
 * Example 3:
 *
 * Input: n = 1
 * Output: [0]
 *
 *
 * Constraints:
 *
 * 	1 <= n <= 1000
 *
 */

package leetcode

// fill the array with pairs - positive and negative number
// if n is odd, add 0
func sumZero(n int) []int {
	res := make([]int, 0, n)

	if n%2 == 1 {
		res = append(res, 0)
	}

	for n > 1 {
		res = append(res, n, -n)
		n -= 2
	}

	return res
}

// find a more elegant rule
func sumZero2(n int) []int {
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = i*2 - n + 1
	}
	return res
}
