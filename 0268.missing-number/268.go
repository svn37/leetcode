/**
 * Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
 * Follow up: Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?
 *
 * Example 1:
 *
 * Input: nums = [3,0,1]
 * Output: 2
 * Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
 *
 * Example 2:
 *
 * Input: nums = [0,1]
 * Output: 2
 * Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
 *
 * Example 3:
 *
 * Input: nums = [9,6,4,2,3,5,7,0,1]
 * Output: 8
 * Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.
 *
 * Example 4:
 *
 * Input: nums = [0]
 * Output: 1
 * Explanation: n = 1 since there is 1 number, so all numbers are in the range [0,1]. 1 is the missing number in the range since it does not appear in nums.
 *
 *
 * Constraints:
 *
 * 	n == nums.length
 * 	1 <= n <= 10^4
 * 	0 <= nums[i] <= n
 * 	All the numbers of nums are unique.
 *
 */

package leetcode

// Given an array containing n distinct numbers taken from 0, 1, 2, ..., n
// Method 1. Each value will be a valid index, mark already visited values.
func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

func missingNumber(nums []int) int {
	beyondRange := (len(nums) + 1) * (-1)

	for i := range nums {
		if nums[i] == 0 || nums[i] == beyondRange {
			continue
		}

		num := abs(nums[i])
		if nums[num-1] == 0 {
			nums[num-1] = beyondRange
		} else {
			nums[num-1] *= -1
		}
	}

	for i := range nums {
		if nums[i] >= 0 {
			return i + 1
		}
	}
	return 0
}

// Method 2. Xor
func missingNumber2(nums []int) int {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor ^= i ^ nums[i]
	}
	// n elements and (n-1) indexes
	return xor ^ len(nums)
}

// Method 3. Gauss' formula
func missingNumber3(nums []int) int {
	n := len(nums)
	sum := 0

	for i := range nums {
		sum += nums[i]
	}
	return n*(n+1)/2 - sum
}
