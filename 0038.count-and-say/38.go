/**
 * The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
 *
 * 	countAndSay(1) = "1"
 * 	countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.
 *
 * To determine how you "say" a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character. Then for each group, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.
 * For example, the saying and conversion for digit string "3322251":
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/23/countandsay.jpg" style="width: 581px; height: 172px;" />
 * Given a positive integer n, return the n^th term of the count-and-say sequence.
 *
 * Example 1:
 *
 * Input: n = 1
 * Output: "1"
 * Explanation: This is the base case.
 *
 * Example 2:
 *
 * Input: n = 4
 * Output: "1211"
 * Explanation:
 * countAndSay(1) = "1"
 * countAndSay(2) = say "1" = one 1 = "11"
 * countAndSay(3) = say "11" = two 1's = "21"
 * countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"
 *
 *
 * Constraints:
 *
 * 	1 <= n <= 30
 *
 */

package leetcode

func countAndSay(n int) string {
	seq := []byte{'1'}
	temp := make([]byte, 0)

	for n != 1 {
		var count byte
		count++
		for i := 1; i <= len(seq); i++ {
			if i == len(seq) || seq[i-1] != seq[i] {
				temp = append(temp, '0'+count, seq[i-1])
				count = 0
			}
			count++
		}
		seq, temp = temp, seq[:0]
		n--
	}

	return string(seq)
}
