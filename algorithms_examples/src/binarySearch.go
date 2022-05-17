package src

import (
	"fmt"
	"math"
	"math/rand"
)

//经典查找算法之：1.二分查找

//二分查找算法三大模版

/*
	模版一. 找一个准确值

	  循环条件：l<=r
	  缩减搜索空间：l = mid +1； r = mid -1

	模版二. 找一个模糊值

	  循环条件：l<r
	  缩减搜索空间：l = mid, r = mid -1 ; l = mid +1； r = mid

	模版三. 万用型

	  循环条件：l<r-1
	  缩减搜索空间：l = mid, r = mid

	模版四. 万能型

	  l := -1
      r := len(nums)

	  循环条件：l<r-1
	  缩减搜索空间：l = mid, r = mid

	注意事项

	  1.查到到边界以后还需要判断返回的值是否存在的问题
	  2.列表是否为空的问题
	  3.选择左端点为m还是右端点为m的问题
	  4.临界值的判断问题

*/

//二分查找算法样例

func BinarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l < r-1 {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m
		} else {
			l = m
		}
	}
	return r
}

//帮助函数

func HelperBinarySearch(nums []int, l int, r int, target int) int {

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

//范例 1：第一次出现的位置

//解题技巧：寻找模糊值，模版二
//分析：找的是边界，m为左

func FirstAppearance(nums []int, target int) int {

	//判空
	if len(nums) == 0 {
		return -1
	}

	l := 0
	r := len(nums) - 1

	//当l == r 时退出，此时返回值就是第一次出现的位置
	for l < r {
		m := (l + r) / 2
		//如果nums[m] < target,此时可以排除nums[m],搜索空间为[m+1,r]
		if target > nums[m] {
			l = m + 1
			//如果nums[m] <= target,此时不能排除nums[m],搜索空间为[l,m]
		} else {
			r = m
		}
	}
	//看返回的值是否正确
	if nums[l] == target {
		return l
	} else {
		return -1
	}
}

//范例 2：最后一次出现的位置

//解题技巧：寻找模糊值，模版二
//分析：找的是边界，m为右

func LastAppearance(nums []int, target int) int {

	//判空
	if len(nums) == 0 {
		return -1
	}

	l := 0
	r := len(nums) - 1

	for l < r {
		m := (r + l + 1) / 2
		//如果nums[m] < target,此时可以排除nums[m],搜索空间为[m,r-1]
		if nums[m] > target {
			r = m - 1
			//如果nums[m] >= target,此时不能排除nums[m],搜索空间为[m,r]
		} else {
			l = m
		}
	}
	//看返回的值是否正确
	if nums[l] == target {
		return l
	} else {
		return -1
	}
}

//范例 3：33. 搜索旋转排序数组
//解题技巧：寻找具体值，模版四
//注意⚠️，此题应多次回顾

func Search(nums []int, target int) int {

	l := -1
	r := len(nums)

	for l+1 != r {
		m := (l + r) / 2

		if nums[m] == target {
			return m
		} else if nums[0] <= nums[m] {
			if nums[l+1] <= target && nums[m] > target {
				r = m
			} else {
				l = m
			}
		} else {
			if nums[m] < target && nums[r-1] >= target {
				l = m
			} else {
				r = m
			}
		}
	}
	return -1
}

//范例 4：154.寻找旋转排序数组中的最小值I
//解题技巧：寻找模糊值，模版二
//注意⚠️，此题应多次回顾

func FindMinI(nums []int) int {
	l := 0
	r := len(nums) - 1
	//搜索模糊值用，l < r
	for l < r {
		m := (l + r) / 2
		//第一步，二分，比较端点，如果左<=中<=右，则有序，最小值是左端点，终止循环，返回左端点
		if nums[l] <= nums[m] && nums[m] <= nums[r] {
			break
		} else if nums[l] <= nums[m] {
			//如果无序，先判断左半边是否需有序号，如果有，则最小值在右半边，搜索空间[m+1,r]
			l = m + 1
		} else {
			//如果左半边无序，则最小值在左半边，搜索空间[l,m]，此时有可能mums[m]就是最值，所以一起搜索
			r = m
		}
	}
	return nums[l]
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

//范例 5: 154.寻找旋转排序数组中的最小值II
//解题技巧：寻找模糊值，模版二
//注意⚠️，此题应多次回顾，二分 + 朴素求解

func FindMinII(nums []int) int {
	l := 0
	r := len(nums) - 1
	//搜索模糊值用，l < r
	for l < r {
		m := (r + l) / 2 //让二分查指的中间元素为右边的一个
		//判断当前值是是否在左半边部分，如果是，则在右半边搜索
		if nums[m] > nums[r] {
			l = m + 1
			//如有右边是有序的，则在[l,m]范围内搜索
		} else if nums[m] < nums[r] {
			r = m
			//如果相等，则将右指针左移一位，去重
		} else if nums[m] == nums[r] {
			r--
		}
	}
	return nums[l]
}

//范例 6: 74. 搜索二维矩阵
//解题技巧：寻找具体值，模版一
//注意⚠️，此题应多次回顾

func searchMatrix(matrix [][]int, target int) bool {

	l := 0
	r := len(matrix) - 1
	flag := false

	for l <= r {
		mT := (l + r) / 2
		//先判断target在哪个列中
		//如果最大元素大于目标值，则在左半边的列表中
		if matrix[mT][len(matrix[0])-1] >= target {
			//然后在当前表中进行二分查找
			//初始化索引
			i := 0
			j := len(matrix[0]) - 1

			for i <= j {
				mE := (i + j) / 2
				if matrix[mT][mE] == target {
					flag = true
					fmt.Printf("它在 %v 行 第 %v列", mT+1, mE+1)
					break
				} else if matrix[mT][mE] < target {
					//在右半边里找
					i = mE + 1
				} else {
					//在左半边里找
					j = mE - 1
				}
			}
			//找不到的话就在[l,mT-1]范围内的表中找
			r = mT - 1
		} else {
			//若目标值大于当前表的最后一个元素，则在[mT+1,r]范围内的表中寻找
			l = mT + 1
		}
	}
	//fmt.Printf("%v 不在当前矩阵中", target)
	return flag
}

//范例 7: 34. 在排序数组中查找元素的第一个和最后一个位置
//解题技巧：寻找模糊值，模版二
//注意⚠️，此题应多次回顾

func searchRange(nums []int, target int) []int {

	var result []int
	//判空
	if len(nums) == 0 {
		return append(result, -1, -1)
	}

	l := 0
	r := len(nums) - 1

	//查找第一次出现的索引
	//当l == r 时退出，此时返回值就是第一次出现的位置
	for l < r {
		m := (l + r) / 2
		//如果nums[m] < target,此时可以排除nums[m],搜索空间为[m+1,r]
		if target > nums[m] {
			l = m + 1
			//如果nums[m] <= target,此时不能排除nums[m],搜索空间为[l,m]
		} else {
			r = m
		}
	}
	//看返回的值是否正确
	if nums[l] == target {
		result = append(result, l)
	} else {
		//没找到,直接返回
		result = append(result, -1, -1)
		return result
	}

	//重新初始化l和r
	l = 0
	r = len(nums) - 1

	for l < r {
		m := l + (r-l+1)/2
		//如果nums[m] < target,此时可以排除nums[m],搜索空间为[m,r-1]
		if nums[m] > target {
			r = m - 1
			//如果nums[m] >= target,此时不能排除nums[m],搜索空间为[m,r]
		} else {
			l = m
		}
	}

	result = append(result, l)

	return result
}

//范例 8: 69. x 的平方根
//解题技巧：寻找具体值，模版一
//注意⚠️，此题应多次回顾

func mySqrt(x int) int {

	l := 0
	r := x
	ans := -1
	for l <= r {
		m := (l + r) / 2
		//如果大于，则在[l,m-1]中继续中搜索
		if m*m > x {
			r = m - 1
			//如果等于，中断
		} else if m*m == x {
			ans = m
			break
			//如果小于，搜索空间[m+1,r]，当恰好是临界值时，不再满足任何一个条件，ans是临界值
		} else {
			ans = m
			l = m + 1
		}
	}
	return ans
}

//范例 9: 81. 搜索旋转排序数组 II
//解题技巧：寻找精确值，模版一
//注意⚠️，此题应多次回顾,也是暴力求解和二分法的结合，问题演变成如何在相同的数中找出不同值
//注意与154题的区别

func SearchII(nums []int, target int) bool {

	if len(nums) == 0 {
		fmt.Println("sorry, the nums is empty")
		return false
	}

	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2

		fmt.Println("index:", m)
		if nums[m] == target {
			//fmt.Println("index:", m)
			return true
			//1.如果无法判断左边升序还是右边升序
		} else if nums[m] == nums[l] && nums[m] == nums[r] {
			l++
			r--
			//如果左边升序，此时右边无序
		} else if nums[l] <= nums[m] {
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
			//如果右边升序，此时左边无序
		} else if nums[m] <= nums[r] {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return false
}

//范例 9: 162. 寻找峰值

func FindPeakElement(nums []int) int {

	n := len(nums)
	idx := rand.Intn(n)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	// 辅助函数，在函数内部被定义

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	for !(get(idx-1) < get(idx) && get(idx) > get(idx+1)) {
		if get(idx) < get(idx+1) {
			idx++
		} else {
			idx--
		}
	}

	return idx
}

//范例10. 167. 两数之和 II - 输入有序数组
//模版二

func TwoSum(numbers []int, target int) []int {
	var result []int
	l := 0
	r := len(numbers) - 1

	for l < r {
		m := (l + r) / 2
		//第一种情况：target是左半边两个数的和
		//注意负数，也就是加数比和大的情况，此时m=r
		if numbers[m] > target {
			r = m
			if numbers[l]+numbers[r] == target {
				return append(result, l+1, r+1)
			}
		} else if target == (numbers[m] + numbers[r]) {
			return append(result, m+1, r+1)
			//第二种情况：target是右半边两个数的和
		} else if target > (numbers[m] + numbers[r]) {
			l = m + 1
			//第三种情况：target的两个加数分别在m的两边
		} else {
			for l < r {
				if numbers[l]+numbers[r] == target {
					return append(result, l+1, r+1)
				} else if numbers[l]+numbers[r] < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	if numbers[l]+numbers[r] != target {
		fmt.Println("找不到")
	}
	return append(result, l+1, r+1)
}

//278. 第一个错误的版本

func FirstBadVersion(n int) int {

	l := 1
	r := n

	for l < r {
		m := (l + r) / 2

		//fmt.Println(m)
		//如果不是错误版本，说明正确版本在右边
		if isBadVersion(m) == false {
			l = m + 1
			//如果是错误版本，错误版本在左边，也可能m就是第一个错误版本
		} else {
			r = m
		}
	}
	return r
}

func isBadVersion(i int) bool {

	if i >= 5 {
		return true
	} else {
		return false
	}
}

//209. 长度最小的子数组

func MinSubArrayLen(target int, nums []int) int {

	//算法一：暴力穷举

	//sum := 0
	////注意这种算法，先遍历每个元素的前1个，再遍历每一个元素的前两个，依次类推
	//count := 0
	//flag := false
	//
	//for count < len(nums) {
	//	//nums := []int{2, 3, 1, 2, 4, 3}
	//	//遍历所有元素
	//	for i := count; i < len(nums); i++ {
	//		sum = 0
	//		//当前累加元素为count个
	//		for j := 0; j <= count; j++ {
	//			sum = sum + nums[i-j]
	//			fmt.Println("count,sum,j:", count, sum, j)
	//			if sum >= target {
	//				flag = true
	//				goto work
	//			}
	//		}
	//		//if flag == true {
	//		//	break
	//		//}
	//	}
	//work:
	//	count++
	//	if flag == true {
	//		break
	//	}
	//}
	//if flag {
	//	return count
	//} else {
	//	return 0
	//}

	//算法二：指针

	min := math.MaxInt32
	sum := 0
	l := 0
	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		//满足窗口内的和大于target时，左边界右移，缩小边界
		for sum >= target {
			if min < r-l+1 {
				min = min
			} else {
				min = r - l + 1
			}
			//窗口左边界右移
			sum -= nums[l]
			l++

		}
	}
	//遍历结束，如果没有找到结果
	if min == math.MaxInt32 {
		return 0
	}
	return min
}
