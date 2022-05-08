package src

//定义节点

//type CatNode struct {
//	no       int
//	name     string
//	preNode  *CatNode
//	nextNode *CatNode
//}
//
////定义接口把方法包含进去
//
//type DoubleCircleOperations interface {
//	InsertCatNode(head *CatNode, newCatNode *CatNode)
//	PrintCatNodeList(head *CatNode)
//	DeleteCatNode(head *CatNode, i int)
//}
//
//func (catNode *CatNode) InsertCatNode(head *CatNode, newCatNode *CatNode) {
//
//	//判断是不是只有头节点,如果是的话，将新节点的数据填入头节点（初始化头节点）
//	if head.nextNode == nil {
//		head.no = newCatNode.no
//		head.name = newCatNode.name
//		head.preNode = head
//		head.nextNode = head //构成一个环形
//
//		fmt.Println("the first cat node is added")
//		return
//	}
//
//	//定义一个临时变量
//	tmp := head
//	for {
//		if tmp.nextNode == head {
//			//此时tmp会到最后一个节点
//			break
//		}
//		tmp = tmp.nextNode
//	}
//
//	//添加节点
//	tmp.nextNode = newCatNode
//	newCatNode.nextNode = head
//	newCatNode.preNode = tmp
//	head.preNode = newCatNode
//
//	fmt.Println("the new node has been added")
//}
//
//func (catNode *CatNode) PrintCatNodeList(head *CatNode) {
//
//	tmp := head
//	//判空
//	if tmp.nextNode == nil {
//		fmt.Println("the circle link list is empty")
//		return
//	}
//
//	//
//	for {
//		//先打印当前节点信息
//		fmt.Printf("[no: %d ,name: %v ,preNode:%v ,nextNode:%v]", tmp.no, tmp.name, tmp.preNode.no, tmp.nextNode.no)
//		if tmp.nextNode == head {
//			break
//		} else {
//			tmp = tmp.nextNode
//		}
//	}
//}
//
//func (catNode *CatNode) DeleteCatNode(head *CatNode, id int) {
//	fmt.Println()
//
//	tmp := head
//	helper := head
//	//判空
//	if tmp.nextNode == nil {
//		fmt.Println("there is no node to delete")
//		return
//	}
//
//	for {
//		//比较是否是要删除的节点
//		if tmp.no == id {
//			//先看一下是不是头节点
//			if tmp == head {
//				//再判断是否是只有一个节点
//				if tmp.nextNode == head {
//					tmp.nextNode = nil
//					tmp.preNode = nil
//					fmt.Println("the head has been deleted")
//					return
//				} else {
//					//本质上没有删除头节点，而是把头节点的数据删除了，并且换成了下一个节点的数据，并将其删除
//					tmp.no = tmp.nextNode.no
//					tmp.name = tmp.nextNode.name
//					tmp.nextNode = tmp.nextNode.nextNode
//					tmp.nextNode.preNode = tmp
//					fmt.Println("the head has been deleted")
//					return
//				}
//			} else {
//				for {
//					//如果不是头节点，让helper指向它的上一个节点
//					if helper.nextNode == tmp {
//						helper.nextNode = tmp.nextNode
//						tmp.nextNode.preNode = helper
//						fmt.Println("the node was delete is:", tmp.no)
//						return
//					}
//					helper = helper.nextNode
//				}
//			}
//		}
//		tmp = tmp.nextNode
//	}
//}
