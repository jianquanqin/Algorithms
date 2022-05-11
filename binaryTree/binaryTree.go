package binaryTree

import (
	"fmt"
	"github.com/isdamir/gotype"
	"math"
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

//求二叉树的最大子树
//找一个叶子节点测试第一次调用本函数，每一次调用都会为参数和变量开辟独立空间，但是全局参数和变量可以共用
//事实上，递归函数只需要考虑一次调用就够了，并且无论采取什么样的base case,它都会递归完整棵树

var maxSum = math.MinInt64
var tmp []int //可以定义一个全局变量，存储每次递归的结果

func FindMaxSubTree(root *gotype.BNode, maxRoot *gotype.BNode) int {

	//1.base case
	if root == nil {
		return 0
	}

	//2.recursion relation

	//求左子树的和
	lmax := FindMaxSubTree(root.LeftChild, maxRoot)
	rmax := FindMaxSubTree(root.RightChild, maxRoot)

	//注意这里的指定数据类型
	sum := lmax + rmax + root.Data.(int)

	//如果所加之和大于前面一棵子树，则更新当前子树为最大子树
	if sum > maxSum {
		maxSum = sum
		tmp = append(tmp, sum)
		maxRoot.Data = root.Data
	}
	fmt.Println(tmp)
	fmt.Println(sum)
	return sum
}

func CreateTree() *gotype.BNode {

	root := &gotype.BNode{Data: 6}
	node1 := &gotype.BNode{Data: 3}
	node2 := &gotype.BNode{Data: -7}
	node3 := &gotype.BNode{Data: -1}
	node4 := &gotype.BNode{Data: 9}
	node5 := &gotype.BNode{Data: 9}

	root.LeftChild = node1
	root.RightChild = node2
	node1.LeftChild = node3
	node1.RightChild = node4
	node3.LeftChild = node5

	return root
}

//求解二叉树的高h

var height = 0

func TreeHeight(root *gotype.BNode) int {

	//如果根节点为空，高度为0
	if root == nil {
		return 0
	}

	//如果根节点左右子树都为空，说明到了叶子节点，高度为1
	if root.LeftChild == nil && root.RightChild == nil {
		return 1
	}

	//子树有左节点或者右节点
	//求左子树的高
	Lh := TreeHeight(root.LeftChild)
	Rh := TreeHeight(root.LeftChild)

	if Lh > Rh {
		height = Lh + 1
	} else {
		height = Rh + 1
	}

	fmt.Println(root.Data, height)

	return height
}

//判断两棵树是否相等
//还是拿叶子节点和上一个节点来构建程序

func IsEqual(root1 *gotype.BNode, root2 *gotype.BNode) bool {

	//判空
	//两个都是空树必相等
	if root1 == nil && root2 == nil {
		return true
	}

	//一个是空树，另一个不是空树，不相等
	if root1 == nil && root2 != nil {
		return false
	}
	if root2 == nil && root1 != nil {
		return false
	}

	//如果两个树都不为空,比较它的根节点值与左右子树的值
	if root1.Data == root2.Data {
		return IsEqual(root1.LeftChild, root2.LeftChild) && IsEqual(root1.RightChild, root2.RightChild)
	}

	//如果不等则返回false
	return false
}

//递归求二叉树的节点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count int

func countNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	//第一种方法
	//rootCount := 1
	//fmt.Println("root:", root.Val)
	//
	//lCount := countNodes(root.Left)
	//rCount := countNodes(root.Right)
	//
	//count = rootCount + lCount + rCount
	//
	//return count

	//第二种方法
	Llevel := BinaryLTreeHeight(root.Left)
	Rlevel := BinaryRTreeHeight(root.Right)
	fmt.Println(Llevel, Rlevel)

	if Llevel == Rlevel {
		count = (1 << Llevel) + countNodes(root.Right)
	} else {
		count = countNodes(root.Left) + (1 << Rlevel)
	}
	return count
}

func CreateATree() *TreeNode {

	root1 := &TreeNode{Val: 1}
	root2 := &TreeNode{Val: 2}
	root3 := &TreeNode{Val: 3}
	root4 := &TreeNode{Val: 4}
	root5 := &TreeNode{Val: 5}
	root6 := &TreeNode{Val: 6}
	//root7 := &TreeNode{Val: 6}

	root := root1
	root.Left = root2
	root.Right = root3
	root2.Left = root4
	root2.Right = root5
	root3.Left = root6
	//root3.Right = root7

	return root
}

func BinaryLTreeHeight(root *TreeNode) (lH int) {

	if root == nil {
		return 0
	}
	h := 0

	for root != nil {
		root = root.Left
		h++
	}
	return h
}

func BinaryRTreeHeight(root *TreeNode) (rH int) {

	if root == nil {
		return 0
	}
	h := 0

	for root != nil {
		root = root.Right
		h++
	}
	return h
}
