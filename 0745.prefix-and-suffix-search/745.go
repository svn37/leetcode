/**
 * Design a special dictionary which has some words and allows you to search the words in it by a prefix and a suffix.
 * Implement the WordFilter class:
 *
 * 	WordFilter(string[] words) Initializes the object with the words in the dictionary.
 * 	f(string prefix, string suffix) Returns the index of the word in the dictionary which has the prefix prefix and the suffix suffix. If there is more than one valid index, return the largest of them. If there is no such word in the dictionary, return -1.
 *
 *
 * Example 1:
 *
 * Input
 * ["WordFilter", "f"]
 * [[["apple"]], ["a", "e"]]
 * Output
 * [null, 0]
 * Explanation
 * WordFilter wordFilter = new WordFilter(["apple"]);
 * wordFilter.f("a", "e"); // return 0, because the word at index 0 has prefix = "a" and suffix = 'e".
 *
 *
 * Constraints:
 *
 * 	1 <= words.length <= 15000
 * 	1 <= words[i].length <= 10
 * 	1 <= prefix.length, suffix.length <= 10
 * 	words[i], prefix and suffix consist of lower-case English letters only.
 * 	At most 15000 calls will be made to the function f.
 *
 */

package leetcode

import "bytes"

type TrieNode struct {
	weight   int
	children []*TrieNode
}

func NewTrie() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 27),
	}
}

func (t *TrieNode) Insert(word string, weight int) {
	node := t
	for _, c := range word {
		if node.children[int(c)-'a'] == nil {
			newNode := NewTrie()
			newNode.weight = weight
			node.children[int(c)-'a'] = newNode
		} else {
			// update weight
			node.children[int(c)-'a'].weight = weight
		}
		node = node.children[int(c)-'a']
	}
}

func (t *TrieNode) Filter(prefix string) int {
	node := t
	for _, c := range prefix {
		if node.children[int(c)-'a'] == nil {
			return -1
		}
		node = node.children[int(c)-'a']
	}
	return node.weight
}

type WordFilter struct {
	trie  *TrieNode
	words []string
}

func Constructor(words []string) WordFilter {
	wf := WordFilter{
		trie:  NewTrie(),
		words: words,
	}

	for weight, word := range words {
		for i := 0; i <= len(word); i++ {
			var b bytes.Buffer

			b.WriteString(word[i:])
			b.WriteRune('{') // goes after 'z' in ascii table
			b.WriteString(word)

			wf.trie.Insert(b.String(), weight)
		}
	}

	return wf
}

func (this *WordFilter) F(prefix string, suffix string) int {
	search := suffix + "{" + prefix
	return this.trie.Filter(search)
}
