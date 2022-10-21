/**
 * Given a string S, consider all duplicated substrings: (contiguous) substrings of S that occur 2 or more times.  (The occurrences may overlap.)
 *
 * Return any duplicated substring that has the longest possible length.  (If S does not have a duplicated substring, the answer is "".)
 *
 *
 *
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">"banana"</span>
 * Output: <span id="example-output-1">"ana"</span>
 *
 *
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">"abcd"</span>
 * Output: <span id="example-output-2">""</span>
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	2 <= S.length <= 10^5
 * 	S consists of lowercase English letters.
 * </ol>
 *
 */

package leetcode

// The obvious method is to build the suffix tree,
// which requires a template of Ukkonen's O(n) algorithm,
// or any other O(log n) algorithm.

// This method uses binary search of the correct length,
// and rabin-karp for scanning the string for substrings of the same length,
// which might be duplicates.
const (
	base1 = 1e9 + 7
	base2 = 16777619
)

func hash(s string, base uint32) uint32 {
	var h uint32
	for i := range s {
		h = (h*base + uint32(s[i]))
	}
	return h
}

// modified rabin-karp, instead of searching for pattern,
// checks if we have seen the same hash before
func test(str string, pos int) int {
	seen1 := make(map[uint32]struct{})
	seen2 := make(map[uint32]struct{})

	var mult1 uint32 = 1 // mult = base^(pos-1)
	var mult2 uint32 = 1
	for i := 0; i < pos-1; i++ {
		mult1 = (mult1 * base1)
		mult2 = (mult2 * base2)
	}

	h1 := hash(str[:pos], base1)
	h2 := hash(str[:pos], base2)
	seen1[h1] = struct{}{}
	seen2[h2] = struct{}{}

	for i := pos; i < len(str); i++ {
		h1 = h1 - mult1*uint32(str[i-pos])
		h1 = h1*base1 + uint32(str[i])

		h2 = h2 - mult2*uint32(str[i-pos])
		h2 = h2*base2 + uint32(str[i])

		if _, ok := seen1[h1]; ok {
			if _, ok := seen2[h2]; ok {
				return i - pos + 1
			}
		}
		seen1[h1] = struct{}{}
		seen2[h2] = struct{}{}
	}
	return -1
}

func longestDupSubstring(str string) string {
	lo, hi := 0, len(str)
	res := 0
	// binary search of the length of the string
	for lo < hi {
		mid := (lo + hi + 1) / 2
		pos := test(str, mid)
		if pos > 0 {
			lo = mid
			res = pos
		} else {
			hi = mid - 1
		}
	}
	return str[res : res+lo]
}
