/**
 * Given an integer array nums, find the sum of the elements between indices i and j (i &le; j), inclusive.
 * Implement the NumArray class:
 *
 * 	NumArray(int[] nums) Initializes the object with the integer array nums.
 * 	int sumRange(int i, int j) Return the sum of the elements of the nums array in the range [i, j] inclusive (i.e., sum(nums[i], nums[i + 1], ... , nums[j]))
 *
 *
 * Example 1:
 *
 * Input
 * ["NumArray", "sumRange", "sumRange", "sumRange"]
 * [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
 * Output
 * [null, 1, -1, -3]
 * Explanation
 * NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
 * numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
 * numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
 * numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
 *
 *
 * Constraints:
 *
 * 	0 <= nums.length <= 10^4
 * 	-10^5 <= nums[i] <= 10^5
 * 	0 <= i <= j < nums.length
 * 	At most 10^4 calls will be made to sumRange.
 *
 */

package leetcode

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums)+1)
	for i := range nums {
		prefixSum[i+1] += nums[i] + prefixSum[i]
	}
	return NumArray{prefixSum}
}

// find the sum of the elements between indices i and j (i ≤ j), inclusive
func (this *NumArray) SumRange(i int, j int) int {
	return this.prefixSum[j+1] - this.prefixSum[i]
}
