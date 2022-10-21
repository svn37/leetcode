/**
 * Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.
 *
 * After doing so, return the array.
 *
 *
 * Example 1:
 * Input: arr = [17,18,5,4,6,1]
 * Output: [18,6,6,6,1,-1]
 *
 *
 * Constraints:
 *
 *
 * 	1 <= arr.length <= 10^4
 * 	1 <= arr[i] <= 10^5
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func replaceElements(arr []int) []int {
	greatest := -1
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], greatest = greatest, max(greatest, arr[i])
	}
	return arr
}
