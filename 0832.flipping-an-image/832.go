/**
 * Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.
 *
 * To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].
 *
 * To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].
 *
 * Example 1:
 *
 *
 * Input: [[1,1,0],[1,0,1],[0,0,0]]
 * Output: [[1,0,0],[0,1,0],[1,1,1]]
 * Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
 * Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
 * Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
 * Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
 * Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
 *
 *
 * Notes:
 *
 *
 * 	1 <= A.length = A[0].length <= 20
 * 	0 <= A[i][j]<font face="sans-serif, Arial, Verdana, Trebuchet MS"> <= </font>1
 *
 *
 */

package leetcode

// Compare the i th and n - i - 1 th in a row.
// The "trick" is that if the values are not the same,
// but you swap and flip, nothing will change.
// So if they are same, we toggle both, otherwise we do nothing.
func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		for j, k := 0, len(A[i])-1; j <= k; j, k = j+1, k-1 {
			if j == k || A[i][j] == A[i][k] {
				newVal := A[i][j] ^ 1
				A[i][j], A[i][k] = newVal, newVal
			}
		}
	}
	return A
}
