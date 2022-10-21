/**
 * Initially on a notepad only one character 'A' is present. You can perform two operations on this notepad for each step:
 *
 * <ol>
 * 	Copy All: You can copy all the characters present on the notepad (partial copy is not allowed).
 * 	Paste: You can paste the characters which are copied last time.
 * </ol>
 *
 *
 *
 * Given a number n. You have to get exactly n 'A' on the notepad by performing the minimum number of steps permitted. Output the minimum number of steps to get n 'A'.
 *
 * Example 1:
 *
 *
 * Input: 3
 * Output: 3
 * Explanation:
 * Intitally, we have one character 'A'.
 * In step 1, we use Copy All operation.
 * In step 2, we use Paste operation to get 'AA'.
 * In step 3, we use Paste operation to get 'AAA'.
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	The n will be in the range [1, 1000].
 * </ol>
 *
 *
 *
 */

package leetcode

func minSteps(n int) int {
	steps := 0
	for d := 2; d <= n; d++ {
		for n%d == 0 {
			steps += d
			n /= d
		}
	}
	return steps
}

/*
 * func minSteps(n int) int {
 *   if n < 2 {
 *     return 0
 *   }
 *   if n >= 2 && n <= 5 {
 *     return n
 *   }
 *   for i := n / 2; i > 2; i-- {
 *     if n%i == 0 {
 *       return (n / i) + minSteps(i)
 *     }
 *   }
 *   return n
 * }
 */
