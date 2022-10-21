/**
 * Given an array nums with n objects colored red, white, or blue, sort them <a href="https://en.wikipedia.org/wiki/In-place_algorithm" target="_blank">in-place</a> so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
 * Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
 * Follow up:
 *
 * 	Could you solve this problem without using the library's sort function?
 * 	Could you come up with a one-pass algorithm using only O(1) constant space?
 *
 *
 * Example 1:
 * Input: nums = [2,0,2,1,1,0]
 * Output: [0,0,1,1,2,2]
 * Example 2:
 * Input: nums = [2,0,1]
 * Output: [0,1,2]
 * Example 3:
 * Input: nums = [0]
 * Output: [0]
 * Example 4:
 * Input: nums = [1]
 * Output: [1]
 *
 * Constraints:
 *
 * 	n == nums.length
 * 	1 <= n <= 300
 * 	nums[i] is 0, 1, or 2.
 *
 */

package leetcode

// counting sort
func sortColors(nums []int) {
	var num0, num1, num2 int

	for i := range nums {
		switch nums[i] {
		case 0:
			num0++
		case 1:
			num1++
		case 2:
			num2++
		}
	}

	for i := 0; i < num0; i++ {
		nums[i] = 0
	}

	for i := num0; i < num0+num1; i++ {
		nums[i] = 1
	}

	for i := num0 + num1; i < len(nums); i++ {
		nums[i] = 2
	}
}

// in-place solution
func sortColors2(nums []int) {
	l, r := 0, len(nums)-1
	for i := 0; i <= r; i++ {
		for nums[i] == 2 && i < r {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}

		for nums[i] == 0 && i > l {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
}

// in-place solution
func sortColors3(nums []int) {
	zero, one, two := -1, -1, -1

	for i := range nums {
		two++
		if nums[i] == 2 {
			nums[two] = 2
			continue
		}
		one++
		if nums[i] == 1 {
			nums[two] = 2
			nums[one] = 1
			continue
		}
		zero++
		if nums[i] == 0 {
			nums[two] = 2
			nums[one] = 1
			nums[zero] = 0
		}
	}
}
