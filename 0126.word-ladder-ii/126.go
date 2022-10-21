/**
 * Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that:
 *
 * <ol>
 * 	Only one letter can be changed at a time
 * 	Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
 * </ol>
 *
 * Note:
 *
 *
 * 	Return an empty list if there is no such transformation sequence.
 * 	All words have the same length.
 * 	All words contain only lowercase alphabetic characters.
 * 	You may assume no duplicates in the word list.
 * 	You may assume beginWord and endWord are non-empty and are not the same.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * Output:
 * [
 *   ["hit","hot","dot","dog","cog"],
 *   ["hit","hot","lot","log","cog"]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * Output: []
 *
 * Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
 *
 *
 *
 *
 *
 */

package leetcode

func findEndWordByBFS(
	beginWord, endWord string,
	wordList []string,
	children map[string][]string,
) bool {
	words := make(map[string]struct{})
	for i := range wordList {
		words[wordList[i]] = struct{}{}
	}

	// if end word is not in the wordList
	if _, ok := words[endWord]; !ok {
		return false
	}

	beginSet := make(map[string]struct{})
	endSet := make(map[string]struct{})

	beginSet[beginWord] = struct{}{}
	endSet[endWord] = struct{}{}

	var reversed, found bool

	for len(beginSet) > 0 && len(endSet) > 0 && !found {
		for word := range beginSet {
			delete(words, word)
		}

		for word := range endSet {
			delete(words, word)
		}

		if len(beginSet) > len(endSet) {
			reversed = !reversed
			beginSet, endSet = endSet, beginSet
		}

		temp := make(map[string]struct{})

		for word := range beginSet {
			w := []rune(word)

			for i := range word {
				oldChar := w[i]

				for newChar := 'a'; newChar <= 'z'; newChar++ {
					if newChar == oldChar {
						continue
					}

					w[i] = newChar
					newWord := string(w)

					if _, ok := endSet[newWord]; ok {
						// form the connection between one graph and another
						if reversed {
							list := children[newWord]
							children[newWord] = append(list, word)
						} else {
							list := children[word]
							children[word] = append(list, newWord)
						}
						found = true
					}

					if _, ok := words[newWord]; ok {
						if reversed {
							list := children[newWord]
							children[newWord] = append(list, word)
						} else {
							list := children[word]
							children[word] = append(list, newWord)
						}

						temp[newWord] = struct{}{}
					}
					w[i] = oldChar
				}
			}
		}
		beginSet = temp
	}

	return found
}

// use DFS (backtracking) to generate the transformation sequences
// according to the children mapping
func buildLaddersByDFS(
	beginWord, endWord string,
	ladder []string,
	ladders *[][]string,
	children map[string][]string,
) {
	if beginWord == endWord {
		*ladders = append(*ladders, ladder)
		return
	}

	// beginWord is guaranteed to be in the map
	for _, word := range children[beginWord] {
		newLadder := append(ladder[:0:0], ladder...)
		newLadder = append(newLadder, word)

		buildLaddersByDFS(word, endWord, newLadder, ladders, children)
	}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	children := make(map[string][]string)
	if !findEndWordByBFS(beginWord, endWord, wordList, children) {
		return nil
	}

	ladders := [][]string{}
	ladder := []string{beginWord}

	buildLaddersByDFS(beginWord, endWord, ladder, &ladders, children)
	return ladders
}
