package BLC

//Test example

//func main() {
//
//	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	v := 5
//
//	result := BLC.SequentialSearch(s, v)
//	if result != -1 {
//		fmt.Println(result)
//	} else {
//		fmt.Println("未找到:", v)
//	}
//}

func SequentialSearch(s []int, value int) int {
	for index, v := range s {
		if v == value {
			return index
		}
	}
	return -1
}
