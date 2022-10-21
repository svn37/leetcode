/**
 * The i-th person has weight people[i], and each boat can carry a maximum weight of limit.
 *
 * Each boat carries at most 2 people at the same time, provided the sum of the weight of those people is at most limit.
 *
 * Return the minimum number of boats to carry every given person.  (It is guaranteed each person can be carried by a boat.)
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: people = <span id="example-input-1-1">[1,2]</span>, limit = <span id="example-input-1-2">3</span>
 * Output: <span id="example-output-1">1</span>
 * Explanation: 1 boat (1, 2)
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: people = <span id="example-input-2-1">[3,2,2,1]</span>, limit = <span id="example-input-2-2">3</span>
 * Output: <span id="example-output-2">3</span>
 * Explanation: 3 boats (1, 2), (2) and (3)
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: people = <span id="example-input-3-1">[3,5,3,4]</span>, limit = <span id="example-input-3-2">5</span>
 * Output: <span id="example-output-3">4</span>
 * Explanation: 4 boats (3), (3), (4), (5)
 *
 * Note:
 *
 *
 * 	1 <= people.length <= 50000
 * 	1 <= people[i] <= limit <= 30000
 *
 * </div>
 * </div>
 * </div>
 *
 */

package leetcode

import "math/rand"

func qsort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivot := rand.Intn(len(a))

	a[right], a[pivot] = a[pivot], a[right]
	for i := range a {
		if a[i] > a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	qsort(a[:left])
	qsort(a[left+1:])
}

func numRescueBoats(people []int, limit int) int {
	qsort(people)

	numBoats := 0
	for len(people) > 0 {
		try := people[0]
		people = people[1:]

		if len(people) == 0 {
			numBoats++
			break
		}

		left := limit - try

		// TODO If the heaviest person can share a boat with the lightest person, then do so
		// this solution was meant to solve another problem
		// when then there is no constraint for how many people a boat can carry
		i := len(people) - 1
		for ; i >= 0; i-- {
			if left-people[i] == 0 {
				break
			} else if left-people[i] < 0 {
				i++
				break
			}
		}
		if i < 0 {
			i = 0
		}

		if i == 0 {
			people = people[1:]
		} else if i == len(people)-1 {
			people = people[:len(people)-1]
		} else if i < len(people)-1 {
			people = append(people[:i], people[i+1:]...)
		}

		numBoats++
	}
	return numBoats
}
