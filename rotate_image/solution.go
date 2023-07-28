package rotate_image

/*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

提示：

n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000

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
