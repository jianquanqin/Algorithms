package src

import (
	"testing"
)

func TestCircleLinkedList(t *testing.T) {

	////初始化头节点
	////环形链表的头节点需要存放数据
	//head := &CatNode{}
	//cat1, cat2, cat3 := &CatNode{no: 1,
	//	name: "jerry",
	//}, &CatNode{no: 2,
	//	name: "tom",
	//}, &CatNode{no: 3,
	//	name: "bob",
	//}
	//
	//cat1.InsertCatNode(head, cat1)
	//cat2.InsertCatNode(head, cat2)
	//cat2.InsertCatNode(head, cat3)
	//head.PrintCatNodeList(head)
	//head.DeleteCatNode(head, 3)
	//cat2.PrintCatNodeList(head)

	//约瑟夫链表测试

	boyList := ConstructBoyList(50)
	PrintBoyList(boyList)

	PlayGame(boyList, 7, 13)
}
