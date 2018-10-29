package gopractice

func containsDuplicate(nums []int) bool {
	cache := map[int]bool{}

	for _, num := range nums {
		if _, ok := cache[num]; ok {
			return true
		} else {
			cache[num] = false
		}
	}
	return false
}
