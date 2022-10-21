/**
 * Given an integer rowIndex, return the rowIndex^th row of the Pascal's triangle.
 * Notice that the row index starts from 0.
 * <img alt="" src="https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif" /><br />
 * <small>In Pascal's triangle, each number is the sum of the two numbers directly above it.</small>
 * Follow up:
 * Could you optimize your algorithm to use only O(k) extra space?
 *
 * Example 1:
 * Input: rowIndex = 3
 * Output: [1,3,3,1]
 * Example 2:
 * Input: rowIndex = 0
 * Output: [1]
 * Example 3:
 * Input: rowIndex = 1
 * Output: [1,1]
 *
 * Constraints:
 *
 * 	0 <= rowIndex <= 40
 *
 */

package leetcode

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i < len(row); i++ {
		for j := i; j >= 1; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}
