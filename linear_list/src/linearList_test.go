package src

import (
	"fmt"
	"testing"
)

//2.将元素插入数组中

func TestInsertDynamicLinearList(t *testing.T) {

	linearList := DynamicLinearList{
		IntList: []int{1, 2, 3, 4, 5},
		MaxSize: 5,
		Len:     5,
	}
	linearList.InsertDynamicLinearList(4, 10)
	if linearList.IntList[3] != 10 {
		t.Fatalf("test failed, expected value;%v,expected value:%v", 10, linearList.IntList[3])
		return
	}
	fmt.Println(linearList)
	t.Logf("pass")
}

func TestInsertLinearList(t *testing.T) {

	linearList := LinearList{
		IntList: [10]int{1, 2, 3, 4, 5},
		Len:     5,
	}
	linearList.InsertLinearList(6, 7)
	if linearList.IntList[5] != 7 {
		t.Fatalf("test failed, expected value;%v,expected value:%v", 7, linearList.IntList[5])
		return
	}
	t.Logf("pass")

}

func TestDeleteLinearList(t *testing.T) {
	linearList := LinearList{
		IntList: [10]int{1, 2, 3, 4, 5},
		Len:     5,
	}
	linearList.DeleteLinearList(5)
	if linearList.IntList[3] != 4 {
		t.Fatalf("test failed, expected value;%v,expected value:%v", 4, linearList.IntList[3])
		return
	}
	fmt.Println(linearList.IntList)
	t.Logf("pass")
}
