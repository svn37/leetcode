/**
 * Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is the smallest in lexicographical order among all possible results.
 * Note: This question is the same as 1081: <a href="https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/" target="_blank">https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/</a>
 *
 * Example 1:
 *
 * Input: s = "bcabc"
 * Output: "abc"
 *
 * Example 2:
 *
 * Input: s = "cbacdcbc"
 * Output: "acdb"
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 10^4
 * 	s consists of lowercase English letters.
 *
 */

package leetcode

type Stack struct {
	items []byte
	idx   int
}

func NewStack(length int) *Stack {
	return &Stack{
		items: make([]byte, length),
	}
}

func (s *Stack) Push(val byte) {
	s.items[s.idx] = val
	s.idx++
}

func (s *Stack) Pop() byte {
	s.idx--
	return s.items[s.idx]
}

func (s *Stack) Peek() byte {
	return s.items[s.idx-1]
}

func (s *Stack) String() string {
	return string(s.items[:s.idx])
}

func (s *Stack) Empty() bool {
	return s.idx == 0
}

func removeDuplicateLetters(s string) string {
	length := 0
	chars := make([]int, 26)
	for _, char := range s {
		if chars[char-'a'] == 0 {
			length++
		}
		chars[char-'a']++
	}

	stack := NewStack(length)
	visited := make(map[byte]bool, length)

	for i := range s {
		chars[s[i]-'a']--

		if visited[s[i]] {
			continue
		}

		for !stack.Empty() && stack.Peek() > s[i] && chars[stack.Peek()-'a'] > 0 {
			visited[stack.Peek()] = false
			stack.Pop()
		}

		stack.Push(s[i])
		visited[s[i]] = true
	}

	return stack.String()
}
