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

 //魔塔游戏
 func magicTower(nums []int) int {
	sum := 1
	
	for _, num := range nums {
		sum += num
	}
	if sum <= 0 {
		return -1
	}

    // 记录最小的数据
	monsters := []int{}
	move := 0
	sum = 1
	//如果扛不住就移动
    // 标记是否移动，移动则最小位置可复用
    isRemove := false
	for i := range nums {
        // 记录最小数据
        sum += nums[i]
		if nums[i] < 0 {
            if isRemove { //发生移动，则复用最小位置
                monsters[0] = nums[i]
            }else {
                // 未发生移动，则新增位置记录最小值
                monsters = append(monsters, nums[i])
            }		
		}

		if sum <= 0 {
			move++
            // 排序取最小值
            sort.Ints(monsters)
			sum -= monsters[0]
            // 最小值取出后，清除最小值，标记最小值位置可复用
			monsters[0] = 0
            isRemove = true
		}else {
            isRemove = false
        }
	}
	return move
}