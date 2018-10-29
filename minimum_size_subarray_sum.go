package gopractice

// https://leetcode.com/problems/minimum-size-subarray-sum

func sum(nums []int, i int, j int) int {
	result := 0
	for _, num := range nums[i:j+1] {
		result += num
	}
	return result
}

func minSubArrayLen(s int, nums []int) int {
	i, j, n := 0, 0, len(nums)
	count := n+1
	for i <= j && j < n {
		if sum(nums, i, j) < s {
			j++
		} else {
			if j - i + 1 < count {
				count = j - i + 1
			}
			i++
		}
	}
	if count < n+1 {
		return count
	}
    return 0
}

func minSubArrayLen2(s int, nums []int) int {
	i, j, n := 0, 0, len(nums)
	count := n+1
	sum := 0

	for i <= j && j < n {
		sum += nums[j]
		j++

		for sum >= s {
			if j - i < count {
				count = j - i
			}
			sum -= nums[i]
			i++
		}
	}
	if count < n+1 {
		return count
	}
	return 0
}