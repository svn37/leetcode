/**
 *
 * Given a string and a string dictionary, find the longest string in the dictionary that can be formed by deleting some characters of the given string. If there are more than one possible results, return the longest word with the smallest lexicographical order. If there is no possible result, return the empty string.
 *
 * Example 1:<br>
 *
 * Input:
 * s = "abpcplea", d = ["ale","apple","monkey","plea"]
 *
 * Output:
 * "apple"
 *
 *
 *
 *
 * Example 2:<br>
 *
 * Input:
 * s = "abpcplea", d = ["a","b","c"]
 *
 * Output:
 * "a"
 *
 *
 *
 * Note:<br>
 * <ol>
 * All the strings in the input will only contain lower-case letters.
 * The size of the dictionary won't exceed 1,000.
 * The length of all the strings in the input won't exceed 1,000.
 * </ol>
 *
 */

package leetcode

import "strings"

func isSubsequence(s, str string) bool {
	j := 0
	for i := range s {
		if s[i] == str[j] {
			j++
		}

		if j == len(str) {
			return true
		}
	}
	return false
}

func findLongestWord(str string, dict []string) string {
	// sorting is less efficient
	maxLenStr := ""
	for _, s := range dict {
		if (len(s) > len(maxLenStr) || (len(s) == len(maxLenStr) && strings.Compare(s, maxLenStr) == -1)) && isSubsequence(str, s) {
			maxLenStr = s
		}
	}
	return maxLenStr
}
