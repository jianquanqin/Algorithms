package src

import (
	"testing"
)

//31. 下一个排列

func TestNextPermutation(t *testing.T) {

	nums := []int{1, 2}
	nextPermutation(nums)
	if nums[0] != 1 {
		t.Fatalf("test failed, expected value;%v,actual value:%v", []int{3, 2, 1}, nums)
		return
	}
	t.Logf("pass")
}
