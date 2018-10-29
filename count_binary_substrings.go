package gopractice

//https://leetcode.com/problems/count-binary-substrings

func countBinarySubstrings(s string) int {
	result := 0

	for i, _ := range(s) {
		countI := 1
		change := false
		for j := i+1; j < len(s); j++ {
			if s[j] == s[i] {
				if change {
					break
				}
				countI++
			} else {
				change = true
				countI--
			}
			if countI == 0 {
				result++
				break
			}
		}
	}
	return result
}

func countBinarySubstrings2(s string) int {
	result, preCount, countI := 0, 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			countI++
		} else {
			preCount = countI
			countI = 1
		}
		if preCount >= countI {
			result++
		}
	}
	return result
}