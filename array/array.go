package array

import (
	"fmt"
	"math"
)

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

func findRepeatNumber(nums []int) int {

	cache := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		cache[nums[i]] = append(cache[nums[i]], i)
		if len(cache[nums[i]]) > 1 {
			return nums[i]
		}
	}

	return math.MinInt
}

//求最大子数组

func maxSubArray(nums []int) int {

	//动态规划
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
	}

	res := math.MinInt
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(nums)
	return res
}

func twoSum(nums []int, target int) []int {

	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[j] == nums[i] {
				res = append(res, i, j)
			}
		}
	}

	return res
}

//先复制再排序

func merge(nums1 []int, m int, nums2 []int, n int) {

	//方法一
	//copy(nums1[m:], nums2)
	//
	//for i := 0; i < len(nums1); i++ {
	//	change := false
	//	for j := 0; j < (len(nums1) - 1); j++ {
	//		if nums1[j] > nums1[j+1] {
	//			nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
	//			change = true
	//		}
	//	}
	//	if change == false {
	//	}
	//}
	//fmt.Println(nums1)

	//方法二

	//tmp := nums1
	//nums1 = []int{}
	//
	//l1 := 0
	//l2 := 0
	//
	//for l2 < n && l1 < m {
	//	if tmp[l1] > nums2[l2] {
	//		nums1 = append(nums1, nums2[l2])
	//		l2++
	//	} else {
	//		nums1 = append(nums1, tmp[l1])
	//		l1++
	//	}
	//}
	//nums1 = append(nums1, tmp[l1:m]...)
	//nums1 = append(nums1, nums2[l2:n]...)
	//
	//fmt.Println(nums1)

	//方法三

	l1 := m - 1
	l2 := n - 1
	p := m + n - 1

	for l1 >= 0 && l2 >= 0 {
		if nums1[l1] > nums2[l2] {
			nums1[p] = nums1[l1]
			p--
			l1--
		} else {
			nums1[p] = nums2[l2]
			p--
			l2--
		}
	}

	if l2 >= 0 {
		copy(nums1[0:l2+1], nums2[0:l2+1])
	}

	fmt.Println(nums1)
}

//350. 两个数组的交集 II

func intersect(nums1 []int, nums2 []int) []int {

	//方法一

	//hashMap := make(map[int]int)
	//var res []int
	//
	//if len(nums1) < len(nums2) {
	//	for _, v := range nums1 {
	//		hashMap[v]++
	//	}
	//	fmt.Println(hashMap)
	//
	//	for _, v := range nums2 {
	//		if hashMap[v] > 0 {
	//			hashMap[v]--
	//			res = append(res, v)
	//		}
	//	}
	//} else {
	//	for _, v := range nums2 {
	//		hashMap[v]++
	//	}
	//
	//	fmt.Println(hashMap)
	//
	//	for _, v := range nums1 {
	//		if hashMap[v] > 0 {
	//			hashMap[v]--
	//			res = append(res, v)
	//		}
	//	}
	//}
	//
	//fmt.Println(hashMap)
	//return res

	//方法二 排序后比较

	bubbleSortLinearList(nums1)
	bubbleSortLinearList(nums2)

	fmt.Println(nums2, nums1)
	var res []int

	l1 := 0
	l2 := 0

	for l1 < len(nums1) && l2 < len(nums2) {
		if nums1[l1] == nums2[l2] {
			res = append(res, nums1[l1])
			l1++
			l2++
		} else {
			if nums1[l1] > nums2[l2] {
				l2++
			} else {
				l1++
			}
		}
	}

	return res
}

func bubbleSortLinearList(nums []int) {

	for i := 0; i < len(nums); i++ {
		//每一次遍历时都会初始化
		//使用flag优化
		isChange := false
		//只需要依次比较n-1个数即可
		for j := 0; j <= len(nums)-2; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isChange = true
			}
		}
		if isChange == false {
			//fmt.Println(nums) //演示排序过程，如果排好提前结束
			return
		}
	}
}

func maxProfit(prices []int) int {

	inprice := math.MaxInt

	max := 0

	for _, v := range prices {
		if v < inprice {
			inprice = v
		} else {
			if v-inprice > max {
				max = v - inprice
			}
		}
	}
	return max
}

func search(nums []int, target int) int {

	l := 0
	r := len(nums) - 1
	count := 0

	for l <= r {
		m := (l + r) / 2

		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			count++
			for j := m + 1; j < len(nums); j++ {
				if nums[j] == target {
					count++
				} else {
					break
				}
			}
			for k := m - 1; k >= 0; k-- {
				if nums[k] == target {
					count++
				} else {
					break
				}
			}
			break
		}
	}

	return count
}

func missingNumber(nums []int) int {

	if nums[len(nums)-1] == len(nums)-1 {
		return len(nums)
	}

	l := 0
	r := len(nums) - 1

	for l < r {

		m := (l + r) / 2

		if nums[m] != m {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

//求最大连续楼层

func maxConsecutive(bottom int, top int, special []int) int {

	tmp := 0

	bubbleSort(special)

	tmp = maxGap(special)
	//fmt.Println(tmp)

	a := special[0] - bottom
	b := top - special[len(special)-1]
	c := max(a, b)

	if c > tmp {
		return c
	}

	return tmp
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func bubbleSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		change := false
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				change = true
			}
		}
		if change == false {
			//fmt.Println(nums) //演示排序过程，如果排好提前结束
			return nums
		}
	}
	return nums
}

func maxGap(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	l := 0
	r := len(nums) - 1

	//special := []int{12, 24, 38, 48}

	for l < r-1 {
		m := (r - l + 1) / 2
		if nums[m]-nums[l] > nums[r]-nums[m] {
			r = m
			//fmt.Println(l, r)
		} else {
			//fmt.Println(l, r)
			l = m
		}
	}
	return nums[r] - nums[l] - 1
}

//剑指 Offer 04. 二维数组中的查找

func findNumberIn2DArray(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		if findInLine(matrix[i], target) == true {
			//fmt.Println(matrix[i])
			return true
		}
	}
	return false
}

func findInLine(nums []int, t int) bool {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (r + l) / 2
		if nums[m] == t {
			return true
		} else {
			if nums[m] > t {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return false
}

//剑指 Offer 11. 旋转数组的最小数字

func minArray(numbers []int) int {

	l := 0
	r := len(numbers) - 1

	for l < r {
		m := (r + l) / 2
		fmt.Println(l, m, r)

		//判断右边，如果升序，最小值一定在左半边
		if numbers[m] < numbers[r] {
			r = m
		} else if numbers[m] > numbers[r] {
			l = m + 1
		} else {
			if numbers[l] == numbers[m] {
				r--
				l++
			} else {
				r = m
			}

		}

	}
	return numbers[l]
}

//剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
//双指针

func exchange(nums []int) []int {

	l := 0
	r := len(nums) - 1

	for l < r {

		if nums[l]%2 == 0 && nums[r]%2 == 1 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		} else if nums[l]%2 == 1 && nums[r]%2 == 0 {
			l++
			r--
		} else if nums[l]%2 == 1 && nums[r]%2 == 1 {
			l++
		} else if nums[l]%2 == 0 && nums[r]%2 == 0 {
			r--
		}
	}
	return nums
}

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]int)
 
	for i :=0; i<len(nums); i++ {
	    if m[nums[i]] == 0 {
		   m[nums[i]] = 1
	    }else {
		   return true
	    }
	}
 
	return false
 
 }