package BLC

//Test example

//func main() {
//	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	v := 5
//
//	result := BinarySearch(s, v)
//	if result != -1 {
//		fmt.Println(result)
//	} else {
//		fmt.Println("未找到:", v)
//	}
//}

func BinarySearch(s []int, val int) int {
	left := 0
	right := len(s) - 1

	for left <= right {
		mid := (left + right) / 2
		if s[mid] == val {
			return mid
			break
		} else if s[mid] > val {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
