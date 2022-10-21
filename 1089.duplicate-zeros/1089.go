/**
 * Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.
 *
 * Note that elements beyond the length of the original array are not written.
 *
 * Do the above modifications to the input array in place, do not return anything from your function.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">[1,0,2,3,0,4,5,0]</span>
 * Output: null
 * Explanation: After calling your function, the input array is modified to: <span id="example-output-1">[1,0,0,2,3,0,0,4]</span>
 *
 *
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">[1,2,3]</span>
 * Output: null
 * Explanation: After calling your function, the input array is modified to: <span id="example-output-2">[1,2,3]</span>
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= arr.length <= 10000
 * 	0 <= arr[i] <= 9
 * </ol>
 */

package leetcode

// Method 1. With a copy (extra space)
func duplicateZeros(arr []int) {
	arrcopy := make([]int, len(arr))
	copy(arrcopy, arr)
	var j int
	for i := range arrcopy {
		if j >= len(arr) {
			break
		}
		arr[j] = arrcopy[i]

		if arrcopy[i] == 0 {
			j++
			if j < len(arr) {
				arr[j] = 0
			}
		}
		j++
	}
}

// Method 2. By inserting zeroes (extra time)
func duplicateZeros2(arr []int) []int {
	countZero := 0
	for i := range arr {
		if arr[i] == 0 {
			countZero++
		}
	}
	newLen := len(arr) + countZero
	for i, j := len(arr)-1, newLen-1; i < j; i, j = i-1, j-1 {
		if j < len(arr) {
			arr[j] = arr[i]
		}
		if arr[i] == 0 {
			j--
			if j < len(arr) {
				arr[j] = arr[i]
			}
		}
	}
	return arr
}
