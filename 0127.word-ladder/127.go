/**
 * Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:
 * <ol>
 * 	Only one letter can be changed at a time.
 * 	Each transformed word must exist in the word list.
 * </ol>
 * Note:
 *
 * 	Return 0 if there is no such transformation sequence.
 * 	All words have the same length.
 * 	All words contain only lowercase alphabetic characters.
 * 	You may assume no duplicates in the word list.
 * 	You may assume beginWord and endWord are non-empty and are not the same.
 *
 * Example 1:
 *
 * Input:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 * Output: 5
 * Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
 * return its length 5.
 *
 * Example 2:
 *
 * Input:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 * Output: 0
 * Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
 *
 *
 *
 */

package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// add words to the set
	words := make(map[string]struct{})
	for i := range wordList {
		words[wordList[i]] = struct{}{}
	}

	if _, ok := words[endWord]; !ok {
		return 0
	}

	// init bidirectional bfs search structures
	beginSet := make(map[string]struct{})
	endSet := make(map[string]struct{})
	visited := make(map[string]struct{})

	// add source and target
	beginSet[beginWord] = struct{}{}
	endSet[endWord] = struct{}{}

	// length of the shortest path
	l := 1

	for len(beginSet) > 0 && len(endSet) > 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

		temp := make(map[string]struct{})

		for wordStr := range beginSet {
			word := []rune(wordStr)

			for i := range word {
				for newChar := 'a'; newChar <= 'z'; newChar++ {
					if word[i] == newChar {
						continue
					}

					oldChar := word[i]
					word[i] = newChar
					newStr := string(word)

					if _, ok := endSet[newStr]; ok {
						return l + 1
					}

					if _, ok := visited[newStr]; !ok {
						if _, ok := words[newStr]; ok {
							temp[newStr] = struct{}{}
							visited[newStr] = struct{}{}
						}
					}
					word[i] = oldChar
				}
			}
		}

		beginSet = temp
		l++
	}
	return 0
}
