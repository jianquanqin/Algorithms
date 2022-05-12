package array

import "fmt"

func containsDuplicate(nums []int) bool {

	cache := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		cache[nums[i]] = append(cache[nums[i]], i)
		if len(cache[nums[i]]) > 1 {
			fmt.Println(cache)
			return true
		}
	}

	fmt.Println(cache)
	return false
}
