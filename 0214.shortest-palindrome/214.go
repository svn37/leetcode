/**
 * Given a string s, you can convert it to a palindrome by adding characters in front of it. Find and return the shortest palindrome you can find by performing this transformation.
 *
 * Example 1:
 * Input: s = "aacecaaa"
 * Output: "aaacecaaa"
 * Example 2:
 * Input: s = "abcd"
 * Output: "dcbabcd"
 *
 * Constraints:
 *
 * 	0 <= s.length <= 5 * 10^4
 * 	s consists of lowercase English letters only.
 *
 */

package leetcode

// Method 1. Brute-force
func expandAroundCenter(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j
}

func makePalindrome(s string, end int) string {
	newStr := s
	for i := end; i < len(s); i++ {
		newStr = string(s[i]) + newStr
	}
	return newStr
}

func shortestPalindrome(s string) string {
	for i := len(s) / 2; i >= 0; i-- {
		begin1, end1 := expandAroundCenter(s, i, i+1)
		if begin1 == 0 {
			return makePalindrome(s, end1)
		}
		begin2, end2 := expandAroundCenter(s, i, i)
		if begin2 == 0 {
			return makePalindrome(s, end2)
		}
	}
	// assert(false)
	// at least i = 0 will invoke makePalindrome,
	// because one char is always a palindrome
	return ""
}

// Method 2. Use KMP prefix array
// We can convert it to an alternative problem "find the longest palindrome substring starts from index 0".
func shortestPalindrome2(str string) string {
	S := []byte(str)
	for i, j := 0, len(S)-1; i < j; i, j = i+1, j-1 {
		S[i], S[j] = S[j], S[i]
	}
	revStr := string(S)

	// run KMP on s
	s := str + "#" + revStr

	// kmp prefix array
	p := make([]int, len(s))
	for i := 1; i < len(p); i++ {
		j := p[i-1]
		for j > 0 && s[i] != s[j] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j
	}

	return revStr[:len(revStr)-p[len(p)-1]] + str
}
