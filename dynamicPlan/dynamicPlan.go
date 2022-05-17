package dynamicPlan

import (
	"fmt"
	"math"
)

//非波那契数列

func fib(n int) int {

	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}

//泰波那契数列
//递归的全局变量要定义在函数之外

func tribonacci(n int) int {

	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	for k, _ := range Cache {
		if k == n {
			return Cache[k]
		}
	}
	Cache[n] = tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	//fmt.Println(Cache)
	return Cache[n]
}

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	for k, _ := range Cache {
		if k == n {
			return Cache[n]
		}
	}

	Cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return Cache[n]

}

//欧几里得算法

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for {
		//取余数，若为0，除数就是二者的最大公约数
		c := a % b
		fmt.Println("c:", c)
		if c == 0 {
			break
		} else {
			//除数等于被除数，余数最为除数重新计算余数
			a = b
			b = c
		}
	}
	return b
}

//746. 使用最小花费爬楼梯

var totalCost int
var Cache = make(map[int]int)
var path []int

//cost = [1,100,1,1,1,100,1,1,100,1]

func minCostClimbingStairs(cost []int) int {

	//方法一：
	//if len(cost) == 0 {
	//	return 0
	//}
	//if len(cost) == 1 {
	//	return cost[0]
	//}
	//
	//if len(cost) == 2 {
	//	if cost[0] > cost[1] {
	//		return cost[1]
	//	} else {
	//		return cost[0]
	//	}
	//}
	//
	//if min(cost) == true {
	//	totalCost += cost[1]
	//	if len(cost) > 1 {
	//		cost = cost[2:]
	//		fmt.Println(cost)
	//		minCostClimbingStairs(cost)
	//	}
	//} else {
	//	totalCost += cost[0] + cost[2]
	//	if len(cost) > 2 {
	//		cost = cost[3:]
	//		fmt.Println(cost)
	//		minCostClimbingStairs(cost)
	//	}
	//}
	//return totalCost

	//方法二
	//这种方法也很聪明
	for i := 2; i < len(cost); i++ {
		fee := min(cost[0], cost[1]) + cost[i]
		cost[0] = cost[1]
		cost[1] = fee
	}
	return min(cost[0], cost[1])
}

//func min(nums []int) bool {
//
//	if nums[0]+nums[2] > nums[1] {
//		return true
//	} else {
//		return false
//	}
//}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//198. 打家劫舍
//递推公式

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	//定义一个动态规划数组
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	//求出递推公式就OK了
	for i := 0; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func robII(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	//如果选了第一家，则不能选择最后一家，范围是[0:n-1]
	//如果未选第一家，则不能选择第一家，范围是[1:n]
	//二者取其大者
	return max(robIII(nums[:n-1]), robIII(nums[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//740. 删除并获得点数
//特别注意这种对于数组的用法

func deleteAndEarn(nums []int) int {

	maxVal := 0
	//遍历数组取出最大值
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	//创建一个固定数组，数组的下标是[0,maxVal],所以数组的最大值要+1，然后把其当map用
	sum := make([]int, maxVal+1)
	//然后把它变成打家劫舍模型
	for _, val := range nums {
		sum[val] += val
	}
	return robIII(sum)
}

func robIII(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	const mod int = 1e9 + 7
	if n <= 2 {
		return n
	}

	p, q, r := 0, 1, 2 //取前三位即可

	for i := 3; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}

//剑指 Offer 63. 股票的最大利润

func maxProfit(prices []int) int {

	// 时间复杂度O(n) 空间复杂度O(1)
	var min, max, maxProfit = math.MaxInt, 0, 0

	for i := 0; i < len(prices); i++ {
		//求最高价格
		if prices[i] > max {
			max = prices[i]
		}
		//如果该价格小于最低价
		//聪明聪明，这样就分离了最高价和最低价
		if prices[i] < min {
			//让该价格等于最低价
			min = prices[i]
			max = 0
		}
		//一次循环，三次分离，牛，重要的是在最小值的初始化上，得设一个很大的值
		if max-min > maxProfit {
			maxProfit = max - min
		}
	}
	return maxProfit
}
