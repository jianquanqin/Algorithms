package src

import "fmt"

const MaxLen = 10

//使用结构体定义固定长度的线性表

type LinearList struct {
	IntList [MaxLen]int
	Len     int
}

//1.插入

func (list *LinearList) InsertLinearList(i int, value int) {
	//判断表是否已经满了
	if len(list.IntList) == list.Len {
		fmt.Println("the list is full")
		return
	}
	//判断插入的位置是都与线性表中的任何一个元素相邻（不允许出现空档）
	if i > list.Len+1 || i < 1 {
		fmt.Println("i is invalid")
		return
	}
	tmp := value
	for k := list.Len - 1; k >= i-1; k-- {
		list.IntList[k+1] = list.IntList[k]
	}
	list.IntList[i-1] = tmp
	list.Len++
}

//2.删除

func (list *LinearList) DeleteLinearList(i int) {
	//判断表是否已经空了了
	if list.Len == 0 {
		fmt.Println("the list is empty")
		return
	}
	//判断所选元素是否在表中
	if i > list.Len || i < 1 {
		fmt.Println("i is invalid")

	}
	for k := i; k < len(list.IntList); k++ {
		list.IntList[k-1] = list.IntList[k]
	}
	list.Len--
}

//3.排序

func BubbleSortLinearList(nums []int) {

	for i := 0; i < len(nums); i++ {
		change := false
		for j := 0; j < (len(nums) - 1); j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				change = true
			}
		}
		if change == false {
			//fmt.Println(nums) //演示排序过程，如果排好提前结束
			return
		}
	}
}

//使用结构体定义长度可变的线性表

type DynamicLinearList struct {
	IntList []int
	MaxSize int
	Len     int
}

//1.增加动态数组长度

func (list *DynamicLinearList) IncreaseDynamicLinearListSize() {
	if list.MaxSize == 0 {
		list.MaxSize++
	}
	if list.Len == list.MaxSize {
		list.MaxSize = list.MaxSize * 2
	}
}

//2.将元素插入动态数组中

func (list *DynamicLinearList) InsertDynamicLinearList(i int, value int) {
	list.IncreaseDynamicLinearListSize()

	if i > list.Len+1 || i < 1 {
		fmt.Println("i is invalid")
		return
	}
	tmp := value
	list.IntList = append(list.IntList, value)

	for k := list.Len - 1; k >= i-1; k-- {
		list.IntList[k+1] = list.IntList[k]
	}
	list.IntList[i-1] = tmp
	list.Len++
}
