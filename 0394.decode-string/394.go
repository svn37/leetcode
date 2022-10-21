/**
 * Given an encoded string, return its decoded string.
 * The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.
 * You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.
 * Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].
 *
 * Example 1:
 * Input: s = "3[a]2[bc]"
 * Output: "aaabcbc"
 * Example 2:
 * Input: s = "3[a2[c]]"
 * Output: "accaccacc"
 * Example 3:
 * Input: s = "2[abc]3[cd]ef"
 * Output: "abcabccdcdcdef"
 * Example 4:
 * Input: s = "abc3[cd]xyz"
 * Output: "abccdcdcdxyz"
 *
 */

package leetcode

import "bytes"

// recursive solution
func decode(s string, start int, b *bytes.Buffer) int {
	var digit int
	for i := start; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit *= 10
			digit += int(s[i] - '0')
		} else if s[i] == '[' {
			oldLen := b.Len()
			i = decode(s, i+1, b)
			// copy, but skip the old data (do not multiply it)
			copied := b.Bytes()[oldLen:]
			for digit > 1 {
				b.Write(copied)
				digit--
			}
			digit-- // it becomes zero
		} else if s[i] == ']' {
			return i
		} else {
			b.WriteByte(s[i])
		}
	}
	// noop, returns only in top-level decode
	return 0
}

func decodeString(s string) string {
	var b bytes.Buffer
	decode(s, 0, &b)
	return b.String()
}

// stack solution
// array represents [multiplier, index]
type Stack struct {
	items [][2]int
}

func (s *Stack) Push(val [2]int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() [2]int {
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func decodeString2(s string) string {
	var b bytes.Buffer
	var digit int
	stack := &Stack{}

	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			digit *= 10
			digit += int(s[i] - '0')
		} else if s[i] == '[' {
			stack.Push([2]int{digit, b.Len()})
			digit = 0
		} else if s[i] == ']' {
			pair := stack.Pop()
			multiplier, oldLen := pair[0], pair[1]
			// copy, but skip the old data (do not multiply it)
			copied := b.Bytes()[oldLen:]
			for multiplier > 1 {
				b.Write(copied)
				multiplier--
			}
		} else {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}
