/**
 * You are given a grid of size N x N, and each cell of this grid has a lamp that is initially turned off.
 * You are also given an array of lamp positions lamps, where lamps[i] = [rowi, coli] indicates that the lamp at grid[rowi][coli] is turned on. When a lamp is turned on, it illuminates its cell and all other cells in the same row, column, or diagonal.
 * Finally, you are given a query array queries, where queries[i] = [rowi, coli]. For the i^th query, determine whether grid[rowi][coli] is illuminated or not. After answering the i^th query, turn off the lamp at grid[rowi][coli] and its 8 adjacent lamps if they exist. A lamp is adjacent if its cell shares either a side or corner with grid[rowi][coli].
 * Return an array of integers ans, where ans[i] should be 1 if the lamp in the i^th query was illuminated, or 0 if the lamp was not.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/19/illu_1.jpg" style="width: 750px; height: 209px;" />
 * Input: N = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,0]]
 * Output: [1,0]
 * Explanation: We have the initial grid with all lamps turned off. In the above picture we see the grid after turning on the lamp at grid[0][0] then turning on the lamp at grid[4][4].
 * The 0^th query asks if the lamp at grid[1][1] is illuminated or not (the blue square). It is illuminated, so set ans[0] = 1. Then, we turn off all lamps in the red square.
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/19/illu_step1.jpg" style="width: 500px; height: 218px;" />
 * The 1^st query asks if the lamp at grid[1][0] is illuminated or not (the blue square). It is not illuminated, so set ans[1] = 1. Then, we turn off all lamps in the red rectangle.
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/19/illu_step2.jpg" style="width: 500px; height: 219px;" />
 *
 * Example 2:
 *
 * Input: N = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,1]]
 * Output: [1,1]
 *
 * Example 3:
 *
 * Input: N = 5, lamps = [[0,0],[0,4]], queries = [[0,4],[0,1],[1,4]]
 * Output: [1,1,0]
 *
 *
 * Constraints:
 *
 * 	1 <= N <= 10^9
 * 	0 <= lamps.length <= 20000
 * 	lamps[i].length == 2
 * 	0 <= lamps[i][j] < N
 * 	0 <= queries.length <= 20000
 * 	queries[i].length == 2
 * 	0 <= queries[i][j] < N
 *
 */

package leetcode

func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	res := make([]int, len(queries))

	if len(queries) == 0 {
		return res
	}

	m1 := make(map[int]int)  // row number to count of lamps
	m2 := make(map[int]int)  // col number to count of lamps
	m3 := make(map[int]int)  // diagonal x-y to count of lamps
	m4 := make(map[int]int)  // diagonal x+y to count of lamps
	m5 := make(map[int]bool) // whether lamp is on/off

	// increment counters while adding lamps
	for i := range lamps {
		x := lamps[i][0]
		y := lamps[i][1]
		m1[x]++
		m2[y]++
		m3[x-y]++
		m4[x+y]++
		m5[N*x+y] = true
	}

	dirs := []int{0, 0, 1, 1, 0, -1, -1, 1, -1, 0}

	// answer queries
	for i := range queries {
		x := queries[i][0]
		y := queries[i][1]

		if m1[x] > 0 || m2[y] > 0 || m3[x-y] > 0 || m4[x+y] > 0 || m5[N*x+y] {
			res[i] = 1
		}

		// switch off the lamps, if any
		for d := 0; d < 9; d++ {
			dx := x + dirs[d]
			dy := y + dirs[d+1]
			if m5[N*dx+dy] {
				m1[dx]--
				m2[dy]--
				m3[dx-dy]--
				m4[dx+dy]--
				m5[N*dx+dy] = false
			}
		}
	}
	return res
}
