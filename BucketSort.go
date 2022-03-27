package BLC

import (
	"fmt"
	"math/rand"
	"time"
)

//Test example

func main() {

	var slice []int
	rand.Seed(time.Now().UnixNano())
	for index := 0; index <= 10; index++ {
		num := rand.Intn(100)
		slice = append(slice, num)
	}
	fmt.Println("生成的数列是：", slice)

	fmt.Println("排好的顺序是：", BucketSort(slice, 5, 87))
}

func BucketSort(s []int, n int, max int) []int {

	buckets := make([][]int, n) //创建n个桶

	for _, val := range s { //把元素先装进各个桶中
		i := Min(val/(max/n), n-1)
		buckets[i] = append(buckets[i], val)
		for j := len(buckets[i]) - 1; j > 0; j-- {
			if buckets[i][j] < buckets[i][j-1] { //插入元素时就进行排序
				buckets[i][j-1], buckets[i][j] = buckets[i][j], buckets[i][j-1]
			} else {
				break
			}
		}
	}

	s = []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			s = append(s, buckets[i][j])
		}
	}
	return s
}

func Min(a, b int) int {
	var i int

	if a > b {
		i = b
	} else {
		i = a
	}
	return i
}