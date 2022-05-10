package stack

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

//20. 有效的括号

func isValid(s string) bool {

	if len(s)%2 == 1 {
		return false
	}

	tmp := "{}()[]"
	m := make(map[byte]byte)
	m[tmp[0]] = tmp[1]
	m[tmp[2]] = tmp[3]
	m[tmp[4]] = tmp[5]
	stack := &ParenthesesStack{}

	for i := 0; i < len(s); i++ {
		if stack.StackSize == 0 {
			stack.PushElement(s[i])
		} else {
			if m[stack.Stack[stack.StackSize-1]] == s[i] {
				stack.PopElement()
			} else {
				stack.PushElement(s[i])
			}
		}
	}
	//栈空说明已经全部配对完毕
	if stack.StackSize == 0 {
		return true
	} else {
		return false
	}
}

//32. 最长有效括号

//暴力求解,时间复杂度n^3,面试先写这个，然后再优化

func longestValidParenthesesI(s string) int {

	var l int

	if len(s)%2 == 1 {
		l = len(s) - 1
	} else {
		l = len(s)
	}

	i := l

	for i > 0 {
		tmp := 0
		for j := i; j <= len(s); j++ {
			//fmt.Println(tmp, j)
			//fmt.Println(s[tmp:j])
			if IsMatched(s[tmp:j]) == false {
				tmp++
			} else {
				return i
			}
		}
		i = i - 2
	}
	return 0
}

func IsMatched(s string) bool {

	tmp := "()"
	m := make(map[byte]byte)
	m[tmp[0]] = tmp[1]
	stack := &ParenthesesStack{}

	var i int

	for i < len(s) {
		//无需匹配，可以直接加入栈新元素
		if s[i] == tmp[0] {
			stack.PushElement(s[i])
			i++
		} else {
			if stack.StackSize > 0 {
				if s[i] == m[stack.Stack[stack.StackSize-1]] {
					stack.PopElement()
					i++
				} else {
					stack.PushElement(s[i])
					i++
				}
			} else {
				stack.PushElement(s[i])
				i++
			}
		}
	}

	//fmt.Println(stack.Stack)
	if stack.StackSize == 0 {
		return true
	}
	return false
}

//定义栈，注意，如下的栈并没有长度限制,仅仅只是利用它的后进先出特性而已，底层仍然是数组
//这种方式其实增加元素时开销很大，使用链表可能会更好一点

type ParenthesesStack struct {
	Stack     []byte
	StackSize int
}

//入栈

func (stack *ParenthesesStack) PushElement(val byte) *ParenthesesStack {
	stack.Stack = append(stack.Stack, val)
	stack.StackSize++
	return stack
}

//出栈

func (stack *ParenthesesStack) PopElement() {

	if stack.StackSize == 0 {
		fmt.Println("stack is empty")
		return
	}
	//注意栈底元素的删除
	if stack.StackSize > 0 {
		stack.StackSize--
		stack.Stack = stack.Stack[0:stack.StackSize]
	}
}

//方法二；动态规划

func longestValidParenthesesII(s string) int {

	count := len(s)
	if count == 0 {
		return 0
	}
	dp := make([]int, count)
	dp[0] = 0
	for i := 1; i < count; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 1 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i > dp[i-1] {
					if s[i-dp[i-1]-1] == '(' {
						if i >= dp[i-1]+2 {
							dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
						} else {
							dp[i] = dp[i-1] + 2
						}
					}
				}
			}
		} else {
			dp[i] = 0
		}
	}
	max := 0
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max
}

//剑指 Offer 09. 用两个栈实现队列

type CQueue struct {
	arr  []interface{}
	size int
}

func Constructor() CQueue {
	var queue = CQueue{}
	return queue
}

func (this *CQueue) AppendTail(value int) {
	this.arr = append(this.arr, value)
	this.size++
}

func (this *CQueue) DeleteHead() int {
	if this.size == 0 {
		return -1
	}
	if this.size == 1 {
		this.size--
		first := this.arr[0].(int)
		this.arr = this.arr[0:0]
		return first
	}
	//first := this.arr[0]
	first := this.arr[0].(int)
	this.arr = this.arr[1:]
	this.size--

	return first
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

//剑指 Offer 30. 包含min函数的栈

type MinStack struct {
	arr  []interface{}
	size int
	min  []interface{}
}

/** initialize your data structure here. */

func ConstructorMinStack() MinStack {
	var minStack = MinStack{}
	return minStack
}

func (this *MinStack) Push(x int) {

	if this.size == 0 {
		this.min = append(this.min, x)
	} else {
		if x < this.min[len(this.min)-1].(int) {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, this.min[len(this.min)-1])
		}
	}

	this.arr = append(this.arr, x)
	this.size++

}

func (this *MinStack) Pop() {

	if this.size == 0 {
		fmt.Println("the stack is empty")
		return
	} else {
		this.size--
		this.arr = this.arr[0:this.size]

		this.min = this.min[0:this.size]
		return
	}
}

func (this *MinStack) Top() int {

	if this.size == 0 {
		fmt.Println("the stack is empty")
		return -1
	} else {
		return this.arr[this.size-1].(int)
	}
}

func (this *MinStack) Min() int {

	if this.size == 0 {
		fmt.Println("the stack is empty")
		return -1
	}

	return this.min[this.size-1].(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
