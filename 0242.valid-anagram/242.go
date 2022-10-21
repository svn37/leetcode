/**
 * Given two strings s and t , write a function to determine if t is an anagram of s.
 *
 * Example 1:
 *
 *
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "rat", t = "car"
 * Output: false
 *
 *
 * Note:<br />
 * You may assume the string contains only lowercase alphabets.
 *
 * Follow up:<br />
 * What if the inputs contain unicode characters? How would you adapt your solution to such case?
 *
 */

package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphabet := make([]int, 26)

	for i := range s {
		alphabet[int(s[i]-'a')]++
		alphabet[int(t[i]-'a')]--
	}

	for i := range alphabet {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}
