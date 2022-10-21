/**
 * Given an array A of positive lengths, return the largest perimeter of a triangle with non-zero area, formed from 3 of these lengths.
 *
 * If it is impossible to form any triangle of non-zero area, return 0.
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
 * Input: <span id="example-input-1-1">[2,1,2]</span>
 * Output: <span id="example-output-1">5</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">[1,2,1]</span>
 * Output: <span id="example-output-2">0</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">[3,2,3,4]</span>
 * Output: <span id="example-output-3">10</span>
 *
 *
 * <div>
 * Example 4:
 *
 *
 * Input: <span id="example-input-4-1">[3,6,2,3]</span>
 * Output: <span id="example-output-4">8</span>
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	3 <= A.length <= 10000
 * 	1 <= A[i] <= 10^6
 * </ol>
 * </div>
 * </div>
 * </div>
 * </div>
 */

package leetcode

import "math/rand"

func qsort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivot := rand.Intn(len(arr))

	arr[right], arr[pivot] = arr[pivot], arr[right]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	qsort(arr[:left])
	qsort(arr[left+1:])
}

func largestPerimeter(A []int) int {
	qsort(A)
	for i := len(A) - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}
