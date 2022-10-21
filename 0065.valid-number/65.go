/**
 * Validate if a given string can be interpreted as a decimal number.
 *
 * Some examples:<br />
 * "0" => true<br />
 * " 0.1 " => true<br />
 * "abc" => false<br />
 * "1 a" => false<br />
 * "2e10" => true<br />
 * " -90e3   " => true<br />
 * " 1e" => false<br />
 * "e3" => false<br />
 * " 6e-1" => true<br />
 * " 99e2.5 " => false<br />
 * "53.5e93" => true<br />
 * " --6 " => false<br />
 * "-+3" => false<br />
 * "95a54e53" => false
 *
 * Note: It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one. However, here is a list of characters that can be in a valid decimal number:
 *
 *
 * 	Numbers 0-9
 * 	Exponent - "e"
 * 	Positive/negative sign - "+"/"-"
 * 	Decimal point - "."
 *
 *
 * Of course, the context of these characters also matters in the input.
 *
 * Update (2015-02-10):<br />
 * The signature of the C++ function had been updated. If you still see your function signature accepts a const char * argument, please click the reload button to reset your code definition.
 *
 */

package leetcode

func isNumber(s string) bool {
	// trim the string
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	s = s[i:]

	j := len(s)
	for j > 0 && s[j-1] == ' ' {
		j--
	}
	s = s[:j]

	pointSeen := false
	eSeen := false
	numSeen := false
	numAfterE := true

	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			numSeen = true
			numAfterE = true
		} else if s[i] == '.' {
			if eSeen || pointSeen {
				return false
			}
			pointSeen = true
		} else if s[i] == 'e' {
			if eSeen || !numSeen {
				return false
			}
			numAfterE = false
			eSeen = true
		} else if s[i] == '-' || s[i] == '+' {
			if i != 0 && s[i-1] != 'e' {
				return false
			}
		} else {
			return false
		}
	}

	return numSeen && numAfterE
}
