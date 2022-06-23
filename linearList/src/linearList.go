package src

import (
	"fmt"
	"strings"
)

const MaxLen = 10000

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

//定义一个结构体当做队列，里面两个属性，一个泛型切片和一个用于表示队列当前长度的整型
type CQueue struct {

arr []interface{}
size int

}


// func Constructor() CQueue {
// var queue = CQueue{} //实例
// 	return queue

// }


func (queue *CQueue) AppendTail(value int)  {
queue.arr = append(queue.arr, value)//因为没有定义队列容量有上限
	queue.size++
}


func (queue *CQueue) DeleteHead() int {

if queue.size == 0 {
		return -1
	}
	if queue.size == 1 {
		queue.size--
		first := queue.arr[0].(int)//标注类型
		queue.arr = queue.arr[0:0]
		return first
	}
	//first := queue.arr[0]
	first := queue.arr[0].(int)
	queue.arr = queue.arr[1:]
	queue.size--

	return first
}

type MinStack struct {

arr  []interface{}
size int
min  []interface{}  

}


/** initialize your data structure here. */
func Constructor() MinStack {
var minStack = MinStack{}
	return minStack
}


func (this *MinStack) Push(x int)  {
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


func (this *MinStack) Pop()  {
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



func ReversePrint(head *ListNode) []int {

	var array []int
	array = append(array,head.Val)
	if head.Next == nil {	
		return array
	}

    tmp := head

    for {
	if tmp.Next == nil {
		break
	}
	tmp = tmp.Next
	array = append(array,tmp.Val)
    }
    return ReverseArr(array)
}

func ReverseArr(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	l :=0
	r :=len(nums)-1

	for l<r {
		nums[l],nums[r] = nums[r],nums[l]
		l ++
		r --
	}

	return nums
}

func ReverseList(head *ListNode) *ListNode {

	if head == nil {
			fmt.Println("the List is empty")
			return head
		}
	
		tmp := head
		cur := head
		var res *ListNode
		
	
		if tmp.Next != nil {
			cur = head.Next
		}
	
		for {
			if cur.Next == nil {
				res = cur
				break
			}
			cur = cur.Next
			tmp = tmp.Next
		}
	
		for {
			cur.Next = tmp
			cur = cur.Next
			if cur == head {
				cur.Next=nil
				break
			}
	
			tmp = head
			for {
				if tmp.Next == cur {
					break
				}
				tmp = tmp.Next
			}
		}
		return res
	}
	
func CopyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	 }
  
	 newHead := &Node{}
	 tmp := head
	  newHead.Val = tmp.Val
  
	  cur := newHead
  
	  if tmp.Next == nil {
		  
		  if tmp.Random == nil {
			  cur.Random = nil
		  }
		  if tmp.Random == tmp {
			  cur.Random = cur
		  }
		  
		  return newHead
	  }
  
	  for {
		  if tmp.Next == nil {
			  break
		  }
		  //新建
		  newNode := &Node{}
		  //初始化
		  newNode.Val = tmp.Next.Val
		  cur.Next = newNode
  
		  tmp = tmp.Next
		  cur = cur.Next
	  }
  
	  tmp = head
	  cur = newHead
  
	  for {
		  
		  if tmp.Random == nil {
			  cur.Random = nil
		  }
		  if tmp.Random == tmp {
			  cur.Random = cur
		  }
  
		  p := head
		  q := newHead
		  for {
			  if tmp.Random == p {
				  cur.Random = q
				  break
			  }
			  p = p.Next
			  q = q.Next
		  }
  
		  fmt.Println(cur.Val)
		if tmp.Next == nil {
			  break
		  }
  
		  tmp = tmp.Next
		  cur = cur.Next
	  }
	  return newHead
}
  
func exchange(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	} 

	l := 0
	r := len(nums)-1

	for l<r {
		if nums[l] % 2 == 1 {
			l ++ 
		}else {
			if nums[r] % 2 == 0 {
				r --
			}else {
				nums[l],nums[r] = nums[r], nums[l]
			}
		}
	}
	return nums
}

func twoSum(nums []int, target int) []int {

	
		left :=0
		right :=len(nums)-1
		for left<right{
			 sum := nums[left]+nums[right]
			if sum==target{
				return []int{nums[left],nums[right]}
			}else if sum>target{
				right--
			}else {
				left++
			}
		}
		return nil
	}


func reverseWords(s string) string {

	var res string
    var i = len(s) - 1
    var j = i
    for i >= 0{
        for i >=0 && s[i] == ' '{
            i--
        }
        j = i
        for i >= 0 && s[i] != ' '{
            i--
        }
        res += s[i+1:j+1] + " "
    }
    return strings.TrimRight(res, " ")
}