/**
 * There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.
 * Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]
 * Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?
 *
 * Example 1:
 *
 * Input: numCourses = 2, prerequisites = [[1,0]]
 * Output: true
 * Explanation: There are a total of 2 courses to take.
 *              To take course 1 you should have finished course 0. So it is possible.
 *
 * Example 2:
 *
 * Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
 * Output: false
 * Explanation: There are a total of 2 courses to take.
 *              To take course 1 you should have finished course 0, and to take course 0 you should
 *              also have finished course 1. So it is impossible.
 *
 *
 * Constraints:
 *
 * 	The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about <a href="https://www.khanacademy.org/computing/computer-science/algorithms/graph-representation/a/representing-graphs" target="_blank">how a graph is represented</a>.
 * 	You may assume that there are no duplicate edges in the input prerequisites.
 * 	1 <= numCourses <= 10^5
 *
 */

package leetcode

// DFS
func buildGraph(edgeList [][]int, n int) [][]int {
	g := make([][]int, n) // adjacency list
	for i := range edgeList {
		startEdge := edgeList[i][1]
		endEdge := edgeList[i][0]
		g[startEdge] = append(g[startEdge], endEdge)
	}
	return g
}

func acyclic(g [][]int, todo, done []bool, i int) bool {
	if todo[i] {
		// cycle
		return false
	}

	if done[i] {
		// not cycle, because done was already investigated
		return true
	}

	todo[i] = true
	done[i] = true

	for j := range g[i] {
		if !acyclic(g, todo, done, g[i][j]) {
			return false
		}
	}

	todo[i] = false
	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := buildGraph(prerequisites, numCourses)
	done := make([]bool, numCourses)
	todo := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !done[i] && !acyclic(g, todo, done, i) {
			return false
		}
	}

	return true
}

// BFS, Kahn's algorithm
func computeIndegrees(g [][]int) []int {
	indegrees := make([]int, len(g))
	for i := range g {
		adjList := g[i]
		for j := range adjList {
			indegrees[adjList[j]]++
		}
	}
	return indegrees
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	g := buildGraph(prerequisites, numCourses)
	indegrees := computeIndegrees(g)
	for range indegrees {
		j := 0
		for j < len(indegrees) {
			if indegrees[j] == 0 {
				break
			}
			j++
		}
		if j == len(indegrees) {
			return false
		}

		indegrees[j]--
		for k := range g[j] {
			indegrees[g[j][k]]--
		}
	}
	return true
}
