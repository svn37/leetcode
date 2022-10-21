/**
 * You are given a string s and an array of strings words of the same length. Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.
 * You can return the answer in any order.
 *
 * Example 1:
 *
 * Input: s = "barfoothefoobarman", words = ["foo","bar"]
 * Output: [0,9]
 * Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
 * The output order does not matter, returning [9,0] is fine too.
 *
 * Example 2:
 *
 * Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
 * Output: []
 *
 * Example 3:
 *
 * Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
 * Output: [6,9,12]
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 10^4
 * 	s consists of lower-case English letters.
 * 	1 <= words.length <= 5000
 * 	1 <= words[i].length <= 30
 * 	words[i] consists of lower-case English letters.
 *
 */

package leetcode

// Method 1. Brute-force, try to find substring by starting at each index of s.
func findSubstring(s string, words []string) []int {
	indices := []int{}

	if len(s) == 0 || len(words) == 0 {
		return indices
	}

	count := make(map[string]int)
	for i := range words {
		count[words[i]]++
	}

	numWords := len(words)
	wordLen := len(words[0])

	for i := 0; i < len(s)-numWords*wordLen+1; i++ {
		seen := make(map[string]int)
		j := 0
		for j < numWords {
			word := s[i+j*wordLen : i+(j+1)*wordLen]
			seen[word]++
			if seen[word] > count[word] {
				// substring starting at i has an extra word (or more words that it is supposed to)
				break
			}
			j++
		}
		if j == numWords {
			indices = append(indices, i)
		}
	}
	return indices
}

// Method 2. Optimize searching for words by moving two pointers, right and left,
// and start searching only len(words[0]) times.
func findSubstring2(s string, words []string) []int {
	indices := []int{}
	if len(s) == 0 || len(words) == 0 {
		return indices
	}
	wordLen := len(words[0])

	// mapping from word to its count
	wordCount := make(map[string]int)
	for i := range words {
		wordCount[words[i]]++
	}

	// number of unique words
	numWords := len(wordCount)
	// position in s from which it doesn't make sense to start searching
	last := len(s) - wordLen + 1

	// start search only wordLen times, compared with previous solution
	// that is the main difference
	for i := 0; i < wordLen; i++ {
		seen := make(map[string]int, numWords)
		foundWords := 0

		left := i
		right := i

		for right < last {
			// move only right pointer
			// 1) until the end, if all words are not found
			// 2) stop if all words are found (that is why the outer for loop)
			for foundWords < numWords && right < last {
				word := s[right : right+wordLen]
				count := wordCount[word]
				if count > 0 {
					seen[word]++
					if seen[word] == count {
						foundWords++
					}
				}
				right += wordLen
			}

			// move only left pointer
			// 1) until left meets right
			// 2) stop if some words are missing
			for foundWords == numWords && left < right {
				word := s[left : left+wordLen]
				curCount := seen[word]
				if curCount > 0 {
					count := wordCount[word]
					// curCount might be larger than global count! in that case, simply decrease curCount
					if curCount == count {
						// if current window includes all the words exactly once
						if (right-left)/wordLen == len(words) {
							indices = append(indices, left)
						}
						foundWords--
					}
					seen[word]--
				}
				left += wordLen
			}
		}
	}
	return indices
}

// Method 3. The same idea as method 2, but with optimizations for word search.
// map each word to unique id
// preprocess s and map each index to word unique id if this word was found at this index.
func findSubstring3(s string, words []string) []int {
	indices := []int{}
	if len(s) == 0 || len(words) == 0 {
		return indices
	}

	// mapping from word to its unique id
	mapping := make(map[string]int)
	// unique id is used as index
	// at most len(words), some entries at the end will be empty
	count := make([]int, len(words))

	// unique id
	id := 0

	// map each word to some index
	// count each word
	for i := range words {
		wordId, ok := mapping[words[i]]
		if !ok {
			mapping[words[i]] = id
			wordId = id
			// increment epoch
			id++
		}
		count[wordId]++
	}

	wordLen := len(words[0])
	// number of unique words
	numWords := len(mapping)

	// map each substring of length wordLen to their unique id,
	// if it occurs in words, or -1 if it doesn't

	// no substring exists after len(s)-wordLen+1
	substrings := make([]int, len(s)-wordLen+1)
	for i := range substrings {
		word := s[i : i+wordLen]
		wordId, ok := mapping[word]
		if ok {
			substrings[i] = wordId
		} else {
			substrings[i] = -1
		}
	}

	// fix the number of linear scans
	for i := 0; i < wordLen; i++ {
		curCount := make([]int, numWords)

		left := i
		right := i
		// its max value is numWords
		// incremented when word was met for count[wordId] times
		seen := 0

		for right < len(substrings) {
			for seen < numWords && right < len(substrings) {
				wordId := substrings[right]
				if wordId != -1 {
					curCount[wordId]++
					if curCount[wordId] == count[wordId] {
						seen++
					}
				}
				right += wordLen
			}

			for seen == numWords && left < right {
				wordId := substrings[left]
				if wordId != -1 {
					if curCount[wordId] == count[wordId] {
						// if current window includes all the words exactly once
						if (right-left)/wordLen == len(words) {
							indices = append(indices, left)
						}
						seen--
					}
					curCount[wordId]--
				}
				left += wordLen
			}
		}
	}
	return indices
}
