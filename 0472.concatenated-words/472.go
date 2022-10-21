/**
 * Given a list of words (without duplicates), please write a program that returns all concatenated words in the given list of words.
 * A concatenated word is defined as a string that is comprised entirely of at least two shorter words in the given array.
 *
 * Example:<br />
 *
 * Input: ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
 *
 * Output: ["catsdogcats","dogcatsdog","ratcatdogcat"]
 *
 * Explanation: "catsdogcats" can be concatenated by "cats", "dog" and "cats"; <br> "dogcatsdog" can be concatenated by "dog", "cats" and "dog"; <br>"ratcatdogcat" can be concatenated by "rat", "cat", "dog" and "cat".
 *
 *
 *
 * Note:<br>
 * <ol>
 * The number of elements of the given array will not exceed 10,000
 * The length sum of elements in the given array will not exceed 600,000.
 * All the input string will only include lower case letters.
 * The returned elements order does not matter.
 * </ol>
 *
 */

package leetcode

import "math/rand"

func qsort(a []string) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivot := rand.Intn(len(a))

	a[right], a[pivot] = a[pivot], a[right]

	for i := range a {
		if len(a[i]) < len(a[right]) {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	qsort(a[:left])
	qsort(a[left+1:])
}

func canForm(word string, prevSet map[string]struct{}) bool {
	if len(prevSet) == 0 {
		return false
	}

	dp := make([]bool, len(word)+1)
	dp[0] = true // something to start with

	for i := 1; i <= len(word); i++ {
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}

			if _, ok := prevSet[word[j:i]]; ok {
				dp[i] = true
				// very important break
				break
			}
		}
	}

	return dp[len(word)]
}

func findAllConcatenatedWordsInADict(words []string) []string {
	qsort(words)

	res := []string{}
	prevSet := make(map[string]struct{})

	for _, word := range words {
		if canForm(word, prevSet) {
			res = append(res, word)
		}
		prevSet[word] = struct{}{}
	}

	return res
}
