/**
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
 * Note:
 *
 * 	The number of elements initialized in nums1 and nums2 are m and n respectively.
 * 	You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
 *
 * Example:
 *
 * Input:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 * Output: [1,2,2,3,5,6]
 *
 *
 * Constraints:
 *
 * 	-10^9 <= nums1[i], nums2[i] <= 10^9
 * 	nums1.length == m + n
 * 	nums2.length == n
 *
 */

package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	// nums1 = [1,2,3,0,0,0], m = 3
	// nums2 = [2,5,6],       n = 3

	// nums1 has free space at the end, after m
	// copy nums1 to this free space (or part of nums1)
	copy(nums1[len(nums1)-m:], nums1[:m])

	// start merging this copied array and nums1
	// the beginning of nums1 is the working area
	i, j := 0, 0
	for k := len(nums1) - m; k < len(nums1) && j < n; i++ {
		if nums1[k] >= nums2[j] {
			nums1[i] = nums2[j]
			j++
		} else {
			nums1[i] = nums1[k]
			k++
		}
	}

	// at the end we have two options:
	// 1) nums2 was merged and j == len(nums2),
	//    so the rest of nums1 already contains sorted array
	// 2) nums2 was not merged completely, so just copy it
	//    to the end of nums1
	if j != len(nums2) {
		copy(nums1[i:], nums2[j:])
	}
}
