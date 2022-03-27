package BLC

import (
	"errors"
	"fmt"
)

//Test example

//func main() {
//
//	stack := &Stack{
//		MaxTop: 5,
//		Top:    -1,
//	}
//	//入栈
//	stack.Push(1)
//	stack.Push(2)
//	stack.Push(3)
//	stack.Push(4)
//	stack.Push(5)
//
//	//出栈
//	val, _ := stack.Pop()
//	fmt.Println(val)
//
//	//遍历栈
//	stack.Range()
//}

//创建栈

type Stack struct {
	MaxTop int //表示最多存放5个元素至栈中
	Top    int //表示初始状态栈为空
	arr    [5]int
}

//注意：不像python语言那样已经有一些对栈的操作方法，golang语言需要我们自己定义这些方法，所以此时限制只有我们的想象力了

//入栈

func (stack *Stack) Push(val int) (err error) {

	//先判断栈是否已经满了
	if stack.Top == stack.MaxTop-1 {
		fmt.Println("the stack is full")
		return errors.New("stack full")
	}
	stack.Top++
	//加入数据
	stack.arr[stack.Top] = val

	return
}

//出栈

func (stack *Stack) Pop() (val int, err error) {
	//先判断栈是否已经空了
	if stack.Top == -1 {
		fmt.Println("the stack is empty")
		return 0, errors.New("stack empty")
	} else {
		val := stack.arr[stack.Top]
		stack.Top--
		return val, nil
	}
}

func (stack *Stack) Range() {
	if stack.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i := stack.Top; i >= 0; i-- {
		fmt.Println(stack.arr[i])
	}
}
