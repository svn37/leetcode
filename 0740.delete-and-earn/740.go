/**
 * Given an array nums of integers, you can perform operations on the array.
 *
 * In each operation, you pick any nums[i] and delete it to earn nums[i] points. After, you must delete every element equal to nums[i] - 1 or nums[i] + 1.
 *
 * You start with 0 points. Return the maximum number of points you can earn by applying such operations.
 *
 * Example 1:
 *
 *
 * Input: nums = [3, 4, 2]
 * Output: 6
 * Explanation:
 * Delete 4 to earn 4 points, consequently 3 is also deleted.
 * Then, delete 2 to earn 2 points. 6 total points are earned.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2, 2, 3, 3, 3, 4]
 * Output: 9
 * Explanation:
 * Delete 3 to earn 3 points, deleting both 2's and the 4.
 * Then, delete 3 again to earn 3 points, and 3 again to earn 3 points.
 * 9 total points are earned.
 *
 *
 *
 *
 * Note:
 *
 *
 * 	The length of nums is at most 20000.
 * 	Each element nums[i] is an integer in the range [1, 10000].
 *
 *
 *
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// We first transform the nums array into a points array that sums up the total number of points for that particular value.
// A value of x will be assigned to index x in points.
// nums: [2, 2, 3, 3, 3, 4] (2 appears 2 times, 3 appears 3 times, 4 appears once)
// points: [0, 0, 4, 9, 4] <- This is the gold in each house!
// The condition that we cannot pick adjacent values is similar to the House Robber question that we
// cannot rob adjacent houses. Simply pass points into the rob function for a quick win 🌝!
func deleteAndEarn(nums []int) int {
	buckets := make([]int, 10001)
	for i := range nums {
		buckets[nums[i]] += nums[i]
	}

	prev, curr := 0, 0
	for _, value := range buckets {
		prev, curr = curr, max(prev+value, curr)
	}
	return curr
}
