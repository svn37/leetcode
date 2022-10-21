/**
 * Suppose you have a random list of people standing in a queue. Each person is described by a pair of integers (h, k), where h is the height of the person and k is the number of people in front of this person who have a height greater than or equal to h. Write an algorithm to reconstruct the queue.
 *
 * Note:<br />
 * The number of people is less than 1,100.
 *
 *
 * Example
 *
 *
 * Input:
 * [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
 *
 * Output:
 * [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
 *
 *
 *
 *
 */

package leetcode

import "sort"

// https://leetcode.com/problems/queue-reconstruction-by-height/discuss/167308/Python-solution -- explanation
// https://leetcode.com/problems/queue-reconstruction-by-height/discuss/124289/golang-solution -- go version
func reconstructQueue(people [][]int) [][]int {
	res := people
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		}
		if people[i][0] == people[j][0] && people[i][1] < people[j][1] {
			return true
		}
		return false
	})

	for i, p := range people {
		j := p[1]
		copy(res[j+1:i+1], res[j:i])
		res[j] = p
	}
	return res
}
