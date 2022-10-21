/**
 * Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.
 *
 *
 * Note:<br />
 *
 * The length of num is less than 10002 and will be &ge; k.
 * The given num does not contain any leading zero.
 *
 *
 *
 *
 * Example 1:
 *
 * Input: num = "1432219", k = 3
 * Output: "1219"
 * Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
 *
 *
 *
 * Example 2:
 *
 * Input: num = "10200", k = 1
 * Output: "200"
 * Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
 *
 *
 *
 * Example 3:
 *
 * Input: num = "10", k = 2
 * Output: "0"
 * Explanation: Remove all the digits from the number and it is left with nothing which is 0.
 *
 *
 */

package leetcode

import "bytes"

// this solution emulates stack
func removeKdigits(num string, k int) string {
	if k >= len(num) || len(num) == 0 {
		return "0"
	}

	N := []byte(num)

	for i := 1; i < len(N) && k > 0; i++ {
		for j := i - 1; j >= 0 && k > 0; j-- {
			if N[j] == '*' {
				continue
			}

			if N[j] > N[i] {
				N[j] = '*'
				k--
			} else {
				break
			}
		}
	}

	var res bytes.Buffer
	for i := range N {
		if N[i] != '*' && i < len(N)-k && !(res.Len() == 0 && N[i] == '0' && i != len(N)-k-1) {
			res.WriteByte(N[i])
		}
	}
	return res.String()
}
