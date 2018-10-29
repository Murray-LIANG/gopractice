
package gopractice

import "strings"


// https://leetcode.com/problems/longest-absolute-file-path/description/

func lengthLongestPath(input string) int {

	result := 0
	dirLength := map[int]int {0 : 0}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		name := strings.TrimLeft(line, "\t")
		depth := (len(line) - len(name))

		if strings.Contains(name, ".") {
			if len(name) + dirLength[depth] > result {
				result = len(name) + dirLength[depth]
			}
		} else {
			dirLength[depth+1] = dirLength[depth] + len(name) + 1
		}
	}
	return result
}