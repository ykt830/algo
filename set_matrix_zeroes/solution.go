package set_matrix_zeroes

/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

提示：
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1


进阶：
一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个仅使用常量空间的解决方案吗？
*/

func setZeroes(matrix [][]int) {
	h, l := make(map[int]bool), make(map[int]bool)
	hLen, lLen := len(matrix), len(matrix[0])

	for i := 0; i < hLen; i++ {
		for j := 0; j < lLen; j++ {
			if matrix[i][j] == 0 {
				h[i] = true
				l[j] = true
			}
		}
	}

	for k := range h {
		for i := 0; i < lLen; i++ {
			matrix[k][i] = 0
		}
	}

	for k := range l {
		for i := 0; i < hLen; i++ {
			matrix[i][k] = 0
		}
	}

	return
}
