package spiral_matrix

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

func spiralOrder(matrix [][]int) []int {
	hLen, lLen := len(matrix), len(matrix[0])
	start := 0
	ret := make([]int, 0)
	for start < hLen && start < lLen {
		ret = append(ret, cycle(matrix, start)...)
		start++
		hLen--
		lLen--
	}

	return ret
}

func cycle(matrix [][]int, start int) []int {
	ret := make([]int, 0)
	hLen, lLen := len(matrix)-start, len(matrix[0])-start
	i, j := start, start
	for ; j < lLen; j++ {
		ret = append(ret, matrix[i][j])
	}
	j--
	i++
	if i >= hLen {
		return ret
	}

	for ; i < hLen; i++ {
		ret = append(ret, matrix[i][j])
	}
	i--
	j--
	if j < start {
		return ret
	}

	for ; j >= start; j-- {
		ret = append(ret, matrix[i][j])
	}
	j++
	i--
	if i <= start {
		return ret
	}

	for ; i > start; i-- {
		ret = append(ret, matrix[i][j])
	}

	return ret
}
