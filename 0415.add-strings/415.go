/**
 * Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.
 *
 * Note:
 * <ol>
 * The length of both num1 and num2 is < 5100.
 * Both num1 and num2 contains only digits 0-9.
 * Both num1 and num2 does not contain any leading zero.
 * You must not use any built-in BigInteger library or convert the inputs to integer directly.
 * </ol>
 *
 */

package leetcode

import "bytes"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addStrings(num1 string, num2 string) string {
	m := len(num1)
	n := len(num2)

	result := make([]int, max(m, n)+1)

	var carry int
	for i, j := m-1, n-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}

		if j >= 0 {
			n2 = int(num2[j] - '0')
		}

		sum := n1 + n2 + carry
		carry = sum / 10
		idx := max(i, j) + 1
		result[idx] = sum % 10
	}

	result[0] = carry

	var b bytes.Buffer

	for i := range result {
		if b.Len() == 0 && result[i] == 0 && i != len(result)-1 {
			continue
		}
		b.WriteByte(byte(result[i] + int('0')))
	}

	return b.String()
}
