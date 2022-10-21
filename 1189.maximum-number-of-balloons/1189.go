/**
 * Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.
 *
 * You can use each character in text at most once. Return the maximum number of instances that can be formed.
 *
 *
 * Example 1:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/09/05/1536_ex1_upd.JPG" style="width: 132px; height: 35px;" />
 *
 *
 * Input: text = "nlaebolko"
 * Output: 1
 *
 *
 * Example 2:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/09/05/1536_ex2_upd.JPG" style="width: 267px; height: 35px;" />
 *
 *
 * Input: text = "loonbalxballpoon"
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: text = "leetcode"
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 	1 <= text.length <= 10^4
 * 	text consists of lower case English letters only.
 *
 */

package leetcode

func maxNumberOfBalloons(text string) int {
	a := make([]int, 26)
	b := "balloon"
	c := 0

	for _, v := range text {
		a[v-'a']++
	}
	for {
		for _, v := range b {
			a[v-'a']--
			if a[v-'a'] < 0 {
				return c
			}
		}
		c++
	}
}
