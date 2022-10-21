/**
 * Write a function that reverses a string. The input string is given as an array of characters char[].
 *
 * Do not allocate extra space for another array, you must do this by modifying the input array <a href="https://en.wikipedia.org/wiki/In-place_algorithm" target="_blank">in-place</a> with O(1) extra memory.
 *
 * You may assume all the characters consist of <a href="https://en.wikipedia.org/wiki/ASCII#Printable_characters" target="_blank">printable ascii characters</a>.
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">["h","e","l","l","o"]</span>
 * Output: <span id="example-output-1">["o","l","l","e","h"]</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">["H","a","n","n","a","h"]</span>
 * Output: <span id="example-output-2">["h","a","n","n","a","H"]</span>
 *
 * </div>
 * </div>
 *
 */

package leetcode

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
