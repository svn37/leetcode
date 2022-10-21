/**
 * Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
 * A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
 * <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/7/73/Telephone-keypad2.svg/200px-Telephone-keypad2.svg.png" style="width: 200px; height: 162px;" />
 *
 * Example 1:
 *
 * Input: digits = "23"
 * Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 * Example 2:
 *
 * Input: digits = ""
 * Output: []
 *
 * Example 3:
 *
 * Input: digits = "2"
 * Output: ["a","b","c"]
 *
 *
 * Constraints:
 *
 * 	0 <= digits.length <= 4
 * 	digits[i] is a digit in the range ['2', '9'].
 *
 */

package leetcode

func digitsToLetters() [][]byte {
	mapping := make([][]byte, 8)

	k := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 3; j++ {
			mapping[i] = append(mapping[i], byte(3*i+j+k+97))
		}
		if i == 5 || i == 7 {
			mapping[i] = append(mapping[i], byte(3*i+3+k+97))
			k = 1
		}
	}

	return mapping
}

func generate(digits string, idx int, comb string, mapping [][]byte, res *[]string) {
	if idx == len(digits) {
		*res = append(*res, comb)
		return
	}

	letters := mapping[int(digits[idx]-'2')]
	for i := range letters {
		generate(digits, idx+1, comb+string(letters[i]), mapping, res)
	}
}

func letterCombinations(digits string) []string {
	mapping := digitsToLetters()

	res := []string{}
	if len(digits) == 0 {
		return res
	}
	generate(digits, 0, "", mapping, &res)
	return res
}
