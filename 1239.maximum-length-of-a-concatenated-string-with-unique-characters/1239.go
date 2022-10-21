/**
 * Given an array of strings arr. String s is a concatenation of a sub-sequence of arr which have unique characters.
 * Return the maximum possible length of s.
 *
 * Example 1:
 *
 * Input: arr = ["un","iq","ue"]
 * Output: 4
 * Explanation: All possible concatenations are "","un","iq","ue","uniq" and "ique".
 * Maximum length is 4.
 *
 * Example 2:
 *
 * Input: arr = ["cha","r","act","ers"]
 * Output: 6
 * Explanation: Possible solutions are "chaers" and "acters".
 *
 * Example 3:
 *
 * Input: arr = ["abcdefghijklmnopqrstuvwxyz"]
 * Output: 26
 *
 *
 * Constraints:
 *
 * 	1 <= arr.length <= 16
 * 	1 <= arr[i].length <= 26
 * 	arr[i] contains only lower case English letters.
 *
 */

package leetcode

import "math/bits"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxLength(arr []string) int {
	// combs contains different combinations of strings
	// which contain unique characters, and are candidates for a longer string
	// character counts are represented by bitsets
	combs := []uint32{0}

	res := 0
	for _, str := range arr {
		var count uint32
		var dup uint32 // check for duplicate characters INSIDE one string
		for _, char := range str {
			dup |= count & (1 << (char - 'a'))
			count |= 1 << (char - 'a')
		}
		if dup > 0 {
			continue
		}
		for i := len(combs) - 1; i >= 0; i-- {
			if combs[i]&count > 0 {
				continue
			}
			combs = append(combs, combs[i]|count)
			res = max(res, bits.OnesCount32(combs[i]|count))
		}
	}
	return res
}
