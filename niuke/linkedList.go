package niuke

type ListNode struct {
	Val int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	
	if pHead == nil {
		return pHead
	}

	tmp := pHead
	cur := pHead

	var res *ListNode

	if tmp.Next != nil {
		cur = pHead.Next
	}

	for {

		if cur.Next == nil {
			res = cur
			break
		}
		cur = cur.Next
		tmp = tmp.Next
	}

	for {

		cur.Next = tmp
		cur = cur.Next

		if cur == pHead {
			cur.Next = nil
			break
		}

		tmp = pHead

		for {
			if tmp.Next == cur {
				break
			}
			tmp = tmp.Next
		}
	}

	return res
}