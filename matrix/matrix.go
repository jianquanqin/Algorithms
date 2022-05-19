package matrix

func matrixReshape(mat [][]int, r int, c int) [][]int {

	n, m := len(mat), len(mat[0])

	if n*m != r*c {
		return mat
	}

	//定义长为r的二维数组
	ans := make([][]int, r)

	for i := range ans {
		//定义长为c的一纬数组
		ans[i] = make([]int, c)
	}
	//遍历整个矩阵
	for i := 0; i < n*m; i++ {

		//整数部分是行，余数部分是列
		//这是个矩阵转换公式，记下来
		ans[i/c][i%c] = mat[i/m][i%m]
	}
	return ans
}

//杨辉三角形

func generate(numRows int) [][]int {

	//先创建长度固定的二维数组
	res := make([][]int, numRows)

	for i := range res {
		//创建每一层的数组
		res[i] = make([]int, i+1)
		//数组的两端点固定都为1
		res[i][0] = 1
		res[i][i] = 1

		//计算每一行空缺的数，注意两个加数与和的关系
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

//36. 有效的数独
//再辨析

func isValidSudoku(board [][]byte) bool {

	//9x9
	var rows, columns [9][9]int
	//3x(3x9)
	var subBoxes [3][3][9]int

	//遍历行,找到某一行
	for i, row := range board {
		//再遍历列（其实就是每一个元素）
		for j, column := range row {
			if column == '.' {
				//如果空就继续
				continue
			}
			//如果不为空，column有可能等于'1','2','3',...,例如，若是'1'，则index = 0，说明第i行第j列的元素是'1'
			//对应的index = 0,1,...,8
			index := column - '1'
			//标记，说明该区域有值
			//最后的结果就是某一行中哪些位置被填入了数据，数据值均为1
			rows[i][index]++
			//标记，说明该区域有值
			//最后的结果就是某一列中哪些位置被填入了数据，数据值均为1
			columns[j][index]++

			subBoxes[i/3][j/3][index]++

			if rows[i][index] > 1 || columns[j][index] > 1 || subBoxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

//73. 矩阵置零
//注意标记数组的命名和使用

func setZeroes(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])

	//定义两个标记数组
	row := make([]bool, m)
	column := make([]bool, n)

	//遍历矩阵找0
	for i, r := range matrix {
		for j, val := range r {
			if val == 0 {
				row[i] = true
				column[j] = true
			}
		}
	}

	//用标记数组更新原矩阵
	for i, r := range matrix {
		for j := range r {
			//更新
			if row[i] || column[j] {
				r[j] = 0
			}
		}
	}
}

func maxValue(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += max(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
