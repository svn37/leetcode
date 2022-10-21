/**
 * Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
 *
 * Example 1:
 *
 * Input: s = "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()".
 *
 * Example 2:
 *
 * Input: s = ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()".
 *
 * Example 3:
 *
 * Input: s = ""
 * Output: 0
 *
 *
 * Constraints:
 *
 * 	0 <= s.length <= 3 * 10^4
 * 	s[i] is '(', or ')'.
 *
 */

package leetcode

// Method 1. Using stack
type Stack struct {
	storage []int
}

func (s *Stack) Push(i int) {
	s.storage = append(s.storage, i)
}

func (s *Stack) Pop() int {
	i := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return i
}

func (s *Stack) Top() int {
	return s.storage[len(s.storage)-1]
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses(str string) int {
	stack := &Stack{}

	// put indexes on the stack
	// strings between these indexes are valid
	for i, char := range str {
		if stack.Empty() || char == '(' || str[stack.Top()] == ')' {
			// string is not valid, push character
			stack.Push(i)
		} else {
			// found "()"
			stack.Pop()
		}
	}

	longest := 0
	end := len(str)

	// find the longest string between invalid indexes
	for !stack.Empty() {
		start := stack.Top()
		longest = max(longest, end-start-1)
		end = stack.Pop()
	}
	return max(longest, end)
}

// Method 2. Dynamic programming.
//  1. if str[i-1] == '(', dp[i] = dp[i-2] + 2
//  2. if str[i-1] == ')' && str[i-dp[i-1]-1] == '(',
//     then dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
func longestValidParentheses2(str string) int {
	if len(str) < 2 {
		return 0
	}

	longest := 0
	dp := make([]int, len(str))

	for i := 1; i < len(str); i++ {
		if str[i] == ')' && i-dp[i-1]-1 >= 0 && str[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
			longest = max(longest, dp[i])
		}
	}
	return longest
}
