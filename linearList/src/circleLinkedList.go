package src

import "fmt"

//定义节点

type CatNode struct {
	no       int
	name     string
	nextNode *CatNode
}

//定义接口把方法包含进去

type CircleOperations interface {
	InsertCatNode(head *CatNode, newCatNode *CatNode)
	PrintCatNodeList(head *CatNode)
	DeleteCatNode(head *CatNode, i int)
}

func (catNode *CatNode) InsertCatNode(head *CatNode, newCatNode *CatNode) {

	//判断是不是只有头节点,如果是的话，将新节点的数据填入头节点（初始化头节点）
	if head.nextNode == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.nextNode = head //构成一个环形

		fmt.Println("the first cat node is added")
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

	//添加节点
	tmp.nextNode = newCatNode
	newCatNode.nextNode = head
	fmt.Println("the new node has been added")
}

func (catNode *CatNode) PrintCatNodeList(head *CatNode) {

	tmp := head
	//判空
	if tmp.nextNode == nil {
		fmt.Println("the circle link list is empty")
		return
	}

	//
	for {
		//先打印当前节点信息
		fmt.Printf("[no: %d ,name: %v, %v]", tmp.no, tmp.name, tmp.nextNode.no)
		if tmp.nextNode == head {
			break
		} else {
			tmp = tmp.nextNode
		}
	}
}

func (catNode *CatNode) DeleteCatNode(head *CatNode, id int) {
	fmt.Println()

	tmp := head
	helper := head
	//判空
	if tmp.nextNode == nil {
		fmt.Println("there is no node to delete")
		return
	}

	for {
		//比较是否是要删除的节点
		if tmp.no == id {
			//先看一下是不是头节点
			if tmp == head {
				//再判断是否是只有一个节点
				if tmp.nextNode == head {
					tmp.nextNode = nil
					fmt.Println("the head has been deleted")
					return
				} else {
					//本质上没有删除头节点，而是把头节点的数据删除了，并且换成了下一个节点的数据，并将其删除
					tmp.no = tmp.nextNode.no
					tmp.name = tmp.nextNode.name
					tmp.nextNode = tmp.nextNode.nextNode
					fmt.Println("the head has been deleted")
					return
				}
			} else {
				for {
					//如果不是头节点，让helper指向它的上一个节点
					if helper.nextNode == tmp {
						helper.nextNode = tmp.nextNode
						fmt.Println("the node was delete is:", tmp.no)
						return
					}
					helper = helper.nextNode
				}
			}
		}
		tmp = tmp.nextNode
	}
}

//使用环形链表求解约瑟夫问题

type BoyNode struct {
	no   int
	next *BoyNode
}

//构建环形链表

func ConstructBoyList(num int) *BoyNode {

	first := &BoyNode{} //定义头节点,注意，此时的头节点需要值

	for i := 1; i <= num; i++ {

		tmp := first
		boy := &BoyNode{no: i}

		//判断头节点是否为空，空的话把第一个元素装入头节点
		if first.no == 0 {
			first = boy
			first.next = first
			fmt.Println(first)
			//依次在后面添加元素
		} else {
			for {
				if tmp.next == first {
					tmp.next = boy
					boy.next = first
					//fmt.Println("added")
					break
				}
				tmp = tmp.next
			}
		}

	}
	return first
}

//打印

func PrintBoyList(first *BoyNode) {

	tmp := first
	if first.no == 0 {
		fmt.Println("the list is empty")
	}
	for {
		fmt.Printf("[no:%d,nextNode.no:%d]", tmp.no, tmp.next.no)
		if tmp.next == first {
			break
		}
		tmp = tmp.next
	}
}

//求解

func PlayGame(first *BoyNode, startNo, count int) {

	if first.next == nil {
		fmt.Println("the boyList is empty")
		return
	}

	//定义两个指针 first 和 tail

	tail := first

	//将tail移动到链尾

	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}

	//一起向前移动，移动到随机开始数数的位置

	for i := 1; i <= startNo-1; i++ {
		first = first.next
		tail = tail.next
	}

	//开始游戏

	for {

		//数数
		for j := 1; j <= count-1; j++ {
			first = first.next
			tail = tail.next
		}

		//删除；重点
		fmt.Println("出列的是：", first.no)
		//不是直接将头节点删除，而是将它变更为别的节点，删掉原来的替身
		first = first.next
		tail.next = first

		//判断是否链表中只剩一个元素
		if first == tail {
			break
		}
	}

	//输出最后一个元素
	fmt.Println("最后出列的是：", first.no)
}
