/**
 * You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.
 * Return the max sliding window.
 *
 * Example 1:
 *
 * Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
 * Output: [3,3,5,5,6,7]
 * Explanation:
 * Window position                Max
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 *  1 [3  -1  -3] 5  3  6  7       3
 *  1  3 [-1  -3  5] 3  6  7       5
 *  1  3  -1 [-3  5  3] 6  7       5
 *  1  3  -1  -3 [5  3  6] 7       6
 *  1  3  -1  -3  5 [3  6  7]      7
 *
 * Example 2:
 *
 * Input: nums = [1], k = 1
 * Output: [1]
 *
 * Example 3:
 *
 * Input: nums = [1,-1], k = 1
 * Output: [1,-1]
 *
 * Example 4:
 *
 * Input: nums = [9,11], k = 2
 * Output: [11]
 *
 * Example 5:
 *
 * Input: nums = [4,-2], k = 2
 * Output: [4]
 *
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 10^5
 * 	-10^4 <= nums[i] <= 10^4
 * 	1 <= k <= nums.length
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k <= 0 {
		return nil
	}

	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	maxLeft[0] = nums[0]
	maxRight[n-1] = nums[n-1]

	// break the array into k-length chunks and calculate minimum
	// to the left and to the right of i in each of them
	for i, j := 1, n-2; i < n; i, j = i+1, j-1 {
		if i%k == 0 {
			maxLeft[i] = nums[i]
		} else {
			maxLeft[i] = max(nums[i], maxLeft[i-1])
		}

		if j%k == 0 {
			maxRight[j] = nums[j]
		} else {
			maxRight[j] = max(nums[j], maxRight[j+1])
		}
	}

	res := make([]int, len(nums)-k+1)

	for i, j := 0, 0; i+k <= n; i, j = i+1, j+1 {
		res[j] = max(maxRight[i], maxLeft[i+k-1])
	}
	return res
}
