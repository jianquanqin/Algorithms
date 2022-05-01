package src

//Test example

//func main() {
//	s := []int{1, 1, 0, 0, 2, 9, 2, 8, 23, 2, 7, 21}
//	fmt.Println(CountSort(s, 23))
//}

func CountSort(s []int, max int) []int {

	//定义一个数组来存放中中间结果
	count := make([]int, max+1)

	//遍历原数组，如果有值，给相应的索引在新的数组中加一
	for i, _ := range count {
		for j, _ := range s {
			if i == s[j] {
				count[i]++
			}
		}
	}
	//清空原数组，遍历打印
	s = []int{}
	for j, v := range count {
		for i := v; i > 0; i-- {
			s = append(s, j)
		}
	}
	return s
}
