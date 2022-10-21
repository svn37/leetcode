/**
 * Design a Skiplist without using any built-in libraries.
 *
 * A Skiplist is a data structure that takes O(log(n)) time to add, erase and search. Comparing with treap and red-black tree which has the same function and performance, the code length of Skiplist can be comparatively short and the idea behind Skiplists are just simple linked lists.
 *
 * For example: we have a Skiplist containing [30,40,50,60,70,90] and we want to add 80 and 45 into it. The Skiplist works this way:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/09/27/1506_skiplist.gif" style="width: 960px; height: 332px;" /><br />
 * <small>Artyom Kalinin [CC BY-SA 3.0], via <a href="https://commons.wikimedia.org/wiki/File:Skip_list_add_element-en.gif" target="_blank" title="Artyom Kalinin [CC BY-SA 3.0 (https://creativecommons.org/licenses/by-sa/3.0)], via Wikimedia Commons">Wikimedia Commons</a></small>
 *
 * You can see there are many layers in the Skiplist. Each layer is a sorted linked list. With the help of the top layers, add , erase and search can be faster than O(n). It can be proven that the average time complexity for each operation is O(log(n)) and space complexity is O(n).
 *
 * To be specific, your design should include these functions:
 *
 *
 * 	bool search(int target) : Return whether the target exists in the Skiplist or not.
 * 	void add(int num): Insert a value into the SkipList.
 * 	bool erase(int num): Remove a value in the Skiplist. If num does not exist in the Skiplist, do nothing and return false. If there exists multiple num values, removing any one of them is fine.
 *
 *
 * See more about Skiplist : <a href="https://en.wikipedia.org/wiki/Skip_list" target="_blank">https://en.wikipedia.org/wiki/Skip_list</a>
 *
 * Note that duplicates may exist in the Skiplist, your code needs to handle this situation.
 *
 *
 *
 * Example:
 *
 *
 * Skiplist skiplist = new Skiplist();
 *
 * skiplist.add(1);
 * skiplist.add(2);
 * skiplist.add(3);
 * skiplist.search(0);   // return false.
 * skiplist.add(4);
 * skiplist.search(1);   // return true.
 * skiplist.erase(0);    // return false, 0 is not in skiplist.
 * skiplist.erase(1);    // return true.
 * skiplist.search(1);   // return false, 1 has already been erased.
 *
 *
 * Constraints:
 *
 *
 * 	0 <= num, target <= 20000
 * 	At most 50000 calls will be made to search, add, and erase.
 *
 */

package leetcode

import (
	"math"
	"math/rand"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Node struct {
	val    int
	levels []*Node
}

func NewNode(val, levels int) *Node {
	return &Node{
		val:    val,
		levels: make([]*Node, levels),
	}
}

func (node *Node) find(target, level int) *Node {
	var prev *Node
	for node != nil && node.val < target {
		prev = node
		node = node.levels[level]
	}

	return prev
}

type Skiplist struct {
	head *Node
}

func Constructor() Skiplist {
	return Skiplist{NewNode(math.MinInt64, 16)}
}

func (this *Skiplist) Search(target int) bool {
	cur := this.head
	for level := 15; level >= 0; level-- {
		node := cur.find(target, level)
		if node == nil {
			continue
		}

		next := node.levels[level]
		if next != nil && next.val == target {
			return true
		}

		cur = node
	}

	return false
}

func (this *Skiplist) Add(num int) {
	levelsNum := min(16, 1+int(math.Log2(1.0/rand.Float64())))
	newNode := NewNode(num, levelsNum)

	cur := this.head
	for level := 15; level >= 0; level-- {
		node := cur.find(num, level)
		if node != nil {
			cur = node
		}

		if level < levelsNum {
			newNode.levels[level] = cur.levels[level]
			cur.levels[level] = newNode
		}
	}
}

func (this *Skiplist) Erase(num int) bool {
	cur := this.head
	var found bool
	for level := 15; level >= 0; level-- {
		node := cur.find(num, level)
		if node == nil {
			continue
		}

		cur = node
		next := node.levels[level]

		if next != nil && next.val == num {
			found = true

			node.levels[level] = next.levels[level]
		}
	}

	return found
}
