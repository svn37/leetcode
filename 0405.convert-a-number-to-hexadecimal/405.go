/**
 *
 * Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, <a href="https://en.wikipedia.org/wiki/Two%27s_complement" target="_blank">two’s complement</a> method is used.
 *
 *
 * Note:
 * <ol>
 * All letters in hexadecimal (a-f) must be in lowercase.
 * The hexadecimal string must not contain extra leading 0s. If the number is zero, it is represented by a single zero character '0'; otherwise, the first character in the hexadecimal string will not be the zero character.
 * The given number is guaranteed to fit within the range of a 32-bit signed integer.
 * You must not use any method provided by the library which converts/formats the number to hex directly.
 * </ol>
 *
 *
 * Example 1:
 *
 * Input:
 * 26
 *
 * Output:
 * "1a"
 *
 *
 *
 * Example 2:
 *
 * Input:
 * -1
 *
 * Output:
 * "ffffffff"
 *
 *
 */

package leetcode

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	chars := "0123456789abcdef"

	str := string(chars[num&0xf])
	num = num >> 4 & 0x0fffffff // -

	for num != 0 {
		str = string(chars[num&0xf]) + str
		num >>= 4
	}
	return str
}

// more simple to understand solution
func toHex2(num int) string {
	if num == 0 {
		return "0"
	}
	chars := "0123456789abcdef"

	res := ""
	for num != 0 {
		res = string(chars[num&15]) + res
		// unsigned shift
		num = int(uint32(num) >> 4)
	}
	return res
}
