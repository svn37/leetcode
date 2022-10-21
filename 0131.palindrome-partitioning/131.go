/**
 * Given a string s, partition s such that every substring of the partition is a palindrome.
 *
 * Return all possible palindrome partitioning of s.
 *
 * Example:
 *
 *
 * Input: "aab"
 * Output:
 * [
 *   ["aa","b"],
 *   ["a","a","b"]
 * ]
 *
 *
 */

package leetcode

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func backtrack(str string, l int, answer []string, res *[][]string) {
	if len(answer) > 0 && l >= len(str) {
		temp := append(answer[:0:0], answer...)
		*res = append(*res, temp)
		return
	}

	for i := l; i < len(str); i++ {
		if isPalindrome(str, l, i) {
			if l == i {
				answer = append(answer, string(str[i]))
			} else {
				answer = append(answer, str[l:i+1])
			}
			backtrack(str, i+1, answer, res)
			answer = answer[:len(answer)-1]
		}
	}
}

func partition(s string) [][]string {
	res := [][]string{}
	backtrack(s, 0, []string{}, &res)
	return res
}
