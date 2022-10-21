/**
 * Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.
 * Examples:
 *
 * s = "leetcode"
 * return 0.
 * s = "loveleetcode"
 * return 2.
 *
 *
 * Note: You may assume the string contains only lowercase English letters.
 *
 */

package leetcode

func firstUniqChar(s string) int {
	chars := [26]int{}
	for i := range s {
		chars[s[i]-'a']++
	}
	for i := range s {
		if chars[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
