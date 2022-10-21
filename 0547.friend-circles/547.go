/**
 * There are N students in a class. Some of them are friends, while some are not. Their friendship is transitive in nature. For example, if A is a direct friend of B, and B is a direct friend of C, then A is an indirect friend of C. And we defined a friend circle is a group of students who are direct or indirect friends.
 * Given a N*N matrix M representing the friend relationship between students in the class. If M[i][j] = 1, then the ith and jth students are direct friends with each other, otherwise not. And you have to output the total number of friend circles among all the students.
 * Example 1:
 *
 * Input:
 * [[1,1,0],
 *  [1,1,0],
 *  [0,0,1]]
 * Output: 2
 * Explanation:The 0th and 1st students are direct friends, so they are in a friend circle.
 * The 2nd student himself is in a friend circle. So return 2.
 *
 *
 * Example 2:
 *
 * Input:
 * [[1,1,0],
 *  [1,1,1],
 *  [0,1,1]]
 * Output: 1
 * Explanation:The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends,
 * so the 0th and 2nd students are indirect friends. All of them are in the same friend circle, so return 1.
 *
 *
 * Constraints:
 *
 * 	1 <= N <= 200
 * 	M[i][i] == 1
 * 	M[i][j] == M[j][i]
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

func (u *UnionFind) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
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
}

func findCircleNum(M [][]int) int {
	m := len(M)
	u := NewUnionFind(m)
	for i := 1; i < m; i++ {
		for j := i; j >= 0; j-- {
			if M[i][j] == 1 {
				u.Union(i, j)
			}
		}
	}
	return u.count
}
