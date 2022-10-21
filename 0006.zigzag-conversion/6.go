/**
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 * Write the code that will take a string and make this conversion given a number of rows:
 *
 * string convert(string s, int numRows);
 *
 *
 * Example 1:
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 *
 * Example 2:
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 * Example 3:
 *
 * Input: s = "A", numRows = 1
 * Output: "A"
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 1000
 * 	s consists of English letters (lower-case and upper-case), ',' and '.'.
 * 	1 <= numRows <= 1000
 *
 */

package leetcode

import "bytes"

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	// gap is always 2 smaller than the number of rows
	gap := numRows - 2

	var b bytes.Buffer
	// t -- no sense naming because it's kinda curRow
	// but we're finding the correct idx in the string, really
	t := 0
	for t != numRows {
		// len(s)+t*2 -- there might be characters after the last vertical column
		for i := t; i < len(s)+t*2; i += numRows + gap {
			// print the character before vertical column
			// i != t -- for first vertical column there is no such character
			// t > 0 -- for first row there are no such characters
			// t < numRows-1 -- for last row there are no such characters
			if i != t && t > 0 && t < numRows-1 {
				b.WriteByte(s[i-t*2])
			}
			// print the character in the vertical column
			if i < len(s) {
				b.WriteByte(s[i])
			}
		}
		t++
	}

	return b.String()
}
