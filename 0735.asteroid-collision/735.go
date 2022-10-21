/**
 * We are given an array asteroids of integers representing asteroids in a row.
 * For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.
 * Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.
 *
 * Example 1:
 *
 * Input: asteroids = [5,10,-5]
 * Output: [5,10]
 * Explanation: The 10 and -5 collide resulting in 10.  The 5 and 10 never collide.
 *
 * Example 2:
 *
 * Input: asteroids = [8,-8]
 * Output: []
 * Explanation: The 8 and -8 collide exploding each other.
 *
 * Example 3:
 *
 * Input: asteroids = [10,2,-5]
 * Output: [10]
 * Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.
 *
 * Example 4:
 *
 * Input: asteroids = [-2,-1,1,2]
 * Output: [-2,-1,1,2]
 * Explanation: The -2 and -1 are moving left, while the 1 and 2 are moving right. Asteroids moving the same direction never meet, so no asteroids will meet each other.
 *
 *
 * Constraints:
 *
 * 	1 <= asteroids <= 10^4
 * 	-1000 <= asteroids[i] <= 1000
 * 	asteroids[i] != 0
 *
 */

package leetcode

type Stack struct {
	storage []int
}

func (s *Stack) Push(val int) {
	s.storage = append(s.storage, val)
}

func (s *Stack) Peek(offset int) int {
	return s.storage[len(s.storage)-offset]
}

func (s *Stack) Pop() int {
	pop := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return pop
}

func (s *Stack) Len() int {
	return len(s.storage)
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func asteroidCollision(asteroids []int) []int {
	if len(asteroids) == 0 {
		return []int{}
	}

	stack := &Stack{}

	for i := range asteroids {
		stack.Push(asteroids[i])

		for stack.Len() > 1 {
			a1 := stack.Peek(1)
			a2 := stack.Peek(2)

			if (a1 < 0) && (a2 > 0) {
				if abs(a1) == abs(a2) {
					stack.Pop()
					stack.Pop()
				} else if abs(a1) > abs(a2) {
					stack.Pop()
					stack.Pop()
					stack.Push(a1)
				} else {
					stack.Pop()
				}
			} else {
				break
			}
		}
	}

	res := make([]int, stack.Len())
	i := stack.Len() - 1
	for !stack.Empty() {
		res[i] = stack.Pop()
		i--
	}
	return res
}
