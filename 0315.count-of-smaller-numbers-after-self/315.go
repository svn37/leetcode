/**
 * You are given an integer array nums and you have to return a new counts array. The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].
 *
 * Example 1:
 *
 * Input: nums = [5,2,6,1]
 * Output: [2,1,1,0]
 * Explanation:
 * To the right of 5 there are 2 smaller elements (2 and 1).
 * To the right of 2 there is only 1 smaller element (1).
 * To the right of 6 there is 1 smaller element (1).
 * To the right of 1 there is 0 smaller element.
 *
 *
 * Constraints:
 *
 * 	0 <= nums.length <= 10^5
 * 	-10^4 <= nums[i] <= 10^4
 *
 */

package leetcode

type Pair struct {
	idx, val int
}

func mergeSort(pairs []*Pair, smaller []int) []*Pair {
	if len(pairs) < 2 {
		return pairs
	}

	mid := len(pairs) / 2

	// https://github.com/go101/go101/wiki/How-to-perfectly-clone-a-slice%3F
	leftCopy := append(pairs[:0:0], pairs[:mid]...)
	rightCopy := append(pairs[:0:0], pairs[mid:]...)

	left := mergeSort(leftCopy, smaller)
	right := mergeSort(rightCopy, smaller)

	i, j := 0, 0
	for i < len(left) || j < len(right) {
		if j == len(right) || (i < len(left) && left[i].val <= right[j].val) {
			pairs[i+j] = left[i]
			// except for this line, it's simply merge sort
			// there are j elements greater than left[i].val
			// in this round of merge
			smaller[left[i].idx] += j
			i++
		} else {
			pairs[i+j] = right[j]
			j++
		}
	}
	return pairs
}

// mergesort solution
func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	pairs := make([]*Pair, len(nums))
	for i := range pairs {
		pairs[i] = &Pair{i, nums[i]}
	}

	smaller := make([]int, len(nums))
	mergeSort(pairs, smaller)

	return smaller
}
