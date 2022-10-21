/**
 * Given a non negative integer number num. For every numbers i in the range 0 &le; i &le; num calculate the number of 1's in their binary representation and return them as an array.
 *
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">2</span>
 * Output: <span id="example-output-1">[0,1,1]</span>
 *
 * Example 2:
 *
 *
 * Input: <span id="example-input-1-1">5</span>
 * Output: [0,1,1,2,1,2]
 *
 *
 * Follow up:
 *
 *
 * 	It is very easy to come up with a solution with run time O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a single pass?
 * 	Space complexity should be O(n).
 * 	Can you do it like a boss? Do it without using any builtin function like __builtin_popcount in c++ or in any other language.
 *
 */

package leetcode

func countBits(num int) []int {
	bits := make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}
