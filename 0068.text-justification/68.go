/**
 * Given an array of words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.
 *
 * You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.
 *
 * Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.
 *
 * For the last line of text, it should be left justified and no extra space is inserted between words.
 *
 * Note:
 *
 *
 * 	A word is defined as a character sequence consisting of non-space characters only.
 * 	Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
 * 	The input array words contains at least one word.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * words = ["This", "is", "an", "example", "of", "text", "justification."]
 * maxWidth = 16
 * Output:
 * [
 *    "This    is    an",
 *    "example  of text",
 *    "justification.  "
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * words = ["What","must","be","acknowledgment","shall","be"]
 * maxWidth = 16
 * Output:
 * [
 *   "What   must   be",
 *   "acknowledgment  ",
 *   "shall be        "
 * ]
 * Explanation: Note that the last line is "shall be    " instead of "shall     be",
 *              because the last line must be left-justified instead of fully-justified.
 *              Note that the second line is also left-justified becase it contains only one word.
 *
 *
 * Example 3:
 *
 *
 * Input:
 * words = ["Science","is","what","we","understand","well","enough","to","explain",
 *          "to","a","computer.","Art","is","everything","else","we","do"]
 * maxWidth = 20
 * Output:
 * [
 *   "Science  is  what we",
 *   "understand      well",
 *   "enough to explain to",
 *   "a  computer.  Art is",
 *   "everything  else  we",
 *   "do                  "
 * ]
 *
 *
 */

package leetcode

import "bytes"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func fullJustify(words []string, maxWidth int) []string {
	numWords := make([]int, 0)
	extraSpaces := make([]int, 0)

	curWidth := 0
	prevIdx := 0
	for i, word := range words {
		if curWidth+len(word) > maxWidth {
			numWords = append(numWords, i-prevIdx)
			// + one space, because we don't count the trailing space
			extraSpaces = append(extraSpaces, maxWidth-curWidth+1)
			curWidth = 0
			prevIdx = i
		}
		// + one space between the words
		curWidth += len(word) + 1
	}

	result := make([]string, len(numWords)+1)
	start := 0

	for i := 0; i <= len(numWords); i++ {
		var str bytes.Buffer
		if i < len(numWords) {
			spaces := extraSpaces[i] / max(numWords[i]-1, 1)
			extra := extraSpaces[i] % max(numWords[i]-1, 1)

			for j := start; j < start+numWords[i]-1; j++ {
				str.WriteString(words[j])
				str.WriteRune(' ')

				s := spaces
				if extra > 0 {
					s++
					extra--
				}
				for s != 0 {
					str.WriteRune(' ')
					s--
				}
			}
			str.WriteString(words[start+numWords[i]-1])

			start += numWords[i]
		} else {
			for j := start; j < len(words); j++ {
				str.WriteString(words[j])
				if j != len(words)-1 {
					str.WriteRune(' ')
				}
			}
		}

		for str.Len() != maxWidth {
			str.WriteRune(' ')
		}

		result[i] = str.String()
	}

	return result
}
