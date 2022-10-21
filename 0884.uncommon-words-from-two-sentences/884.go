/**
 * We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)
 *
 * A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.
 *
 * Return a list of all uncommon words.
 *
 * You may return the list in any order.
 *
 *
 *
 * <ol>
 * </ol>
 *
 * <div>
 * Example 1:
 *
 *
 * Input: A = <span id="example-input-1-1">"this apple is sweet"</span>, B = <span id="example-input-1-2">"this apple is sour"</span>
 * Output: <span id="example-output-1">["sweet","sour"]</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: A = <span id="example-input-2-1">"apple apple"</span>, B = <span id="example-input-2-2">"banana"</span>
 * Output: <span id="example-output-2">["banana"]</span>
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	0 <= A.length <= 200
 * 	0 <= B.length <= 200
 * 	A and B both contain only spaces and lowercase letters.
 * </ol>
 * </div>
 * </div>
 *
 */

package leetcode

func addWordsToMap(A []byte, m map[string]bool) {
	prev := 0
	for i := 0; i <= len(A); i++ {
		if (i == len(A) || A[i] == ' ') && i-prev > 0 {
			word := string(A[prev:i])
			if _, ok := m[word]; ok {
				m[word] = false
			} else {
				m[word] = true
			}
			prev = i + 1
		}
	}
}

func uncommonFromSentences(a string, b string) []string {
	m := make(map[string]bool)
	addWordsToMap([]byte(a), m)
	addWordsToMap([]byte(b), m)

	res := make([]string, 0)
	for word, val := range m {
		if val {
			res = append(res, word)
		}
	}
	return res
}
