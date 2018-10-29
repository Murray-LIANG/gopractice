package gopractice

// https://leetcode.com/problems/compare-version-numbers

import "strings"
import "strconv"

func compareVersion(version1 string, version2 string) int {
	max := func (x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	list1, list2 := strings.Split(version1, "."), strings.Split(version2, ".")
	
	size := max(len(list1), len(list2))
	for i := 0 ; i < size; i++ {
		int1, int2 := 0, 0
		if i < len(list1) {
			int1, _ = strconv.Atoi(list1[i])
		}
		if i < len(list2) {
			int2, _ = strconv.Atoi(list2[i])
		}
		if int1 < int2 {
			return -1
		} else if int1 > int2 {
			return 1
		}
	}
	return 0
}