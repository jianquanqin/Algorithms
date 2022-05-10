package binaryTree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {

	//data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//root := ArrayToTree(data, 0, len(data)-1)
	//
	//gotype.PrintTreeMidOrder(root)
	//PrintTreeLayer(root)

	////求二叉树的最大子树
	//root := CreateTree()
	//maxRoot := &gotype.BNode{}
	//FindMaxSubTree(root, maxRoot)
	//fmt.Println(maxSum, maxRoot.Data)

	//求二叉树的高
	//h := TreeHeight(root)
	//fmt.Println(h)

	//判断两棵树是否相等
	root1 := CreateTree()
	root2 := CreateTree()

	fmt.Println(IsEqual(root1, root2))
}
