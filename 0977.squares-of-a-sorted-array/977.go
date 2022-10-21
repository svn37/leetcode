/**
 * Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">[-4,-1,0,3,10]</span>
 * Output: <span id="example-output-1">[0,1,9,16,100]</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">[-7,-3,2,3,11]</span>
 * Output: <span id="example-output-2">[4,9,9,49,121]</span>
 *
 *
 *
 *
 * <span>Note:</span>
 *
 * <ol>
 * 	<span>1 <= A.length <= 10000</span>
 * 	-10000 <= A[i] <= 10000
 * 	A is sorted in non-decreasing order.
 * </ol>
 * </div>
 * </div>
 */

package leetcode

func sortedSquares(A []int) []int {
	right := 0
	for right = range A {
		if A[right] >= 0 {
			break
		}
	}
	res := make([]int, 0, len(A))

	left := right - 1
	for left >= 0 && right < len(A) {
		if -A[left] < A[right] {
			res = append(res, A[left]*A[left])
			left--
		} else {
			res = append(res, A[right]*A[right])
			right++
		}
	}

	for left >= 0 {
		res = append(res, A[left]*A[left])
		left--
	}
	for right < len(A) {
		res = append(res, A[right]*A[right])
		right++
	}
	return res
}
