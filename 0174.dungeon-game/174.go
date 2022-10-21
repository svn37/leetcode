/**
 * <style type="text/css">table.dungeon, .dungeon th, .dungeon td {
 *   border:3px solid black;
 * }
 *
 *  .dungeon th, .dungeon td {
 *     text-align: center;
 *     height: 70px;
 *     width: 70px;
 * }
 * </style>
 * The demons had captured the princess (P) and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid out in a 2D grid. Our valiant knight (K) was initially positioned in the top-left room and must fight his way through the dungeon to rescue the princess.
 *
 * The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.
 *
 * Some of the rooms are guarded by demons, so the knight loses health (negative integers) upon entering these rooms; other rooms are either empty (0's) or contain magic orbs that increase the knight's health (positive integers).
 *
 * In order to reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.
 *
 *
 *
 * Write a function to determine the knight's minimum initial health so that he is able to rescue the princess.
 *
 * For example, given the dungeon below, the initial health of the knight must be at least 7 if he follows the optimal path RIGHT-> RIGHT -> DOWN -> DOWN.
 *
 * <table class="dungeon">
 * 	<tbody>
 * 		<tr>
 * 			<td>-2 (K)</td>
 * 			<td>-3</td>
 * 			<td>3</td>
 * 		</tr>
 * 		<tr>
 * 			<td>-5</td>
 * 			<td>-10</td>
 * 			<td>1</td>
 * 		</tr>
 * 		<tr>
 * 			<td>10</td>
 * 			<td>30</td>
 * 			<td>-5 (P)</td>
 * 		</tr>
 * 	</tbody>
 * </table>
 *
 *
 *
 * Note:
 *
 *
 * 	The knight's health has no upper bound.
 * 	Any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.
 *
 *
 */

package leetcode

import "math"

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

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return -1
	}
	memo := make([]int, len(dungeon[0])+1)
	for i := range memo {
		memo[i] = math.MaxInt64
	}
	memo[len(dungeon[0])-1] = 1

	for i := len(dungeon) - 1; i >= 0; i-- {
		for j := len(dungeon[0]) - 1; j >= 0; j-- {
			memo[j] = min(memo[j], memo[j+1]) - dungeon[i][j]
			memo[j] = max(1, memo[j])
		}
	}

	return memo[0]
}
