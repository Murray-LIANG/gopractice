package gopractice

// https://leetcode.com/problems/subarray-product-less-than-k

func numSubarrayProductLessThanK(nums []int, k int) int {
	result := 0

	var product int64
	product = 1
	i := 0
	for j, num := range nums {
		product *= int64(num)
		for ; i <= j && product >= int64(k); i++ {
			product /= int64(nums[i])
		}
		result += j-i+1
	}
	return result 
}