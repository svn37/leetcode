/**
 * Given N axis-aligned rectangles where N > 0, determine if they all together form an exact cover of a rectangular region.
 *
 * Each rectangle is represented as a bottom-left point and a top-right point. For example, a unit square is represented as [1,1,2,2]. (coordinate of bottom-left point is (1, 1) and top-right point is (2, 2)).
 *
 * <div style="float:right"><img src="https://assets.leetcode.com/uploads/2018/10/22/rectangle_perfect.gif" style="width: 249px; height: 250px;" /></div>
 *
 * Example 1:
 *
 *
 * rectangles = [
 *   [1,1,3,3],
 *   [3,1,4,2],
 *   [3,2,4,4],
 *   [1,3,2,4],
 *   [2,3,3,4]
 * ]
 *
 * Return true. All 5 rectangles together form an exact cover of a rectangular region.
 *
 *
 *
 *
 * <div style="clear:both"> </div>
 *
 * <div style="float:right"><img src="https://assets.leetcode.com/uploads/2018/10/22/rectangle_separated.gif" style="width: 249px; height: 250px;" /></div>
 *
 * Example 2:
 *
 *
 * rectangles = [
 *   [1,1,2,3],
 *   [1,3,2,4],
 *   [3,1,4,2],
 *   [3,2,4,4]
 * ]
 *
 * Return false. Because there is a gap between the two rectangular regions.
 *
 *
 *
 *
 * <div style="clear:both"> </div>
 *
 * <div style="float:right"><img src="https://assets.leetcode.com/uploads/2018/10/22/rectangle_hole.gif" style="width: 249px; height: 250px;" /></div>
 *
 * Example 3:
 *
 *
 * rectangles = [
 *   [1,1,3,3],
 *   [3,1,4,2],
 *   [1,3,2,4],
 *   [3,2,4,4]
 * ]
 *
 * Return false. Because there is a gap in the top center.
 *
 *
 *
 *
 * <div style="clear:both"> </div>
 *
 * <div style="float:right"><img src="https://assets.leetcode.com/uploads/2018/10/22/rectangle_intersect.gif" style="width: 249px; height: 250px;" /></div>
 *
 * Example 4:
 *
 *
 * rectangles = [
 *   [1,1,3,3],
 *   [3,1,4,2],
 *   [1,3,2,4],
 *   [2,2,4,4]
 * ]
 *
 * Return false. Because two of the rectangles overlap with each other.
 *
 *
 *
 */

package leetcode

import (
	"fmt"
	"math"
)

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

func addOrRemove(m map[string]struct{}, x, y int) bool {
	str := fmt.Sprintf("%d %d", x, y)
	if _, ok := m[str]; ok {
		delete(m, str)
		return true
	}
	m[str] = struct{}{}
	return false
}

func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) == 0 || len(rectangles[0]) == 0 {
		return false
	}

	// corner points of the big rectangle
	x1 := math.MaxInt64
	x2 := math.MinInt64
	y1 := math.MaxInt64
	y2 := math.MinInt64

	area := 0
	m := make(map[string]struct{})

	for _, rectangle := range rectangles {
		x1 = min(x1, rectangle[0])
		y1 = min(y1, rectangle[1])
		x2 = max(x2, rectangle[2])
		y2 = max(y2, rectangle[3])

		area += (rectangle[2] - rectangle[0]) * (rectangle[3] - rectangle[1])

		addOrRemove(m, rectangle[0], rectangle[1])
		addOrRemove(m, rectangle[0], rectangle[3])
		addOrRemove(m, rectangle[2], rectangle[1])
		addOrRemove(m, rectangle[2], rectangle[3])
	}

	return len(m) == 4 && addOrRemove(m, x1, y1) && addOrRemove(m, x1, y2) && addOrRemove(m, x2, y1) && addOrRemove(m, x2, y2) && area == (x2-x1)*(y2-y1)
}
