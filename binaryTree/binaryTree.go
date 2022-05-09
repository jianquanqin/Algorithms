package binaryTree

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//二叉树的链式表示

type BNode struct {
	Data interface{} //这里用接口替代泛型
	LC   *BNode
	RC   *BNode
}

//将数组存放到二叉树中
//二叉树的从上到下，从左到右构建

func ArrayToTree(arr []int, start int, end int) *gotype.BNode {

	var root *gotype.BNode

	if end >= start {

		root = gotype.NewBNode()

		mid := (end + start + 1) / 2
		root.Data = arr[mid]

		//递归调用自身
		root.LeftChild = ArrayToTree(arr, start, mid-1)
		root.RightChild = ArrayToTree(arr, mid+1, end)
	}
	return root
}

//层序遍历二叉树
//链表这种数据结构头节点非常重要
//时间复杂度是O(n)
//空间复杂度，某一层的最大值

//data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func PrintTreeLayer(root *gotype.BNode) {
	if root == nil {
		fmt.Println("二叉树为空")
		return
	}

	//新建一个变量，用来调用方法
	var p *gotype.BNode

	//新建一个队列，用来存放遍历到的节点
	queue := gotype.NewSliceQueue()
	//先把根结送入队列，其实就定义了这是一个存放二叉树的队列
	queue.EnQueue(root)

	for queue.Size() > 0 {

		//在接口里放入类型，接口的类型就确定了
		//初始化p的值

		p = queue.DeQueue().(*gotype.BNode)

		fmt.Print(p.Data, "")

		//如果孩子不为空，则将左孩子送入队列中
		if p.LeftChild != nil {
			queue.EnQueue(p.LeftChild)
		}
		if p.RightChild != nil {
			queue.EnQueue(p.RightChild)
		}
	}
}

//
func PrintAtLevel(root *gotype.BNode, level int) int {
	return 1
}
