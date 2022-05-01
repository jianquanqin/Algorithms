package src

//Test example

//func main() {
//	nums := []int{2, 14, 11, 15, 35, 1, 9, 6, 100, 26, 7, 8}
//	a := SelectSort(nums)
//	fmt.Println(a)
//}

func SelectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		//fmt.Println(nums)//演示过程
	}
	return nums
}
