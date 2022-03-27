package BLC

//Test example

//func main() {
//	nums := []int{14, 11, 15, 35, 1, 9, 6, 8, 23, 2, 7, 21}
//	fmt.Println(ShellSort(nums))
//}

func InsertSort1(nums []int, gap int) []int {
	for i := gap; i < len(nums); i++ {
		for j := i - gap; j >= 0; j-- {
			if nums[j] > nums[j+gap] {
				nums[j], nums[j+gap] = nums[j+gap], nums[j]
			}
		}
		//fmt.Println(nums) //演示过程
	}
	return nums
}
func ShellSort(nums []int) []int {
	d := len(nums) / 2 //先把数列分为两组
	for d >= 1 {
		InsertSort1(nums, d) //再将对应位置的元素比较，小的放前面
		d /= 2               //然后再次对半分，直至d=1，此时数列中最多有一对元素顺序相反，调整一下即可
	}
	return nums
}
