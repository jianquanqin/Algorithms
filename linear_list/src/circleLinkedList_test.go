package src

import (
	"fmt"
	"testing"
)

func TestCircleLinkedList(t *testing.T) {

	//初始化头节点
	//环形链表的头节点需要存放数据
	head := &CatNode{}
	cat1 := &CatNode{no: 1,
		name: "jerry",
	}
	cat2 := &CatNode{no: 2,
		name: "tom",
	}

	//var a CircleOperations
	//a.InsertCatNode(head, cat1)

	head.InsertCatNode(head, cat1)
	head.InsertCatNode(head, cat2)

	fmt.Println(head)
	head.PrintCatNodeList(head)
}
