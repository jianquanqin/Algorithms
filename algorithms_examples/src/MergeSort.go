package src

//Test example

//func main() {
//	s := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
//	fmt.Println(MergeSort(s))
//
//}

func MergeSort(slice []int) []int {

	//求数列的长度
	length := len(slice)
	//如果数列长度无或者只有一个元素，则无序排序
	if length <= 1 {
		return slice
	} else {
		//将数列二分，各自递归
		mid := length / 2
		ls := MergeSort(slice[:mid])
		rs := MergeSort(slice[mid:])
		return Merge(ls, rs) //合并
	}
}
func Merge(ls []int, rs []int) []int {

	//定义两个索引
	i := 0
	j := 0

	//建立一个空slice，用来保存新生成的数组
	var slice []int
	//s1 := []int{1, 3, 20, 5, 2, 12, 10, 11}

	for i < len(ls) && j < len(rs) { //保证不越界

		if ls[i] < rs[j] { //比较两个数列中的元素，谁小谁入slice
			slice = append(slice, ls[i])
			i++
		} else if ls[i] > rs[j] { //比较两个数列中的元素，谁小谁入slice
			slice = append(slice, rs[j])
			j++
		} else { // 相等则都入slice
			slice = append(slice, ls[i])
			i++
			slice = append(slice, rs[j])
			j++
		}
	}
	for i < len(ls) { //若还有剩余则一起入slice
		slice = append(slice, ls[i])
		i++
	}
	for j < len(rs) {
		slice = append(slice, rs[j])
		j++
	}
	return slice
}
