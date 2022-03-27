package BLC

//Test example

//func main() {
//	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
//	fmt.Println(QuickSort(arr))
//}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0]          //第一个数据
	low := make([]int, 0, 0)     //比我小的数据
	hight := make([]int, 0, 0)   //比我大的数据
	mid := make([]int, 0, 0)     //与我一样大的数据
	mid = append(mid, splitdata) //加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}

//Test example
//
//func main() {
//
//	nums := []int{17, 8, 9, 4, 3, 2, 1, 6, 23, 28, 13, 23, 77, 43}
//	result := QuickSort1(nums, 0, len(nums)-1)
//	fmt.Println(result)
//}

func QuickSort1(nums []int, left int, right int) []int {

	if left < right {
		mid := Partition(nums, left, right)
		QuickSort1(nums, left, mid-1)
		QuickSort1(nums, mid+1, right)
		return nums
	}
	return nums
}

func Partition(nums []int, left int, right int) int {
	tmp := nums[left]
work:
	for right > 0 {
		if left < right && nums[right] < tmp {
			nums[left] = nums[right]
			for left < len(nums)-1 {
				if left < right && nums[left] > tmp {
					nums[right] = nums[left]
					continue work
				}
				left++
				if left == right {
					nums[right] = tmp
				}
			}
		}
		right--
		if left == right {
			nums[right] = tmp
		}
	}
	return left
}
