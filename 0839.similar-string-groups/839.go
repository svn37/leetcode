/**
 * Two strings X and Y are similar if we can swap two letters (in different positions) of X, so that it equals Y. Also two strings X and Y are similar if they are equal.
 * For example, "tars" and "rats" are similar (swapping at positions 0 and 2), and "rats" and "arts" are similar, but "star" is not similar to "tars", "rats", or "arts".
 * Together, these form two connected groups by similarity: {"tars", "rats", "arts"} and {"star"}.  Notice that "tars" and "arts" are in the same group even though they are not similar.  Formally, each group is such that a word is in the group if and only if it is similar to at least one other word in the group.
 * We are given a list strs of strings where every string in strs is an anagram of every other string in strs. How many groups are there?
 *
 * Example 1:
 *
 * Input: strs = ["tars","rats","arts","star"]
 * Output: 2
 *
 * Example 2:
 *
 * Input: strs = ["omv","ovm"]
 * Output: 1
 *
 *
 * Constraints:
 *
 * 	1 <= strs.length <= 100
 * 	1 <= strs[i].length <= 1000
 * 	sum(strs[i].length) <= 2 * 10^4
 * 	strs[i] consists of lowercase letters only.
 * 	All words in strs have the same length and are anagrams of each other.
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

// it works because of the condition 5
// all words in strings have the same length and are anagrams of each other
func similar(a, b string) bool {
	n := 0
	for i := range a {
		if a[i] != b[i] {
			n++
			if n > 2 {
				return false
			}
		}
	}
	return true
}

func numSimilarGroups(strings []string) int {
	u := NewUnionFind(len(strings))

	for i := range strings {
		for j := i + 1; j < len(strings); j++ {
			if similar(strings[i], strings[j]) {
				u.Union(i, j)
			}
		}
	}

	return u.count
}
