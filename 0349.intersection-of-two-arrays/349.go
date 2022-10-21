/**
 * Given two arrays, write a function to compute their intersection.
 *
 * Example 1:
 *
 *
 * Input: nums1 = <span id="example-input-1-1">[1,2,2,1]</span>, nums2 = <span id="example-input-1-2">[2,2]</span>
 * Output: <span id="example-output-1">[2]</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: nums1 = <span id="example-input-2-1">[4,9,5]</span>, nums2 = <span id="example-input-2-2">[9,4,9,8,4]</span>
 * Output: <span id="example-output-2">[9,4]</span>
 * </div>
 *
 * Note:
 *
 *
 * 	Each element in the result must be unique.
 * 	The result can be in any order.
 *
 *
 *
 *
 */

package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, num := range nums1 {
		m[num] = struct{}{}
	}

	res := []int{}
	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			res = append(res, num)
			delete(m, num)
		}
	}
	return res
}
