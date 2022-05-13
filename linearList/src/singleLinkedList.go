package src

import (
	"fmt"
	"log"
)

//1.先定义节点

type LNode struct {
	Data  string
	Order int
	Next  *LNode
}

//1.从队尾插入

func InsertLinkedList(head *LNode, newNode *LNode) {

	//1.创建辅助节点
	tmp := head

	//2.遍历链表，找到最后一个节点
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}

	//3.将新节点加入到链尾
	tmp.Next = newNode

}

//2.显示链表

func ListLinkedList(head *LNode) {
	//1.创建辅助节点
	tmp := head
	//2.判断表是否为空
	if tmp.Next == nil {
		fmt.Println("list is empty")
		return
	}

	//3.遍历链表，找到最后一个节点
	for {
		//从第一个节点开始打印
		fmt.Printf("[%s,%v]->", tmp.Next.Data, tmp.Next.Order)
		//打印之后指向下一个节点
		tmp = tmp.Next
		//判断下下一个节点是否为空
		if tmp.Next == nil {
			break
		}
	}
}

//3.按照条件插入

func InsertLinkedListWithCondition(head *LNode, newNode *LNode) {

	//1.创建辅助节点
	tmp := head
	flag := true

	//2.遍历链表，找到最后一个节点
	for {
		if tmp.Next == nil { //空表直接加
			break
		} else if tmp.Next.Order > newNode.Order { //
			break
		} else if tmp.Next.Order == newNode.Order {
			flag = false
			break
		}
		tmp = tmp.Next
	}

	if !flag {
		fmt.Println("can not insert, the order existed")
		log.Fatalln()
	} else {
		newNode.Next = tmp.Next
		tmp.Next = newNode
		return
	}
}

//4.删除

func DeleteLinkedList(head *LNode, id int) {

	//1.创建辅助节点
	tmp := head
	flag := false

	//2.遍历链表，找到最后一个节点
	for {

		if tmp.Next == nil {
			break
		} else if tmp.Next.Order == id {
			flag = true
			break
		}
		tmp = tmp.Next
	}

	if flag {
		tmp.Next = tmp.Next.Next
	} else {
		fmt.Println("the id didn't exit")
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	nums1 := Range(l1)
	nums2 := Range(l2)

	nums := Sum(nums1, nums2)

	list := ConstructList(nums)
	return list
}

func Range(l *ListNode) []int {

	tmp := l
	var res []int

	for {

		res = append(res, tmp.Val)

		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	return res
}

func ConstructList(nums []int) *ListNode {

	firstNode := &ListNode{}
	flag := false

	for i := 0; i < len(nums); i++ {

		node := &ListNode{Val: nums[i]}
		tmp := firstNode

		if firstNode.Val == 0 && flag == false {
			firstNode = node
			flag = true
			//fmt.Println("first node added")
		} else {
			for {
				if tmp.Next == nil {
					tmp.Next = node
					//fmt.Println("node added")
					break
				}
				tmp = tmp.Next
			}
		}
	}

	//fmt.Println(firstNode)
	return firstNode
}

func Sum(nums1 []int, nums2 []int) []int {

	var test []int

	min := func(nums1 []int, nums2 []int) int {

		if len(nums1) > len(nums2) {
			return len(nums2)
		} else {
			return len(nums1)
		}
	}
	l := min(nums1, nums2)

	max := func(nums1 []int, nums2 []int) []int {

		if len(nums1) > len(nums2) {
			return nums1
		} else {
			return nums2
		}
	}
	m := max(nums1, nums2)

	in := 0
	sum := 0
	for i := 0; i < len(m); i++ {
		if i < l {
			sum = nums1[i] + nums2[i] + in
			if sum >= 10 {
				in = 1
				test = append(test, sum-10)
			} else {
				in = 0
				test = append(test, sum)
			}
		} else {
			sum = m[i] + in
			if sum >= 10 {
				in = 1
				test = append(test, sum-10)
			} else {
				in = 0
				test = append(test, sum)
			}
		}
	}
	if in == 1 {
		test = append(test, in)
	}
	//fmt.Println(test)
	return test
}

func ListList(head *ListNode) {
	//1.创建辅助节点
	tmp := head
	//2.判断表是否为空
	if tmp.Next == nil {
		fmt.Printf("[%v]->", tmp.Val)
		fmt.Println("list is empty")
		return
	}

	//3.遍历链表，找到最后一个节点
	for {
		//从第一个节点开始打印
		fmt.Printf("[%v]->", tmp.Val)
		//判断下下一个节点是否为空
		if tmp.Next == nil {
			break
		}
		//打印之后指向下一个节点
		tmp = tmp.Next
	}
}

//19. 删除链表的倒数第 N 个结点

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	tmp := head
	count := 1
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
		count++
	}

	tmp = head
	if count-n == 0 {
		if tmp.Next != nil {
			head = tmp.Next
			head.Next = tmp.Next.Next
		} else {
			newhead := head.Next
			head.Next = nil
			head = newhead
		}
	} else {
		for i := 1; i < count-n; i++ {
			tmp = tmp.Next
		}
		tmp.Next = tmp.Next.Next
	}
	return head
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//剑指 Offer 35. 复杂链表的复制

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	tmp := head
	newHead := &Node{}
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

//从尾到头打印链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reversePrint(head *ListNode) []int {

	var cache []int
	if head == nil {
		fmt.Println("the List is empty")
		return cache
	}

	tmp := head

	for {

		cache = append(cache, tmp.Val)

		if tmp.Next == nil {
			break
		}

		tmp = tmp.Next
	}

	l := 0
	r := len(cache) - 1
	for l < r {
		cache[l], cache[r] = cache[r], cache[l]
		l++
		r--
	}
	return cache
}

//反转链表

func reverseList(head *ListNode) *ListNode {
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
			cur.Next = nil
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
