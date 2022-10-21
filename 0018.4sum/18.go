/**
 * Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
 * Notice that the solution set must not contain duplicate quadruplets.
 *
 * Example 1:
 * Input: nums = [1,0,-1,0,-2,2], target = 0
 * Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 * Example 2:
 * Input: nums = [], target = 0
 * Output: []
 *
 * Constraints:
 *
 * 	0 <= nums.length <= 200
 * 	-10^9 <= nums[i] <= 10^9
 * 	-10^9 <= target <= 10^9
 *
 */

package leetcode

import "math/rand"

func qsort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivot := rand.Intn(len(arr))

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	qsort(arr[:left])
	qsort(arr[left+1:])
}

// The idea is to reduce k-sum problem to (k-1)-sum problem.
// While solving each problem, try to apply lots of optimizations.
// O(n^(k-1))

func twoSum(nums []int, target int, res *[][]int, z1, z2 int) {
	if 2*nums[0] > target || 2*nums[len(nums)-1] < target {
		return
	}

	lo, hi := 0, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum == target {
			*res = append(*res, []int{z1, z2, nums[lo], nums[hi]})

			lo++
			for lo < hi && nums[lo] == nums[lo-1] {
				lo++
			}

			hi--
			for lo < hi && nums[hi] == nums[hi+1] {
				hi--
			}
		} else if sum < target {
			lo++
		} else {
			hi--
		}
	}
}

func threeSum(nums []int, target int, res *[][]int, z1 int) {
	max := nums[len(nums)-1]
	if 3*nums[0] > target || 3*max < target {
		return
	}

	for i := 0; i < len(nums)-2; i++ {
		z := nums[i]
		if i > 0 && z == nums[i-1] {
			// avoid duplicates
			continue
		}

		if z+2*max < target {
			// z is too small
			continue
		}

		if 3*z > target {
			// z is too large
			break
		}

		if 3*z == target {
			// z is the last possible candidate
			if i+2 < len(nums) && nums[i+2] == z {
				*res = append(*res, []int{z1, z, z, z})
			}
			break
		}

		twoSum(nums[i+1:], target-z, res, z1, z)
	}
}

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)

	if len(nums) < 4 {
		return res
	}

	qsort(nums)

	max := nums[len(nums)-1]
	if 4*nums[0] > target || 4*max < target {
		return res
	}

	for i := 0; i < len(nums)-3; i++ {
		z := nums[i]
		if i > 0 && z == nums[i-1] {
			// avoid duplicates
			continue
		}

		if z+3*max < target {
			// z is too small
			continue
		}

		if 4*z > target {
			// z is too large
			break
		}

		if 4*z == target {
			// z is the last possible candidate
			// check if the array contains such duplicates
			if i+3 < len(nums) && nums[i+3] == z {
				res = append(res, []int{z, z, z, z})
			}
			break
		}

		threeSum(nums[i+1:], target-z, &res, z)
	}

	return res
}
