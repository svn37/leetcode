/**
 * Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
 * Return the quotient after dividing dividend by divisor.
 * The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.
 * Note:
 *
 * 	Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [-2^31,  2^31 - 1]. For this problem, assume that your function returns 2^31 - 1 when the division result overflows.
 *
 *
 * Example 1:
 *
 * Input: dividend = 10, divisor = 3
 * Output: 3
 * Explanation: 10/3 = truncate(3.33333..) = 3.
 *
 * Example 2:
 *
 * Input: dividend = 7, divisor = -3
 * Output: -2
 * Explanation: 7/-3 = truncate(-2.33333..) = -2.
 *
 * Example 3:
 *
 * Input: dividend = 0, divisor = 1
 * Output: 0
 *
 * Example 4:
 *
 * Input: dividend = 1, divisor = 1
 * Output: 1
 *
 *
 * Constraints:
 *
 * 	-2^31 <= dividend, divisor <= 2^31 - 1
 * 	divisor != 0
 *
 */

package leetcode

import "math"

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

// without multiplication, division and mod operators
// we can use << to multiply the number by two

// the quotient (result) of a division is just the number of times
// that we can subtract the divisor from the dividend without making it negative

// 29           / 5
// 9 (29-5*2*2) / 5
// 4 (9-5*1)    / 5 -> stop (4 < 5)
// result is 2*2 (2nd iteration) + 1 (3rd iteration) = 5

// the outer loop reduces n by at least half each iteration. So, it has O(log N) iterations.
// the inner loop has at most log N iterations.
// the overall complexity is O((log N)^2)
func divide(dividend int, divisor int) int {
	// cannot divide by zero
	// -2147483648 and -1 case, the only one when the division result overflows
	if divisor == 0 || (dividend == math.MinInt32 && divisor == -1) {
		return math.MaxInt32
	}

	// save the sign of the result
	// we'll do binary shifts on |n|
	sign := 1
	if (dividend < 0) != (divisor < 0) {
		sign = -1
	}

	dividend = abs(dividend)
	divisor = abs(divisor)

	res := 0

	for dividend >= divisor {
		temp := divisor
		mult := 1
		for dividend >= temp<<1 {
			// multiply by two
			temp <<= 1
			mult <<= 1
		}
		dividend -= temp
		// the result of division is at least mult
		res += mult
	}

	return res * sign
}
