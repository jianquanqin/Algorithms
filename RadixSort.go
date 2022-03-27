package BLC

//Test example

//func main() {
//
//	var slice []int
//	//生成一个随机数组
//	rand.Seed(time.Now().UnixNano())
//	for index := 0; index <= 20; index++ {
//		num := rand.Intn(1000)
//		slice = append(slice, num)
//	}
//
//	fmt.Println("生成的数列是：", slice)
//	fmt.Println("排好的顺序是：", RadixSort(slice))
//
//}

func RadixSort(s []int) []int {
	max := Maxvalue(s)

	it := 0

	array := []int{1, 10, 100, 1000, 10000, 100000}

	for array[it] <= max {

		buckets := make([][]int, 10) //创建10个桶
		for _, val := range s {      //把元素先装进各个桶中
			digit := (val / array[it]) % 10
			buckets[digit] = append(buckets[digit], val)
		}

		s = []int{}
		for i := 0; i < 10; i++ {
			for j := 0; j < len(buckets[i]); j++ {
				s = append(s, buckets[i][j])
			}
		}

		it++
	}
	return s
}

//求一个数列中的最大值

func Maxvalue(s []int) int {

	maxvalue := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > maxvalue {
			maxvalue = s[i]
		}
	}
	return maxvalue
}
