/**
 * Given a list of words, each word consists of English lowercase letters.
 *
 * Let's say word1 is a predecessor of word2 if and only if we can add exactly one letter anywhere in word1 to make it equal to word2.  For example, "abc" is a predecessor of "abac".
 *
 * A word chain is a sequence of words [word_1, word_2, ..., word_k] with k >= 1, where word_1 is a predecessor of word_2, word_2 is a predecessor of word_3, and so on.
 *
 * Return the longest possible length of a word chain with words chosen from the given list of words.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">["a","b","ba","bca","bda","bdca"]</span>
 * Output: <span id="example-output-1">4
 * Explanation: one of </span>the longest word chain is "a","ba","bda","bdca".
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= words.length <= 1000
 * 	1 <= words[i].length <= 16
 * 	words[i] only consists of English lowercase letters.
 * </ol>
 *
 * <div>
 *
 * </div>
 */

package leetcode

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestStrChain(words []string) int {
	m := make(map[string]int)
	// sort.Strings doesn't sort by length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	res := 0
	for _, word := range words {
		best := 0
		for i := 0; i < len(word); i++ {
			prev := word[:i] + word[i+1:]
			best = max(best, m[prev]+1)
		}
		m[word] = best
		res = max(res, best)
	}
	return res
}
