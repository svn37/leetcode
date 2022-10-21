/**
 * Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.
 * An interleaving of two strings s and t is a configuration where they are divided into non-empty substrings such that:
 *
 * 	s = s1 + s2 + ... + sn
 * 	t = t1 + t2 + ... + tm
 * 	|n - m| <= 1
 * 	The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
 *
 * Note: a + b is the concatenation of strings a and b.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/09/02/interleave.jpg" style="width: 561px; height: 203px;" />
 * Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 * Output: true
 *
 * Example 2:
 *
 * Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 * Output: false
 *
 * Example 3:
 *
 * Input: s1 = "", s2 = "", s3 = ""
 * Output: true
 *
 *
 * Constraints:
 *
 * 	0 <= s1.length, s2.length <= 100
 * 	0 <= s3.length <= 200
 * 	s1, s2, and s3 consist of lower-case English letters.
 *
 */

package leetcode

// Method 1. Dynamic programming.
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// 2D dynamic programming
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	for i := 0; i < len(s1)+1; i++ {
		for j := 0; j < len(s2)+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i-1]
			} else {
				dp[i][j] = (dp[i][j-1] && s2[j-1] == s3[i+j-1]) || (dp[i-1][j] && s1[i-1] == s3[i+j-1])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

// Method 2. 1D dynamic programming.
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// 1D dynamic programming
	dp := make([]bool, len(s2)+1)
	dp[0] = true
	for j := 1; j < len(dp); j++ {
		dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
		// can break early, but uglier
	}

	for i := 1; i < len(s1)+1; i++ {
		dp[0] = dp[0] && s1[i-1] == s3[i-1]

		for j := 1; j < len(s2)+1; j++ {
			dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) || (dp[j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[len(s2)]
}

// Method 3. Breadth-first search.
type Queue struct {
	items []Cell
}

func (q *Queue) Push(val Cell) {
	q.items = append(q.items, val)
}

func (q *Queue) PopFront() Cell {
	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

type Cell struct {
	x, y int
}

// if s1 = "aab" and s2 = "abc" and s3 = "aaabcb"
// graph looks like this
// we can only move down and right

// o--a--o--b--o--c--o
// |     |     |     |
// a     a     a     a
// |     |     |     |
// o--a--o--b--o--c--o
// |     |     |     |
// a     a     a     a
// |     |     |     |
// o--a--o--b--o--c--o
// |     |     |     |
// b     b     b     b
// |     |     |     |
// o--a--o--b--o--c--o

func isInterleave3(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	queue := &Queue{}
	visited := make(map[Cell]struct{})

	queue.Push(Cell{0, 0})

	for !queue.Empty() {
		cell := queue.PopFront()

		// successfully interleaved the string
		if cell.x == len(s1) && cell.y == len(s2) {
			return true
		}

		if _, ok := visited[cell]; ok {
			continue
		}
		visited[cell] = struct{}{}

		// down
		if cell.x < len(s1) && s1[cell.x] == s3[cell.x+cell.y] {
			queue.Push(Cell{cell.x + 1, cell.y})
		}

		// right
		if cell.y < len(s2) && s2[cell.y] == s3[cell.x+cell.y] {
			queue.Push(Cell{cell.x, cell.y + 1})
		}
	}

	return false
}
