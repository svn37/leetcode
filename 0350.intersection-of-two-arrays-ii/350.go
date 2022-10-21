/**
 * Given two arrays, write a function to compute their intersection.
 *
 * Example 1:
 *
 *
 * Input: nums1 = <span id="example-input-1-1">[1,2,2,1]</span>, nums2 = <span id="example-input-1-2">[2,2]</span>
 * Output: <span id="example-output-1">[2,2]</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: nums1 = <span id="example-input-2-1">[4,9,5]</span>, nums2 = <span id="example-input-2-2">[9,4,9,8,4]</span>
 * Output: <span id="example-output-2">[4,9]</span>
 * </div>
 *
 * Note:
 *
 *
 * 	Each element in the result should appear as many times as it shows in both arrays.
 * 	The result can be in any order.
 *
 *
 * Follow up:
 *
 *
 * 	What if the given array is already sorted? How would you optimize your algorithm?
 * 	What if nums1's size is small compared to nums2's size? Which algorithm is better?
 * 	What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
 *
 *
 */

package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}

	res := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			res = append(res, num)
		}
		m[num]--
	}
	return res
}
