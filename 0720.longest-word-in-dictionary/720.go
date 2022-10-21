/**
 * Given a list of strings words representing an English Dictionary, find the longest word in words that can be built one character at a time by other words in words.  If there is more than one possible answer, return the longest word with the smallest lexicographical order.  If there is no answer, return the empty string.
 *
 * Example 1:<br />
 *
 * Input:
 * words = ["w","wo","wor","worl", "world"]
 * Output: "world"
 * Explanation:
 * The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
 *
 *
 *
 * Example 2:<br />
 *
 * Input:
 * words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
 * Output: "apple"
 * Explanation:
 * Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
 *
 *
 *
 * Note:
 * All the strings in the input will only contain lowercase letters.
 * The length of words will be in the range [1, 1000].
 * The length of words[i] will be in the range [1, 30].
 *
 */

package leetcode

import "sort"

type Trie struct {
	children [26]*Trie
	isWord   bool
}

func NewTrie() *Trie {
	return &Trie{}
}

func (trie *Trie) Insert(word string) {
	node := trie
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

func (trie *Trie) StartsWith(prefix string) (*Trie, bool) {
	node := trie
	for _, char := range prefix {
		nextNode := node.children[char-'a']
		if nextNode == nil {
			return nil, false
		}
		node = nextNode
	}
	return node, true
}

func longestWord(words []string) string {
	sort.Strings(words)
	trie := NewTrie()

	longest := ""
	for _, word := range words {
		if node, ok := trie.StartsWith(word[:len(word)-1]); ok {
			node.Insert(word[len(word)-1:])
			if len(longest) < len(word) {
				longest = word
			}
		}
	}
	return longest
}
