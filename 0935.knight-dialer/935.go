/**
 * The chess knight has a unique movement, it may move two squares vertically and one square horizontally, or two squares horizontally and one square vertically (with both forming the shape of an L). The possible movements of chess knight are shown in this diagaram:
 * A chess knight can move as indicated in the chess diagram below:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/18/chess.jpg" style="width: 402px; height: 402px;" />
 * We have a chess knight and a phone pad as shown below, the knight can only stand on a numeric cell (i.e. blue cell).
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/18/phone.jpg" style="width: 242px; height: 322px;" />
 * Given an integer n, return how many distinct phone numbers of length n we can dial.
 * You are allowed to place the knight on any numeric cell initially and then you should perform n - 1 jumps to dial a number of length n. All jumps should be valid knight jumps.
 * As the answer may be very large, return the answer modulo 10^9 + 7.
 *
 * Example 1:
 *
 * Input: n = 1
 * Output: 10
 * Explanation: We need to dial a number of length 1, so placing the knight over any numeric cell of the 10 cells is sufficient.
 *
 * Example 2:
 *
 * Input: n = 2
 * Output: 20
 * Explanation: All the valid number we can dial are [04, 06, 16, 18, 27, 29, 34, 38, 40, 43, 49, 60, 61, 67, 72, 76, 81, 83, 92, 94]
 *
 * Example 3:
 *
 * Input: n = 3
 * Output: 46
 *
 * Example 4:
 *
 * Input: n = 4
 * Output: 104
 *
 * Example 5:
 *
 * Input: n = 3131
 * Output: 136006598
 * Explanation: Please take care of the mod.
 *
 *
 * Constraints:
 *
 * 	1 <= n <= 5000
 *
 */

package leetcode

const mod = 1e9 + 7

func knightDialer(N int) int {
	// init unique numbers for each key for 0 hops
	keys := make([]int, 10)
	for i := range keys {
		keys[i] = 1
	}

	for N > 1 {
		keys[0], keys[1], keys[2], keys[3], keys[4],
			keys[5], keys[6], keys[7], keys[8], keys[9] = (keys[4]+keys[6])%mod, // 0
			(keys[6]+keys[8])%mod, // 1
			(keys[7]+keys[9])%mod, // 2
			(keys[4]+keys[8])%mod, // 3
			((keys[3]+keys[9])%mod+keys[0])%mod, // 4
			// if hops > 0, knight can't print even 5 itself,
			// because it must make at least one hop
			0, // 5
			((keys[1]+keys[7])%mod+keys[0])%mod, // 6
			(keys[2]+keys[6])%mod, // 7
			(keys[1]+keys[3])%mod, // 8
			(keys[2]+keys[4])%mod // 9

		N--
	}

	sum := 0
	for i := range keys {
		sum += keys[i]
		sum %= mod
	}
	return sum
}
