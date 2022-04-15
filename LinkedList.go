package src

import (
	"fmt"
)

//func main() {
//	head := &src.HeroNode{}
//
//	hero1 := &src.HeroNode{1, "宋江", "及时雨", nil}
//	hero2 := &src.HeroNode{2, "卢俊义", "玉麒麟", nil}
//	hero3 := &src.HeroNode{3, "吴用", "智多星", nil}
//	hero4 := &src.HeroNode{4, "公孙胜", "入云龙", nil}
//
//	src.InsertHeroNodeWithOrder(head, hero2)
//	src.InsertHeroNodeWithOrder(head, hero1)
//	src.InsertHeroNodeWithOrder(head, hero3)
//	src.InsertHeroNodeWithOrder(head, hero4)
//	src.ListNodes(head)
//
//	src.DeleteHeroNode(head, 3)
//	src.ListNodes(head)
//}

type HeroNode struct {
	No       int64
	Name     string
	Nickname string
	Next     *HeroNode
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {

	tmp := head //define a temporary variable

	for {
		//when the up to the last node,break
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	tmp.Next = newHeroNode //add a new node
}

func InsertHeroNodeWithOrder(head *HeroNode, newHeroNode *HeroNode) {

	tmp := head //define a temporary variable
	isNotExisted := true

	for {
		//when the up to the last node,break
		if tmp.Next == nil {
			break
		} else if tmp.Next.No > newHeroNode.No {
			break
		} else if tmp.Next.No == newHeroNode.No {
			isNotExisted = false
			break
		}
		tmp = tmp.Next

	}
	if !isNotExisted {

		fmt.Printf("sorry, the %d existed", newHeroNode.No)
		return
	} else {
		newHeroNode.Next = tmp.Next
		tmp.Next = newHeroNode
	}
}

func ListNodes(head *HeroNode) {

	tmp := head //auxiliary node
	if tmp.Next == nil {
		fmt.Println("the linked table is null")
		return
	}

	for {
		fmt.Printf("[%v,%s,%s] -> ",
			tmp.Next.No,
			tmp.Next.Name,
			tmp.Next.Nickname,
		)

		tmp = tmp.Next

		if tmp.Next == nil {
			break
		}
	}
}

func DeleteHeroNode(head *HeroNode, no int64) {
	tmp := head

	flag := false

	for {
		if tmp.Next == nil {
			break
		} else if tmp.Next.No == no {
			flag = true
			break
		}
		tmp = tmp.Next
	}

	if flag {
		tmp.Next = tmp.Next.Next
	} else {
		fmt.Println("the no didn't existed")
	}
}
