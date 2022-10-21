/**
 * Convert a non-negative integer num to its English words representation.
 *
 * Example 1:
 * Input: num = 123
 * Output: "One Hundred Twenty Three"
 * Example 2:
 * Input: num = 12345
 * Output: "Twelve Thousand Three Hundred Forty Five"
 * Example 3:
 * Input: num = 1234567
 * Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
 * Example 4:
 * Input: num = 1234567891
 * Output: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"
 *
 * Constraints:
 *
 * 	0 <= num <= 2^31 - 1
 *
 */

package leetcode

import "strings"

var (
	thousands      []string = []string{"", "Thousand", "Million", "Billion"}
	lessThanTwenty []string = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens           []string = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
)

func helper(num int) string {
	if num == 0 {
		return ""
	} else if num < 20 {
		return lessThanTwenty[num] + " "
	} else if num < 100 {
		return tens[num/10] + " " + helper(num%10)
	} else {
		return lessThanTwenty[num/100] + " Hundred " + helper(num%100)
	}
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	words := ""
	i := 0

	for num > 0 {
		if num%1000 != 0 {
			words = helper(num%1000) + thousands[i] + " " + words
		}
		num /= 1000
		i++
	}

	return strings.Trim(words, " ")
}
