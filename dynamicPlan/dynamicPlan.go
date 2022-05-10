package dynamicPlan

import "fmt"

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

var Cache = make(map[int]int)

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
	for k, _ := range cache {
		if k == n {
			return cache[n]
		}
	}

	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return cache[n]

}

var cache = make(map[int]int)

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
