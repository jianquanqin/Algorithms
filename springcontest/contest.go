package springcontest

import (
	"sort"
)



func PurchasePlans(nums []int, target int) int {
	sort.Ints(nums)//使用sort包直接排序
	var l, r = 0, len(nums)-1
	var res = 0
	for l < r {
	    if nums[l] + nums[r] > target{
		   r--
		   continue
	    }
	    res += (r - l)//这一步很重要，说明中间这几个都OK
	    l++
	}
	return res % 1000000007
 }