/**
 *
 * In this problem, a rooted tree is a directed graph such that, there is exactly one node (the root) for which all other nodes are descendants of this node, plus every node has exactly one parent, except for the root node which has no parents.
 *
 * The given input is a directed graph that started as a rooted tree with N nodes (with distinct values 1, 2, ..., N), with one additional directed edge added.  The added edge has two different vertices chosen from 1 to N, and was not an edge that already existed.
 *
 * The resulting graph is given as a 2D-array of edges.  Each element of edges is a pair [u, v] that represents a directed edge connecting nodes u and v, where u is a parent of child v.
 *
 * Return an edge that can be removed so that the resulting graph is a rooted tree of N nodes.  If there are multiple answers, return the answer that occurs last in the given 2D-array.
 * Example 1:<br />
 *
 * Input: [[1,2], [1,3], [2,3]]
 * Output: [2,3]
 * Explanation: The given directed graph will be like this:
 *   1
 *  / \
 * v   v
 * 2-->3
 *
 *
 * Example 2:<br />
 *
 * Input: [[1,2], [2,3], [3,4], [4,1], [1,5]]
 * Output: [4,1]
 * Explanation: The given directed graph will be like this:
 * 5 <- 1 -> 2
 *      ^    |
 *      |    v
 *      4 <- 3
 *
 *
 * Note:<br />
 * The size of the input 2D-array will be between 3 and 1000.
 * Every integer represented in the 2D-array will be between 1 and N, where N is the size of the input array.
 *
 */

package leetcode

type UnionFind struct {
	count  int
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{
		count:  n,
		parent: parent,
		rank:   make([]int, n),
	}
}

func (u *UnionFind) Find(p int) int {
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]] // path compression by halving
		p = u.parent[p]
	}
	return p
}

func (u *UnionFind) Union(p, q int) bool {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return false
	}
	if u.rank[rootQ] > u.rank[rootP] {
		u.parent[rootP] = rootQ
	} else {
		u.parent[rootQ] = rootP
		if u.rank[rootQ] == u.rank[rootP] {
			u.rank[rootP]++
		}
	}
	u.count--
	return true
}

func findRedundantConnection(edges [][]int, skipEdge []int) []int {
	u := NewUnionFind(len(edges))
	for i := range edges {
		if skipEdge != nil && edges[i][0] == skipEdge[0] && edges[i][1] == skipEdge[1] {
			continue
		}
		if !u.Union(edges[i][0]-1, edges[i][1]-1) {
			return edges[i]
		}
	}
	return nil
}

//  1. Check whether there is a node having two parents.
//     If so, store them as candidates A and B, and set the second edge invalid.
//  2. Perform normal union find.
//     If the tree is now valid
//     simply return candidate B
//     else if candidates not existing
//     we find a circle, return current edge;
//     else
//     remove candidate A instead of B.
func findRedundantDirectedConnection(edges [][]int) []int {
	incoming := make([]int, len(edges)+1)
	var candidateA, candidateB []int

	for i := range edges {
		if incoming[edges[i][1]] != 0 {
			candidateA = []int{edges[i][0], edges[i][1]}
			candidateB = []int{incoming[edges[i][1]], edges[i][1]}
			break
		}
		incoming[edges[i][1]] = edges[i][0]
	}

	if candidateA == nil {
		return findRedundantConnection(edges, nil)
	} else {
		res := findRedundantConnection(edges, candidateA)
		if res == nil {
			return candidateA
		}
		return candidateB
	}
}
