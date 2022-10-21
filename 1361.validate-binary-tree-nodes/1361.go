/**
 * You have n binary tree nodes numbered from 0 to n - 1 where node i has two children leftChild[i] and rightChild[i], return true if and only if all the given nodes form exactly one valid binary tree.
 * If node i has no left child then leftChild[i] will equal -1, similarly for the right child.
 * Note that the nodes have no values and that we only use the node numbers in this problem.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/08/23/1503_ex1.png" style="width: 195px; height: 287px;" />
 *
 * Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
 * Output: true
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/08/23/1503_ex2.png" style="width: 183px; height: 272px;" />
 *
 * Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]
 * Output: false
 *
 * Example 3:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/08/23/1503_ex3.png" style="width: 82px; height: 174px;" />
 *
 * Input: n = 2, leftChild = [1,0], rightChild = [-1,-1]
 * Output: false
 *
 * Example 4:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/08/23/1503_ex4.png" style="width: 470px; height: 191px;" />
 *
 * Input: n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1]
 * Output: false
 *
 *
 * Constraints:
 *
 * 	1 <= n <= 10^4
 * 	leftChild.length == rightChild.length == n
 * 	-1 <= leftChild[i], rightChild[i] <= n - 1
 *
 */

package leetcode

func dfs(cur int, leftChild []int, rightChild []int, m map[int]struct{}) bool {
	if cur == -1 {
		return true
	}
	if _, ok := m[cur]; ok {
		return false
	}
	m[cur] = struct{}{}
	return dfs(leftChild[cur], leftChild, rightChild, m) && dfs(rightChild[cur], leftChild, rightChild, m)
}

// the idea is:
// - find the root of the tree (there can be only one root)
// - check if there is a cycle
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	m := make(map[int]int)
	// root xor helps to find the root
	root := 0
	for i := 0; i < n; i++ {
		root ^= i
	}
	for i := range leftChild {
		if leftChild[i] >= 0 {
			if _, ok := m[leftChild[i]]; ok {
				return false
			}
			m[leftChild[i]] = i
			root ^= leftChild[i]
		}

		if rightChild[i] >= 0 {
			if _, ok := m[rightChild[i]]; ok {
				return false
			}
			m[rightChild[i]] = i
			root ^= rightChild[i]
		}
	}

	if len(m) != n-1 {
		return false
	}

	// check if there is a cycle
	return dfs(root, leftChild, rightChild, make(map[int]struct{}))
}
