/**
 * You are given a series of video clips from a sporting event that lasted T seconds.  These video clips can be overlapping with each other and have varied lengths.
 * Each video clip clips[i] is an interval: it starts at time clips[i][0] and ends at time clips[i][1].  We can cut these clips into segments freely: for example, a clip [0, 7] can be cut into segments [0, 1] + [1, 3] + [3, 7].
 * Return the minimum number of clips needed so that we can cut the clips into segments that cover the entire sporting event ([0, T]).  If the task is impossible, return -1.
 *
 * Example 1:
 *
 * Input: clips = <span id="example-input-1-1">[[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]]</span>, T = <span id="example-input-1-2">10</span>
 * Output: <span id="example-output-1">3</span>
 * Explanation:
 * We take the clips [0,2], [8,10], [1,9]; a total of 3 clips.
 * Then, we can reconstruct the sporting event as follows:
 * We cut [1,9] into segments [1,2] + [2,8] + [8,9].
 * Now we have segments [0,2] + [2,8] + [8,10] which cover the sporting event [0, 10].
 *
 * Example 2:
 *
 * Input: clips = <span id="example-input-2-1">[[0,1],[1,2]]</span>, T = <span id="example-input-2-2">5</span>
 * Output: <span id="example-output-2">-1</span>
 * Explanation:
 * We can't cover [0,5] with only [0,1] and [1,2].
 *
 * Example 3:
 *
 * Input: clips = <span id="example-input-3-1">[[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]]</span>, T = <span id="example-input-3-2">9</span>
 * Output: <span id="example-output-3">3</span>
 * Explanation:
 * We can take clips [0,4], [4,7], and [6,9].
 *
 * Example 4:
 *
 * Input: clips = <span id="example-input-4-1">[[0,4],[2,8]]</span>, T = <span id="example-input-4-2">5</span>
 * Output: <span id="example-output-4">2</span>
 * Explanation:
 * Notice you can have extra video after the event ends.
 *
 *
 * Constraints:
 *
 * 	1 <= clips.length <= 100
 * 	0 <= clips[i][0] <= clips[i][1] <= 100
 * 	0 <= T <= 100
 *
 */

package leetcode

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// We track our current stitching position (st). For each iteration, we check all overlapping clips, and pick the one that advances our stitching position the furthest.
// We sort our clips by the starting point. Since clips are sorted, we need to only analyze each clip once. For each round, we check all overlapping clips (clips[i][0] <= st) and advance our stitching position as far as we can (end = max(end, clips[i][1])).
// Return -1 if we cannot advance our stitching position (st == end).
func videoStitching(clips [][]int, T int) int {
	res := 0
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})

	for i, st, end := 0, 0, 0; st < T; res++ {
		for ; i < len(clips) && clips[i][0] <= st; i++ {
			end = max(end, clips[i][1])
		}
		if st == end {
			return -1
		}
		st = end
	}
	return res
}
