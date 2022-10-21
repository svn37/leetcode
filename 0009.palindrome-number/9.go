/**
 * Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
 * Follow up: Could you solve it without converting the integer to a string?
 *
 * Example 1:
 *
 * Input: x = 121
 * Output: true
 *
 * Example 2:
 *
 * Input: x = -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
 *
 * Example 3:
 *
 * Input: x = 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 *
 * Example 4:
 *
 * Input: x = -101
 * Output: false
 *
 *
 * Constraints:
 *
 * 	-2^31 <= x <= 2^31 - 1
 *
 */

package leetcode

// the same problem as reverse integer, except we don't need to reverse the whole number.
// half the number is enough, which also avoids overflows if num > math.MaxInt64
func isPalindrome(num int) bool {
	// if the number is negative, or its last digit is 0, return immediately
	if num < 0 || (num != 0 && num%10 == 0) {
		return false
	}

	var reverse int
	for num > reverse {
		reverse *= 10
		reverse += num % 10
		num /= 10
	}

	// if digits length is odd, need to divide by 10
	return num == reverse || num == reverse/10
}
