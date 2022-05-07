package src

import "fmt"

//定义节点

type CatNode struct {
	no       int
	name     string
	nextNode *CatNode
}

type CircleOperations interface {
	InsertCatNode(head *CatNode, newCatNode *CatNode)
	PrintCatNodeList(head *CatNode)
}

func (catNode *CatNode) InsertCatNode(head *CatNode, newCatNode *CatNode) {

	//判断是不是需要添加第一个节点
	if head.nextNode == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.nextNode = head //构成一个环形

		fmt.Println("the first cat is added")
		return
	}

	//定义一个临时变量
	tmp := head
	for {
		if tmp.nextNode == head {
			//此时tmp会到最后一个节点
			break
		}
		tmp = tmp.nextNode
	}

	tmp.nextNode = newCatNode
	newCatNode.nextNode = head
}

func (catNode *CatNode) PrintCatNodeList(head *CatNode) {

	tmp := head

	if tmp.nextNode == nil {
		fmt.Println("the circle link list is empty")
		return
	}

	for {
		if tmp.nextNode == head {
			break
		}
		fmt.Printf("[no: %d ,name: %v, ->]", tmp.no, tmp.name)
		tmp = tmp.nextNode
		if tmp.nextNode == head {
			break
		}
	}
}
