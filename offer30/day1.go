package main

import "fmt"

type CQueue struct {
	arr []interface{}
	size int
}


func main() {





}

func CreateQueue() CQueue {
	var queue = CQueue{}
	return queue

	
}

func (queue *CQueue)appendTail(val interface{})  {

	queue.arr = append(queue.arr, val)
	queue.size ++

}

func (queue *CQueue)deleteHead() int{

	first := 0
	if queue.size == 0 {
		first = -1
		return first
	}

	if queue.size == 1 {
		first = queue.arr[0].(int)
		queue.arr = queue.arr[0:0]
		queue.size --
			
	}else {
		first = queue.arr[0].(int)
		queue.arr = queue.arr[1:]
		queue.size --		
	}
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
	