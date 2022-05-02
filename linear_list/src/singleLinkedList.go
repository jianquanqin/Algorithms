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
