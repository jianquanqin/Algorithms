package main


//双指针 难！！

func  main() {



}

type ListNode struct {
	Val int
	Next *ListNode

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{   //递归终止条件
	    return l2
	} 
	if l2==nil{    //递归终止条件
	    return l1 
	}
	if l1.Val>l2.Val{  //交换两节点
	    l1,l2=l2,l1
	}
	l1.Next=mergeTwoLists(l1.Next,l2)  //递归
	return l1
 }

 //公共节点
 //笔记，避免在链表上直接进行循环，要使用其他数据结构，如map

 func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//定义了一个map
	vis := map[*ListNode]bool{}

	//遍历链表A，装入map，所有key都为true
	for tmp := headA; tmp != nil; tmp = tmp.Next {
	    vis[tmp] = true
	}
	//遍历链表B，如果在map中发现同样的key对应的value为true，则说明两个链表相交
	for tmp := headB; tmp != nil; tmp = tmp.Next {
	    if vis[tmp] {
		   return tmp
	    }
	}
	//没找到返回nil
	return nil
 }