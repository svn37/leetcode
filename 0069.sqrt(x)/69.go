/**
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 *              the decimal part is truncated, 2 is returned.
 *
 *
 */

package leetcode

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	lo, hi := 0, x
	for hi-lo > 1 {
		middle := (hi - lo) / 2
		guess := lo + middle
		if guess*guess > x {
			hi -= middle
		} else if guess*guess < x {
			lo += middle
		} else {
			return guess
		}
	}
	return lo
}
