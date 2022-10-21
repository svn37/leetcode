/**
 * Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.
 *
 * <div>
 * Example 1:
 *
 *
 * Input: nums = <span id="example-input-1-1">[1,2,3,1]</span>, k = <span id="example-input-1-2">3</span>
 * Output: <span id="example-output-1">true</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: nums = <span id="example-input-2-1">[1,0,1,1]</span>, k = <span id="example-input-2-2">1</span>
 * Output: <span id="example-output-2">true</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: nums = <span id="example-input-3-1">[1,2,3,1,2,3]</span>, k = <span id="example-input-3-2">2</span>
 * Output: <span id="example-output-3">false</span>
 *
 * </div>
 * </div>
 * </div>
 *
 */

package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]struct{})
	// sliding window
	for i := range nums {
		if i > k {
			delete(m, nums[i-k-1])
		}
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
	}
	return false
}
