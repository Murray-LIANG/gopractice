package gopractice

// https://leetcode.com/problems/combination-sum

import "sort"

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	var backtrace func([]int, int, int)

	backtrace = func(tempResult []int, remain int, startIndex int) {
		if remain < 0 {
			return
		} else if remain == 0 {
			copy := []int{}
			for _, candidate := range tempResult {
				copy = append(copy, candidate)
			}
			result = append(result, copy)
		} else {
			for i := startIndex; i < len(candidates); i++ {
				tempResult = append(tempResult, candidates[i])
				backtrace(tempResult, remain - candidates[i], i)
				tempResult = tempResult[:len(tempResult)-1]
			}
		}
	}

	sort.Sort(IntSlice(candidates))
	backtrace([]int{}, target, 0)
	return result
}