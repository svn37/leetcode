/**
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 * Constraints:
 *
 * 	Each string consists only of '0' or '1' characters.
 * 	1 <= a.length, b.length <= 10^4
 * 	Each string is either "0" or doesn't contain any leading zero.
 *
 */

package leetcode

import "bytes"

func addBinary(a string, b string) string {
	var ab bytes.Buffer

	i, j, carry := len(a)-1, len(b)-1, 0
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		ab.WriteByte(byte(sum%2) + '0')
		carry = sum / 2
	}
	if carry != 0 {
		ab.WriteByte(byte(carry) + '0')
	}
	str := ab.Bytes()
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}
