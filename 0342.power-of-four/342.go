/**
 * Given an integer n, return true if it is a power of four. Otherwise, return false.
 * An integer n is a power of four, if there exists an integer x such that n == 4^x.
 *
 * Example 1:
 * Input: n = 16
 * Output: true
 * Example 2:
 * Input: n = 5
 * Output: false
 * Example 3:
 * Input: n = 1
 * Output: true
 *
 * Constraints:
 *
 * 	-2^31 <= n <= 2^31 - 1
 *
 *
 * Follow up: Could you solve it without loops/recursion?
 */

package leetcode

/*
 * func isPowerOfFour(num int) bool {
 *   if num < 0 {
 *     return false
 *   }
 *   var zeroSeq int
 *   for num&1 == 0 && num != 0 {
 *     zeroSeq++
 *     num >>= 1
 *   }
 *   if num != 1 {
 *     return false
 *   }
 *   return zeroSeq%2 == 0
 * }
 */

// power of 4 numbers have those 3 common features
// First,greater than 0.
// Second,only have one '1' bit in their binary notation,so we use x&(x-1) to delete the lowest '1',
// and if then it becomes 0,it prove that there is only one '1' bit.
// Third,the only '1' bit should be locate at the odd location,for example,16.
// It's binary is 00010000.So we can use '0x55555555' to check if the '1' bit is in the right place
//
//	0x55555555 is a hexametrical numbe (1010101010101010101010101010101 in binary) with a length of 32,
//
// to make sure the 1 locates in the odd location
func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && num&0x55555555 != 0
}
