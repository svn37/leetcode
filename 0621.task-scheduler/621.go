/**
 * Given a characters array tasks, representing the tasks a CPU needs to do, where each letter represents a different task. Tasks could be done in any order. Each task is done in one unit of time. For each unit of time, the CPU could complete either one task or just be idle.
 * However, there is a non-negative integer n that represents the cooldown period between two same tasks (the same letter in the array), that is that there must be at least n units of time between any two same tasks.
 * Return the least number of units of times that the CPU will take to finish all the given tasks.
 *
 * Example 1:
 *
 * Input: tasks = ["A","A","A","B","B","B"], n = 2
 * Output: 8
 * Explanation:
 * A -> B -> idle -> A -> B -> idle -> A -> B
 * There is at least 2 units of time between any two same tasks.
 *
 * Example 2:
 *
 * Input: tasks = ["A","A","A","B","B","B"], n = 0
 * Output: 6
 * Explanation: On this case any permutation of size 6 would work since n = 0.
 * ["A","A","A","B","B","B"]
 * ["A","B","A","B","A","B"]
 * ["B","B","B","A","A","A"]
 * ...
 * And so on.
 *
 * Example 3:
 *
 * Input: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
 * Output: 16
 * Explanation:
 * One possible solution is
 * A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> idle -> idle -> A -> idle -> idle -> A
 *
 *
 * Constraints:
 *
 * 	1 <= task.length <= 10^4
 * 	tasks[i] is upper-case English letter.
 * 	The integer n is in the range [0, 100].
 *
 */

package leetcode

import "sort"

// The key is to find out how many idles do we need.
// Let's first look at how to arrange them. it's not hard to figure out
// that we can do a "greedy arrangement": always arrange task with most frequency first.
// E.g. we have following tasks : 3 A, 2 B, 1 C. and we have n = 2. According to what
// we have above, we should first arrange A, and then B and C.
// Imagine there are "slots" and we need to arrange tasks by putting them into "slots"
func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}

	chars := make([]int, 26)
	for _, char := range tasks {
		chars[char-'A']++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(chars)))

	// remainder is the number of the most frequent elements
	remainder := 1
	for i := 1; i < len(chars); i++ {
		if chars[i-1] == chars[i] {
			remainder++
		} else {
			break
		}
	}

	// strange-looking formula
	// a - - a - - a
	// - is free slots
	// a b - a b - a b
	// ab at the end is remainder
	// all the other elements must fit because their length is decreasing in the chars array
	res := (chars[0]-1)*(n+1) + remainder
	if len(tasks) > res {
		return len(tasks)
	}
	return res
}
