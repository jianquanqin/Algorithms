package src

import "fmt"

//Test example

//func main() {
//	nums := []int{14, 11, 15, 35, 1, 9,6,8}
//	a := InsertSort(nums)
//	fmt.Println(a)
//}

func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ { //从数列的第二个数开始（选出的数）
		for j := i - 1; j >= 0; j-- { //依次将选出的数从右往左与其比较，小的话交换位置
			if nums[j] > nums[j+1] { //
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		fmt.Println(nums) //演示过程
	}
	return nums
}
