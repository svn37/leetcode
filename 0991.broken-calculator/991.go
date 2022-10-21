/**
 * On a broken calculator that has a number showing on its display, we can perform two operations:
 *
 *
 * 	Double: Multiply the number on the display by 2, or;
 * 	Decrement: Subtract 1 from the number on the display.
 *
 *
 * Initially, the calculator is displaying the number X.
 *
 * Return the minimum number of operations needed to display the number Y.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: X = <span id="example-input-1-1">2</span>, Y = <span id="example-input-1-2">3</span>
 * Output: <span id="example-output-1">2</span>
 * Explanation: Use double operation and then decrement operation {2 -> 4 -> 3}.
 *
 *
 * Example 2:
 *
 *
 * Input: X = <span id="example-input-2-1">5</span>, Y = <span id="example-input-2-2">8</span>
 * Output: <span id="example-output-2">2</span>
 * Explanation: Use decrement and then double {5 -> 4 -> 8}.
 *
 *
 * Example 3:
 *
 *
 * Input: X = <span id="example-input-3-1">3</span>, Y = <span id="example-input-3-2">10</span>
 * Output: <span id="example-output-3">3</span>
 * Explanation:  Use double, decrement and double {3 -> 6 -> 5 -> 10}.
 *
 *
 * Example 4:
 *
 *
 * Input: X = <span id="example-input-4-1">1024</span>, Y = <span id="example-input-4-2">1</span>
 * Output: <span id="example-output-4">1023</span>
 * Explanation: Use decrement operations 1023 times.
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= X <= 10^9
 * 	1 <= Y <= 10^9
 * </ol>
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func brokenCalc(X int, Y int) int {
	if X >= Y {
		return X - Y
	}

	var count int

	for {
		if Y%2 != 0 {
			Y++
			count++
		}

		if Y/2 > X {
			Y /= 2
			count++
		} else {
			return 1 + count + min(X*2-Y, X-Y/2)
		}
	}
}
