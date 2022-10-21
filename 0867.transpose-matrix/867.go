/**
 * Given a matrix A, return the transpose of A.
 *
 * The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.
 *
 * <br>
 * <img src="https://assets.leetcode.com/uploads/2019/10/20/hint_transpose.png" width="700"/>
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">[[1,2,3],[4,5,6],[7,8,9]]</span>
 * Output: <span id="example-output-1">[[1,4,7],[2,5,8],[3,6,9]]</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">[[1,2,3],[4,5,6]]</span>
 * Output: <span id="example-output-2">[[1,4],[2,5],[3,6]]</span>
 *
 *
 *
 *
 * <span>Note:</span>
 *
 * <ol>
 * 	<span>1 <= A.length <= 1000</span>
 * 	<span>1 <= A[0].length <= 1000</span>
 * </ol>
 * </div>
 * </div>
 */

package leetcode

func transpose(A [][]int) [][]int {
	m := len(A)
	n := len(A[0])

	B := make([][]int, n)
	for i := range B {
		B[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			B[i][j] = A[j][i]
		}
	}

	return B
}
