/**
 * Given an array of strings strs, group the anagrams together. You can return the answer in any order.
 * An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
 *
 * Example 1:
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 * Example 2:
 * Input: strs = [""]
 * Output: [[""]]
 * Example 3:
 * Input: strs = ["a"]
 * Output: [["a"]]
 *
 * Constraints:
 *
 * 	1 <= strs.length <= 10^4
 * 	0 <= strs[i].length <= 100
 * 	strs[i] consists of lower-case English letters.
 *
 */

package leetcode

import "bytes"

// Use alphabet array as key
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		counter := [26]int{}
		for _, char := range str {
			counter[char-'a']++
		}
		m[counter] = append(m[counter], str)
	}

	res := make([][]string, 0, len(m))
	for _, anagrams := range m {
		res = append(res, anagrams)
	}
	return res
}

// More conventional counting sort
func countingSort(str string) string {
	counter := make([]int, 26)
	for _, char := range str {
		counter[char-'a']++
	}
	var newStr bytes.Buffer
	for i := 0; i < 26; i++ {
		for counter[i] > 0 {
			newStr.WriteByte(byte(i) + 'a')
			counter[i]--
		}
	}
	return newStr.String()
}

func groupAnagrams2(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		sortedStr := countingSort(str)
		anagrams := m[sortedStr]
		m[sortedStr] = append(anagrams, str)
	}

	res := make([][]string, 0, len(m))
	for _, anagrams := range m {
		res = append(res, anagrams)
	}
	return res
}
