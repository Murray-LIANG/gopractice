// https://leetcode.com/problems/find-pivot-index/description/

package gopractice

func pivotIndex(nums []int) int {
	// the solution is: for each number in nums, check whether
	// sum(nums[:index]) == (sum(nums) - nums[index]) / 2
	total := 0
	for _, num := range nums {
		total += num
	}

	left := 0
	for index, num := range nums {
		if 2 * left == (total - num) {
			return index
		}
		left += num
	}
	return -1 
}