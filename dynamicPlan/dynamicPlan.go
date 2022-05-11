package dynamicPlan

import (
	"fmt"
)

//非波那契数列

func fib(n int) int {

	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
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
