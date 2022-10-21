/**
 * Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.
 * Notice that you may not slant the container.
 *
 * Example 1:
 * <img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg" style="width: 600px; height: 287px;" />
 * Input: height = [1,8,6,2,5,4,8,3,7]
 * Output: 49
 * Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
 *
 * Example 2:
 *
 * Input: height = [1,1]
 * Output: 1
 *
 * Example 3:
 *
 * Input: height = [4,3,2,1,4]
 * Output: 16
 *
 * Example 4:
 *
 * Input: height = [1,2,1]
 * Output: 2
 *
 *
 * Constraints:
 *
 * 	n = height.length
 * 	2 <= n <= 3 * 10^4
 * 	0 <= height[i] <= 3 * 10^4
 *
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Two pointer approach.
// The area formed between the lines will always be limited by the height of the shorter line.
// Further, the farther the lines, the more will be the area obtained.

// 1 .The widest container (using first and last line) is a good candidate, because of its width.
// Its water level is the height of the smaller one of first and last line.
// 2. All other containers are less wide and thus would need a higher water level in order to hold more water.
// 3. The smaller one of first and last line doesn't support a higher water level and can thus be safely removed from further consideration.
func maxArea(height []int) int {
	maxArea := 0

	left, right := 0, len(height)-1
	for left < right {
		maxArea = max(maxArea, min(height[left], height[right])*(right-left))

		if height[left] < height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		} else {
			left++
			right--
		}
	}

	return maxArea
}
