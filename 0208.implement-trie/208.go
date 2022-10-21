/**
 * Implement a trie with insert, search, and startsWith methods.
 *
 * Example:
 *
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // returns true
 * trie.search("app");     // returns false
 * trie.startsWith("app"); // returns true
 * trie.insert("app");
 * trie.search("app");     // returns true
 *
 *
 * Note:
 *
 *
 * 	You may assume that all inputs are consist of lowercase letters a-z.
 * 	All inputs are guaranteed to be non-empty strings.
 *
 *
 */

package leetcode

type Trie struct {
	children [26]*Trie
	isWord   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, char := range word {
		nextNode := node.children[char-'a']
		if nextNode == nil {
			nextNode = &Trie{}
			node.children[char-'a'] = nextNode
		}
		node = nextNode
	}
	node.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, char := range word {
		nextNode := node.children[char-'a']
		if nextNode == nil {
			return false
		}
		node = nextNode
	}
	return node.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, char := range prefix {
		nextNode := node.children[char-'a']
		if nextNode == nil {
			return false
		}
		node = nextNode
	}
	// the only difference from Search() method
	return true
}
