/**
 * You are given an array of strings words and a string chars.
 *
 * A string is good if it can be formed by characters from chars (each character can only be used once).
 *
 * Return the sum of lengths of all good strings in words.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: words = <span id="example-input-1-1">["cat","bt","hat","tree"]</span>, chars = <span id="example-input-1-2">"atach"</span>
 * Output: <span id="example-output-1">6</span>
 * Explanation:
 * The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.
 *
 *
 * Example 2:
 *
 *
 * Input: words = <span id="example-input-2-1">["hello","world","leetcode"]</span>, chars = <span id="example-input-2-2">"welldonehoneyr"</span>
 * Output: <span id="example-output-2">10</span>
 * Explanation:
 * The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.
 *
 *
 *
 *
 * <span>Note:</span>
 *
 * <ol>
 * 	1 <= words.length <= 1000
 * 	1 <= words[i].length, chars.length <= 100
 * 	All strings contain lowercase English letters only.
 * </ol>
 */

package leetcode

func countWord(word string, count []int) int {
	for _, char := range word {
		count[char-'a']--
		if count[char-'a'] < 0 {
			return 0
		}
	}
	return len(word)
}

func countCharacters(words []string, chars string) int {
	// count of each character in chars
	count := make([]int, 26)
	for _, char := range chars {
		count[char-'a']++
	}

	res := 0
	temp := make([]int, 26)
	for i := range words {
		copy(temp, count)
		res += countWord(words[i], temp)
	}
	return res
}
