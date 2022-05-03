package src

//经典排序算法之：1.冒泡排序
//时间复杂度为O(n^2)，最快1趟，最慢n趟
//优化：若某一趟未发生排序，则说明已经排好，使用flag优化

func BubbleSortLinearList(nums []int) {

	for i := 0; i < len(nums); i++ {
		//每一次遍历时都会初始化
		//使用flag优化
		isChange := false
		//只需要依次比较n-1个数即可
		for j := 0; j < len(nums)-2; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isChange = true
			}
		}
		if isChange == false {
			//fmt.Println(nums) //演示排序过程，如果排好提前结束
			return
		}
	}
}

//范例 1 ：31. 下一个排列

func nextPermutation(nums []int) {

	//找到调换位,然后再重排，再将次小的与之调换
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j > 0; j-- {
			//此处的规则可以强制i走在j之前，够让当前元素与所有后面的元素相比较
			if nums[i] < nums[j] && j > i {
				nums[j], nums[i] = nums[i], nums[j]

				//使用冒泡排序
				BubbleSortLinearList(nums[i+1:])
				return
			}
		}
	}
	BubbleSortLinearList(nums)
	return
}
