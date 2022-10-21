/**
 * Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum &ge; s. If there isn't one, return 0 instead.
 *
 * Example:
 *
 *
 * Input: s = 7, nums = [2,3,1,2,4,3]
 * Output: 2
 * Explanation: the subarray [4,3] has the minimal length under the problem constraint.
 *
 * <div class="spoilers">Follow up:</div>
 *
 * <div class="spoilers">If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n). </div>
 *
 */

package leetcode

// Method 1 Two pointers, O(n)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSubArrayLen(s int, nums []int) int {
	left, right := 0, 0
	sum := 0
	minVal := len(nums) + 1

	for right < len(nums) {
		sum += nums[right]

		for sum >= s {
			minVal = min(minVal, right-left+1)

			sum -= nums[left]
			left++
		}
		right++
	}
	if minVal == len(nums)+1 {
		return 0
	}
	return minVal
}

// Method 2. O(n log n)
func bsearch(arr []int, lo, hi int, target int) int {
	for lo <= hi {
		mid := (lo + hi) / 2
		if arr[mid] >= target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func minSubArrayLen2(s int, nums []int) int {
	prefixSum := make([]int, len(nums)+1)
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	minVal := len(nums) + 1

	for i := range prefixSum {
		end := bsearch(prefixSum, i+1, len(prefixSum)-1, prefixSum[i]+s)
		if end == len(prefixSum) {
			break
		}
		minVal = min(minVal, end-i)
	}
	if minVal == len(nums)+1 {
		return 0
	}
	return minVal
}
