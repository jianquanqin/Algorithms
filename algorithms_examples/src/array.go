package src

//数独

func solveSudoku(board [][]byte) {
	tmp := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] == '.' {
				for _, t := range tmp {
					for _, u := range board[i] {
						if t != u {
							for k := 0; k < 9; k++ {
								w := board[k][j]
								if t != w {

								}
							}
						}
					}
				}
			}
		}
	}
}

// 39. 组合总和

func combinationSum(candidates []int, target int) [][]int {

	var result [][]int

	bubbleSortLinearList(candidates)
	r := len(candidates) - 1

	for r > 0 {
		if candidates[len(candidates)-1] > target {
			r--
		}
	}
	candidates = candidates[0 : r+1]
	for i, _ := range candidates {
		if candidates[i] == target {
			result = append(result, []int{candidates[i]})
		}
	}
	l := 0
	if candidates[r] != target {

		if candidates[l]+candidates[r] == target {
			result = append(result, []int{candidates[l], candidates[r]})
		} else if candidates[l]+candidates[r] > target {
			r--
		} else {
			l++
		}
	}

	return result
}

func bubbleSortLinearList(nums []int) {

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
