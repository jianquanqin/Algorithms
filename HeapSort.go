package BLC

import "fmt"

// Test example

//func main() {
//	//将下列随机数列顺序输出
//	s := []int{2, 8, 5, 6, 3, 9, 7, 1, 25, 23, 14, 0}
//	HeadSort(s)
//}

func HeadSort(s []int) {
	n := len(s)

	//将任意数列构造成堆

	for i := (n - 2) / 2; i >= 0; i-- {
		sift(s, i, n-1)
	}
	fmt.Println("排好的堆是：", s)

	//将堆从小到大排序
	for j := n - 1; j >= 0; j-- {
		s[0], s[j] = s[j], s[0]
		sift(s, 0, j-1)
	}
	fmt.Println("堆中元素按照顺序输出：", s)

}

//定义节点整理函数，将给定节点整理成堆

func sift(s []int, low int, high int) {

	//定义两个指针
	i := low     //第一个指针指向根节点
	j := 2*i + 1 //第二个指针指向根节点的左孩子

	tmp := s[low] //将根节点的数先取出存放

	for j <= high { //保证叶子节点不越界
		if j+1 <= high && s[j] < s[j+1] { //在右孩子有且大于左孩子的情况下，j越界会退出for循环，左孩子大时会执行下一个if语句
			j = j + 1 //第二个指针指向右孩子
		}
		if s[j] > tmp { //右孩子大于根节点
			s[i] = s[j] //将右孩子填充至根节点
			i = j       //将右孩子作为根节点，遍历下一层1
			j = 2*i + 1
			s[i] = tmp //将根节点填充到右孩子的位置，重新开始遍历
		} else { //右孩子小于根节点，不做改变
			s[i] = tmp
			break
		}
	}
	//fmt.Println(s) 测试代码
}

//Topk 问题算法

//Test example

//func main() {
//	//将下列随机数列顺序最大的七位
//	s := []int{2, 8, 5, 6, 3, 9, 7, 1, 25, 23, 14, 0}
//	Topk(s, 7)
//}

func Topk(s []int, k int) {
	heap := s[0:k]

	//建堆
	for i := (k - 2) / 2; i >= 0; i-- {
		sift(heap, i, k-1)
	}
	fmt.Println("新建的小根堆是：", heap)

	//遍历列表中所有剩余元素，如果大于堆顶，则纳入重新调整
	for j := k; j < len(s); j++ {
		if s[j] > heap[0] {
			heap[0] = s[j]
			Sift(heap, 0, k-1)
		}
	}
	fmt.Println("遍历后重新调整过的堆是：", heap)

	//挨个输出
	for h := k - 1; h >= 0; h-- {
		s[0], s[h] = s[h], s[0]
		Sift(s, 0, h-1)
	}
	fmt.Println("堆中元素按照顺序输出：", heap)

}

//定义节点整理函数，将给定节点整理成小根堆

func Sift(s []int, low int, high int) {

	//定义两个指针
	i := low     //第一个指针指向根节点
	j := 2*i + 1 //第二个指针指向根节点的左孩子

	tmp := s[low] //将根节点的数先取出存放

	for j <= high { //保证叶子节点不越界
		if j+1 <= high && s[j] > s[j+1] { //在右孩子有且大于左孩子的情况下，j越界会退出for循环，左孩子小时会执行下一个if语句
			j = j + 1 //第二个指针指向右孩子
		}
		if s[j] < tmp { //右孩子小于根节点
			s[i] = s[j] //将右孩子填充至根节点
			i = j       //将右孩子作为根节点，遍历下一层1
			j = 2*i + 1
			s[i] = tmp //将根节点填充到右孩子的位置，重新开始遍历
		} else { //右孩子大于根节点，不做改变
			s[i] = tmp
			break
		}
	}
	//fmt.Println(s) 测试代码
}
