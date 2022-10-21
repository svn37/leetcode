/**
 * Alice has a hand of cards, given as an array of integers.
 * Now she wants to rearrange the cards into groups so that each group is size W, and consists of W consecutive cards.
 * Return true if and only if she can.
 * Note: This question is the same as 1296: <a href="https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/">https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/</a>
 *
 * Example 1:
 *
 * Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
 * Output: true
 * Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]
 *
 * Example 2:
 *
 * Input: hand = [1,2,3,4,5], W = 4
 * Output: false
 * Explanation: Alice's hand can't be rearranged into groups of 4.
 *
 *
 * Constraints:
 *
 * 	1 <= hand.length <= 10000
 * 	0 <= hand[i] <= 10^9
 * 	1 <= W <= hand.length
 *
 */

package leetcode

import (
	"math"
	"math/rand"
)

func qsort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivot := rand.Intn(len(a))
	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
}

// this function modifies the input in-place, the same could be done by
// copying the array, if n is not large
func isNStraightHand(hand []int, W int) bool {
	// no need to even try
	if len(hand)%W != 0 {
		return false
	}

	qsort(hand)

	// minInt64 will be dummy value
	for i := range hand {
		if hand[i] != math.MinInt64 {
			w := 1
			prev := hand[i]
			// stop if W consecutive elements were found
			for j := i + 1; j < len(hand) && w != W; j++ {
				if hand[j] == math.MinInt64 {
					continue
				} else if hand[j]-prev == 1 {
					prev = hand[j]
					hand[j] = math.MinInt64
					w++
					// this sequence will be never be completed, return
				} else if hand[j]-prev > 1 {
					return false
				}
			}
			// loop finished, but W consecutive numbers weren't found
			if w != W {
				return false
			}

			hand[i] = math.MinInt64
		}
	}
	return true
}
