/**
 * We are given a 2-dimensional grid. "." is an empty cell, "#" is a wall, "@" is the starting point, ("a", "b", ...) are keys, and ("A", "B", ...) are locks.
 *
 * We start at the starting point, and one move consists of walking one space in one of the 4 cardinal directions.  We cannot walk outside the grid, or walk into a wall.  If we walk over a key, we pick it up.  We can't walk over a lock unless we have the corresponding key.
 *
 * For some <font face="monospace">1 <= K <= 6</font>, there is exactly one lowercase and one uppercase letter of the first K letters of the English alphabet in the grid.  This means that there is exactly one key for each lock, and one lock for each key; and also that the letters used to represent the keys and locks were chosen in the same order as the English alphabet.
 *
 * Return the lowest number of moves to acquire all keys.  If it's impossible, return -1.
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">["@.a.#","###.#","b.A.B"]</span>
 * Output: <span id="example-output-1">8</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">["@..aA","..B#.","....b"]</span>
 * Output: <span id="example-output-2">6</span>
 *
 * </div>
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= grid.length <= 30
 * 	1 <= grid[0].length <= 30
 * 	grid[i][j] contains only '.', '#', '@', 'a'-'f' and 'A'-'F'
 * 	The number of keys is in [1, 6].  Each key has a different letter and opens exactly one lock.
 * </ol>
 * </div>
 *
 */

package leetcode

import "fmt"

type Queue struct {
	items []*State
}

func (q *Queue) Push(val *State) {
	q.items = append(q.items, val)
}

func (q *Queue) Poll() *State {
	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

type State struct {
	keys, x, y int
}

func max(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

func shortestPathAllKeys(grid []string) int {
	var numKeys uint8
	x, y := -1, -1
	m := len(grid)
	n := len(grid[0])

	// preprocess the grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			char := grid[i][j]
			// find the starting point
			if char == '@' {
				x = i
				y = j
			}
			// find how many locks and keys the grid contains
			if char >= 'a' && char <= 'f' {
				numKeys = max(numKeys, char-'a'+1)
			}
		}
	}

	start := &State{0, x, y}
	queue := &Queue{}
	visited := make(map[string]struct{})

	queue.Push(start)
	str := fmt.Sprintf("0 %d %d", x, y)
	visited[str] = struct{}{}

	dirs := []int{0, 1, 0, -1, 0}
	size, level := 0, 0

	for !queue.Empty() {
		if size == 0 {
			size = queue.Len()
			level++
		}

		state := queue.Poll()

		for i := 0; i < 4; i++ {
			dx := state.x + dirs[i]
			dy := state.y + dirs[i+1]
			// copy the keys, because they are used in four directions
			keys := state.keys

			if dx >= 0 && dx < m && dy >= 0 && dy < n {
				char := grid[dx][dy]
				// wall
				if char == '#' {
					continue
				}

				if char >= 'a' && char <= 'f' {
					// set the bit, we found the key
					keys |= 1 << (char - 'a')
				}

				// we found all the keys
				if keys == (1<<numKeys)-1 {
					return level
				}

				if char >= 'A' && char <= 'F' && keys>>(char-'A')&1 == 0 {
					// found the lock, but don't have the key
					continue
				}

				str := fmt.Sprintf("%d %d %d", state.keys, dx, dy)
				if _, ok := visited[str]; !ok {
					visited[str] = struct{}{}
					queue.Push(&State{keys, dx, dy})
				}
			}
		}

		size--
	}

	return -1
}
