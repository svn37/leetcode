/**
 * Write a function to find the longest common prefix string amongst an array of strings.
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 * Input: strs = ["flower","flow","flight"]
 * Output: "fl"
 *
 * Example 2:
 *
 * Input: strs = ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Constraints:
 *
 * 	0 <= strs.length <= 200
 * 	0 <= strs[i].length <= 200
 * 	strs[i] consists of only lower-case English letters.
 *
 */

package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	word := strs[0]
	for _, str := range strs {
		i := 0
		for i < len(word) && i < len(str) && word[i] == str[i] {
			i++
		}
		word = word[:i]
	}
	return word
}

/*
 * func longestCommonPrefix(strs []string) string {
 *   if len(strs) == 0 {
 *     return ""
 *   }
 *
 *   var curChar byte
 *   curIdx := 0
 *
 *   // simply check character at position curIdx in every string,
 *   // stop if any string reached the end, or the character doesn't match
 *   for {
 *     for i, str := range strs {
 *       if len(str) <= curIdx || (curChar != 0 && curChar != str[curIdx]) {
 *         return strs[i][:curIdx]
 *       } else if curChar == 0 {
 *         curChar = str[curIdx]
 *       }
 *     }
 *     curIdx++
 *     curChar = 0
 *   }
 * }
 */
