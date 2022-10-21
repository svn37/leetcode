/**
 * You have a list of words and a pattern, and you want to know which words in words matches the pattern.
 *
 * A word matches the pattern if there exists a permutation of letters p so that after replacing every letter x in the pattern with p(x), we get the desired word.
 *
 * (Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.)
 *
 * Return a list of the words in words that match the given pattern.
 *
 * You may return the answer in any order.
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: words = <span id="example-input-1-1">["abc","deq","mee","aqq","dkd","ccc"]</span>, pattern = <span id="example-input-1-2">"abb"</span>
 * Output: <span id="example-output-1">["mee","aqq"]</span>
 * <span>Explanation: </span>"mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
 * "ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation,
 * since a and b map to the same letter.
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= words.length <= 50
 * 	1 <= pattern.length = words[i].length <= 20
 *
 * </div>
 *
 */

package leetcode

func comparePatterns(word, pattern string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := range word {
		char1, ok1 := m1[word[i]]
		char2, ok2 := m2[pattern[i]]
		if (ok1 && char1 != pattern[i]) || (ok2 && char2 != word[i]) {
			return false
		}
		m1[word[i]] = pattern[i]
		m2[pattern[i]] = word[i]
	}
	return true
}

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	for _, word := range words {
		if comparePatterns(word, pattern) {
			res = append(res, word)
		}
	}
	return res
}
