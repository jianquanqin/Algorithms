package main

import "fmt"

func main() {

	
}

type ListNode struct {

	Val int

	Next *ListNode

}

func reversePrint(head *ListNode) []int {
	var array []int
		if head == nil {	
			return array
		}
	
	    tmp := head
	    array = append(array,tmp.Val)
	
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
	

//题目2: 反转链表
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var res = head

	var tmp = head
	var cur = tmp.Next


	//找到最后一个节点
	for {
		if cur.Next == nil {
			res = cur //让最后一个节点成为新的头节点
			cur.Next = tmp
			cur = cur.Next
			break
		}
		cur = cur.Next
		tmp = tmp.Next
	}

	for {
		//辅助节点回到原来的位置
		tmp = head

		for {
			if tmp.Next == cur {
				cur.Next = tmp
				cur = cur.Next
				break
			}
			tmp = tmp.Next
		}
		//如果辅助节点仍然在头节点
		if tmp == head {
			cur.Next = nil
			break
		}

	}
	return res
}

//题目3，复杂链表的复制 ‼️重要

type Node struct {
	Val int
	Next *Node
	Random *Node

}

func copyRandomList(head *Node) *Node {

	if head == nil {
	    return nil
	}
	
	//新建
	newHead := &Node{}
	tmp := head
	 newHead.Val = tmp.Val
 
	 cur := newHead
 
	 
	 //处理第一个节点的情况
	 if tmp.Next == nil {
		 
		 if tmp.Random == nil {
			 cur.Random = nil
		 }
		 if tmp.Random == tmp {
			 cur.Random = cur
		 }
		 
		 return newHead
	 }


	 //造节点
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

	 
	 //初始化
	 tmp = head
	 cur = newHead


 
	 for {
		 
		//两种情况
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