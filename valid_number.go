// https://leetcode.com/problems/valid-number
package gopractice

import "strings"

func isNumber(s string) bool {
	s = strings.TrimSpace(s)

	pointSeen, eSeen, numberSeen, numberAfterE := false, false, false, true
	
	for index, ch := range s {
		if '0' <= ch && ch <= '9' {
			numberSeen = true
			numberAfterE = true
		} else if ch == '.' {
			if eSeen || pointSeen {
				return false
			}
			pointSeen = true
		} else if ch == 'e' {
			if eSeen || !numberSeen {
				return false
			}
			eSeen = true
			numberAfterE = false
		} else if ch == '-' || ch == '+' {
			if index != 0 && s[index-1] != 'e' {
				return false
			}
		} else {
			return false
		}
	}
	return numberSeen && numberAfterE
}