package gopractice

// https://leetcode.com/problems/summary-ranges

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	nums = append(nums, nums[len(nums)-1])
	result := []string{}
	startI := -1
	gap := -1

	for index, num := range nums {
		if startI == -1 || num - index != gap {
			if startI != -1 {
				if nums[startI] != nums[index-1] {
					result = append(result,
						fmt.Sprintf("%v->%v", nums[startI], nums[index-1]))
				} else {
					result = append(result, fmt.Sprintf("%v", nums[startI]))
				}
			}
			startI = index
			gap = num - index
		}
	}
	return result
}