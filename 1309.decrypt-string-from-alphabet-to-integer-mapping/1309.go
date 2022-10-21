/**
 * Given a string s formed by digits ('0' - '9') and '#' . We want to map s to English lowercase characters as follows:
 *
 * 	Characters ('a' to 'i') are represented by ('1' to '9') respectively.
 * 	Characters ('j' to 'z') are represented by ('10#' to '26#') respectively.
 *
 * Return the string formed after mapping.
 * It's guaranteed that a unique mapping will always exist.
 *
 * Example 1:
 *
 * Input: s = "10#11#12"
 * Output: "jkab"
 * Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".
 *
 * Example 2:
 *
 * Input: s = "1326#"
 * Output: "acz"
 *
 * Example 3:
 *
 * Input: s = "25#"
 * Output: "y"
 *
 * Example 4:
 *
 * Input: s = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
 * Output: "abcdefghijklmnopqrstuvwxyz"
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 1000
 * 	s[i] only contains digits letters ('0'-'9') and '#' letter.
 * 	s will be valid string such that mapping is always possible.
 *
 */

package leetcode

import (
	"bytes"
	"strconv"
)

func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func freqAlphabets(s string) string {
	var b bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		var num int
		if s[i] == '#' {
			num, _ = strconv.Atoi(s[i-2 : i])
			i -= 2
		} else {
			num = int(s[i] - '0')
		}
		b.WriteByte(byte(num + 96))
	}
	return string(reverse(b.Bytes()))
}
