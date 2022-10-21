/**
 * Given a string s containing only digits, return all possible valid IP addresses that can be obtained from s. You can return them in any order.
 * A valid IP address consists of exactly four integers, each integer is between 0 and 255, separated by single dots and cannot have leading zeros. For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses and "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
 *
 * Example 1:
 * Input: s = "25525511135"
 * Output: ["255.255.11.135","255.255.111.35"]
 * Example 2:
 * Input: s = "0000"
 * Output: ["0.0.0.0"]
 * Example 3:
 * Input: s = "1111"
 * Output: ["1.1.1.1"]
 * Example 4:
 * Input: s = "010010"
 * Output: ["0.10.0.10","0.100.1.0"]
 * Example 5:
 * Input: s = "101023"
 * Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 * Constraints:
 *
 * 	0 <= s.length <= 3000
 * 	s consists of digits only.
 *
 */

package leetcode

func validRemainingStr(strLeft, digits int) bool {
	min := digits
	max := digits * 3

	if strLeft > max || strLeft < min {
		return false
	}

	return true
}

func validNumber(num string) bool {
	// if the first character in a two- or more digit number
	// is zero, the string is not valid
	if len(num) > 1 && num[0] == 48 {
		return false
	}

	if len(num) == 3 {
		n := 0
		for _, char := range num {
			n *= 10
			n += int(char - '0')
		}

		if n > 255 {
			return false
		}
	}
	return true
}

func restore(s string, b, e int, digitsLeft int, comb string, res *[]string) {
	for i := b; i < e; i++ {
		if i+1 >= len(s) {
			if validRemainingStr(0, digitsLeft-1) && validNumber(s[b:]) {
				*res = append(*res, comb+s[b:])
			}
			return
		}

		if validRemainingStr(len(s[i+1:]), digitsLeft-1) && validNumber(s[b:i+1]) {
			restore(s, i+1, i+4, digitsLeft-1, comb+s[b:i+1]+".", res)
		}
	}
}

func restoreIpAddresses(s string) []string {
	res := []string{}
	if !validRemainingStr(len(s), 4) {
		return res
	}
	restore(s, 0, 3, 4, "", &res)
	return res
}
