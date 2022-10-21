/**
 * An array is monotonic if it is either monotone increasing or monotone decreasing.
 *
 * An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].
 *
 * Return true if and only if the given array A is monotonic.
 *
 *
 *
 * <ol>
 * </ol>
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">[1,2,2,3]</span>
 * Output: <span id="example-output-1">true</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">[6,5,4,4]</span>
 * Output: <span id="example-output-2">true</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">[1,3,2]</span>
 * Output: <span id="example-output-3">false</span>
 *
 *
 * <div>
 * Example 4:
 *
 *
 * Input: <span id="example-input-4-1">[1,2,4,5]</span>
 * Output: <span id="example-output-4">true</span>
 *
 *
 * <div>
 * Example 5:
 *
 *
 * Input: <span id="example-input-5-1">[1,1,1]</span>
 * Output: <span id="example-output-5">true</span>
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= A.length <= 50000
 * 	-100000 <= A[i] <= 100000
 * </ol>
 * </div>
 * </div>
 * </div>
 * </div>
 * </div>
 *
 */

package leetcode

func isMonotonic(A []int) bool {
	var up bool // if array is increasing
	n := len(A)

	if A[0] < A[n-1] {
		up = true
	}

	for i := 1; i < n; i++ {
		if A[i-1] != A[i] && A[i]-A[i-1] > 0 != up {
			return false
		}
	}
	return true
}
