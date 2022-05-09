package src

import (
	"fmt"
	"testing"
)

func TestDuLinkedListFunction(t *testing.T) {

	head := &DuLNode{} //头节点为空

	node1 := &DuLNode{Data: "宋江", Order: 1, Pre: nil, Next: nil}
	node2 := &DuLNode{Data: "花荣", Order: 2, Pre: nil, Next: nil}
	node3 := &DuLNode{Data: "吴用", Order: 3, Pre: nil, Next: nil}
	node4 := &DuLNode{Data: "李逵", Order: 5, Pre: nil, Next: nil}
	node5 := &DuLNode{Data: "李逵", Order: 4, Pre: nil, Next: nil}

	////test尾插
	//InsertDuLinkedList(head, node1)
	//InsertDuLinkedList(head, node2)
	//InsertDuLinkedList(head, node3)
	//InsertDuLinkedList(head, node4)

	//test逆序打印
	//ReverseListDuLinkedList(head)

	//test顺序插入
	InsertDuLinkedListWithCondition(head, node4)
	InsertDuLinkedListWithCondition(head, node1)
	InsertDuLinkedListWithCondition(head, node3)
	InsertDuLinkedListWithCondition(head, node5)
	InsertDuLinkedListWithCondition(head, node2)

	//test顺序打印
	ListDuLinkedList(head)
	fmt.Println()

	//test删除
	DeleteDuLinkedList(head, 1)
	ListDuLinkedList(head)
}
