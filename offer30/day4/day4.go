package main

func main() {
}

func findNumberIn2DArray(matrix [][]int, target int) bool {

	//标记行列
	i, j := len(matrix)-1, 0
	//
	for i >= 0 && j < len(matrix[0]) {
		//从最后一行一个一个开始遍历
		if target > matrix[i][j] {
			j++
			//当前行中没有的话找上一行
		} else if target < matrix[i][j] {
			i--
		} else {
			return true
		}
	}
	return false
}