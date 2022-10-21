/**
 * Given a string s, determine if it is valid.
 * A string s is valid if, starting with an empty string t = "", you can transform t into s after performing the following operation any number of times:
 *
 * 	Insert string "abc" into any position in t. More formally, t becomes tleft + "abc" + tright, where t == tleft + tright. Note that tleft and tright may be empty.
 *
 * Return true if s is a valid string, otherwise, return false.
 *
 * Example 1:
 *
 * Input: s = "aabcbc"
 * Output: true
 * Explanation:
 * "" -> "<u>abc</u>" -> "a<u>abc</u>bc"
 * Thus, "aabcbc" is valid.
 * Example 2:
 *
 * Input: s = "abcabcababcc"
 * Output: true
 * Explanation:
 * "" -> "<u>abc</u>" -> "abc<u>abc</u>" -> "abcabc<u>abc</u>" -> "abcabcab<u>abc</u>c"
 * Thus, "abcabcababcc" is valid.
 *
 * Example 3:
 *
 * Input: s = "abccba"
 * Output: false
 * Explanation: It is impossible to get "abccba" using the operation.
 *
 * Example 4:
 *
 * Input: s = "cababc"
 * Output: false
 * Explanation: It is impossible to get "cababc" using the operation.
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 2 * 10^4
 * 	s consists of letters 'a', 'b', and 'c'
 *
 */

package leetcode

type Stack struct {
	storage []rune
}

func (s *Stack) Push(i rune) {
	s.storage = append(s.storage, i)
}

func (s *Stack) Pop() rune {
	i := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return i
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func isValid(S string) bool {
	stack := &Stack{}
	for _, char := range S {
		if char == 'c' {
			ok := !stack.Empty() && stack.Pop() == 'b' && !stack.Empty() && stack.Pop() == 'a'
			if !ok {
				return false
			}
		} else {
			stack.Push(char)
		}
	}
	return stack.Empty()
}
