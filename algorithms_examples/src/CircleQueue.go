package src

import (
	"errors"
	"fmt"
)

//Test example

//func main() {
//
//	cq := &CircleQueue{MaxSize: 5, Head: 0, Tail: 0}
//
//	//添加元素
//	cq.Push(1)
//	cq.Push(2)
//	cq.Push(3)
//	cq.Push(4)
//
//	//显示队列
//	cq.Show()
//	//取出元素
//	cq.Pop()
//
//	//队列中还剩多少个元素
//	count := cq.Count()
//	fmt.Println(count)
//
//	//显示队列
//	cq.Show()
//
//}

type CircleQueue struct {
	MaxSize int
	Array   [5]int //定义一个固定数组
	Head    int
	Tail    int
}

//添加元素

func (cq *CircleQueue) Push(val int) (err error) {
	if cq.IsFull() {
		fmt.Println("the queue is full, cannot add a new number")
		return errors.New("queue is full")
	}
	cq.Array[cq.Tail] = val
	cq.Tail++

	return nil
}

//取出元素

func (cq *CircleQueue) Pop() (val int, err error) {
	if cq.IsEmpty() {
		fmt.Println("the queue is empty")
		return 0, errors.New("the queue is empty")
	}

	val = cq.Array[cq.Head]
	cq.Head++

	return val, nil
}

//队列是否为满

func (cq *CircleQueue) IsFull() bool {
	return (cq.Tail+1)%cq.MaxSize == cq.Head
}

//队列是否为空

func (cq *CircleQueue) IsEmpty() bool {
	return cq.Head == cq.Tail
}

//计算对列中的元素

func (cq *CircleQueue) Count() int {

	count := (cq.Tail + cq.MaxSize - cq.Head) % cq.MaxSize
	return count
}

//显示队列

func (cq *CircleQueue) Show() {
	if cq.Count() == 0 {
		fmt.Println("queue is empty")
	}

	tmp := cq.Head
	for i := 0; i < cq.Count(); i++ {
		fmt.Printf("array[%d]=%d\t", tmp, cq.Array[tmp])
		tmp = (tmp + 1) % cq.MaxSize
	}
}
