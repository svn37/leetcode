/**
 * Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
 * Follow up: The overall run time complexity should be O(log (m+n)).
 *
 * Example 1:
 *
 * Input: nums1 = [1,3], nums2 = [2]
 * Output: 2.00000
 * Explanation: merged array = [1,2,3] and median is 2.
 *
 * Example 2:
 *
 * Input: nums1 = [1,2], nums2 = [3,4]
 * Output: 2.50000
 * Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 *
 * Example 3:
 *
 * Input: nums1 = [0,0], nums2 = [0,0]
 * Output: 0.00000
 *
 * Example 4:
 *
 * Input: nums1 = [], nums2 = [1]
 * Output: 1.00000
 *
 * Example 5:
 *
 * Input: nums1 = [2], nums2 = []
 * Output: 2.00000
 *
 *
 * Constraints:
 *
 * 	nums1.length == m
 * 	nums2.length == n
 * 	0 <= m <= 1000
 * 	0 <= n <= 1000
 * 	1 <= m + n <= 2000
 * 	-10^6 <= nums1[i], nums2[i] <= 10^6
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://leetcode.com/problems/median-of-two-sorted-arrays/discuss/2481/Share-my-O(log(min(mn)))-solution-with-explanation
// log(min(m,n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	// ensure that m <= n
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}

	// binary search of nums1, such that both parts are equal, and
	// nums2[j-1] <= nums1[i] && nums1[i-1] <= nums2[j]
	imin, imax, half := 0, m, (m+n+1)/2

	for imin <= imax {
		i := (imin + imax) / 2
		j := half - i

		if i < m && nums2[j-1] > nums1[i] {
			// i is too small, search in the right part
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// i is too big, search in the left part
			imax = i - 1
		} else {
			// i is found
			var maxOfLeft int
			if i == 0 {
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				maxOfLeft = nums1[i-1]
			} else {
				maxOfLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}

			var minOfRight int
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				minOfRight = min(nums1[i], nums2[j])
			}

			return float64(maxOfLeft+minOfRight) / 2
		}
	}

	return 0.0
}

// brute-force solution, merge the arrays and find the median
// m+n
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	res := make([]int, 0, len(nums1)+len(nums2))

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i], nums2[j])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	if i < len(nums1) {
		res = append(res, nums1[i:]...)
	} else if j < len(nums2) {
		res = append(res, nums2[j:]...)
	}

	if len(res)%2 == 1 {
		return float64(res[len(res)/2])
	}

	return float64(res[len(res)/2]+res[len(res)/2-1]) / 2
}
