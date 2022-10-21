/**
 * Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
 *
 * For example, given the following triangle
 *
 *
 * [
 *      [2],
 *     [3,4],
 *    [6,5,7],
 *   [4,1,8,3]
 * ]
 *
 *
 * The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 *
 * Note:
 *
 * Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
 *
 */

package leetcode

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// more verbose solution
func minimumTotal(triangle [][]int) int {
	memo := make([]int, len(triangle))
	minVal := math.MaxInt64

	for i := range triangle {
		var prev int
		for j := range triangle[i] {
			temp := memo[j]
			if j > 0 && j < len(triangle[i])-1 {
				memo[j] = min(triangle[i][j]+prev, triangle[i][j]+memo[j])
			} else if j == 0 {
				memo[j] = triangle[i][j] + memo[j]
			} else {
				memo[j] = triangle[i][j] + prev
			}
			prev = temp

			// only calculate minVal when it is the last row
			if i == len(triangle)-1 {
				minVal = min(memo[j], minVal)
			}
		}
	}

	return minVal
}

// more concise solution
// bottom-up approach
func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	dp := triangle[n-1]

	for layer := n - 2; layer >= 0; layer-- {
		for i := 0; i <= layer; i++ {
			dp[i] = min(dp[i], dp[i+1]) + triangle[layer][i]
		}
	}
	return dp[0]
}
