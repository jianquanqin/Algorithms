package main

func main() {}

//判断是否有重复元素

func containsDuplicate(nums []int) bool {

	m := make(map[int]bool)

	for _,v := range nums {

		if m[v] == true {
			return true
		}else {
			m[v] = true
		}

	}
	return false

}

//求最大子序列和

func maxSubArray(nums []int) int {

	

}