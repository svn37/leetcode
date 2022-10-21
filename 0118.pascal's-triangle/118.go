/**
 * Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
 *
 * <img alt="" src="https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif" style="height:240px; width:260px" /><br />
 * <small>In Pascal's triangle, each number is the sum of the two numbers directly above it.</small>
 *
 * Example:
 *
 *
 * Input: 5
 * Output:
 * [
 *      [1],
 *     [1,1],
 *    [1,2,1],
 *   [1,3,3,1],
 *  [1,4,6,4,1]
 * ]
 *
 *
 */

package leetcode

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	triangle := make([][]int, 1)
	triangle[0] = []int{1}

	for i := 1; i < numRows; i++ {
		triangle = append(triangle, []int{})
		triangle[i] = append(triangle[i], 1)
		for j := 1; j < len(triangle[i-1]); j++ {
			triangle[i] = append(triangle[i], triangle[i-1][j]+triangle[i-1][j-1])
		}
		triangle[i] = append(triangle[i], 1)
	}
	return triangle
}
