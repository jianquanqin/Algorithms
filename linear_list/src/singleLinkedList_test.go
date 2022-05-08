package src

import (
	"fmt"
	"testing"
)

func TestLinkedListFunction(t *testing.T) {

	head := &LNode{} //头节点为空

	node1 := &LNode{Data: "宋江", Order: 1, Next: nil}
	node2 := &LNode{Data: "卢俊义", Order: 3, Next: nil}
	node3 := &LNode{Data: "吴用", Order: 2, Next: nil}
	node4 := &LNode{Data: "李逵", Order: 4, Next: nil}

	////test
	//InsertLinkedList(head, node1)
	//InsertLinkedList(head, node2)
	//InsertLinkedList(head, node3)
	//InsertLinkedList(head, node4)

	//test
	InsertLinkedListWithCondition(head, node3)
	InsertLinkedListWithCondition(head, node1)
	InsertLinkedListWithCondition(head, node2)
	InsertLinkedListWithCondition(head, node4)

	//test
	DeleteLinkedList(head, 3)

	ListLinkedList(head)

	fmt.Println("---")
}

func TestLeetCode(t *testing.T) {

	//nums1 := []int{1}
	//nums1 := []int{}
	//nums2 := []int{0}

	//list1 := ConstructList(nums1)
	//list2 := ConstructList(nums2)

	//result := addTwoNumbers(list1, list2)
	//ListList(result)

	//list := ConstructList(nums1)
	//ListList(list)
	//removeNthFromEnd(list, 1)
	//
	//fmt.Println(list)
	//ListList(list)
}
