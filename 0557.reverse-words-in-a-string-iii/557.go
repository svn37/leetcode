/**
 * Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.
 *
 * Example 1:<br />
 *
 * Input: "Let's take LeetCode contest"
 * Output: "s'teL ekat edoCteeL tsetnoc"
 *
 *
 *
 * Note:
 * In the string, each word is separated by single space and there will not be any extra space in the string.
 *
 */

package leetcode

func reverseWord(str []byte) {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}

func reverseWords(s string) string {
	S := []byte(s)
	prev := 0
	for i := 0; i <= len(S); i++ {
		if i == len(S) || S[i] == ' ' {
			reverseWord(S[prev:i])
			prev = i + 1
		}
	}
	return string(S)
}
