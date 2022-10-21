/**
 * Given an array of strings products and a string searchWord. We want to design a system that suggests at most three product names from products after each character of searchWord is typed. Suggested products should have common prefix with the searchWord. If there are more than three products with a common prefix return the three lexicographically minimums products.
 * Return list of lists of the suggested products after each character of searchWord is typed.
 *
 * Example 1:
 *
 * Input: products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
 * Output: [
 * ["mobile","moneypot","monitor"],
 * ["mobile","moneypot","monitor"],
 * ["mouse","mousepad"],
 * ["mouse","mousepad"],
 * ["mouse","mousepad"]
 * ]
 * Explanation: products sorted lexicographically = ["mobile","moneypot","monitor","mouse","mousepad"]
 * After typing m and mo all products match and we show user ["mobile","moneypot","monitor"]
 * After typing mou, mous and mouse the system suggests ["mouse","mousepad"]
 *
 * Example 2:
 *
 * Input: products = ["havana"], searchWord = "havana"
 * Output: [["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
 *
 * Example 3:
 *
 * Input: products = ["bags","baggage","banner","box","cloths"], searchWord = "bags"
 * Output: [["baggage","bags","banner"],["baggage","bags","banner"],["baggage","bags"],["bags"]]
 *
 * Example 4:
 *
 * Input: products = ["havana"], searchWord = "tatiana"
 * Output: [[],[],[],[],[],[],[]]
 *
 *
 * Constraints:
 *
 * 	1 <= products.length <= 1000
 * 	There are no repeated elements in products.
 * 	1 <= &Sigma; products[i].length <= 2 * 10^4
 * 	All characters of products[i] are lower-case English letters.
 * 	1 <= searchWord.length <= 1000
 * 	All characters of searchWord are lower-case English letters.
 *
 */

package leetcode

import (
	"sort"
	"strings"
)

type Trie struct {
	children [26]*Trie
	words    []string
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
		node.words = append(node.words, word)
	}
}

// hardcoded 3 candidates
func (trie *Trie) StartsWith(prefix string) [][]string {
	node := trie
	words := make([][]string, len(prefix))
	for i, char := range prefix {
		nextNode := node.children[char-'a']
		if nextNode == nil {
			break
		}
		node = nextNode
		if len(node.words) < 3 {
			words[i] = node.words
		} else {
			words[i] = node.words[:3]
		}
	}
	return words
}

// Method 1. Build a trie of words, with a slight change - each node includes candidates
// could use indices into the array
func suggestedProducts(products []string, searchWord string) [][]string {
	// sort the products so that results are put in the trie in lexicographical order during trie build
	// if products are allowed to change, it wouldn't work
	sort.Strings(products)
	trie := NewTrie()
	for _, product := range products {
		trie.Insert(product)
	}
	return trie.StartsWith(searchWord)
}

func bsearch(a []string, target string) int {
	lo, hi := 0, len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if strings.Compare(a[mid], target) >= 0 {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// Method 2. Use binary search in the sorted array each time we input a new character
func suggestedProducts2(products []string, searchWord string) [][]string {
	// sort the products
	sort.Strings(products)

	res := make([][]string, len(searchWord))
	for i := range searchWord {
		prefix := searchWord[:i+1]
		k := bsearch(products, prefix)
		for k > 0 && prefix == products[k-1] { // in case there are more than 1 cur in products
			k-- // find the first one
		}
		if k < 0 {
			break
		}
		for limit := k + 3; k < len(products) && k < limit && strings.HasPrefix(products[k], prefix); k++ {
			res[i] = append(res[i], products[k])
		}
	}
	return res
}
