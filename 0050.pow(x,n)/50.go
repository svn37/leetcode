/**
 * Implement <a href="http://www.cplusplus.com/reference/valarray/pow/" target="_blank">pow(x, n)</a>, which calculates x raised to the power n (i.e. x^<span style="font-size:10.8333px">n</span>).
 *
 * Example 1:
 *
 * Input: x = 2.00000, n = 10
 * Output: 1024.00000
 *
 * Example 2:
 *
 * Input: x = 2.10000, n = 3
 * Output: 9.26100
 *
 * Example 3:
 *
 * Input: x = 2.00000, n = -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 * Constraints:
 *
 * 	-100.0 < x < 100.0
 * 	-2^31 <= n <= 2^31-1
 * 	-10^4 <= x^n <= 10^4
 *
 */

package leetcode

func myPow(x float64, n int) float64 {
	result := 1.0

	absN := n
	if n < 0 {
		absN *= -1
	}

	for absN > 0 {
		if absN&1 == 1 {
			result *= x
		}
		absN >>= 1
		x *= x
	}

	if n < 0 {
		return 1 / result
	}
	return result
}
