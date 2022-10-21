/**
 * Implement the MapSum class:
 *
 * 	MapSum() Initializes the MapSum object.
 * 	void insert(String key, int val) Inserts the key-val pair into the map. If the key already existed, the original key-value pair will be overridden to the new one.
 * 	int sum(string prefix) Returns the sum of all the pairs' value whose key starts with the prefix.
 *
 *
 * Example 1:
 *
 * Input
 * ["MapSum", "insert", "sum", "insert", "sum"]
 * [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
 * Output
 * [null, null, 3, null, 5]
 * Explanation
 * MapSum mapSum = new MapSum();
 * mapSum.insert("apple", 3);
 * mapSum.sum("ap");           // return 3 (<u>ap</u>ple = 3)
 * mapSum.insert("app", 2);
 * mapSum.sum("ap");           // return 5 (<u>ap</u>ple + <u>ap</u>p = 3 + 2 = 5)
 *
 *
 * Constraints:
 *
 * 	1 <= key.length, prefix.length <= 50
 * 	key and prefix consist of only lowercase English letters.
 * 	1 <= val <= 1000
 * 	At most 50 calls will be made to insert and sum.
 *
 */

package leetcode

type TrieNode struct {
	weight   int
	children []*TrieNode
}

func NewTrie() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 26),
	}
}

func (t *TrieNode) Insert(word string, weight int) {
	node := t
	for _, c := range word {
		if node.children[int(c)-'a'] == nil {
			newNode := NewTrie()
			node.children[int(c)-'a'] = newNode
		}
		// update weight
		node.children[int(c)-'a'].weight += weight

		node = node.children[int(c)-'a']
	}
}

func (t *TrieNode) Weight(prefix string) int {
	node := t
	for _, c := range prefix {
		if node.children[int(c)-'a'] == nil {
			return 0
		}
		node = node.children[int(c)-'a']
	}
	return node.weight
}

type MapSum struct {
	trie  *TrieNode
	words map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		trie:  NewTrie(),
		words: make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	delta := val
	if weight, ok := this.words[key]; ok {
		delta -= weight
	} else {
		this.words[key] = val
	}

	this.trie.Insert(key, delta)
}

func (this *MapSum) Sum(prefix string) int {
	return this.trie.Weight(prefix)
}
