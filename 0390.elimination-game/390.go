/**
 *
 * There is a list of sorted integers from 1 to n. Starting from left to right, remove the first number and every other number afterward until you reach the end of the list.
 *
 * Repeat the previous step again, but this time from right to left, remove the right most number and every other number from the remaining numbers.
 *
 * We keep repeating the steps again, alternating left to right and right to left, until a single number remains.
 *
 * Find the last number that remains starting with a list of length n.
 *
 * Example:
 *
 * Input:
 * n = 9,
 * <u>1</u> 2 <u>3</u> 4 <u>5</u> 6 <u>7</u> 8 <u>9</u>
 * 2 <u>4</u> 6 <u>8</u>
 * <u>2</u> 6
 * 6
 *
 * Output:
 * 6
 *
 *
 */

package leetcode

// correct and efficient solution, the solution is more explanative
/*
 * func lastRemaining(n int) int {
 *   if n == 1 {
 *     return 1
 *   } else {
 *     return 2 * (1 + n/2 - lastRemaining(n/2))
 *   }
 * }
 */

// eliminate [1...n] first from left to right, then alternate
func leftToRight(n int) int {
	if n == 1 {
		return 1
	}
	// scan from left to right is simple, the length of array doesn't matter
	// [1, 2, 3, 4] -> 2 * [1, 2]
	// [1, 2, 3, 4, 5] -> 2 * [1, 2]
	return 2 * rightToLeft(n/2)
}

// eliminate [1...n] first from right to left, then alternate
func rightToLeft(n int) int {
	if n == 1 {
		return 1
	}
	// if the length of array is even, we will get only odd number
	// [1, 2, 3, 4] -> [1, 3] = 2 * [1, 2] - 1
	if n%2 == 0 {
		return 2*leftToRight(n/2) - 1
		// else if the length of array is odd, we will get only even number
		// [1, 2, 3, 4, 5] -> [2, 4] = 2 * [1, 2]
	} else {
		return 2 * leftToRight(n/2)
	}
}

func lastRemaining(n int) int {
	return leftToRight(n)
}
