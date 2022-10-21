/**
 * Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
 *
 * 	Integers in each row are sorted from left to right.
 * 	The first integer of each row is greater than the last integer of the previous row.
 *
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/05/mat.jpg" style="width: 322px; height: 242px;" />
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]], target = 3
 * Output: true
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg" style="width: 322px; height: 242px;" />
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]], target = 13
 * Output: false
 *
 * Example 3:
 *
 * Input: matrix = [], target = 0
 * Output: false
 *
 *
 * Constraints:
 *
 * 	m == matrix.length
 * 	n == matrix[i].length
 * 	0 <= m, n <= 100
 * 	-10^4 <= matrix[i][j], target <= 10^4
 *
 */

package leetcode

func bsearch(a []int, target int) bool {
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if a[mid] == target {
			return true
		} else if a[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	lo, hi := 0, len(matrix)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if matrix[mid][0] == target {
			return true
		} else if mid+1 < len(matrix) && matrix[mid][0] < target && matrix[mid+1][0] > target {
			lo = mid
			break
		} else if matrix[mid][0] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return bsearch(matrix[lo], target)
}
