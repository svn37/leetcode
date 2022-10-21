/**
 * You are given a string s, and an array of pairs of indices in the string pairs where pairs[i] = [a, b] indicates 2 indices(0-indexed) of the string.
 * You can swap the characters at any pair of indices in the given pairs any number of times.
 * Return the lexicographically smallest string that s can be changed to after using the swaps.
 *
 * Example 1:
 *
 * Input: s = "dcab", pairs = [[0,3],[1,2]]
 * Output: "bacd"
 * Explaination:
 * Swap s[0] and s[3], s = "bcad"
 * Swap s[1] and s[2], s = "bacd"
 *
 * Example 2:
 *
 * Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
 * Output: "abcd"
 * Explaination:
 * Swap s[0] and s[3], s = "bcad"
 * Swap s[0] and s[2], s = "acbd"
 * Swap s[1] and s[2], s = "abcd"
 * Example 3:
 *
 * Input: s = "cba", pairs = [[0,1],[1,2]]
 * Output: "abc"
 * Explaination:
 * Swap s[0] and s[1], s = "bca"
 * Swap s[1] and s[2], s = "bac"
 * Swap s[0] and s[1], s = "abc"
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 10^5
 * 	0 <= pairs.length <= 10^5
 * 	0 <= pairs[i][0], pairs[i][1] < s.length
 * 	s only contains lower case English letters.
 *
 */

package leetcode

import (
	"bytes"
	"math/rand"
)

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

func qsort(arr []byte) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivot := rand.Intn(len(arr))

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	qsort(arr[:left])
	qsort(arr[left+1:])
}

func smallestStringWithSwaps(str string, pairs [][]int) string {
	u := NewUnionFind(len(str))

	for _, pair := range pairs {
		u.Union(pair[0], pair[1])
	}

	m := make(map[int][]byte)
	for i := range str {
		idx := u.Find(i)
		m[idx] = append(m[idx], str[i])
	}

	for i := range m {
		qsort(m[i])
	}

	var res bytes.Buffer
	for i := range str {
		idx := u.Find(i)

		char := m[idx][0]
		m[idx] = m[idx][1:]

		res.WriteByte(char)
	}
	return res.String()
}
