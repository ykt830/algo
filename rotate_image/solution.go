package rotate_image

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

提示：
m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-109 <= matrix[i][j] <= 109
每行的所有元素从左到右升序排列
每列的所有元素从上到下升序排列
-109 <= target <= 109
*/

func rotate(matrix [][]int) {
	start := 0
	length := len(matrix) - 1
	for start <= length {
		cycle(matrix, start)
		start++
		length--
	}
}

func cycle(matrix [][]int, start int) {
	length := len(matrix) - start
	for i := start; i < length-1; i++ {
		h, l := start, i
		tmp := matrix[h][l]
		for j := 0; j < 4; j++ {
			newH := l
			newL := len(matrix) - 1 - h

			matrix[newH][newL], tmp = tmp, matrix[newH][newL]
			h, l = newH, newL
		}
	}
}
