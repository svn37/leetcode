/**
 * All DNA is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T', for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.
 * Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.
 *
 * Example 1:
 * Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
 * Output: ["AAAAACCCCC","CCCCCAAAAA"]
 * Example 2:
 * Input: s = "AAAAAAAAAAAAA"
 * Output: ["AAAAAAAAAA"]
 *
 * Constraints:
 *
 * 	0 <= s.length <= 10^5
 * 	s[i] is 'A', 'C', 'G', or 'T'.
 *
 */

package leetcode

func findRepeatedDnaSequences(s string) []string {
	words := make(map[uint32]int)

	res := []string{}
	// need only 4 chars instead of 26, but it's simpler
	chars := make([]uint32, 26)

	// chars['A'-'A'] = 0
	chars['C'-'A'] = 1
	chars['G'-'A'] = 2
	chars['T'-'A'] = 3

	for i := 0; i < len(s)-9; i++ {
		var v uint32
		for j := i; j < i+10; j++ {
			v <<= 2
			v |= chars[int(s[j]-'A')]
		}

		if words[v] == 1 {
			res = append(res, s[i:i+10])
		}
		words[v]++
	}
	return res
}
