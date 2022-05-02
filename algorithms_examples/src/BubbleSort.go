package src

import "fmt"

//Test example

//func main() {
//	nums := []int{2, 14, 11, 15, 35, 1, 9, 6, 100, 26, 7, 8}
//	a := BubbleSort(nums)
//	fmt.Println(a)
//}

func BubbleSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		change := false
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				change = true
			}
		}
		if change == false {
			fmt.Println(nums) //演示排序过程，如果排好提前结束
			return nums
		}
	}
	return nums
}
