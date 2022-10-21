/**
 * Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
 * Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
 *
 * Example 1:
 * Input: num1 = "2", num2 = "3"
 * Output: "6"
 * Example 2:
 * Input: num1 = "123", num2 = "456"
 * Output: "56088"
 *
 * Constraints:
 *
 * 	1 <= num1.length, num2.length <= 200
 * 	num1 and num2 consist of digits only.
 * 	Both num1 and num2 do not contain any leading zero, except the number 0 itself.
 *
 */

package leetcode

import "bytes"

// https://leetcode.com/problems/multiply-strings/discuss/17605/Easiest-JAVA-Solution-with-Graph-Explanation
func multiply(num1 string, num2 string) string {
	m := len(num1)
	n := len(num2)

	memo := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1 := i + j
			p2 := i + j + 1
			sum := mul + memo[p2]

			memo[p1] += sum / 10
			memo[p2] = sum % 10
		}
	}

	var b bytes.Buffer

	for _, digit := range memo {
		if b.Len() == 0 && digit == 0 {
			continue
		}
		b.WriteByte(byte(digit) + '0')
	}

	if b.Len() == 0 {
		return "0"
	}

	return b.String()
}
