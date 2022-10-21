/**
 * Design a data structure that supports adding new words and finding if a string matches any previously added string.
 * Implement the WordDictionary class:
 *
 * 	WordDictionary() Initializes the object.
 * 	void addWord(word) Adds word to the data structure, it can be matched later.
 * 	bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.
 *
 *
 * Example:
 *
 * Input
 * ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
 * [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
 * Output
 * [null,null,null,null,false,true,true,true]
 * Explanation
 * WordDictionary wordDictionary = new WordDictionary();
 * wordDictionary.addWord("bad");
 * wordDictionary.addWord("dad");
 * wordDictionary.addWord("mad");
 * wordDictionary.search("pad"); // return False
 * wordDictionary.search("bad"); // return True
 * wordDictionary.search(".ad"); // return True
 * wordDictionary.search("b.."); // return True
 *
 *
 * Constraints:
 *
 * 	1 <= word.length <= 500
 * 	word in addWord consists lower-case English letters.
 * 	word in search consist of  '.' or lower-case English letters.
 * 	At most 50000 calls will be made to addWord and search.
 *
 */

package leetcode

type TrieNode struct {
	children []*TrieNode
	isWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 26),
	}
}

func (node *TrieNode) Insert(word string) {
	for _, char := range word {
		c := int(char - 'a')
		newNode := node.children[c]
		if newNode == nil {
			newNode = NewTrieNode()
			node.children[c] = newNode
		}
		node = newNode
	}
	node.isWord = true
}

func (node *TrieNode) Search(word string) bool {
	for i, char := range word {
		c := int(char - 'a')
		if char != '.' {
			node = node.children[c]
			if node == nil {
				return false
			}
		} else {
			for _, charNode := range node.children {
				if charNode != nil && charNode.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
	}
	return node.isWord
}

type WordDictionary struct {
	trie *TrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		trie: NewTrieNode(),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search(word)
}
